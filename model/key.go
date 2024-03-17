package model

var KeyConfig *Key

type Key struct {
	Secret string `json:"secret"`
}
