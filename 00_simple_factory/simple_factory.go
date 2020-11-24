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

// apiMapper is a mapper using key to change API
var apiMapper = map[int]API{
	1:&hiAPI{},
	2:&helloAPI{},
}

// GetAPI return a instance of API
func GetAPI(key int) API{
	api,exist := apiMapper[key]
	if !exist{
		return nil
	}
	return api
}

//Shaper is interface
type Shaper interface{
	Draw()
}

//Rectangle is one of Shaper implement
type Rectangle struct{}

// Draw a rectangle
func (rectangle *Rectangle) Draw(){
	fmt.Println("Inside Rectangle::draw() method.")
}

//Square is one of Shaper implement
type Square struct{}

//Draw a square
func (square *Square) Draw(){
	fmt.Println("Inside Square ::draw() method.")
}

//Circle is one of Shaper implement
type Circle struct{
}

//Draw a circle
func (circle *Circle) Draw() {
	fmt.Println("Inside Circle  ::draw() method.")
} 

//shapeMapper is a mapper using key to change Shaper
var shapeMapper = map[string]Shaper{
	"rectangle":&Rectangle{},
	"circle":&Circle{},
	"square":&Square{},
}

//GetShaper return a instance of Shaper
func GetShaper( key string) Shaper{
	shape,exist := shapeMapper[key]
	if !exist{
		return nil
	}
	return shape
	
}
