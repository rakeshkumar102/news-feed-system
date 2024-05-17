package services

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/users/configs"
	"github.com/pranay999000/users/functions"
	"github.com/pranay999000/users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		c.Bind(&newUser)

		var result bson.M

		filter := bson.M{"email": newUser.Email}
		checkEmail := userCollection.FindOne(context.TODO(), filter).Decode(&result)

		if checkEmail != nil {
			if checkEmail == mongo.ErrNoDocuments {
				userResponse, err := userCollection.InsertOne(context.TODO(), newUser)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"success": false,
						"message": "Unable to create user",
					})
					return
				}

				newUser.ID = userResponse.InsertedID.(primitive.ObjectID)
				newUser.Password = ""
				functions.CreateUserNode(newUser.ID.Hex())

				c.JSON(http.StatusCreated, gin.H{
					"success": true,
					"user": newUser,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Email already exists",
			})
		}
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authUser models.User
		c.Bind(&authUser)

		var result bson.M

		filter := bson.M{"email": authUser.Email}
		checkEmail := userCollection.FindOne(context.TODO(), filter).Decode(&result)

		if checkEmail != nil {
			if checkEmail == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"message": "Email not found",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
				})
			}
		} else {
			if result["password"] == authUser.Password {
				token, err := configs.GenerateJWT(authUser.Email, authUser.Name, authUser.ID.String())

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"success": false,
						"message": "Unable to generate token",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"success": true,
					"token": token,
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"message": "Password does not match",
				})
			}
		}
	}
}
