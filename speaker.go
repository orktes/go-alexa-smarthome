package smarthome

import "fmt"

type SetMuteRequest struct {
	Mute bool `json:"mute"`
}

type SetVolumeRequest struct {
	Volume int `json:"volume"`
}

type AdjustVolumeRequest struct {
	VolumeDelta  int  `json:"volume"`
	VolumeDefaut bool `json:"volumeDefault"`
}

type speaker struct {
	sm *Smarthome
}

func (s *speaker) SetMute(endpoint Endpoint, payload SetMuteRequest) (resp EndpointResponse, err error) {
	return s.sm.setPropertyStatesAndCreateEndpointResponse(
		endpoint,
		map[string]map[string]interface{}{
			"Alexa.Speaker": map[string]interface{}{
				"mute": payload.Mute,
			},
		},
		map[string]map[string]interface{}{
			"Alexa.Speaker": map[string]interface{}{
				"mute": payload.Mute,
			},
			"Alexa.EndpointHealth": map[string]interface{}{
				"connectivity": map[string]interface{}{
					"value": "OK",
				},
			},
		})
}

func (s *speaker) SetVolume(endpoint Endpoint, payload SetVolumeRequest) (resp EndpointResponse, err error) {
	return s.sm.setPropertyStatesAndCreateEndpointResponse(
		endpoint,
		map[string]map[string]interface{}{
			"Alexa.Speaker": map[string]interface{}{
				"volume": payload.Volume,
			},
		},
		map[string]map[string]interface{}{
			"Alexa.Speaker": map[string]interface{}{
				"volume": payload.Volume,
			},
			"Alexa.EndpointHealth": map[string]interface{}{
				"connectivity": map[string]interface{}{
					"value": "OK",
				},
			},
		})
}

func (s *speaker) AdjustVolume(endpoint Endpoint, payload *AdjustVolumeRequest) (resp EndpointResponse, err error) {
	device := s.sm.GetDevice(endpoint.EndpointID)
	if device == nil {
		return resp, fmt.Errorf("Could not find endpoint with id %s", endpoint.EndpointID)
	}

	capability := device.GetCapabilityHandler("Alexa.Speaker")
	if capability == nil {
		return resp, fmt.Errorf("%s doesnt implement Speaker", endpoint.EndpointID)
	}

	var volume int
	if prop, ok := capability.propertyHandlers["volume"]; ok {
		val, err := prop.GetValue()
		if err != nil {
			return resp, err
		}

		volume = val.(int)
	}

	volume += payload.VolumeDelta
	if volume > 100 {
		volume = 100
	} else if volume < 0 {
		volume = 0
	}

	return s.sm.setPropertyStatesAndCreateEndpointResponse(
		endpoint,
		map[string]map[string]interface{}{
			"Alexa.Speaker": map[string]interface{}{
				"volume": volume,
			},
		},
		map[string]map[string]interface{}{
			"Alexa.Speaker": map[string]interface{}{
				"volume": volume,
			},
			"Alexa.EndpointHealth": map[string]interface{}{
				"connectivity": map[string]interface{}{
					"value": "OK",
				},
			},
		})
}
