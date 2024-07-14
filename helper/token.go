package helper

import (
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("0123456789")

func RandomOTP() string {
    b := make([]rune, 6)
    for i := range b {
        b[i] = letters[randomizer.Intn(len(letters))]
    }
    return string(b)
}