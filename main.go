package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-requests/models"
	"golang-requests/utils"
)

func main() {
	content := models.Person{
		FirstName: "Benzion",
		LastName:  "Yehezkel",
		Online:    false,
	}
	jsonData, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	reqData := bytes.NewBuffer(jsonData)
	result, err := utils.CreatePostRequest("https://www.postman-echo.com/post", "application/json", reqData)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
