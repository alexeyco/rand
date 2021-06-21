package rand_test

import (
	"log"

	"github.com/alexeyco/rand"
)

func ExampleRandom_Int64() {
	n, err := rand.New().Int64(100, 10)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(n)
}

func ExampleRandom_String() {
	s, err := rand.New().String(rand.LettersAndNumbers, 10, 10)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(s)
}
