package names

import (
	"fmt"
	"math/rand"
)

var adjs = []string{
	"eager",
	"lazy",
	"quaint",
	"sparkling",
	"clever",
	"mushy",
	"scary",
	"clumsy",
	"silly",
	"witty",
	"happy",
}

var names = []string{
	"giraffe",
	"panda",
	"mole",
	"squirrel",
	"raccoon",
	"badger",
	"skunk",
	"chipmunk",
	"toad",
}

func Generate() string {
	return fmt.Sprintf("%s-%s-%d", adjs[rand.Int()%len(adjs)], names[rand.Int()%len(names)], rand.Int31()%10000)
}
