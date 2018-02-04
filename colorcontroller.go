package smarthome

type SetColorRequest struct {
	Color map[string]interface{} `json:"color"`
}

type colorController struct {
	sm *Smarthome
}

func (cc *colorController) SetColor(endpoint Endpoint, payload SetColorRequest) (resp EndpointResponse, err error) {
	return cc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.ColorController",
		"color",
		payload.Color,
	)
}
