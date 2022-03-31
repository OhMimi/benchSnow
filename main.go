package main

import (
	"fmt"

	bwSnowflake "github.com/bwmarrin/snowflake"
	ftSnowflake "github.com/fish-tennis/snowflake"
)

func main() {

}

func NewBW() *bwSnowflake.Node {
	// Create a new Node with a Node number of 1
	node, err := bwSnowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return node
}

func NewFt() *ftSnowflake.SnowFlake {
	workerId := uint16(1)
	return ftSnowflake.NewSnowFlake(workerId)
}
