package world

import (
	"github.com/eirikwang/GoSpace/entity"
	"fmt"
	"math/rand"
	"github.com/eirikwang/GoSpace/core"
)

type Map struct {
	Tiles [][]uint32
}

func (m *Map) IsEntityAllowed(entity *entity.Entity) {
	fmt.Println(entity)
}

func GenerateMap(v *core.Vertex) Map {
	m := make([][]uint32, v.X)
	for i := range m {
		m[i] = make([]uint32, v.Y)
		for j := range m[i] {
			m[i][j] = rand.Uint32()%2
		}
	}
	return Map{m}
}

