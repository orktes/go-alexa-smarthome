package smarthome

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Smarthome struct {
	devices map[string]Device
	auth    AuthorizationInterface

	controllers map[string]interface{}

	sync.Mutex
}

func New(auth AuthorizationInterface) (sh *Smarthome) {

	sh = &Smarthome{
		auth:    auth,
		devices: map[string]Device{},
	}

	sh.controllers = map[string]interface{}{
		"Alexa":                      &alexa{sh},
		"Alexa.Authorization":        &authorization{sh},
		"Alexa.Discovery":            &discovery{sh},
		"Alexa.PowerController":      &powerController{sh},
		"Alexa.BrightnessController": &brightnessController{sh},
	}

	return
}

func (s *Smarthome) AddDevice(d Device) {
	s.Lock()
	defer s.Unlock()

	s.devices[d.ID()] = d
}

func (s *Smarthome) GetDevice(id string) Device {
	s.Lock()
	defer s.Unlock()

	return s.devices[id]
}

func (s *Smarthome) Handle(req *Request) *Response {
	namespace := req.Directive.Header.Namespace
	controller, ok := s.controllers[namespace]
	if !ok {
		// TODO send unknown namespace error
		println("TODO")
		return nil
	}

	name := req.Directive.Header.Name

	t := reflect.TypeOf(controller)
	m, ok := t.MethodByName(name)
	if !ok {
		println("TODO", "no such method", name)
		// TODO send unknown name error
		return nil
	}

	vals := []reflect.Value{reflect.ValueOf(controller)}

	if req.Directive.Endpoint != nil && m.Type.NumIn() >= 2 {
		argumentType := m.Type.In(1)
		if argumentType == reflect.TypeOf(*req.Directive.Endpoint) {
			vals = append(vals, reflect.ValueOf(*req.Directive.Endpoint))
		}
	}

	if m.Type.NumIn() > len(vals) {
		argumentType := m.Type.In(m.Type.NumIn() - 1)
		val := reflect.New(argumentType)
		err := json.Unmarshal(req.Directive.Payload, val.Interface())
		if err != nil {
			panic("TODO correct way to respond")
		}
		vals = append(vals, val.Elem())
	}

	res := m.Func.Call(vals)

	var payload interface{}
	var context interface{}

	resNamespace := "Alexa"
	resName := "Response"

	// Well this is an odd place for this. Bad API design from amazon or my own brain fart?
	if name == "ReportState" {
		resName = "StateReport"
	}

	if req.Directive.Endpoint == nil {
		resName = name + ".Response"
		payload = res[0].Interface()
		resNamespace = req.Directive.Header.Namespace

	} else {
		context = res[0].Interface()
		payload = map[string]interface{}{}
	}

	errReturn := res[1].Interface()
	if errReturn != nil {
		if err, ok := errReturn.(*sherror); ok {
			payload = map[string]interface{}{
				"type":    err.Type,
				"message": err.Message,
			}
		} else {
			payload = map[string]interface{}{
				"type":    "UNKNOWN",
				"message": fmt.Sprintf("%s", errReturn),
			}
		}
		context = nil
	}

	return &Response{
		Event: Event{
			Header: Header{
				Namespace:        resNamespace,
				Name:             resName,
				CorrelationToken: req.Directive.Header.CorrelationToken,
				PayloadVersion:   "3",
				MessageID:        uuid.New().String(),
			},
			Payload:  payload,
			Endpoint: req.Directive.Endpoint,
		},
		Context: context,
	}
}

func (s *Smarthome) setPropertyStatesAndCreateEndpointResponse(endpoint Endpoint, set map[string]map[string]interface{}, respond map[string]map[string]interface{}) (resp EndpointResponse, err error) {
	device := s.GetDevice(endpoint.EndpointID)
	if device == nil {
		return resp, fmt.Errorf("Could not find endpoint with id %s", endpoint.EndpointID)
	}

	for capabilityName, properties := range set {
		c := device.GetCapabilityHandler(capabilityName)
		for propName, propValue := range properties {
			if c != nil && c.propertyHandlers[propName] != nil {
				if err := c.propertyHandlers[propName].SetValue(propValue); err != nil {
					return resp, err
				}
			}
		}
	}

	for capabilityName, properties := range respond {
		for propName, propValue := range properties {
			resp.Properties = append(resp.Properties, Property{
				Namespace:                 capabilityName,
				Name:                      propName,
				Value:                     propValue,
				TimeOfSample:              zuluTime{time.Now()},
				UncertaintyInMilliseconds: 0,
			})
		}
	}

	return resp, nil
}
