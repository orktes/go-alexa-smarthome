package smarthome

import (
	"encoding/json"
	"testing"
)

func TestHandleAuth(t *testing.T) {
	reqJSON := `
	{
		"directive": {
			"header": {
				"namespace": "Alexa.Authorization",
				"name": "AcceptGrant",
				"payloadVersion": "3",
				"messageId": "1bd5d003-31b9-476f-ad03-71d471922820",
				"correlationToken": "dFMb0z+PgpgdDmluhJ1LddFvSqZ/jCc8ptlAKulUj90jSqg=="
			},
			"payload": {
				"grant": {
					"type": "OAuth2.AuthorizationCode",
					"code": "ANUbUKCJqlBOpMhwYWxU"
				},
				"grantee": {
					"type": "BearerToken",
					"token": "access-token-from-skill"
				}
			}
		}
	}
	`

	authFN := AuthorizationFunc(func(req AcceptGrantRequest) error {
		if req.Grant.Code != "ANUbUKCJqlBOpMhwYWxU" {
			t.Error("Wrong code received")
		}

		if req.Grantee.Token != "access-token-from-skill" {
			t.Error("Fox bauh")
		}

		return nil
	})

	sh := New(authFN)

	req := &Request{}
	json.Unmarshal([]byte(reqJSON), req)

	res := sh.Handle(req)

	if res.Event.Header.Name != "AcceptGrant.Response" {
		t.Error("Failed")
	}
}

func TestHandlePowerControllerTurnOn(t *testing.T) {
	reqJSON := `
	{
		"directive": {
			"header": {
				"namespace": "Alexa.PowerController",
				"name": "TurnOn",
				"payloadVersion": "3",
				"messageId": "1bd5d003-31b9-476f-ad03-71d471922820",
				"correlationToken": "dFMb0z+PgpgdDmluhJ1LddFvSqZ/jCc8ptlAKulUj90jSqg=="
			},
			"endpoint": {
				"scope": {
					"type": "BearerToken",
					"token": "access-token-from-skill"
				},
				"endpointId": "endpoint-001",
				"cookie": {}
			},
			"payload": {}
		}
	}
	`
	sh := New(nil)

	req := &Request{}
	json.Unmarshal([]byte(reqJSON), req)

	res := sh.Handle(req)

	if res.Event.Header.Name != "Response" {
		t.Error("Failed")
	}

	if len(res.Context.(EndpointResponse).Properties) != 0 {
		t.Error("Nothing should be returned")
	}
}
