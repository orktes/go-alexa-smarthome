package smarthome

import "time"

type alexa struct {
	sh *Smarthome
}

func (a *alexa) ReportState(endpoint Endpoint) (resp EndpointResponse, err error) {
	device := a.sh.GetDevice(endpoint.EndpointID)
	if device == nil {
		return
	}

	for _, capability := range device.GetCapabilityHandlers() {
		for name, propHander := range capability.propertyHandlers {
			val, err := propHander.GetValue()
			if err != nil {
				// TODO log error
				continue
			}

			resp.Properties = append(resp.Properties, Property{
				Namespace:                 capability.Interface,
				Name:                      name,
				Value:                     val,
				TimeOfSample:              zuluTime{time.Now()},
				UncertaintyInMilliseconds: 0,
			})
		}
	}
	return
}
