package models

type CarInfo struct {
	Model           string   `json: model`
	Engine          string   `json: engine`
	Year            string   `json: year`
	Coolant         string   `json: coolant`
	BrakeFluid      string   `json: brakeFluid`
	Description     []string `json: description`
	ChronicIssues   []string `json: chronicIssues`
	MaintenanceTips []string `json: maintenanceTips`
	//Images          []string `json: images`
	// Peças compartilhasdas com outros carros
}
