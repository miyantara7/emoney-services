package entity

type ListBillerResponse struct {
	Code       int       `json:"code"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	ListBiller []*Biller `json:"data"`
}

type DetailBillerResponse struct {
	Code         int     `json:"code"`
	Status       string  `json:"status"`
	Message      string  `json:"message"`
	DetailBiller *Biller `json:"data"`
}

type Biller struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Fee         int    `json:"fee"`
}

type BillerRequest struct {
	ID string `json:"id"`
}

type EMoney struct {
	BaseModel
	NoKartu    string `gorm:"column:no_kartu;size:20"`
	UserId     string `gorm:"column:user_id;size:100"`
	DataUserId string `gorm:"column:data_user_id;size:100"`
	Balance    *int   `gorm:"column:balance"`
}

func (EMoney) TableName() string {
	return "t_e_money"
}

type TransactionHistory struct {
	BaseModel
	UserId      string
	NoKartu     string
	CreatedDate string
	UpdateDate  string
	Setor       string
	Tarik       string
	Balance     string
}

func (TransactionHistory) TableName() string {
	return "t_trx_history"
}
