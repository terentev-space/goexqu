package handler_service_entities

//Format: {"urls": ["", ""], "data": "{}", "options": {}}
type WebhookMessageEntity struct {
	URLs    []string `json:"urls"`
	Data    string   `json:"data"`
	Options struct {
	} `json:"options"`
}
