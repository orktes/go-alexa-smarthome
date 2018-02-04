package smarthome

type powerController struct {
	sm *Smarthome
}

func (pc *powerController) TurnOn(endpoint Endpoint) (resp EndpointResponse, err error) {
	return pc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.PowerController",
		"powerState",
		"ON",
	)
}

func (pc *powerController) TurnOff(endpoint Endpoint) (resp EndpointResponse, err error) {
	return pc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.PowerController",
		"powerState",
		"OFF",
	)
}
