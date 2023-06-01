package handlers

import (
	"go-microservices/auth-service/services"
)

type handler struct {
	service services.Service
}

func NewAuthHandler(service services.Service) *handler {
	return &handler{service: service}
}
