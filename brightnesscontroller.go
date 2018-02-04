package smarthome

type SetBrightnessRequest struct {
	Brightness int `json:"brightness"`
}

type AdjustBrightnessRequest struct {
	BrightnessDelta int `json:"brightnessDelta"`
}

type brightnessController struct {
	sm *Smarthome
}

func (bc *brightnessController) SetBrightness(endpoint Endpoint, payload *SetBrightnessRequest) (resp EndpointResponse, err error) {
	return bc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.BrightnessController",
		"brightness",
		payload.Brightness,
	)
}

func (bc *brightnessController) AdjustBrightness(endpoint Endpoint, payload *AdjustBrightnessRequest) (resp EndpointResponse, err error) {
	var brightness int

	if val, err := bc.sm.getValueForProperty(endpoint, "Alexa.BrightnessController", "brightness"); err != nil {
		brightness = val.(int)
	}

	brightness += payload.BrightnessDelta
	if brightness > 100 {
		brightness = 100
	} else if brightness < 0 {
		brightness = 0
	}

	return bc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.BrightnessController",
		"brightness",
		brightness,
	)
}
