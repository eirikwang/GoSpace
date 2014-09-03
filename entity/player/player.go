package player

import (
	"core"
)
type Player struct {
}

func (entity *Player) GetPosition() (v *core.Vertex){
	return &core.Vertex{0,0}
}



