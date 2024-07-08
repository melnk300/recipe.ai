package database

import (
	"context"
	"fmt"
	"github.com/melnk300/recipe.ai/server/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeConnection() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?ssl=false", config.Config.Database.Username, config.Config.Database.Password,
		config.Config.Database.Host, config.Config.Database.Port)
	con, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("Database is unavailable")
	}

	return con
}

func UseCollection(coll string) *mongo.Collection {
	con := MakeConnection().Database(config.Config.Database.DatabaseName)

	return con.Collection(coll)
}
