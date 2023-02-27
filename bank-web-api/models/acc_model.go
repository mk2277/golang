package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accounts struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	AccountNo int                `json:"accountno" `
	PhoneNo   int                `json:"phoneno" `
	Balance   float64            `json:"balance" `
	ChangeBal float64            `json:"changebal"`
}
