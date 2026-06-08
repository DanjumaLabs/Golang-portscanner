package main

import (
	"math/rand"
	"time"
)

func stealthDelay() {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
}
