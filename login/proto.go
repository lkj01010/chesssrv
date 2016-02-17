package login

type ReqLogin struct {
	Account string `json:"account"`
	PswMd5 string `json:"pswmd5"`
}

