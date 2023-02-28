package database

import(
	"log"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client{
client,err:=mongo.NewClient(options.Client().ApplyURI("mongo://localhost:27017"))

if err!=nil{
	log.Fatal(err)
}

 ctx,cancel := context.withTimeOut(context.Backgound(),10*time.Second)

defer cancel()

err=client.Connect(ctx)
if err!=nil{
	log.Fatal(err)
}

err=client.Ping(context.TODO(),nil)
if err!=nil{
	log.Fatal("Failed to connect to mongodb")
	return nil
}

fmt.Println("Sucessfully connected tomongoDB")
client
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection{}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection{}
