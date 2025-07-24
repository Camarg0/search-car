package services

// todos esses imports servem pra lidar com requisicoes http e json
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Camarg0/search-car-api/internal/models"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

// definição de structs para Requisição e Resposta

type OpenAIMessage struct {
	// cada mensagem tem um papel: pode ser user, assistant ou developer (documentação explica)
	Role string `json:"role"`
	// conteudo da mensagem
	Content string `json:"content"`
}

type OpenAIRequest struct {
	// modelo do GPT utilizado
	Model string `json:"model"`
	// slice de mensagens enviadas
	Messages []OpenAIMessage `json:"messages"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message OpenAIMessage `json:"message"`
	} `json:"choices"`
}

// Função que recebe o modelo do carro fornecido pelo usuário e busca a informação na OpenAI, retornando o json já convertido no model carInfo
func GetCarInfoFromOpenAI(carModel string) (*models.CarInfo, error) {
	rawPrompt, err := GetPrompt("../../prompts/car_info_prompt.txt")

	if err != nil {
		return nil, err
	}

	// Sprintf é usada para formatar a string retornando pra uma variavel, mas sem imprimir no console
	prompt := fmt.Sprintf(rawPrompt, carModel)

	requestBody := OpenAIRequest{
		Model: "gpt-4.1-nano",
		Messages: []OpenAIMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	// Marshaling do prompt - Serialização da request para JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Inicializo a minha nova http request
	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("nao foi possivel achar a chave secreta da api")
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Inicializo o meu client do httpRequest para enviar a requisição
	client := &http.Client{} // struct do client utilizando todas as opções default
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Fechar o corpo da resposta final e a conexão para evitar qualquer leaking de memória
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// A variável err não pega erros de http, somente se houver uma falha no transporte
	if response.StatusCode != http.StatusOK {
		fmt.Println("Erro da OpenAI:", string(body))
		return nil, fmt.Errorf("OpenAI retornou erro %d: %s", response.StatusCode, string(body))
	}

	var openAiResp = OpenAIResponse{}

	// Deserializando o json e alocando o resultado para a minha Response
	err = json.Unmarshal(body, &openAiResp)
	if err != nil {
		return nil, err
	}

	// Validando o meu número de structs de choices dentro da minha Response
	if len(openAiResp.Choices) == 0 {
		return nil, fmt.Errorf("nenhum resultado encontrado pela IA")
	}

	rawContent := openAiResp.Choices[0].Message.Content

	// Limpar a string do json
	cleanJson, err := strconv.Unquote(`"` + rawContent + `"`)
	if err != nil {
		cleanJson = rawContent
	}

	var carInfo models.CarInfo
	// Deserializar o json no model carInfo
	err = json.Unmarshal([]byte(cleanJson), &carInfo)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter o json do resultado para o model carInfo %v", err)
	}

	return &carInfo, nil
}

func GetPrompt(filePath string) (string, error) {
	prompt, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo de prompt: %w", err)
	}

	return string(prompt), nil
}
