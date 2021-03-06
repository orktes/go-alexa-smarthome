# go-alexa-smarthome


[![GoDoc](https://godoc.org/github.com/orktes/go-alexa-smarthome?status.svg)](https://godoc.org/github.com/orktes/go-alexa-smarthome)


```go

package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    smarthome "github.com/orktes/go-alexa-smarthome"
)

type mockPropertyHandler struct {
	val interface{}
}

func (mockHandler *mockPropertyHandler) GetValue() (interface{}, error) {
    fmt.Printf("Getting value %+v\n", mockHandler.val)
	return mockHandler.val, nil
}

func (mockHandler *mockPropertyHandler) SetValue(val interface{}) error {
	fmt.Printf("Received value %+v\n", val)
	mockHandler.val = val
	return nil
}

func (mockHandler *mockPropertyHandler) UpdateChannel() <-chan interface{} {
	return nil
}

func main() {
	sm := smarthome.New(smarthome.AuthorizationFunc(func (req smarthome.AcceptGrantRequest) error {
	    return nil
	}))

	abstractTestDevice := smarthome.NewAbstractDevice(
		"1",
		"Fake switch",
		"Homeautomation",
		"Just a fake light",
	)
	abstractTestDevice.AddDisplayCategory("SWITCH")
	capability := abstractTestDevice.NewCapability("PowerController")
	capability.AddPropertyHandler("powerState", &mockPropertyHandler{val: "ON"})

	sm.AddDevice(abstractTestDevice)

	lambda.Start(sm.Handle)
}

```

## TODO

### Controllers

 - [x] Authorization 
 - [ ] CameraStreamController 
 - [ ] ChannelController 
 - [x] ColorTemperatureController 
 - [x] Discovery 
 - [ ] InputController 
 - [x] PercentageController 
 - [x] PowerController 
 - [ ] SceneController 
 - [x] StateReport 
 - [ ] TemperatureSensor 
 - [x] BrightnessController 
 - [x] ChangeReport 
 - [x] ColorController 
 - [ ] DeferredResponse 
 - [x] ErrorResponse 
 - [ ] LockController 
 - [ ] PlaybackController 
 - [ ] PowerLevelController 
 - [x] Speaker 
 - [ ] StepSpeaker 
 - [ ] ThermostatController

### High level APIs (device abstractions)
- [ ] Light
- [ ] Switch
- [ ] TV
- [ ] Amplifier
- [ ] Speaker
- [ ] Temperature sensor
