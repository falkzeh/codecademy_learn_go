package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	leftSafely := rand.Intn(5)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn && openedVault < 70 {
		isHeistOn = false
		fmt.Println("Vault cannot be opened!")
	}

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Heist failed...")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside..")
		case 2:
			isHeistOn = false
			fmt.Println("There is a traitor...")
		case 3:
			isHeistOn = false
			fmt.Println("Alarm went off...")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("We stole ", amtStolen)
	}

}
