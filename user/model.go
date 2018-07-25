package user

type AuthResponse struct {
	Token string `json:"token"`
	User Model   `json:"user"`
}

type Model struct {
	ID       string `json:"id"`
	Phone    string `json:"-"`
	Name 	 string  `json:"name"`
	Avatar 	 string  `json:"avatar"`
	FBId     string  `json:"-"`
	Code     int    `json:"-"`
	Secret string `json:"-"`
}


type List struct {
	Objects []Model `json:"objects"`
	Total int		`json:"total"`
}
