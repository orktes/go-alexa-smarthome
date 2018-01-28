package smarthome

type DiscoverResponse struct {
	DiscoveryEndpoints []DiscoveryEndpoint `json:"endpoints"`
}

type DiscoveryEndpoint struct {
	EndpointID        string       `json:"endpointId"`
	ManufacturerName  string       `json:"manufacturerName"`
	FriendlyName      string       `json:"friendlyName"`
	Description       string       `json:"description"`
	DisplayCategories []string     `json:"displayCategories,omitempty"`
	Cookie            Cookie       `json:"cookie"`
	Capabilities      []Capability `json:"capabilities,omitempty"`
}

type Capability struct {
	Type                       string                      `json:"type"`
	Interface                  string                      `json:"interface"`
	Version                    string                      `json:"version"`
	Properties                 *Properties                 `json:"properties,omitempty"`
	SupportsDeactivation       *bool                       `json:"supportsDeactivation,omitempty"`
	ProactivelyReported        *bool                       `json:"proactivelyReported,omitempty"`
	CameraStreamConfigurations []CameraStreamConfiguration `json:"cameraStreamConfigurations,omitempty"`
}

type CameraStreamConfiguration struct {
	Protocols          []string     `json:"protocols"`
	Resolutions        []Resolution `json:"resolutions"`
	AuthorizationTypes []string     `json:"authorizationTypes"`
	VideoCodecs        []string     `json:"videoCodecs"`
	AudioCodecs        []string     `json:"audioCodecs"`
}

type Resolution struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

type Properties struct {
	Supported           []Supported `json:"supported"`
	ProactivelyReported bool        `json:"proactivelyReported"`
	Retrievable         bool        `json:"retrievable"`
}

type Supported struct {
	Name string `json:"name"`
}

type Cookie struct {
	Detail1 *string `json:"detail1,omitempty"`
	Detail2 *string `json:"detail2,omitempty"`
}

type DiscoverRequest struct {
	Scope Scope `json:"scope"`
}

type discovery struct {
	sm *Smarthome
}

func (d *discovery) Discover(req DiscoverRequest) (interface{}, error) {
	d.sm.Lock()
	defer d.sm.Unlock()

	res := DiscoverResponse{}

	for _, device := range d.sm.devices {
		res.DiscoveryEndpoints = append(
			res.DiscoveryEndpoints,
			DiscoveryEndpoint{
				EndpointID:        device.ID(),
				Cookie:            device.Cookie(),
				Capabilities:      device.Capabilities(),
				ManufacturerName:  device.ManufacturerName(),
				FriendlyName:      device.FriendlyName(),
				Description:       device.Description(),
				DisplayCategories: device.DisplayCategories(),
			},
		)
	}

	return res, nil
}
