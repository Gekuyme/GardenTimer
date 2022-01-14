package entity

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type Garden struct {
	Trees map[uuid.UUID]*Tree
	//Author uuid.UUID
}

func NewGarden() *Garden {
	return &Garden{
		Trees: map[uuid.UUID]*Tree{},
	}
}
func (c *Garden) Add(d *Tree) {
	c.Trees[d.Id] = d
	fmt.Println("Tree added to Garden")
}
