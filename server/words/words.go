package words

import (
	"math/rand"
)

var words = []string{
	"Pencil",
	"Apple",
	"Headphones",
	"Laptop",
	"PS4",
	"Plane",
	"Boat",
	"Car",
	"Fish",
	"Dog",
	"Cat",
}

func randShuffle(a []string) []string {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func GetRandomWords(numberOfWords int) []string {
	shuffledWords := randShuffle(words)
	shiftOffset := numberOfWords

	// if numberOfWords of words is greater then
	if shiftOffset > len(words) {
		shiftOffset = len(words)
	}

	shiftedArray := shuffledWords[:shiftOffset]

	return shiftedArray
}
