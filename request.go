package smarthome

import "encoding/json"

type Request struct {
	Directive Directive `json:"directive"`
}

type Endpoint struct {
	Scope      Scope  `json:"scope"`
	EndpointID string `json:"endpointId"`
	Cookie     Cookie `json:"cookie"`
}

type Directive struct {
	Header   Header          `json:"header"`
	Payload  json.RawMessage `json:"payload"`
	Endpoint *Endpoint       `json:"Endpoint"`
}
