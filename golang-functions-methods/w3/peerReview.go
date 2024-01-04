package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	fmt.Println("Please enter name of animal and requested information about it")
	for {
		fmt.Println(">")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')

		line = line[:len(line)-1]

		words := strings.Split(line, " ")
		if len(words) < 2 {
			fmt.Println("Wrong input")
			continue
		}

		var animal = animals[words[0]]
		if words[1] == "eat" {
			fmt.Println(animal.Eat())
		} else if words[1] == "move" {
			fmt.Println(animal.Move())
		} else if words[1] == "speak" {
			fmt.Println(animal.Speak())
		} else {
			fmt.Println("No such action")
		}
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}
