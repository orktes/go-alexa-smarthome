package smarthome

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	res, _ := sh.Handle(req)

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

	res, _ := sh.Handle(req)

	if res.Event.Header.Name != "Response" {
		t.Error("Failed")
	}

	if res.Context != nil && len(res.Context.(EndpointResponse).Properties) != 0 {
		t.Error("Nothing should be returned")
	}
}

func TestHandleErrors(t *testing.T) {
	cases := []struct {
		namespace     string
		name          string
		expectedError error
	}{
		{
			namespace: "Alexa.FailController",
			name:      "TurnOn",
			expectedError: fmt.Errorf(
				"Controller Alexa.FailController has not been implemented",
			),
		},
		{
			namespace: "Alexa.PowerController",
			name:      "Fail",
			expectedError: fmt.Errorf(
				"Method Fail has not been implemented for Alexa.PowerController",
			),
		},
	}

	for _, c := range cases {
		reqJSON := fmt.Sprintf(`
		{
			"directive": {
				"header": {
					"namespace": "%s",
					"name": "%s",
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
		`, c.namespace, c.name)

		sh := New(nil)

		req := &Request{}
		json.Unmarshal([]byte(reqJSON), req)

		_, err := sh.Handle(req)
		if !reflect.DeepEqual(err, c.expectedError) {
			t.Fatalf("Expected: \"%s\", got: \"%s\"", c.expectedError, err)
		}
	}
}
