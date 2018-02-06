package smarthome

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
	return s.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.Speaker",
		"mute",
		payload.Mute,
	)

}

func (s *speaker) SetVolume(endpoint Endpoint, payload SetVolumeRequest) (resp EndpointResponse, err error) {
	return s.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.Speaker",
		"volume",
		payload.Volume,
	)

}

func (s *speaker) AdjustVolume(endpoint Endpoint, payload *AdjustVolumeRequest) (resp EndpointResponse, err error) {

	var volume int
	if val, err := s.sm.getValueForProperty(endpoint, "Alexa.Speaker", "volume"); err == nil {
		volume = val.(int)
	}

	volume += payload.VolumeDelta
	if volume > 100 {
		volume = 100
	} else if volume < 0 {
		volume = 0
	}

	return s.sm.setPropAndEndpointHealthResponse(
		endpoint,
		"Alexa.Speaker",
		"volume",
		volume,
	)
}
