package chat

type MessageFromUser struct {
	Action string `json:"action"`
	Data map[string]interface{} `json:"data"`
}

type MessageToUser struct {
	Action string `json:"action"`
	Data interface{} `json:"data"`
}


