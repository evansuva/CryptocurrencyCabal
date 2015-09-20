package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"regexp"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
  
  vpub, vpriv := generateVanityAddress("Yolo")
  
  fmt.Printf("This is a private key in hex:\t[%s]\n",
		hex.EncodeToString(vpriv.Serialize()))

  fmt.Printf("This is a public key in hex:\t[%s]\n",
	  hex.EncodeToString(vpub.SerializeCompressed()))

  vaddr := generateAddr(vpub)

  // Output the bitcoin address derived from the public key
  fmt.Printf("This is the associated Bitcoin address:\t[%s]\n", vaddr.String())
  
}

func generateVanityAddress(pattern string)(*btcec.PublicKey, *btcec.PrivateKey) {
  
  for {
      pub, priv := generateKeyPair()

      matched, _ := regexp.MatchString(pattern, generateAddr(pub).String())
      
      if matched {
	return pub, priv
      }
      
    }
}
  
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