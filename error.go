package smarthome

type sherror struct {
	Type string `json:"type"`
	Err  error  `json:"message"`
}

func (err *sherror) Error() string {
	return err.Err.Error()
}

func newError(typ string, mes error) *sherror {
	return &sherror{Type: typ, Err: mes}
}
