package main

import "fmt"

type Car struct {
    Brand string
    Speed int
}

func (c Car) Info() string {
    return fmt.Sprintf("%s going %d km/h", c.Brand, c.Speed)
}

func main() {
    car := Car{"Tesla", 150}
    fmt.Println(car.Info())
}
