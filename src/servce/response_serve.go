package servce

var code = map[string]uint{
	"SECCEED": 200,
	"ERROR":   400,
}

// ResResult 响应体
type ResResult struct {
	code uint
	msg  string
	data interface{}
}

func (response *ResResult) GetData() interface{} {
	return response.data
}

func ResSucceed(data interface{}) ResResult {
	return ResResult{
		code: code["SECCEED"],
		msg:  "succeed",
		data: data,
	}
}

func ResErr(msg string, state string) ResResult {
	return ResResult{
		code: code[state],
		msg:  msg,
		data: nil,
	}
}
