package engine

import (
	"math/rand"
	"sync"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986
func GenerateID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type PlayerMap struct {
	sync.RWMutex
	internal map[int]*Player
}

func makePlayerMap() *PlayerMap {
	return &PlayerMap{internal: make(map[int]*Player)}
}

func (c *PlayerMap) Load(key int) (value *Player, ok bool) {
	c.RLock()
	defer c.RUnlock()
	result, ok := c.internal[key]
	return result, ok
}

func (c *PlayerMap) Store(key int, value *Player) {
	c.Lock()
	defer c.Unlock()
	c.internal[key] = value
}

func (c *PlayerMap) LoadOrStore(key int, value *Player) (actual *Player, loaded bool) {
	c.Lock()
	defer c.Unlock()
	actual, ok := c.internal[key]
	if !ok {
		c.internal[key] = value
		return value, false
	}
	return actual, true
}

func (c *PlayerMap) Length() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.internal)
}

func (c *PlayerMap) Delete(key int) {
	c.Lock()
	defer c.Unlock()
	delete(c.internal, key)
}

type GameMap struct {
	sync.RWMutex
	internal map[string]*Game
}

func makeGameMap() *GameMap {
	return &GameMap{internal: make(map[string]*Game)}
}

func (c *GameMap) Load(key string) (value *Game, ok bool) {
	c.RLock()
	defer c.RUnlock()
	result, ok := c.internal[key]
	return result, ok
}

func (c *GameMap) Store(key string, value *Game) {
	c.Lock()
	defer c.Unlock()
	c.internal[key] = value
}

func (c *GameMap) LoadOrStore(key string, value *Game) (actual *Game, loaded bool) {
	c.Lock()
	defer c.Unlock()
	actual, ok := c.internal[key]
	if !ok {
		c.internal[key] = value
		return value, false
	}
	return actual, true
}

func (c *GameMap) Length() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.internal)
}

func (c *GameMap) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.internal, key)
}
