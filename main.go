package main

import (
	"fmt"
	"math/rand"
)

const even = 0
const odd = 1

type PlayerInfo struct {
	id  int
	num int
}

func Player(id int, channel chan<- PlayerInfo) {
	//send che manda id e numero
	num := rand.Intn(6)
	fmt.Printf("[Player %d] my number is %d \n", id, num)

	channel <- PlayerInfo{id, num}
}

func Match(result chan<- int, channel chan PlayerInfo) {

	//mess che contiene id e numero
	pl1 := <-channel
	pl2 := <-channel

	sum := pl1.num + pl2.num

	choice_pl1 := rand.Intn(2)

	if choice_pl1 == even {
		fmt.Printf("[player %d] is even\n", pl1.id)
		fmt.Printf("[player %d] is odd\n", pl2.id)
	} else {
		fmt.Printf("[player %d] is odd\n", pl1.id)
		fmt.Printf("[player %d] is even\n", pl2.id)
	}

	//diciamo al main chi ha vinto
	if (sum%2 == 0 && choice_pl1 == even) || (sum%2 != 0 && choice_pl1 == odd) {
		fmt.Printf("Player %d and Player %d: FIGHT!\nWinner: %d\n", pl1.id, pl2.id, pl1.id)
		result <- pl1.id
	} else {
		fmt.Printf("Player %d and Player %d: FIGHT!\nWinner: %d\n", pl1.id, pl2.id, pl2.id)
		result <- pl2.id
	}
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func round(m int, players []int) []int {
	matchRoom := make([]chan PlayerInfo, m/2)
	results := make(chan int)

	for i := 0; i < m; i += 2 {
		matchRoom[i/2] = make(chan PlayerInfo)

		go Player(players[i], matchRoom[i/2])
		go Player(players[i+1], matchRoom[i/2])

		go Match(results, matchRoom[i/2])
	}

	var nextRoundPlayers []int

	for i := 0; i < m/2; i++ { // m/2 è il numero di matches in un round
		winner := <-results
		nextRoundPlayers = append(nextRoundPlayers, winner)
	}

	return nextRoundPlayers
}

func setup() (int, []int) {
	var m int

	fmt.Print("Insert number of players: ")
	fmt.Scanln(&m)

	for !isPowerOfTwo(m) {
		fmt.Print("Not valid! Insert number of players (pow of 2): ")
		fmt.Scanln(&m)
	}

	var players []int

	for i := 0; i < m; i++ {
		players = append(players, i)
	}

	return m, players
}

func main() {

	m, players := setup()

	for i := 1; m >= 2; i++ {
		fmt.Printf("\nSTARTING ROUND %d (remaining players: %d)\n", i, len(players))
		players = round(m, players)

		m = m / 2
	}

	fmt.Printf("the winner of the tournament is: %d", players[0])

}
