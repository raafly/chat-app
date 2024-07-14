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
	var letter = []rune("1234567890")

	randomString := func (length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letter[seed.Intn(len(letter))]
		}

		return string(b)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("random num", randomString(6))
	}

}

// var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randomString(length int) string {
//     b := make([]rune, length)
//     for i := range b {
//         b[i] = letters[randomizer.Intn(len(letters))]
//     }
//     return string(b)
// }

func TestRandomStringSecond(t *testing.T) {
	fmt.Println("random ", RandomOTP())
}

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = [...]int{1,2,3,4,5,6,7,8,9,0}

func RandomOTP() []int {
    b := make([]int, 6)
    for i := range b {
        b[i] = letters[randomizer.Intn(len(letters))]
    }
    return b
}