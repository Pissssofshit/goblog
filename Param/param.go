package param

type LoginParam struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"pass_word" binding:"required"`
}

type ResponseStruct struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

func NewResponseStruct() ResponseStruct {
	val := ResponseStruct{Code: 0, Msg: "success"}
	val.Data = make(map[string]interface{})
	return val
}
