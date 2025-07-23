package services

// todos esses imports servem pra lidar com requisicoes http e json
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		Message OpenAIMessage
	}
}

// Função que recebe o modelo do carro fornecido pelo usuário e busca a informação na OpenAI
func GetCarInfoFromOpenAI(carModel string) (string, error) {
	// Sprintf é usada para formatar a string retornando pra uma variavel, mas sem imprimir no console
	// Uso de crase invertida para string crua
	prompt := fmt.Sprintf(`
Tenho um sistema de busca de informações de carro. O carro que estou buscando informações é um "%s".
Preciso que você gere um JSON com as seguintes informações, de maneira estruturada. Precisa ser um texto interessante de ler, e sem informações muito técnicas. Caso haja termos técnicos, explicar de maneira geral sucintamente o que é:
- Modelo do carro (será passado por parâmetro via API pra OpenAI)
- Ano do carro
- Modelo do motor do carro
- Líquido de arrefecimento recomendado para o carro (inclusive com a sua respectiva cor)
- Fluido de freio recomendado para o carro
- Descrição breve do carro
- Problemas crônicos daquele modelo para ficar de olho
- Dicas de manutenção específicas para aquele modelo de carro
`, carModel)

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
		return "", err
	}

	// Inicializo a minha nova http request
	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_API_KEY"))

	// Inicializo o meu client do httpRequest para enviar a requisição
	client := &http.Client{} // struct do client utilizando todas as opções default
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Fechar o corpo da resposta final e a conexão para evitar qualquer leaking de memória
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var openAiResp = OpenAIResponse{}
	// Deserializando o json e alocando o resultado para a minha Response
	err = json.Unmarshal(body, &openAiResp)
	if err != nil {
		return "", err
	}

	// Validando o meu número de structs de choices dentro da minha Response
	if len(openAiResp.Choices) == 0 {
		return "", fmt.Errorf("nenhum resultado encontrado pela IA")
	}

	return openAiResp.Choices[0].Message.Content, nil
}
