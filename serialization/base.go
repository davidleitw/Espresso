package serialization

type BaseResponse interface {
	BuildResponse()
	GetResStatus() int
	GetResData() interface{}
	GetResMessage() string
}

type Response map[string]interface{}

func BuildResponse(status int, data interface{}, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

// type Response struct {
// 	status  int    `swaggertype: "int" example:"status code"`
// 	message string `swaggertype: "string example:"message"`
// 	data    interface{}
// }

// func BuildResponse(status int, data interface{}, message string) Response {
// 	return Response{
// 		status:  status,
// 		message: message,
// 		data:    data,
// 	}
// }

func GetResStatus(res Response) interface{} {
	//return res.status
	return res["status"]
}

func GetResData(res Response) interface{} {
	//return res.data
	return res["data"]
}

func GetResMessage(res Response) interface{} {
	//return res.message
	return res["message"]
}
