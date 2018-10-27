package coderunner

import "sync/atomic"

type idGen struct {
	id int32
}

func (g *idGen) genID() int32 {
	return atomic.AddInt32(&g.id, 1)
}
