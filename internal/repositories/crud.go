package repositories

import (
	"context"
	"fmt"
	"time"

	"5g-v2x-data-management-service/internal/config"
	"5g-v2x-data-management-service/internal/infrastructures/database"
)

// CRUDRepository ...
type CRUDRepository struct {
	MONGO  *database.MongoDatabase
	config *config.Config
}

// NewCRUDRepository ...
func NewCRUDRepository(m *database.MongoDatabase, c *config.Config) *CRUDRepository {
	return &CRUDRepository{
		MONGO:  m,
		config: c,
	}
}

// Create ...
func (crud *CRUDRepository) Create(collectionName string, m interface{}) (string, error) {
	fmt.Println("AAAAA")
	collection := crud.MONGO.Client.Database(crud.config.DatabaseName).Collection(collectionName)
	fmt.Println("BBBBB")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println("CCCCC")
	defer cancel()
	result, err := collection.InsertOne(ctx, m)
	fmt.Println("DDDDD")

	if err != nil {
		return "", err
	}
	fmt.Println("EEEEE")

	id := fmt.Sprintf("%s", result.InsertedID)
	fmt.Println("FFFFF")
	fmt.Println(id)
	return id, nil
}
