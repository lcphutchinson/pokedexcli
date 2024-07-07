package pokedex

import (
	"sync"
	"github.com/lcphutchinson/pokejson"
)

type Pokedex struct {
	entries	map[string]pokejson.Pokemon
	mu	*sync.RWMutex
}

func (p Pokedex) Add(mon pokejson.Pokemon) bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, ok := p.entries[mon.Name]
	if ok {
		return false
	}
	p.entries[mon.Name] = mon
	return true
}

func (p Pokedex) Get(mon string) (pokejson.Pokemon, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	data, ok := p.entries[mon]
	if !ok {
		return pokejson.Pokemon{}, false
	}
	return data, true
}

func (p Pokedex) List() []string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var lst []string
	for _, mon := range p.entries {
		lst = append(lst, mon.Name)
	}
	return lst
}

func NewDex() Pokedex {
	return Pokedex{
		entries:	map[string]pokejson.Pokemon{},
		mu:		&sync.RWMutex{},
	}
}
