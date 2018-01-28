package smarthome

import "strings"

type Device interface {
	ID() string
	Cookie() Cookie
	Capabilities() []Capability
	ManufacturerName() string
	FriendlyName() string
	Description() string
	DisplayCategories() []string

	GetCapabilityHandler(typ string) *CapabilityHandler
	GetCapabilityHandlers() []*CapabilityHandler
}

type AbstractDevice struct {
	id                string
	cookie            Cookie
	manufacturerName  string
	friendlyName      string
	description       string
	displayCategories []string

	capabilityHandlers []*CapabilityHandler
}

func NewAbstractDevice(id string, name string, manafactureName string, description string) *AbstractDevice {
	ad := &AbstractDevice{
		id:               id,
		friendlyName:     name,
		manufacturerName: manafactureName,
		description:      description,
	}

	ad.NewCapability("")

	return ad
}

func (ad *AbstractDevice) NewCapability(interf string) *CapabilityHandler {
	interfaceNameParts := []string{"Alexa"}

	if interf != "" {
		interfaceNameParts = append(interfaceNameParts, interf)
	}

	ch := &CapabilityHandler{
		Type:        "AlexaInterface",
		Interface:   strings.Join(interfaceNameParts, "."),
		Version:     "3",
		Retrievable: true,
	}

	ad.capabilityHandlers = append(ad.capabilityHandlers, ch)

	return ch
}

func (ad *AbstractDevice) GetCapabilityHandlers() []*CapabilityHandler {
	return ad.capabilityHandlers
}

func (ad *AbstractDevice) GetCapabilityHandler(typ string) *CapabilityHandler {
	for _, c := range ad.capabilityHandlers {
		if c.Interface == typ {
			return c
		}
	}

	return nil
}

func (ad *AbstractDevice) ID() string {
	return ad.id
}

func (ad *AbstractDevice) Cookie() Cookie {
	return ad.cookie
}

func (ad *AbstractDevice) ManufacturerName() string {
	return ad.manufacturerName
}

func (ad *AbstractDevice) FriendlyName() string {
	return ad.friendlyName
}

func (ad *AbstractDevice) Description() string {
	return ad.description
}

func (ad *AbstractDevice) DisplayCategories() []string {
	return ad.displayCategories
}

func (ad *AbstractDevice) AddDisplayCategory(cat string) {
	ad.displayCategories = append(ad.displayCategories, cat)
}

func (ad *AbstractDevice) Capabilities() []Capability {
	capabilities := make([]Capability, 0, len(ad.capabilityHandlers))

	for _, handler := range ad.capabilityHandlers {
		supported := make([]Supported, 0, len(handler.propertyHandlers))

		for name, _ := range handler.propertyHandlers {
			supported = append(supported, Supported{
				Name: name,
			})
		}

		var properties *Properties

		if len(supported) > 0 {
			properties = &Properties{
				ProactivelyReported: handler.ProactivelyReported,
				Retrievable:         handler.Retrievable,
				Supported:           supported,
			}
		}

		capabilities = append(capabilities, Capability{
			Type:                       handler.Type,
			Interface:                  handler.Interface,
			Version:                    handler.Version,
			Properties:                 properties,
			CameraStreamConfigurations: handler.CameraStreamConfigurations,
		})
	}

	return capabilities
}

type CapabilityHandler struct {
	Type      string
	Interface string
	Version   string

	ProactivelyReported bool
	Retrievable         bool

	CameraStreamConfigurations []CameraStreamConfiguration

	propertyHandlers map[string]PropertyHandler
}

func (ch *CapabilityHandler) AddPropertyHandler(name string, ph PropertyHandler) {
	if ch.propertyHandlers == nil {
		ch.propertyHandlers = map[string]PropertyHandler{}
	}

	ch.propertyHandlers[name] = ph
}

type PropertyHandler interface {
	GetValue() (interface{}, error)
	SetValue(val interface{}) error
	UpdateChannel() <-chan interface{}
}

type PropertyHandlerSetterFunc func(val interface{}) error

func (fn PropertyHandlerSetterFunc) GetValue() (interface{}, error) {
	return nil, nil
}
func (fn PropertyHandlerSetterFunc) SetValue(val interface{}) error {
	return fn(val)
}
func (_ PropertyHandlerSetterFunc) UpdateChannel() <-chan interface{} {
	return nil
}
