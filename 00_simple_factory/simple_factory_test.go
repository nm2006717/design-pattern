package simplefactory

import "testing"

//TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := GetAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := GetAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Type2 test fail")
	}
}

func TestCircle(t *testing.T) {
	Shape := GetShaper("circle")
	Shape.Draw()
}

func TestRectangle(t *testing.T) {
	Shape := GetShaper("rectangle")
	Shape.Draw()
}

func TestSquare(t *testing.T) {
	Shape := GetShaper("square")
	Shape.Draw()
}
