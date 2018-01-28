package smarthome

type Response struct {
	Event   Event       `json:"event"`
	Context interface{} `json:"context,omitempty"`
}

type Event struct {
	Header   Header      `json:"header,omitempty"`
	Payload  interface{} `json:"payload,omitempty"`
	Endpoint *Endpoint   `json:"endpoint,omitempty"`
}
