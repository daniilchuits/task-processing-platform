package transportation

type JWT struct {
	Token string `json:"token"`
}

func NewJWT(jwt string) JWT {
	return JWT{Token: jwt}
}
