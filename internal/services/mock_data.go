package services

import "github.com/Camarg0/search-car-api/internal/models"

const car = "civic-vti"

var mockCarData = map[string]models.CarInfo{ //map["civic-vti"]Struct (chave:valor)
	car: {
		Model:       "car",
		Year:        "1993",
		Engine:      "B16A2",
		Coolant:     "Verde",
		BrakeFluid:  "DOT4",
		Description: "Lindo carro",
		ChronicIssues: []string{
			"Teto solar emperrado",
			"Funcionamento do carro em baixa rotação",
		},
		MaintenanceTips: []string{
			"Mantenha os giros altos de vez em quando",
			"Trocar óleo a cada 5000km",
		},
		// Images: []string{
		// 	"https://quatrorodas.abril.com.br/wp-content/uploads/2016/11/5772db9f0e216345751b7a61qr-682-cclassic-civic-0278-tif.jpeg?crop=1&resize=1212,909",
		// },
	},
}

func GetMockedCarInfo(car string) (*models.CarInfo, bool) {
	mockedCar, exists := mockCarData[car]

	if !exists {
		return nil, false
	}

	return &mockedCar, true
}
