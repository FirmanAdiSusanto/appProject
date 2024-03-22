package handler

type LoginResponse struct {
	Fullname string `json:"fullname"`
	Token    string `json:"token"`
}
