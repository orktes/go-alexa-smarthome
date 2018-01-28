package smarthome

import (
	"encoding/json"
	"reflect"
	"sync"

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
		"Alexa":                 &alexa{sh},
		"Alexa.Authorization":   &authorization{sh},
		"Alexa.Discovery":       &discovery{sh},
		"Alexa.PowerController": &powerController{sh},
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
		resName = "ErrorResponse"
		payload = errReturn
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
