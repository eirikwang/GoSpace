package entity
import (
	"github.com/eirikwang/GoSpace/core"
)

type Entity interface{
	GetPosition() core.Vertex
}
