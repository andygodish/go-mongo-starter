package models

type Organization struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string 
}      