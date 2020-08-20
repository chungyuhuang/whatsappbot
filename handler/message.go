package handler

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
)

type WhatsAppMessageHandler struct{}

func (WhatsAppMessageHandler) HandleError(err error) {
	fmt.Printf("%+v\n", err)
}

func (WhatsAppMessageHandler) HandleTextMessage(message whatsapp.TextMessage) {
	fmt.Printf("%+v\n", message)
}
