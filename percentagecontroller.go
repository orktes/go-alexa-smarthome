package smarthome

type SetPercentageRequest struct {
	Percentage int `json:"percentage"`
}

type AdjustPercentageRequest struct {
	PercentageDelta int `json:"percentageDelta"`
}

type percentageController struct {
	sm *Smarthome
}

func (bc *percentageController) SetPercentage(endpoint Endpoint, payload *SetPercentageRequest) (resp EndpointResponse, err error) {
	return bc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.PercentageController",
		"percentage",
		payload.Percentage,
	)
}

func (bc *percentageController) AdjustPercentage(endpoint Endpoint, payload *AdjustPercentageRequest) (resp EndpointResponse, err error) {
	var percentage int

	if val, err := bc.sm.getValueForProperty(endpoint, "Alexa.PercentageController", "percentage"); err != nil {
		percentage = val.(int)
	}

	percentage += payload.PercentageDelta
	if percentage > 100 {
		percentage = 100
	} else if percentage < 0 {
		percentage = 0
	}

	return bc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.PercentageController",
		"percentage",
		percentage,
	)
}
