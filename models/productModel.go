package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Product struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name"`
    Price       float64            `bson:"price"`
    Description string             `bson:"description"`
    InStock     bool               `bson:"inStock"`
    CreatedAt   time.Time          `bson:"createdAt"`
    UpdatedAt   time.Time          `bson:"updatedAt"`
}
