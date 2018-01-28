package smarthome

import (
	"encoding/json"
	"time"
)

type zuluTime struct {
	time.Time
}

func (zulu *zuluTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(zulu.Format(time.RFC3339))
}
