package apirole

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const RoleAnonymous = "anonymous"

// Roles a model
type Roles struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Display     string             `json:"display"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

// RoleUser a model
type RoleUser struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID int                `json:"userId" bson:"userId"`
	RoleID string             `json:"roleId" bson:"roleId"`
}

// Policy a model
type Policy struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Policy string             `json:"-" bson:"ptype"`
	RoleId string             `json:"roleId" bson:"v0"`
	Path   string             `json:"path" bson:"v1"`
	Method string             `json:"method" bson:"v2"`
	V3     string             `json:"v3" bson:"v3"`
	V4     string             `json:"v4" bson:"v4"`
	V5     string             `json:"v5" bson:"v5"`
}
