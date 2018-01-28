package smarthome

type EndpointResponse struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	Namespace                 string      `json:"namespace"`
	Name                      string      `json:"name"`
	Value                     interface{} `json:"value"`
	TimeOfSample              zuluTime    `json:"timeOfSample"`
	UncertaintyInMilliseconds int64       `json:"uncertaintyInMilliseconds"`
}
