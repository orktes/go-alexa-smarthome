package smarthome

const ACCEPT_GRANT_FAILED = "ACCEPT_GRANT_FAILED"

type AcceptGrantRequest struct {
	Grant   Grant `json:"grant"`
	Grantee Scope `json:"grantee"`
}

type AcceptGrantResponse struct{}

type Grant struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type Scope struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type authorization struct {
	sm *Smarthome
}

func (a *authorization) AcceptGrant(req AcceptGrantRequest) (AcceptGrantResponse, error) {
	if err := a.sm.auth.AcceptGrant(req); err != nil {
		return AcceptGrantResponse{}, newError(ACCEPT_GRANT_FAILED, err)
	}

	return AcceptGrantResponse{}, nil
}

type AuthorizationInterface interface {
	AcceptGrant(req AcceptGrantRequest) error
}

type AuthorizationFunc func(req AcceptGrantRequest) error

func (af AuthorizationFunc) AcceptGrant(req AcceptGrantRequest) error {
	return af(req)
}
