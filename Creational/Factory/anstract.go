package main
import (
  "fmt"
)

const (
  typeCar = iota
  typeBicycle 
)

type Veichle interface {
  getName() string
  getTopSpeed() int
  getNumberOfWheels() int
}

type car struct{
  Brand string
}

func (cr *car) getName() string {
  return cr.Brand
}

func (cr *car) getTopSpeed() int {
  return 120
}

func (cr *car) getNumberOfWheels() int {
  return 4
}

type bicycle struct{
  Brand string
}

func (by *bicycle) getName() string {
  return by.Brand
}

func (by *bicycle) getTopSpeed() int {
  return 20
}

func (by *bicycle) getNumberOfWheels() int {
  return 2
}


func New(VeichleType int) Veichle {

  if VeichleType == typeCar {
    return &car{Brand:"Honda"}
  }

  return &bicycle{Brand:"cyclone"}
}

func main() {
  newCar := New(typeCar)
  newCycle := New(typeBicycle)

  fmt.Println(newCar.getName())
  fmt.Println(newCycle.getName())
}