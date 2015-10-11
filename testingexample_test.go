package main

import "testing"

// Game contains the state of a bowling game.
type Game struct {
	rolls   []int
	current int
}

func TestSomething(t *testing.T) {
	// test stuff here...
	// t.Fail()

	t.Error("I'm in a bad mood.")
}


// NewGame allocates and starts a new game of bowling.
func NewGame() *Game {
	game := new(Game)
	game.rolls = make([]int, maxThrowsPerGame)
	return game
}


// Roll rolls the ball and knocks down the number of pins specified by pins.
func (self *Game) Roll(pins int) {
	self.rolls[self.current] = pins
	self.current++
}


// Score calculates and returns the player's current score.
func (self *Game) Score() (sum int) {
	for throw, frame := 0, 0; frame < framesPerGame; frame++ {
		if self.isStrike(throw) {
			sum += self.strikeBonusFor(throw)
			throw += 1
		} else if self.isSpare(throw) {
			sum += self.spareBonusFor(throw)
			throw += 2
		} else {
			sum += self.framePointsAt(throw)
			throw += 2
		}
	}
	return sum
}


// Score calculates and returns the player's current score.
func (self *Game) Score() (sum int) {
	for throw, frame := 0, 0; frame < framesPerGame; frame++ {
		if self.isStrike(throw) {
			sum += self.strikeBonusFor(throw)
			throw += 1
		} else if self.isSpare(throw) {
			sum += self.spareBonusFor(throw)
			throw += 2
		} else {
			sum += self.framePointsAt(throw)
			throw += 2
		}
	}
	return sum
}
