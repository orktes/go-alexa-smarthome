package smarthome

type Header struct {
	Namespace        string `json:"namespace"`
	Name             string `json:"name"`
	PayloadVersion   string `json:"payloadVersion"`
	MessageID        string `json:"messageId"`
	CorrelationToken string `json:"correlationToken"`
}
