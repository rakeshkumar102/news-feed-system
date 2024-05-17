package services

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "12"))
		defer cancel()

		opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit)).SetProjection(bson.M{"password": 0})
		filter := bson.D{}

		cursor, err := userCollection.Find(ctx, filter, opts)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		defer cursor.Close(ctx)
		
		var userList []models.User
		for cursor.Next(ctx) {
			var user models.User

			if err = cursor.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
				})
				return
			}

			userList = append(userList, user)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"users": userList,
		})

	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		user_id := c.Param("user_id")

		objectId, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid user id",
			})
			return
		}

		filter := bson.M{"_id": objectId}

		var result bson.M

		opts := options.FindOne().SetProjection(bson.M{"password": 0})

		cursor := userCollection.FindOne(ctx, filter, opts).Decode(&result)
		
		if cursor != nil {
			if cursor == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{
					"success": false,
					"message": "user not found",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"user": result,
			})
		}

	}
}

func GetUsersByIds() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()

		var requestIds []string
		c.Bind(&requestIds)
		var oids []primitive.ObjectID

		for _, val := range requestIds {
			id, err := primitive.ObjectIDFromHex(val)

			if err != nil {
				continue
			}

			oids = append(oids, id)
		}

		var userList []models.User
		opts := options.FindOne().SetProjection(bson.M{"password": 0})

		for _, ids := range oids {
			var user models.User
			userCollection.FindOne(ctx, bson.M{"_id": ids}, opts).Decode(&user)

			userList = append(userList, user)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"users": userList,
		})
	}
}