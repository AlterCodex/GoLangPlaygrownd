package main

import "fmt"

type Animal struct {
	movement string
	name     string
	voice    string
	eaten    string
}

func (a Animal) eat() {
	fmt.Printf("%s\n", a.eaten)
}

func (a Animal) move() {
	fmt.Printf("%s\n", a.movement)
}

func (a Animal) speak() {
	fmt.Printf("%s\n", a.voice)
}

func main() {
	snake := Animal{movement: "slither",
		name:  "snake",
		voice: "hss",
		eaten: "mice"}
	cow := Animal{movement: "walk",
		name:  "cow",
		voice: "moo",
		eaten: "grass"}
	bird := Animal{movement: "fly",
		name:  "bird",
		voice: "peep",
		eaten: "worms"}
	animals := make(map[string]Animal)
	animals[snake.name] = snake
	animals[cow.name] = cow
	animals[bird.name] = bird

	animal := ""
	action := ""
	for true {
		fmt.Printf("Insert the Animal and Action you want to use: \n")
		fmt.Printf("cntrl + c to Exit \n")
		n, _ := fmt.Scanf("%s%s%*c", &animal, &action)
		if n != 2 {
			print("insuficient Args")
			return
		}
		switch action {
		case "eat":
			animals[animal].eat()
		case "move":
			animals[animal].move()
		case "speak":
			animals[animal].speak()
		}
	}

}
