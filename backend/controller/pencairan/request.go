package pencairan

type DataRequest struct {
	Branch            string `json:"branch"`
	ChannelingCompany string `json:"channeling_company"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	PPK               string `json:"ppk"`
}

type DataPostPPK struct {
	PPK []string `json:"ppk"`
}
