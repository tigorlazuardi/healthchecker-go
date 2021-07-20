package pkg

type PublishMessage struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Extra   map[string]interface{} `json:"extra"`
}
