package main

import (
	"github.com/spaceuptech/space-api-go/api"
	"github.com/spaceuptech/space-api-go/api/utils"
	"fmt"
)

func main() {
	api, err := api.Init("books-app", "localhost", "8081", false)
	if(err != nil) {
		fmt.Println(err)
	}
	db := api.Mongo()
	condition := utils.Cond("id", "==", 1)
	inc := map[string]interface{}{"likes": 1}
	resp, err := db.Update("books").Where(condition).Inc(inc).Apply()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		if resp.Status == 200 {
			fmt.Println("Success")
		} else {
			fmt.Println("Error Processing Request:", resp.Error)
		}
	}
}