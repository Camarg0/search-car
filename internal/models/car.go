package models

type CarInfo struct {
	Model           string   `json:"model"`
	Year            string   `json:"year"`
	Engine          string   `json:"engine"`
	Coolant         string   `json:"coolant"`
	BrakeFluid      string   `json:"brakeFluid"`
	Description     string   `json:"description"`
	ChronicIssues   []string `json:"chronicIssues"`
	MaintenanceTips []string `json:"maintenanceTips"`
	SharedParts     []string `json:"sharedParts"`
	//Images          []string `json:"images"`
}
