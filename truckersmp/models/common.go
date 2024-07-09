package models

type UsualResponseWrapper struct {
	Error      bool    `json:"error"`
	Descriptor *string `json:"descriptor"`
	Response   any     `json:"response"`
}
