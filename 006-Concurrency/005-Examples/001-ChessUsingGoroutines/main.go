package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This sample program demonstrates how to use an unbuffered
// channel to simulate a move of chess between two goroutines.

var quit chan bool

/*
The first line within the for-loop of the goroutine player blocks until an int is taken
out of the move channel. With a 5 percent probability, the player loses the game and sets quit to true.

After outputting that the player has moved and posting the player’s turn,
it puts an int back into the move channel, freeing the other player to move.
*/
func main() {
	rand.Seed(time.Now().UnixNano())
	move := make(chan int)
	quit = make(chan bool)
	// Launch two players.
	go player("Bobby Fischer", move)
	go player("Boris Spassky", move)
	// Start the move
	move <- 1
	<-quit // Blocks until quit assigned a value
}

// player simulates a person moving in chess.
func player(name string, move chan int) {
	// This function takes data out of the move channel
	// and puts data back into the move channel
	for {
		// Wait for turn to play
		turn := <-move // blocks until move assigned a value (every second)
		// Pick a random number and see if we lose the move
		n := rand.Intn(100)
		if n <= 5 && turn >= 5 {
			fmt.Printf("Player %s was check mated and loses!", name)
			quit <- true
			return
		}
		// Display and then increment the total move count by one.
		fmt.Printf("Player %s has moved. Turn %d.\n", name, turn)
		turn++
		// Yield the turn back to the opposing player
		time.Sleep(1 * time.Second)
		move <- turn
	}
}

/*
Player Boris Spassky has moved. Turn 1.
Player Bobby Fischer has moved. Turn 2.
Player Boris Spassky has moved. Turn 3.
Player Bobby Fischer has moved. Turn 4.
Player Boris Spassky has moved. Turn 5.
Player Bobby Fischer has moved. Turn 6.
Player Boris Spassky has moved. Turn 7.
Player Bobby Fischer has moved. Turn 8.
Player Boris Spassky has moved. Turn 9.
Player Bobby Fischer has moved. Turn 10.
Player Boris Spassky has moved. Turn 11.
Player Bobby Fischer has moved. Turn 12.
Player Boris Spassky has moved. Turn 13.
Player Bobby Fischer has moved. Turn 14.
Player Boris Spassky has moved. Turn 15.
Player Bobby Fischer has moved. Turn 16.
Player Boris Spassky has moved. Turn 17.
Player Bobby Fischer has moved. Turn 18.
Player Boris Spassky has moved. Turn 19.
Player Bobby Fischer has moved. Turn 20.
Player Boris Spassky has moved. Turn 21.
Player Bobby Fischer has moved. Turn 22.
Player Boris Spassky has moved. Turn 23.
Player Bobby Fischer has moved. Turn 24.
Player Boris Spassky has moved. Turn 25.
Player Bobby Fischer was check mated and loses!
*/
