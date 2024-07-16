package main

// import (
// 	"net"
// 	"sync"

// 	"github.com/google/uuid"
// )

// type World struct {
// 	connections sync.Map
// }

// func NewWorld() *World {
// 	return &World{
// 		connections: sync.Map{},
// 	}
// }

// func (w *World) AddConnection(conn net.Conn) string {
// 	id := uuid.New().String()
// 	w.connections.Store(id, conn)

// 	return id
// }
