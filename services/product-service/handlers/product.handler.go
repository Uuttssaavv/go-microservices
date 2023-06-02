package handlers

import "go-microservices/services/product-service/services"

type handler struct {
	services services.Service
}

func NewProductHandler(service services.Service) *handler {
	return &handler{services: service}
}
