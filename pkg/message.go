package pkg

type Publisher interface {
	// Never close the given msgChan
	Publish(msgChan chan<- PublishMessage)
	Name() string
}

type PublishMessage struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Extra   map[string]interface{} `json:"extra"`
}
