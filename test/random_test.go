package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println(random.Int())
	fmt.Println(random.Int())
	fmt.Println(random.Int())
}

func TestRandomString(t *testing.T) {
	seed := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMOPQRSTUVWXYZ")

	randomString := func (length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letter[seed.Intn(len(letter))]
		}

		return string(b)
	}

	fmt.Println("random string 5 ", randomString(10))
}

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[randomizer.Intn(len(letters))]
    }
    return string(b)
}

func TestRandomStringSecond(t *testing.T) {
	fmt.Println("random ", randomString(10))
}