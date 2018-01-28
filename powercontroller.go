package smarthome

import "time"

type powerController struct {
	sm *Smarthome
}

func (pc *powerController) TurnOn(endpoint Endpoint) (resp EndpointResponse, err error) {
	device := pc.sm.GetDevice(endpoint.EndpointID)
	c := device.GetCapabilityHandler("Alexa.PowerController")
	if c != nil && c.propertyHandlers["powerState"] != nil {
		if err := c.propertyHandlers["powerState"].SetValue("ON"); err != nil {
			return resp, err
		}
	}

	resp.Properties = append(resp.Properties, Property{
		Namespace:                 "Alexa.PowerController",
		Name:                      "powerState",
		Value:                     "ON",
		TimeOfSample:              zuluTime{time.Now()},
		UncertaintyInMilliseconds: 0,
	})

	// TODO really check this
	resp.Properties = append(resp.Properties, Property{
		Namespace: "Alexa.EndpointHealth",
		Name:      "connectivity",
		Value: map[string]interface{}{
			"value": "OK",
		},
		TimeOfSample:              zuluTime{time.Now()},
		UncertaintyInMilliseconds: 0,
	})

	return resp, nil
}

func (pc *powerController) TurnOff(endpoint Endpoint) (resp EndpointResponse, err error) {
	device := pc.sm.GetDevice(endpoint.EndpointID)
	c := device.GetCapabilityHandler("Alexa.PowerController")
	if c != nil && c.propertyHandlers["powerState"] != nil {
		if err := c.propertyHandlers["powerState"].SetValue("OFF"); err != nil {
			return resp, err
		}
	}

	resp.Properties = append(resp.Properties, Property{
		Namespace:                 "Alexa.PowerController",
		Name:                      "powerState",
		Value:                     "ON",
		TimeOfSample:              zuluTime{time.Now()},
		UncertaintyInMilliseconds: 200,
	})

	// TODO really check this
	resp.Properties = append(resp.Properties, Property{
		Namespace: "Alexa.EndpointHealth",
		Name:      "connectivity",
		Value: map[string]interface{}{
			"value": "OK",
		},
		TimeOfSample:              zuluTime{time.Now()},
		UncertaintyInMilliseconds: 200,
	})

	return resp, nil
}
