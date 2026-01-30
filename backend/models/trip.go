package models

import(
	"time"
)

type Trip struct {
	ID 								int64	 		`json:"id"`
	LodgingLocation 	string 		`json:"lodging_location" binding:"required,min=1,max=200"`
	TripDescription 	string 		`json:"trip_description binding:required"`
	ArrivalDate 			string		`json:"arrival_date" binding:"required,min=1,max=200"`
	DepartureDate 		string 		`json:"departure_date" binding:"required, min=1,max=200"`
	TripReview				string		`json:"trip_review" binding:"required,min=1,max=200"`
  Status          	string    `json:"status" binding:"required,oneof=upcoming ongoing completed"`
  TouristID       	int       `json:"tourist_id" binding:"required,gt=0"`
	CreatedAt 				time.Time `binding:"required"`
	UpdatedAt 				time.Time `binding:"required"`
}

