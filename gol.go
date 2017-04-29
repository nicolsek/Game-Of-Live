package main

/* Rules! */

/*

	1. Any live cell with fewer than two live neighbors dies, as if caused by underpopulation.
	2. Any live cell with two or three live neighbors lives on to the next generation.
	3. Any live cell with more than three live neighbors dies, as if by overpopulation.
	4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

	(Rules taken from Wikipedia)

*/

// main ... Main functions that begins the simulation of life and acts as a state based loop.
func main() {
	//Creating and passing a reference to the GOL object.
	GOL := makeGOL()

	//While GOL is simulating.
	for GOL.isSimulating {
		//Simulating using the reference to the GOL object.
		simulate(GOL)
		//Draw the GOL cells.
		draw(GOL)
	}

}

// GOL ... A struct that holds the required data for simulating and managing the Game of Life.
type GOL struct {
	cells        [256][256]int
	isSimulating bool
}

// makeGOL ... Creates the GOL object and sets the running to true.
func makeGOL() *GOL {
	GOL := new(GOL)
	GOL.isSimulating = true

	return GOL
}

// simulate ... Begins simulating the GOL and using the object alters the values and checks the rules.
func simulate(gol *GOL) {

}

func draw(gol *GOL) {

}
