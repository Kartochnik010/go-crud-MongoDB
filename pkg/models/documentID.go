package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentID struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Surname      string             `json:"surname"`
	Name         string             `json:"name"`
	MiddleName   string             `json:"middleName"`
	DateOfBirth  string             `json:"dateOfBirth"`
	Iin          string             `json:"iin"`
	SerialNumber string             `json:"serialNumber"`
	HomeRegion   string             `json:"homeRegion"`
	Nationality  string             `json:"nationality"`
	LinkToImage  string             `json:"linkToImage"`
}
