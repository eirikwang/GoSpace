package entity
import (
	"core"
)

type Entity interface{
	GetPosion() *core.Vertex
}
