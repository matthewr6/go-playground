package main

import (
    "fmt"
)

type Human struct {
    height int
    weight int
    age int
    metric bool
}

func (h Human) bmi() int {
    // so apparently you use pointers if you don't want change
    if h.metric {
        return (h.weight / (h.height * h.height))
    }
    return (703 * h.weight / (h.height * h.height))
}

func (h *Human) happybirthday() {
    // has to be pointer to change in main
    h.age++
}

func main() {
    somebody := &Human{1, 2, 3, false}
    fmt.Println(somebody)
    fmt.Printf("somebody's age is %v.\n", somebody.age)
    somebodyage := &somebody.age
    fmt.Println("Happy birthday, somebody.")
    *somebodyage++
    fmt.Printf("somebody's age is now %v.\n", somebody.age)
    fmt.Printf("somebody's BMI is %v.\n", somebody.bmi())
    fmt.Println("Happy birthday, somebody.")
    somebody.happybirthday()
    fmt.Printf("somebody's age is now %v.\n", somebody.age)
}