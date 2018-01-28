# go-alexa-smarthome

```go

package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    smarthome "github.com/orktes/go-alexa-smarthome"
)

type mockProperyHandler struct {
	val interface{}
}

func (mbph mockProperyHandler) GetValue() (interface{}, error) {
	return mbph.val, nil
}

func (mbph mockProperyHandler) SetValue(val interface{}) error {
	fmt.Printf("Received value %+v\n", val)
	mbph.val = val
	return nil
}

func (mbph mockProperyHandler) UpdateChannel() <-chan interface{} {
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
	capability.AddPropertyHandler("powerState", &mockProperyHandler{val: "ON"})

	sm.AddDevice(abstractTestDevice)

	lambda.Start(sm.Handle)
}




```