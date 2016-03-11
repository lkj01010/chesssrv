package dao

type Args struct {
	Id     string
	Int    int
	String string
}

type Reply struct {
	Code   string `json:"code"`
	Int    int `json:"int"`
	String string `json:"string"`
}
