package entity

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Tree struct {
	Id             uuid.UUID
	Purpose        string
	Duration       int
	ReposeCount    int
	ReposeDuration int
	Garden         *Garden
}

func NewTree(garden *Garden, name string, duration int, count int, resDuration int) *Tree {
	return &Tree{
		Id:             uuid.NewV4(),
		Purpose:        name,
		Duration:       duration,
		ReposeCount:    count,
		ReposeDuration: resDuration,
		Garden:         garden,
	}
}


func (t *Tree) Start() {
	for i := 0; i < t.ReposeCount; i++ {
		fmt.Printf("Start Working, Stage [%d] \n\n",i+1)
		orig := t.Duration
		for j := 0; j < t.Duration; j++ {
			orig -= 1
			fmt.Printf("Tree Duration: %d.s\n", orig)
			if j == (t.Duration - 1) {
				fmt.Printf("End working, be repose %ds \n\n",t.ReposeDuration)
				//time.Sleep(time.Duration(t.ReposeDuration) * time.Second)
				for k := 0; k < t.ReposeDuration; k++ {
					fmt.Printf("Repose duration: %d \n",k)
					time.Sleep(time.Second)
				}
			}
			time.Sleep(time.Second)
		}
	}
	fmt.Println("Working End, Tree Alive!!!")
}
