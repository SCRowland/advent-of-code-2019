package challenge13

import "intcode"

// Arcade is an arcade
type Arcade struct {
	program *intcode.Program
}

func (a *Arcade) Run() {
	for {
		select {
		case <-a.program.Final:
			return
		default:
		}
		//		across, down, thing := a.GetOutput()
	}
}

func (a *Arcade) GetOutput() (across, down, thing int64) {

	across = <-a.program.Output
	down = <-a.program.Output
	thing = <-a.program.Output

	return across, down, thing
}

// NewArcade creates a new instance of an arcade
func NewArcade() *Arcade {
	arcade := &Arcade{
		program: intcode.NewIntCodeProgram(puzzleInput),
	}

	go arcade.program.Execute()

	return arcade
}
