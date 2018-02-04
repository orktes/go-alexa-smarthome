package smarthome

import "time"

type SetColorTemperatureRequest struct {
	ColorTemperatureInKelvin int `json:"colorTemperatureInKelvin"`
}

type colorTemperatureController struct {
	sm *Smarthome
}

func (cc *colorTemperatureController) action(endpoint Endpoint, action string) (resp EndpointResponse, err error) {
	val, err := cc.sm.invokeAction(endpoint, "Alexa.ColorTemperatureController", action, nil)
	if err != nil {
		return resp, err
	}

	resp.Properties = append(resp.Properties, Property{
		Namespace:                 "Alexa.ColorTemperatureController",
		Name:                      "colorTemperatureInKelvin",
		Value:                     val,
		TimeOfSample:              zuluTime{time.Now()},
		UncertaintyInMilliseconds: 0,
	})
	resp.Properties = append(resp.Properties, Property{
		Namespace: "Alexa.EndpointHealth",
		Name:      "connectivity",
		Value: map[string]interface{}{
			"value": "OK",
		},
		TimeOfSample:              zuluTime{time.Now()},
		UncertaintyInMilliseconds: 0,
	})

	return
}

func (cc *colorTemperatureController) DecreaseColorTemperature(endpoint Endpoint) (resp EndpointResponse, err error) {
	return cc.action(endpoint, "DecreaseColorTemperature")
}

func (cc *colorTemperatureController) IncreaseColorTemperature(endpoint Endpoint) (resp EndpointResponse, err error) {
	return cc.action(endpoint, "IncreaseColorTemperature")
}

func (cc *colorTemperatureController) SetColorTemperature(endpoint Endpoint, payload SetColorTemperatureRequest) (resp EndpointResponse, err error) {
	return cc.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.ColorTemperatureController",
		"colorTemperatureInKelvin",
		payload.ColorTemperatureInKelvin,
	)
}
