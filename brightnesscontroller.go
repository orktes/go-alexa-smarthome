package smarthome

import "fmt"

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
	return bc.sm.setPropertyStatesAndCreateEndpointResponse(
		endpoint,
		map[string]map[string]interface{}{
			"Alexa.BrightnessController": map[string]interface{}{
				"brightness": payload.Brightness,
			},
		},
		map[string]map[string]interface{}{
			"Alexa.BrightnessController": map[string]interface{}{
				"brightness": payload.Brightness,
			},
			"Alexa.EndpointHealth": map[string]interface{}{
				"connectivity": map[string]interface{}{
					"value": "OK",
				},
			},
		})
}

func (bc *brightnessController) AdjustBrightness(endpoint Endpoint, payload *AdjustBrightnessRequest) (resp EndpointResponse, err error) {
	device := bc.sm.GetDevice(endpoint.EndpointID)
	if device == nil {
		return resp, fmt.Errorf("Could not find endpoint with id %s", endpoint.EndpointID)
	}

	capability := device.GetCapabilityHandler("Alexa.BrightnessController")
	if capability == nil {
		return resp, fmt.Errorf("%s doesnt implement BrightnessController", endpoint.EndpointID)
	}

	var brightness int
	if prop, ok := capability.propertyHandlers["brightness"]; ok {
		val, err := prop.GetValue()
		if err != nil {
			return resp, err
		}

		brightness = val.(int)
	}

	brightness += payload.BrightnessDelta
	if brightness > 100 {
		brightness = 100
	} else if brightness < 0 {
		brightness = 0
	}

	return bc.sm.setPropertyStatesAndCreateEndpointResponse(
		endpoint,
		map[string]map[string]interface{}{
			"Alexa.BrightnessController": map[string]interface{}{
				"brightness": brightness,
			},
		},
		map[string]map[string]interface{}{
			"Alexa.BrightnessController": map[string]interface{}{
				"brightness": brightness,
			},
			"Alexa.EndpointHealth": map[string]interface{}{
				"connectivity": map[string]interface{}{
					"value": "OK",
				},
			},
		})
}
