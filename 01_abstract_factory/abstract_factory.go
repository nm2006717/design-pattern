package abstractfactory

import "fmt"

// Shaper is a interface
type Shaper interface {
	Draw()
}

//Color is a interface
type color interface {
	Fill()
}

// Circle is one of Shaper implement
type Circle struct{}

//Draw a cicle
func (c *Circle) Draw() {
	fmt.Println("Circle Draw() method")
}

//Square is one of Shaper implement
type Square struct{}

//Draw a square
func (c *Square) Draw() {
	fmt.Println("Square Draw() method")
}

// Red is one of Color implemment
type Red struct{}

//Fill a red
func (red *Red) Fill() {
	fmt.Println("Red Fill() method")
}

// Green is one of Color implement
type Green struct{}

//Fill a green
func (green *Green) Fill() {
	fmt.Println("Green Fill() method")
}

type Abstractfactory interface {
	GetShape(shapeName string) Shaper
	GetColor(colorName string) Color
}

//ShapeFactory 模型工厂的类
type ShapeFactory struct{}

//ColorFactory 色彩工厂的类
type ColorFactory struct{}

//GetShape 模型工厂实例获取模型子类的方法
func (sh ShapeFactory) GetShape(shapeName string) Shape {
	switch shapeName {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

//GetColor 模型工厂实例不需要色彩方法
func (sh ShapeFactory) GetColor(colorName string) Color {
	return nil
}

//GetShape 色彩工厂实例不需要获取模型方法
func (cf ColorFactory) GetShape(shapeName string) Shape {
	return nil
}

//GetColor 色彩工厂实例，获取具体色彩子类
func (cf ColorFactory) GetColor(colorName string) Color {
	switch colorName {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

//FactoryProducer 超级工厂类，用于获取工厂实例
type FactoryProducer struct{}

//GetFactory 获取工厂方法
func (fp FactoryProducer) GetFactory(factoryname string) AbstractFactory {
	switch factoryname {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}

//NewFactoryProducer 创建FactoryProducer类
func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}
