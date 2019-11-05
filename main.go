package main

import "fmt"

func main() {
	printWelcome()
	name := getName("What is your name?")
	printGreeting(name)
	fmt.Println("Let's go on an adventure!")
	travel()
}

func printWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getName("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			travelToRandomPlanet()
		} else if choice == "N" {
			planetName := getName("Name the planet you would like to visit.")
			travelToPlanet(planetName)
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

func responseToPrompt(prompt) string {
	var name string
	fmt.Println(prompt)
	fmt.Scan(&name)
	return name
}

func responseToPrompt(prompt) string{
	var name string
	fmt.Println(prompt)
    fmt.Scan(&name)
    return name
}

func responseToPrompt(prompt) string {
	var name string
	fmt.Println(prompt)
	fmt.Scan(&name)
	return name
}

func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}


func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
