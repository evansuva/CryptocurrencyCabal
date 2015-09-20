/*
** spend.go
** 
** cs4501 Fall 2015
** Problem Set 1
*/

/*
For this program to execute correctly the following needs to be provided:

- An internet connection
- A private key
- A receiving address
- The raw json of the funding transaction
- The index into that transaction that funds your private key.

The program will formulate a transaction from these provided parameters
and then it will dump the newly formed tx to the command line as well as
try to broadcast the transaction into the bitcoin network. The raw hex dumped
by the program can parsed into a 'semi' human readable format using services
like: https://blockchain.info/decode-tx

*/

package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/wire"
)

var toaddress = flag.String("toaddress", "", "The address to send Bitcoin to.")
var privkey = flag.String("privkey", "", "The private key of the input tx.")
var txid = flag.String("txid", "", "The transaction id corresponding to the funding Bitcoin transaction.")
var vout = flag.Int("vout", -1, "The index into the funding transaction.")
var amount = flag.Int64("amount", -1, "The number of satoshis for the transaction" )

// Use 10,000 satoshi as the standard transaction fee. 
const TX_FEE = 10000

type requiredArgs struct {
	txid      *wire.ShaHash
	vout      uint32
	toAddress btcutil.Address
	privKey   *btcec.PrivateKey
  amount    int64
}

// getArgs parses command line args and asserts that a private key and an
// address are present and correctly formatted.
func getArgs() requiredArgs {
	flag.Parse()
	if *toaddress == "" || *privkey == "" || *txid == "" || *vout == -1 || *amount == -1 {
		fmt.Println("\nThis program generates a bitcoin trans action that moves coins from an input to an output.\n" +
			"You must provide a key, a receiving address, a transaction id (the hash\n" +
			"of a tx), the amount to be transferred (in satoshis), and the index into\n" +
      "the outputs of that tx that fund your address. Use http://blockchain.info/pushtx\n" +
			"to send the raw transaction.\n")
		flag.PrintDefaults()
		fmt.Println("")
		os.Exit(0)
	}

	pkBytes, err := hex.DecodeString(*privkey)
	if err != nil {
		log.Fatal(err)
	}

	// PrivKeyFromBytes returns public key as a separate result, but can ignore it here.
	key, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	addr, err := btcutil.DecodeAddress(*toaddress, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}

	txid, err := wire.NewShaHashFromStr(*txid)
	if err != nil {
		log.Fatal(err)
	}

	args := requiredArgs{
		txid:      txid,
		vout:      uint32(*vout),
		toAddress: addr,
		privKey:   key,
    amount:    int64(*amount),
	}

	return args
}

type BlockChainInfoTxOut struct {
	Value     int    `json:"value"`
	ScriptHex string `json:"script"`
}

type blockChainInfoTx struct {
	Ver     int                   `json:"ver"`
	Hash    string                `json:"hash"`
	Outputs []BlockChainInfoTxOut `json:"out"`
}

// Uses the txid of the target funding transaction and asks blockchain.info's
// api for information (in json) relaated to that transaction.
func lookupTxid(hash *wire.ShaHash) *blockChainInfoTx {

	url := "https://blockchain.info/rawtx/" + hash.String()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(fmt.Errorf("Tx Lookup failed: %v", err))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(fmt.Errorf("TxInfo read failed: %s", err))
	}

	//fmt.Printf("%s\n", b)
	txinfo := &blockChainInfoTx{}
	err = json.Unmarshal(b, txinfo)
	if err != nil {
		log.Fatal(err)
	}

	if txinfo.Ver != 1 {
		log.Fatal(fmt.Errorf("Blockchain.info's response seems bad: %v", txinfo))
	}

	return txinfo
}

// getFundingParams pulls the relevant transaction information from the json returned by blockchain.info
// To generate a new valid transaction all of the parameters of the TxOut we are
// spending from must be used.
func getFundingParams(rawtx *blockChainInfoTx, vout uint32) (*wire.TxOut, *wire.OutPoint) {
	blkChnTxOut := rawtx.Outputs[vout]

	hash, err := wire.NewShaHashFromStr(rawtx.Hash)
	if err != nil {
		log.Fatal(err)
	}

	// Then convert it to a btcutil amount
	amnt := btcutil.Amount(int64(blkChnTxOut.Value))

	if err != nil {
		log.Fatal(err)
	}

	outpoint := wire.NewOutPoint(hash, vout)

	subscript, err := hex.DecodeString(blkChnTxOut.ScriptHex)
	if err != nil {
		log.Fatal(err)
	}

	oldTxOut := wire.NewTxOut(int64(amnt), subscript)

	return oldTxOut, outpoint
}

func main() {
	// Pull the required arguments off of the command line.
	reqArgs := getArgs()

	// Get the bitcoin tx from blockchain.info's api
	rawFundingTx := lookupTxid(reqArgs.txid)

	// Get the parameters we need from the funding transaction
	oldTxOut, outpoint := getFundingParams(rawFundingTx, reqArgs.vout)

	// Formulate a new transaction from the provided parameters
	tx := wire.NewMsgTx()

	// Create the TxIn
	txin := createTxIn(outpoint)
	tx.AddTxIn(txin)

	// Create the TxOut
	txout := createTxOut(oldTxOut.Value, reqArgs.toAddress)
	txrem := createTxRemainder(oldTxOut.Value)
	tx.AddTxOut(txout)
  tx.AddTxOut(txrem)

	// Generate a signature over the whole tx.
	sig := generateSig(tx, reqArgs.privKey, oldTxOut.PkScript)
	tx.TxIn[0].SignatureScript = sig

	// Dump the bytes to stdout
	dumpHex(tx)

	// Send the transaction to the network
  broadcastTx(tx)
}

// createTxIn pulls the outpoint out of the funding TxOut and uses it as a reference
// for the txin that will be placed in a new transaction.
func createTxIn(outpoint *wire.OutPoint) *wire.TxIn {
	// The second arg is the txin's signature script, which we are leaving empty
	// until the entire transaction is ready.
	txin := wire.NewTxIn(outpoint, []byte{})
	return txin
}

// createTxOut generates a TxOut that can be added to a transaction. 
func createTxOut(inCoin int64, addr btcutil.Address) *wire.TxOut {
	// Pay the minimum network fee so that nodes will broadcast the tx.
	outCoin := *amount
  if outCoin == 0 {
    log.Fatal("You probably don't want to transfer 0 satoshis")
  }
  if outCoin > inCoin - TX_FEE {
    log.Fatal("Can't complete transaction--the request amount" +
    " is larger than available funds after transaction fee")
  }
	// Take the address and generate a PubKeyScript out of it
	script, err := txscript.PayToAddrScript(addr)
	if err != nil {
		log.Fatal(err)
	}
	txout := wire.NewTxOut(outCoin, script)
	return txout
}

func createTxRemainder(inCoin int64) *wire.TxOut {
  remainder := inCoin - *amount - TX_FEE

  // Put the remainder back in our wallet
	pkBytes, err := hex.DecodeString(*privkey)
	if err != nil {
		log.Fatal(err)
	}
  // Get pubkey object
	_, pubkey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

  // Get an address object from WIF string
  changeAddress := generateAddr(pubkey)

  script, err := txscript.PayToAddrScript(changeAddress)
  if err != nil {
    log.Fatal(err)
  }
  remtx := wire.NewTxOut(remainder, script)
  return remtx
}

//taken from address.go
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


// generateSig requires a transaction, a private key, and the bytes of the raw
// scriptPubKey. It will then generate a signature over all of the outputs of
// the provided tx. This is the last step of creating a valid transaction.
func generateSig(tx *wire.MsgTx, privkey *btcec.PrivateKey, scriptPubKey []byte) []byte {

	// The all important signature. Each input is documented below.
	scriptSig, err := txscript.SignatureScript(
		tx,                   // The tx to be signed.
		0,                    // The index of the txin the signature is for.
		scriptPubKey,         // The other half of the script from the PubKeyHash.
		txscript.SigHashAll,  // The signature flags that indicate what the sig covers.
		privkey,              // The key to generate the signature with.
		true,                 // The compress sig flag. This saves space on the blockchain.
	)
	if err != nil {
		log.Fatal(err)
	}

	return scriptSig
}

// dumpHex dumps the raw bytes of a Bitcoin transaction to stdout. This is the
// format that Bitcoin wire's protocol accepts, so you could connect to a node,
// send them these bytes, and if the tx was valid, the node would forward the
// tx through the network.
func dumpHex(tx *wire.MsgTx) {
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	tx.Serialize(buf)
	hexstr := hex.EncodeToString(buf.Bytes())
	fmt.Println("Here is your raw bitcoin transaction:")
	fmt.Println(hexstr)
}

type sendTxJson struct {
	RawTx string `json:"rawtx"`
}

// broadcastTx tries to send the transaction using an api that will broadcast
// a submitted transaction on behalf of the user.
//
// The transaction is broadcast to the bitcoin network using this API:
//    https://github.com/bitpay/insight-api
// 
func broadcastTx(tx *wire.MsgTx) {
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	tx.Serialize(buf)
	hexstr := hex.EncodeToString(buf.Bytes())

	url := "https://insight.bitpay.com/api/tx/send"
	contentType := "application/json"

	fmt.Printf("Sending transaction to: %s\n", url)
	sendTxJson := &sendTxJson{RawTx: hexstr}
	j, err := json.Marshal(sendTxJson)
	if err != nil {
		log.Fatal(fmt.Errorf("Broadcasting the tx failed: %v", err))
	}
	buf = bytes.NewBuffer(j)
	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		log.Fatal(fmt.Errorf("Broadcasting the tx failed: %v", err))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The sending api responded with:\n%s\n", b)
}
