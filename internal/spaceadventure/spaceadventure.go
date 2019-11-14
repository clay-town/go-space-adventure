package spaceadventure

import "fmt"

//Start is available to the outside world
func Start(planetarySystem PlanetarySystem) {
	printWelcome(planetarySystem)
	printGreeting(responseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	travel()
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Println("Welcome to the", planetarySystem.Name)
	fmt.Println("There are 8 planets to explore.")
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func promptForRandomOrSpecificDestination() bool {
	switch choice := responseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)"); choice {
	case "Y", "y":
		return true
	case "N", "n":
		return false
	default:
		fmt.Println("Sorry, I didn't get that.")
		promptForRandomOrSpecificDestination()
	}
	return false
}

func travel() {
	switch tf := promptForRandomOrSpecificDestination(); tf {
	case true:
		travelToRandomPlanet()
	case false:
		travelToPlanet(responseToPrompt("Name the planet you would like to visit"))
	}
}

func responseToPrompt(prompt string) string {
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
