package entity

type BillerResponse struct {
	Code       string    `json:"code"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	ListBiller []*Biller `json:"data"`
}

type Biller struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Fee         string `json:"fee"`
}
