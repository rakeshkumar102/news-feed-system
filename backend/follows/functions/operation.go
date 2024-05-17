package functions

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pranay999000/follows/configs"
	"github.com/pranay999000/follows/models"
)

func BasicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func CheckVertesExists(user_id string, channel chan bool) {
	url := "http://localhost:2480/command/UserGraph/sql"
	method := "POST"

	user_byte := []byte(user_id)
	var reqBody = []byte(`{"command": "select from Follows where user_id = :user_id", "parameters": {"user_id": "`)
	reqBody = append(reqBody, user_byte...)
	var end = []byte(`",}}`)
	reqBody = append(reqBody, end...)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Basic " + BasicAuth("root", "password"))

	if err != nil {
		log.Fatalln(err)
		channel <- false
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		channel <- false
	}
	defer res.Body.Close()
	
	var result models.User

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
		channel <- false
		return
	}

    if err := json.Unmarshal(body, &result); err != nil {
        log.Fatalln(err)
		channel <- false
		return
    }

	channel <- len(result.Result) != 0

}

func CheckEdgeExists(user_id string, following_user_id string, channel chan bool) {
	url := "http://localhost:2480/command/UserGraph/sql"
	method := "POST"

	following_user_byte := []byte(following_user_id)

	var reqBody = []byte(`{"command": "select expand( `)
	var directionIn = []byte(`in()`)
	reqBody = append(reqBody, directionIn...)
	var mid = []byte(` ) from Follows where user_id = :user_id", "parameters": {"user_id": "`)
	reqBody = append(reqBody, mid...)
	reqBody = append(reqBody, following_user_byte...)
	var end = []byte(`",}}`)
	reqBody = append(reqBody, end...)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Fatalln(err)
		channel <- false
		return
	}

	req.Header.Add("Authorization", "Basic " + BasicAuth("root", "password"))

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		channel <- false
		return
	}
	defer res.Body.Close()

	var result models.User

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
		channel <- false
		return
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalln(err)
		channel <- false
		return
	}

	for _, r := range result.Result {
		if r.UserId == user_id {
			channel <- true
			return
		}
	}
	channel <- false
}

func GetEdge(user_id string, following_user_id string, channel chan string) {
	url := "http://localhost:2480/command/UserGraph/sql"
	method := "POST"

	following_user_byte := []byte(following_user_id)

	var reqBody = []byte(`{"command": "select expand( `)
	var directionIn = []byte(`in()`)
	reqBody = append(reqBody, directionIn...)
	var mid = []byte(` ) from Follows where user_id = :user_id", "parameters": {"user_id": "`)
	reqBody = append(reqBody, mid...)
	reqBody = append(reqBody, following_user_byte...)
	var end = []byte(`",}}`)
	reqBody = append(reqBody, end...)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Fatalln(err)
		channel <- ""
		return
	}

	req.Header.Add("Authorization", "Basic " + BasicAuth("root", "password"))

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		channel <- ""
		return
	}
	defer res.Body.Close()

	var result models.User

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
		channel <- ""
		return
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalln(err)
		channel <- ""
		return
	}

	for _, r := range result.Result {
		if r.UserId == user_id {
			channel <- r.OutFollowing[0]
			return
		}
	}
	channel <- ""
}

func CreateVertex(user_id string) (*http.Response, error) {
		orientDBUrl := "http://localhost:2480/command/UserGraph/sql"
		method := "POST"

		channel := make(chan bool, 1)

		go CheckVertesExists(user_id, channel)

		checkVertex := <- channel

		if checkVertex {
			return nil, fmt.Errorf("vertex already exists") 
		}

		var reqBody = []byte(`{"command": "create vertex Follows set user_id = :user_id", "parameters": {"user_id": "`)
		user_byte := []byte(user_id)
		reqBody = append(reqBody, user_byte...)
		var end = []byte(`",}}`)
		reqBody = append(reqBody, end...)

		client := &http.Client{}
		req, err := http.NewRequest(method, orientDBUrl, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		if err != nil {
			log.Fatalln(err)
			return nil, fmt.Errorf("Internal Server Error")
		}

		req.Header.Add("Authorization", "Basic " + BasicAuth("root", "password"))

		res, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
			return nil, fmt.Errorf("Failed to get orientDB response")
		}

		defer res.Body.Close()

		return res, nil
}

func GetUserData(raw models.User) (*models.UserData, error) {
	var ids []string

	for _, r := range raw.Result {
		ids = append(ids, r.UserId)
	}

	userUrl, err := configs.EnvMap("users")

	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(ids)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", userUrl + "ids", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var result models.UserData

	body, err  := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}