package controllers

import (
	"net/http"

	"example.com/go-practice/models"
	"example.com/go-practice/services"
	"github.com/gin-gonic/gin"
)




type UserController struct{
	UserService services.UserService
}



func NewUserController(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}


func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User 
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} 

	err := uc.UserService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Succussfull added User"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	userName := ctx.Param("name")
	userDetails, err := uc.UserService.GetUser(&userName)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, userDetails)
}

func (uc *UserController) GetAll(ctx *gin.Context)  {
	userList, err := uc.UserService.GetAll()

	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, userList)
}

func (uc *UserController) UpdateUser(ctx *gin.Context){
	var user models.User
	if err := ctx.ShouldBindJSON(&user) ; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message" : "Successfully Updated"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")	
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Succesfully Deleted User"})
}



func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroutes := rg.Group("/user")
	userroutes.POST("/create", uc.CreateUser)
	userroutes.GET("/:name", uc.GetUser)
	userroutes.GET("/all", uc.GetAll)
	userroutes.PUT("", uc.UpdateUser)
	userroutes.DELETE("/:name", uc.DeleteUser)
}