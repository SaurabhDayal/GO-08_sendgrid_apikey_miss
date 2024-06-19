package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email       string             `json:"email"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	Number      string             `json:"number"`
	CountryCode string             `json:"countryCode"`
	Occupation  string             `json:"occupation" bson:"occupation"`
}
