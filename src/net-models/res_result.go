package net_models

var Code = map[string]uint{
	"SECCEED": 200,
	"ERROR":   400,
}

// ResResult 响应体
type ResResult struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (response *ResResult) GetData() interface{} {
	return response.Data
}

func ResSucceed(data interface{}) ResResult {
	return ResResult{
		Code: Code["SECCEED"],
		Msg:  "succeed",
		Data: data,
	}
}

func ResErr(msg string, state string) ResResult {
	return ResResult{
		Code: Code[state],
		Msg:  msg,
		Data: nil,
	}
}

func Pending(data interface{}) ResResult {
	return ResResult{
		Code: 0,
		Msg:  "pending",
		Data: data,
	}
}
