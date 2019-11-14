package main

import "github.com/clay-town/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Solar System", Planets: []spaceadventure.Planet{
				spaceadventure.Planet{Name: "Arrakis", Description: "A desert planet, it's only resource is Melange."},
			},
		},
	)
}
