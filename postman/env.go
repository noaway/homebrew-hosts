package postman

import (
	"encoding/json"
)

// PostmanEnv struct PostmanEnv
type PostmanEnv struct {
	Name   string         `json:"name"`
	Values []PostmanValue `json:"values"`
}

func (pe *PostmanEnv) JsonBytes() ([]byte, error) {
	return json.Marshal(pe)
}

func (pe *PostmanEnv) JsonString() (string, error) {
	et, err := pe.JsonBytes()
	if err != nil {
		return "", err
	}
	return string(et), nil
}

type PostmanValue struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
}
