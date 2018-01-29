package smarthome

type powerController struct {
	sm *Smarthome
}

func (pc *powerController) TurnOn(endpoint Endpoint) (resp EndpointResponse, err error) {
	return pc.sm.setPropertyStatesAndCreateEndpointResponse(
		endpoint,
		map[string]map[string]interface{}{
			"Alexa.PowerController": map[string]interface{}{
				"powerState": "ON",
			},
		},
		map[string]map[string]interface{}{
			"Alexa.PowerController": map[string]interface{}{
				"powerState": "ON",
			},
			"Alexa.EndpointHealth": map[string]interface{}{
				"connectivity": map[string]interface{}{
					"value": "OK",
				},
			},
		})
}

func (pc *powerController) TurnOff(endpoint Endpoint) (resp EndpointResponse, err error) {
	return pc.sm.setPropertyStatesAndCreateEndpointResponse(
		endpoint,
		map[string]map[string]interface{}{
			"Alexa.PowerController": map[string]interface{}{
				"powerState": "OFF",
			},
		},
		map[string]map[string]interface{}{
			"Alexa.PowerController": map[string]interface{}{
				"powerState": "OFF",
			},
			"Alexa.EndpointHealth": map[string]interface{}{
				"connectivity": map[string]interface{}{
					"value": "OK",
				},
			},
		})
}
