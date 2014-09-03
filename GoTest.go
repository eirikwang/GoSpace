package main

import (
	"fmt"
	"github.com/eirikwang/GoSpace/core"
	//"code.google.com/p/draw2d/"
	"github.com/eirikwang/GoSpace/entity/player"
	"github.com/eirikwang/GoSpace/entity"
	//"github.com/go-gl/gl"
	"github.com/eirikwang/GoSpace/world"
)

var (
	array = [...]int32{
	1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16}
	slice = array[:]
)

func main() {
	fmt.Printf("Hello world!")
	var ent entity.Entity
	value := &core.Vertex{10, 10}
	v := world.GenerateMap(value)
	fmt.Println(v)
	ent = &player.Player{}
	fmt.Println(ent.GetPosition())


}
