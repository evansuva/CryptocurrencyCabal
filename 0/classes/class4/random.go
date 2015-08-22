package main

import (
    "fmt"
    "crypto/rand"
)

func main() {
    b := make([]byte, 16)
    for {
	_, err := rand.Read(b)
        if err != nil {
  	    // error
        } else {
	    fmt.Printf("%x", b)
	}
    }
}
