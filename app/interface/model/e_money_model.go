package model

type LoginResponse struct {
	Username string
	UserId   string
}

type CreateUserReq struct {
	Username    string
	Password    string
	Email       string
	NoIdentitas string
	UserId      string
	TglLahir    string
	Nama        string
}

type DetailUserReq struct {
	Username string
	UserId   string
}

type DataUser struct {
	Username    string
	Nama        string
	UserId      string
	Email       string
	NoIdentitas string
	TglLahir    string
}

type GetBalance struct {
	UserId   string
	UserName string
	NoKartu  string
}

type CreateEmoney struct {
	UserId   string
	UserName string
	Saldo    string
}

type GetTrxHist struct {
	UserId   string
	UserName string
	NoKartu  string
}
