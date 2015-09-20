package main

import (
  "fmt"
  "math/big"
)

func main() {
  // p is the big int representation of the modulus for the secp256k1 curve
  var p big.Int
  p.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)

  // now we find the value of 2^256 - 2^32 - 2^9 - 2^8 - 2^7 - 2^6 - 2^4 - 1
  diff_256 := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
  diff_32 := new(big.Int).Exp(big.NewInt(2), big.NewInt(32), nil)
  diff_9 := new(big.Int).Exp(big.NewInt(2), big.NewInt(9), nil)
  diff_8 := new(big.Int).Exp(big.NewInt(2), big.NewInt(8), nil)
  diff_7 := new(big.Int).Exp(big.NewInt(2), big.NewInt(7), nil)
  diff_6 := new(big.Int).Exp(big.NewInt(2), big.NewInt(6), nil)
  diff_4 := new(big.Int).Exp(big.NewInt(2), big.NewInt(4), nil)
  diff_1 := big.NewInt(1)

  diff_256.Sub(diff_256, diff_32)
  diff_256.Sub(diff_256, diff_9)
  diff_256.Sub(diff_256, diff_8)
  diff_256.Sub(diff_256, diff_7)
  diff_256.Sub(diff_256, diff_6)
  diff_256.Sub(diff_256, diff_4)
  diff_256.Sub(diff_256, diff_1)

  if diff_256.Cmp(&p) == 0 {
    fmt.Printf("difference: %v\n", diff_256.String())
    fmt.Printf("modulus: %v\n", p.String())
    fmt.Println("The modulus given for the secp256k1 curve and the large prime are equal")
  }

}
