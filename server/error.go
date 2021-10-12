package server

type Err struct {
	ErrCode 	int			`json:"errcode"`
	ErrMsg 		string 		`json:"errmsg"`
}
