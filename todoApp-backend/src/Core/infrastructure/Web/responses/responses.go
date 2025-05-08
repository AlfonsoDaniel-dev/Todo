package responses

type Response struct {
	MessageType string      `json:"messageType"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(messageType string, Message string, Data interface{}) Response {
	return Response{
		MessageType: messageType,
		Message:     Message,
		Data:        Data,
	}
}
