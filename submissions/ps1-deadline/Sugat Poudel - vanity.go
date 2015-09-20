package main

import (
  "encoding/hex"
  "fmt"
  "log"
  "github.com/btcsuite/btcd/btcec"
  "github.com/btcsuite/btcd/chaincfg"
  "github.com/btcsuite/btcutil"
  "regexp"
)

// generateKeyPair creates a key pair based on the eliptic curve
// secp256k1. The key pair returned by this function is two points
// on this curve. Bitcoin requires that the public and private keys
// used to generate signatures are generated using this curve.

func generateKeyPair() (*btcec.PublicKey, *btcec.PrivateKey) {

	// Generate a private key, use the curve secpc256k1 and kill the program on
	// any errors
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		// There was an error. Log it and bail out
		log.Fatal(err)
	}

	return priv.PubKey(), priv
}

// generateAddr computes the associated bitcon address from the provided
// public key. We compute ripemd160(sha256(b)) of the pubkey and then
// shimmy the hashed bytes into btcsuite's AddressPubKeyHash type
func generateAddr(pub *btcec.PublicKey) *btcutil.AddressPubKeyHash {

	net := &chaincfg.MainNetParams

	// Serialize the public key into bytes and then run ripemd160(sha256(b)) on it
	b := btcutil.Hash160(pub.SerializeCompressed())

	// Convert the hashed public key into the btcsuite type so that the library
	// will handle the base58 encoding when we call addr.String()
	addr, err := btcutil.NewAddressPubKeyHash(b, net)
	if err != nil {
		log.Fatal(err)
	}

	return addr
}


func generateVanityAddress(pattern string) (*btcec.PublicKey, *btcec.PrivateKey) {
  var pub *btcec.PublicKey
  var priv *btcec.PrivateKey
  var addr *btcutil.AddressPubKeyHash

  valid := false

  for valid != true {
    pub, priv = generateKeyPair()
    addr = generateAddr(pub)
    fmt.Println(addr.String())
    valid, _ = regexp.MatchString(pattern, addr.String())
  }

  return pub, priv
}


func main() {

    pub, priv := generateVanityAddress("(s|S)(u|U)(g|G)(a|A)(t|T)")
    fmt.Printf("vanity address: %s\n", generateAddr(pub).String())

  	// To use this address we must store our private key somewhere. Everything
  	// else can be recovered from the private key.
  	fmt.Printf("private key: [%s]\n",
  		hex.EncodeToString(priv.Serialize()))

  	fmt.Printf("public key: [%s]\n",
  		hex.EncodeToString(pub.SerializeCompressed()))

}
