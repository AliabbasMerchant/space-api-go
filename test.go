package main

import (
	"github.com/spaceuptech/space-api-go/api"
	"github.com/spaceuptech/space-api-go/api/model"
	"fmt"
)

// New initialised a new instance of the API object
func New(project, host, port string, sslEnabled bool) (*api.API, error) {
	return api.Init(project, host, port, sslEnabled)
}

func main() {
	api, err := New("grpc", "localhost", "8081", false)
	if(err != nil) {
		fmt.Println(err)
	}
	api.SetToken("my_secret")
	// db := api.MySQL()
	// resp, err := db.Get("books").Apply()
	// fmt.Println(resp.Status)
	// fmt.Println(err)
	// service := api.Service("service")
	// service.RegisterFunc("echo_func", Echo)
	// service.Start()
	db := api.MySQL()
	db.LiveQuery("books").Subscribe(func(liveData *model.LiveData, changeType string) () {
		fmt.Println(changeType)
		// v := make([]interface{})
		var v []interface{}
		liveData.Unmarshal(&v)
		fmt.Println(v)
	}, func(err error) () {
		fmt.Println(err)
	})
	for {}
}
// func Echo(params, auth *model.Message, fn service.CallBackFunction) {
// 	var i interface{}
// 	params.Unmarshal(&i)
// 	fn("response", i)
// }