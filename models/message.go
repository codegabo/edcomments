package models

// Message structura para responder con mensajes de sistema al cliente del API
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
