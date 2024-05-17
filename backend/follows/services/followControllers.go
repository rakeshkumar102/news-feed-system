package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/follows/functions"
	"github.com/pranay999000/follows/models"
)

func ConnectOrientDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := "http://localhost:2480/database/UserGraph"
		method := "GET"

		client := &http.Client {}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		
		req.Header.Add("Authorization", "Basic " + functions.BasicAuth("root", "password"))

		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		defer res.Body.Close()

		var respBody gin.H

		decoder := json.NewDecoder(res.Body)
		jsonErr := decoder.Decode(&respBody)

		if jsonErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": respBody,
		})
	}
}

func GetFollows() gin.HandlerFunc {
	return func(c *gin.Context) {
		orientDBUrl := "http://localhost:2480/command/UserGraph/sql"
		method := "POST"

		user_id := c.Param("user_id")
		direction := c.Param("direction")

		var b  = []byte(`{"command": "select expand( `)
		var directionIn = []byte(`in()`)
		var directionOut = []byte(`out()`)
		
		if direction == "following" {
			b = append(b, directionOut...)
		} else {
			b = append(b, directionIn...)
		}

		var mid = []byte(` ) from follows where user_id = :user_id","parameters": {"user_id": "`)

		b = append(b, mid...)
		user_byte := []byte(user_id)
		b = append(b, user_byte...)
		var end = []byte(`",}}`)
		b = append(b, end...)

		client := &http.Client{}
		req, err := http.NewRequest(method, orientDBUrl, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		req.Header.Add("Authorization", "Basic " + functions.BasicAuth("root", "password"))

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		defer res.Body.Close()

		var result models.User

		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}

		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalln(err)
		}

		users, err := functions.GetUserData(result)

		if err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users.Users,
		})
	}
}

func CreateUserVertex() gin.HandlerFunc{
	return func(c *gin.Context) {
		user_id := c.Param("user_id")


		res, err := functions.CreateVertex(user_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})	
		}

		var respBody gin.H

		decoder := json.NewDecoder(res.Body)
		jsonErr := decoder.Decode(&respBody)

		if jsonErr != nil {
			log.Fatalln(jsonErr)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": respBody,
		})
	}
}

func CreateFollowEdge() gin.HandlerFunc {
	return func(c *gin.Context) {
		orientDBUrl := "http://localhost:2480/command/UserGraph/sql"
		method := "POST"

		user_id := c.Param("user_id")
		following_user_id := c.Param("following_user_id")

		channel := make(chan bool, 2)
		
		go functions.CheckVertesExists(user_id, channel)
		go functions.CheckVertesExists(following_user_id, channel)
		
		checkUser := <- channel
		checkFollowingUser := <- channel
		
		if !checkUser || !checkFollowingUser {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Unable to find user",
			})
			return
		}

		channelCheckEdge := make(chan bool, 1)

		go functions.CheckEdgeExists(user_id, following_user_id, channelCheckEdge)

		checkEdge := <- channelCheckEdge

		if checkEdge {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Edge already exists",
			})
			return
		}

		user_byte := []byte(user_id)
		following_byte := []byte(following_user_id)
		
		var reqBody = []byte(`{"command": "create edge Following from ( select from Follows where user_id = :user_id ) to ( select from Follows where user_id = :following_user_id )", "parameters": {"user_id": "`)
		reqBody = append(reqBody, user_byte...)
		var mid = []byte(`", "following_user_id": "`)
		reqBody = append(reqBody, mid...)
		reqBody = append(reqBody, following_byte...)
		var end = []byte(`",}}`)
		reqBody = append(reqBody, end...)
		
		client := &http.Client{}
		req, err := http.NewRequest(method, orientDBUrl, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		req.Header.Add("Authorization", "Basic " + functions.BasicAuth("root", "password"))

		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		defer res.Body.Close()

		var result models.User

		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}

		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalln(err.Error())
		}
		
		c.JSON(http.StatusOK, gin.H{
			"response": result,
		})
		

	}
}

func DeleteEdge() gin.HandlerFunc {
	return func(c *gin.Context) {
		orientDBUrl := "http://localhost:2480/command/UserGraph/sql"
		method := "POST"

		user_id := c.Param("user_id")
		following_user_id := c.Param("following_user_id")

		channel := make(chan bool, 2)

		go functions.CheckVertesExists(user_id, channel)
		go functions.CheckVertesExists(following_user_id, channel)

		checkUser := <- channel
		checkFollowingUser := <- channel

		if !checkUser || !checkFollowingUser {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Unable to find user",
			})
			return
		}

		channelCheckEdge := make(chan string, 1)

		go functions.GetEdge(user_id, following_user_id, channelCheckEdge)

		checkEdge := <- channelCheckEdge

		if checkEdge == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Edge does not exists",
			})
			return
		}

		var reqBody = []byte(`{"command": "delete edge `)
		edge_byte := []byte(checkEdge)
		reqBody = append(reqBody, edge_byte...)
		var end = []byte(`"}`)
		reqBody = append(reqBody, end...)

		client := &http.Client{}
		req, err := http.NewRequest(method, orientDBUrl, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		req.Header.Add("Authorization", "Basic " + functions.BasicAuth("root", "password"))

		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		defer res.Body.Close()

		var result models.User

		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err.Error())
		}

		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalln(err.Error())
		}
		
		
		c.JSON(http.StatusOK, gin.H{
			"response": result,
		})
	}
}