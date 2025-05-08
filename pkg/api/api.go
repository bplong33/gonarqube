// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    api, err := UnmarshalAPI(bytes)
//    bytes, err = api.Marshal()

package api

import "encoding/json"

func UnmarshalAPI(data []byte) (API, error) {
	var r API
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *API) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type API struct {
	WebServices []WebService `json:"webServices"`
}

type WebService struct {
	Path        string   `json:"path"`
	Since       *string  `json:"since,omitempty"`
	Description *string  `json:"description,omitempty"`
	Actions     []Action `json:"actions"`
}

type Action struct {
	Key                string      `json:"key"`
	Description        string      `json:"description"`
	Since              string      `json:"since"`
	DeprecatedSince    *string     `json:"deprecatedSince,omitempty"`
	Internal           bool        `json:"internal"`
	Post               bool        `json:"post"`
	HasResponseExample bool        `json:"hasResponseExample"`
	Changelog          []Changelog `json:"changelog"`
	Params             []Param     `json:"params,omitempty"`
}

type Changelog struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}

type Param struct {
	Key                string   `json:"key"`
	Description        *string  `json:"description,omitempty"`
	Required           bool     `json:"required"`
	Internal           bool     `json:"internal"`
	MaximumLength      *int64   `json:"maximumLength,omitempty"`
	Since              *string  `json:"since,omitempty"`
	DefaultValue       *string  `json:"defaultValue,omitempty"`
	MaximumValue       *int64   `json:"maximumValue,omitempty"`
	ExampleValue       *string  `json:"exampleValue,omitempty"`
	PossibleValues     []string `json:"possibleValues,omitempty"`
	MinimumLength      *int64   `json:"minimumLength,omitempty"`
	DeprecatedKey      *string  `json:"deprecatedKey,omitempty"`
	DeprecatedKeySince *string  `json:"deprecatedKeySince,omitempty"`
	DeprecatedSince    *string  `json:"deprecatedSince,omitempty"`
	MaxValuesAllowed   *int64   `json:"maxValuesAllowed,omitempty"`
}
