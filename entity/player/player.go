package player

import (
	"github.com/eirikwang/GoSpace/core"
)
type Player struct {
}

func (entity *Player) GetPosition() core.Vertex{
	return core.Vertex{0,0}
}



