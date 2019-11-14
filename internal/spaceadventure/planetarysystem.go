package spaceadventure

//PlanetarySystem struct is a struct of Planet structs
type PlanetarySystem struct {
	Name    string
	Planets []Planet
}

//NumberOfPlanets returns the length of the Planets struct
func (ps PlanetarySystem) NumberOfPlanets() int {
	return len(ps.Planets)
}
