package controllers

import (
	"projeto_turismo_jp/repositories"
	"projeto_turismo_jp/websocket"
)


type TouristController struct {
	repo repositories.TouristRepository
	hub *websocket.Hub
}

func NewTouristController(repo repositories.TouristRepository, hub *websocket.Hub) *TouristController {
	return &TouristController{
		repo: repo,
		hub: hub,
	}
}

type TripController struct {
	repo repositories.TripRepository
	hub *websocket.Hub
}

func NewTripController(repo repositories.TripRepository, hub *websocket.Hub) *TripController {
	return &TripController{
		repo: repo,
		hub: hub,
	}
}

