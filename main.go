package main

import (
	"context"
	"fmt"
	"log"

	"example.com/go-practice/controllers"
	"example.com/go-practice/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


var (
	server 			*gin.Engine
	userservice 	services.UserService
	usercontroller 	controllers.UserController
	ctx 			context.Context
	userCollection 	*mongo.Collection
	mongoClient 	*mongo.Client
	err 			error
)



func init(){
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://anji2006:Anji@2006@cluster0.gufvgcd.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	mongoClient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo connection is Established")

	userCollection = mongoClient.Database("userdb").Collection("user")
	userservice = services.NewUserService(userCollection, ctx)
	usercontroller = controllers.NewUserController(userservice)
	server = gin.Default()
}


func main(){
	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9089"))
}