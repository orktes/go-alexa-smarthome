package smarthome

type sherror struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	err     error  `json:"error"`
}

func (err *sherror) Error() string {
	return err.Message
}

func newError(typ string, mes error) *sherror {
	return &sherror{Type: typ, err: mes, Message: mes.Error()}
}
