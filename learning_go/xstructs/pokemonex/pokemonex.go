package pokemonex

import (
	"time"
)

type User struct {
	name      string
	createdAt time.Time
	op        []Pokemon
}

type Pokemon struct {
	name  string
	tp    string
	power int
}

func (u1 User) Equals(u2 User) bool {
	// name equality
	if u1.name != u2.name {
		return false
	}
}

// Attack user(1) launches an attack on user(2).
func Attack(u1, u2 User) int {
	var i int
	for i < len(u1.op) && i < len(u2.op) {
		if u1.op[i].power > u2.op[i].power {
			return 1
		} else if u1.op[i].power == u2.op[i].power {
			return 0
		}
	}
	return -1
}

// PokeCompare compares whether two users have the same pokemons
func PokeCompare(u1, u2 User) bool {
}

func DealPokemonCards(...User) {

}
