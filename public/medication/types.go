package medication

type AIResponse struct {
	DoctorName  string `json:"doctor_name"`
	Date        string `json:"date"`
	Medications []struct {
		ID          uint    `json:"id"`
		DrugName    string `json:"drug_name"`
		Amount      string `json:"amount_per_dose"`
		Schedule    string `json:"schedule"`
		Duration    string `json:"duration"`
		Instructions string `json:"instructions"`
	} `json:"medications"`
}