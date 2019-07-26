package main

import (
	"fmt"

	"./ramukaka"
)

func main() {
	ramukaka.Teach("What is your name")
	ramukaka.Teach("Where do you stay")
	ramukaka.Teach("which country")
	ramukaka.Teach("I love my country")
	fmt.Printf("Teaching Done\n")

	//ramukaka.Show()
	oops, _ := ramukaka.Ask("What is your name")
	if !oops {
		fmt.Printf("1 found\n")
	} else {
		fmt.Printf("1 not found\n")

	}

	oops, _ = ramukaka.Ask("which country")
	if !oops {
		fmt.Printf("3 found\n")
	} else {
		fmt.Printf("3 not found\n")

	}
	ramukaka.ChangeMyMind("Where do you stay", "Where do I stay")
	oops, _ = ramukaka.Ask("Where do you stay")
	if !oops {
		fmt.Printf("2 found\n")
	} else {
		fmt.Printf("2 not found\n")

	}

	oops, _ = ramukaka.Ask("Where do I stay")
	if !oops {
		fmt.Printf("4 found\n")
	} else {
		fmt.Printf("4 not found\n")

	}

	ramukaka.ChangeMyMind("Where do I stay", "Where do I stay in india")
	oops, _ = ramukaka.Ask("Where do I stay")
	if !oops {
		fmt.Printf("4 found\n")
	} else {
		fmt.Printf("4 not found\n")

	}

	oops, _ = ramukaka.Ask("Where do I stay in india")
	if !oops {
		fmt.Printf("5 found\n")
	} else {
		fmt.Printf("5 not found\n")

	}
	temp := &ramukaka.Ramukaka
	ramukaka.Show(temp)

	/*oops, _ = ramukaka.Forget("Where do I stay in india")

	oops, _ = ramukaka.Ask("Where do I stay in india")
	if !oops {
		fmt.Printf("5 found\n")
	} else {
		fmt.Printf("5 not found\n")

	}*/
}
