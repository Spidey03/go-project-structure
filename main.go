package main

import (
	"fmt"

	"github.com/Spidey03/go-project-structure/desert"
	"github.com/Spidey03/go-project-structure/desert/castle"
	"github.com/Spidey03/go-project-structure/space"
)

func main() {
	fmt.Println("Main Says: Hi, there!!")
	fmt.Println("Astronaut Says: ", space.Speak())
	fmt.Println("Knight Says: ", castle.Speak())
	fmt.Println("Cowboy Says: ", desert.Speak())
}
