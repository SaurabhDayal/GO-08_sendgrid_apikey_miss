package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email       string             `json:"email" json:"email"`
	FirstName   string             `json:"firstName" bson:"firstName"`
	LastName    string             `json:"lastName" bson:"lastName"`
	Number      string             `json:"number" bson:"number"`
	CountryCode string             `json:"countryCode" bson:"countryCode"`
	Occupation  string             `json:"occupation" bson:"occupation"`
}
