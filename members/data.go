package members

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DAO struct {
	Collection *mongo.Collection
}

func New(ctx context.Context) *DAO {
	host := os.Getenv("APP_DB_HOST")
	port := os.Getenv("APP_DB_PORT")
	username := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	db := os.Getenv("APP_DB_NAME")

	newDAO := DAO{
		Collection: connectDB(ctx, host, port, username, password, db, "members"),
	}

	return &newDAO
}

func connectDB(ctx context.Context, host string, port string, username string, password string, db string, collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	clientOptions := options.Client().ApplyURI(uri)
	internalContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	client, err := mongo.Connect(internalContext, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB!  =D")

	return client.Database(db).Collection(collection)
}

func (dao *DAO) Add(member *Member, ctx context.Context) error {
	internalContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := dao.Collection.InsertOne(internalContext, member)
	if err != nil {
		return err
	}

	member.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (dao *DAO) GetByID(id primitive.ObjectID, ctx context.Context) (*Member, error) {
	internalContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result := Member{}
	err := dao.Collection.FindOne(internalContext, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
