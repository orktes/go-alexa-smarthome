# go-alexa-smarthome

```go

package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    smarthome "github.com/orktes/go-alexa-smarthome"
)

type mockPropertyHandler struct {
	val interface{}
}

func (mockHandler mockPropertyHandler) GetValue() (interface{}, error) {
	return mockHandler.val, nil
}

func (mockHandler mockPropertyHandler) SetValue(val interface{}) error {
	fmt.Printf("Received value %+v\n", val)
	mockHandler.val = val
	return nil
}

func (mockHandler mockPropertyHandler) UpdateChannel() <-chan interface{} {
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