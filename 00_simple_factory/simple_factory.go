package simplefactory

import "fmt"

//API is interface
type API interface {
	Say(name string) string
}

//hiAPI is one of API implement
type hiAPI struct{}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API implement
type helloAPI struct{}

//Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

//NewAPI return Api instance by type
// func NewAPI(t int) API {
// 	if t == 1 {
// 		return &hiAPI{}
// 	} else if t == 2 {
// 		return &helloAPI{}
// 	}
// 	return nil
// }

// apiMapper is a mapper about to key to change API
var apiMapper = map[int]API{
	1:&hiAPI{},
	2:&helloAPI{},
}

// GetAPI return a instance of API
func GetAPI(key int) API{
	api,exist := apiMapper[key]
	if !exist{

	}
	return api
}
