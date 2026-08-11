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
		result <- pl1.id
	} else {
		result <- pl2.id
	}
}

func main() {
	m := 4

	channels := make([]chan PlayerInfo, 2)
	result := make(chan int)

	for i := 0; i < m; i += 2 {
		channels[i] = make(chan PlayerInfo)

		go Player(i, channels[i/2])
		go Player(i+1, channels[i/2])

		go Match(result, channels[i/2])
	}
}
