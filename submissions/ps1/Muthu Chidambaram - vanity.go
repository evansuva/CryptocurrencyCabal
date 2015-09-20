/*
** NAME: Muthu Chidambaram
** ID: mc4xf
** vanity.go
**
** cs4501 Fall 2015
** Problem Set 1
*/

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

func generateVanityAddress(pattern string) (*btcec.PublicKey, *btcec.PrivateKey, *btcutil.AddressPubKeyHash) {
	//Generate vanity address that matches a pattern
	pub, priv := generateKeyPair()
	addr := generateAddr(pub)
	matched, err := regexp.MatchString(pattern, addr.String())
	
	for !matched && err == nil {
		pub, priv = generateKeyPair()
		addr = generateAddr(pub)
		matched, err = regexp.MatchString(pattern, addr.String())
	}
	
	if err != nil {
		log.Fatal(err)
	}

	return pub, priv, addr
}

func main() {

	/*// In order to recieve coins we must generate a public/private key pair.
	pub, priv := generateKeyPair()

	// To use this address we must store our private key somewhere. Everything
	// else can be recovered from the private key.
	fmt.Printf("This is a private key in hex:\t[%s]\n",
		hex.EncodeToString(priv.Serialize()))

	fmt.Printf("This is a public key in hex:\t[%s]\n",
		hex.EncodeToString(pub.SerializeCompressed()))

	addr := generateAddr(pub)

	// Output the bitcoin address derived from the public key
	fmt.Printf("This is the associated Bitcoin address:\t[%s]\n", addr.String())*/
	pub, priv, addr := generateVanityAddress("[0-9]fed[0-9]")
	fmt.Printf("This is the private key in hex:\t[%s]\n", hex.EncodeToString(priv.Serialize()))
	fmt.Printf("This is the public key in hex:\t[%s]\n", hex.EncodeToString(pub.SerializeCompressed()))
	fmt.Printf("This is the vanity address:\t[%s]\n", addr.String())
}
