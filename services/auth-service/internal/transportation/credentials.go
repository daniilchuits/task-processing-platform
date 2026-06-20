package transportation

type Credentials struct {
	Id       int    `json:"id,omitempty"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
