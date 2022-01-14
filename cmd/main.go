package main

import (
	"Garden/entity"
	"fmt"
)

func main() {
	garden := entity.NewGarden()
	fmt.Println("Garden: ", garden)
	tree := entity.NewTree(garden, "learning", 5, 2, 5)
	garden.Add(tree)
	tree.Start()
}
