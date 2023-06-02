package handlers

import (
	"go-microservices/services/auth-service/services"
)

type handler struct {
	service services.Service
}

func NewAuthHandler(service services.Service) *handler {
	return &handler{service: service}
}
