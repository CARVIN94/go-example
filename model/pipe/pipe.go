package pipe

// Pipe 数据传递格式
type Model struct {
	State   int         `json:"state"`
	Message interface{} `json:"message"`
}
