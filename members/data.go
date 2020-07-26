package members

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type DAO struct {
	collection *mongo.Collection
}

func New() *DAO {
	host := os.Getenv("APP_DB_HOST")
	port := os.Getenv("APP_DB_PORT")
	username := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	db := os.Getenv("APP_DB_NAME")

	newDAO := DAO{
		collection: connectDB(host, port, username, password, db, "members"),
	}

	return &newDAO
}

func connectDB(host string, port string, username string, password string, db string, collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to MongoDB!  =D")

	return client.Database(db).Collection(collection)
}

func (dao *DAO) Add(member *Member, ctx context.Context) error {
	result, err := dao.collection.InsertOne(ctx, member)
	if err != nil {
		return err
	}
	member.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
