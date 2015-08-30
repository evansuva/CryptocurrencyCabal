---
title: "PS1: Bitcoin Transactions"
date: '2015-08-31'
nocomment: false
menu: "hidden" 
---

# Problem Set 1: Bitcoin Transactions

   <div class="due">
Due: Tuesday, 15 September at 8:29pm
   </div>

<!--more-->

## Purpose

The goal of this assignment is for everyone in the class to understand
how keys, addresses, and transactions work in bitcoin.  In addition,
this assignment should help everyone get up-to-speed with the software
tools (including the Go programming language) we will use in later
assignments.

### Collaboration Policy

For this assignment, everyone should submit their own assignment and
should writeup their own answers to the questions as well as execute all
the required transactions with your own keys.

You may, and are encouraged to, discuss all of the problems with anyone
else you want (both on-line using the course web site or any other means
you choose, and in person), and it is okay to share code with others so
long as you understand everything in all of the code you use.

### Submission

Submit your answers as a single PDF file using [this
link](https://www.dropbox.com/request/KV26Qd6UrcB9j6nBLVw8).  The name
of your file should be `<your email ID>-ps1.pdf`.

Your submission should include clearly marked answers for all the
problems (highlighted in yellow).

# Blockchain Sleuthing

You should have received some bitcoin at the address you submitted in
the registration form.  (If you did not receive any bitcoin, contact the
course staff right away!)
 
From that transaction, you have an address that was used to send you the
bitcoin.  For these questions, your goal is to figure out what you can
about the way bitcoin was distributed to the class.  For these
questions, manual sleuthing should be sufficient (although ambitious
students may find ways to automate this).

   <div class="problem"> 

**Problem 1.** Answer the following questions about the transaction
  where you received the bitcoin.  If you received more than one
  transfer, include all the transaction IDs).

a. What is the transaction ID?

b. What was the transaction fee for the transaction?  (Give your answer
in BTC, as well as current approximate US dollar value.)

c. What was the total value of all the transactions in the block
containing your transfer?  (Note:
[https://blockchain.info](https://blockchain.info) provides this info
conveniently, although you could compute it yourself)

d. How long did it take from when the transaction was received until it
had 3 confirmations?  (Include an explanation of how you estimated this
in your answer.)
    </div>

Bitcoin advocates often taut its "anonymity", but bitcoin transactions
are actually publicly visible.  

   <div class="problem"> 
**Problem 2.** See how much can you figure out about the way bitcoin was
transferred to students in the class, starting from your transactions.

a. Identify the bitcoin addresses of what are likely to be other
students in the class (you could potentially find 27, but it is enough
to find 3). 

b. Trace back the source of the bitcoin as far as you can.  Bonus points
if you can figure out from which exchange the bitcoin was purchased and
when.

c. (Bonus) Can you learn anything about where the send of the bitcoin is
located geographically?  (In this case, you have external information to
know I'm in Charlottesville, but what could you learn about the sender's
probable location just from the information in the blockchain?)

   </div>

Hints:

- Start by looking at the transaction that sent bitcoin to your
  receiving address.  You can search for this by searching for your
  recieving address at
  [https://blockexplorer.com](https://blockexplorer.com),
  [https://insight.bitpay.com/](https://insight.bitpay.com/),
  [https://blockchain.info](https://blockchain.info), or many other
  sites that provide information about the bitcoin blockchain.  

- You can go forward, by following what happened with the "change" from
  that transaction.

- You can go backward, by following transactions to the sending address.

   <div class="problem">

**Problem 3.** Suppose a malicious developer wanted to distribute a
  bitcoin wallet implementation that would steal as much bitcoin as
  possible from its users with a little chance as possible of getting
  caught.  (a) Explain things a malicious developed might do to create
  an evil wallet.  (b) How confident are you your money is safe in the
  wallet you are using, and what would you do to increase your
  confidence if you were going to store all of your income in it?

   </div>


# Getting GOing

You are free to use any programming language and open source bitcoin
libraries and openly-licensed code you want for this assignment, but
must follow the license requirements of any code you use and credit this
code in your submission.  

The directions we provide use the [BTC
Suite](https://github.com/btcsuite) library for bitcoin, implemented in
the [Go](https://golang.org/).  

Most of you do not have any experience using Go, but it is not a
difficult language to learn coming from experience with Java (which all
of you have), and although its not [my favorite programming
language](http://rust-class.org) it is a language that nearly everyone
finds enjoyable to use and it is becoming widely used in industry
(especially at Google, where it was developed).  The main reason we are
using it for this, though, is because the [BTC
library](https://github.com/btcsuite) is the best bitcoin library we are
aware of, and it is written in Go.

If you are comfortable learning a new programming language by diving
right into moderately complex programs and figuring out things as you
go, you should be able to jump right into this assignment.  If you
prefer a more structured introduction to Go, there are many tutorials
available, including the [Tour of
Go](https://tour.golang.org/welcome/1).  The [Go by
Example](https://gobyexample.com/) site is very helpful.  For more
documentation, visit [https://golang.org/doc/](https://golang.org/doc/).

<!--
# Going Gets Good

After you have installed go, you can set up the BTC library by doing:

```shell
> go get github.com/btcsuite/btcec
```
-->

# Obtain the Starting Code

Before continuing with this assignment, you should [set up git and your
github account]({{< relref "tools/github.md" >}}) and follow the
directions there to set up your private repository containing the
starting code for ps1.  (It may seem like overkill to use git for this
assignment since you will not need to write much code or work with
teammates on this one.  But, it is good to get experience using git and
will become necessary to work effectively with teammates for later
projects.)

Once you have finished setting up your `ps1` repository, it should
contain the files:
 
- [keypair.go](https://github.com/CryptocurrencyCabal/ps1/blob/master/keypair.go):
  code for generating a bitcoin key pair (including its public address).
- [spend.go](https://github.com/CryptocurrencyCabal/ps1/blob/master/spend.go): code for generating a bitcoin transaction.

# Elliptic Curve Cryptography

The btcsuite library includes,
[btcec](https://github.com/btcsuite/btcd/tree/master/btcec), an implementation of the
ECDSC algorithm using the
[secp256k1](https://en.bitcoin.it/wiki/Secp256k1) elliptic curve used by
bitcoin.

Examine the
[btcec.go](https://github.com/btcsuite/btcd/blob/master/btcec/btcec.go) code.

   <div class="problem">
**Problem 4.** Find the <span class="math">_y_<sup>2</sup> =
  _x_<sup>3</sup> + 7</span> curve in this code.  (For your answer, just
  submit the line numbers for the code that tests if a given point is on the curve.)
   </div>

Elliptic curves for cryptography needs really big numbers.  The modulus for the `secp256k1` curve is found on [line 929](https://github.com/btcsuite/btcd/blob/BTCD_0_11_1_BETA/btcec/btcec.go#L909):
```
 secp256k1.P = fromHex("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F")
```
This should be the value <span class="math">2<sup>256</sup> - 2<sup>32</sup> - 2<sup>9</sup> - 2<sup>8</sup> - 2<sup>7</sup> - 2<sup>6</sup> - 2<sup>4</sup> - 1.  

   <div class="problem">

**Problem 5.** Verify that the modulus used as `secp256k1.P` in
  `btcec.go` is correct.  You can do this either using
  [math/big](https://golang.org/pkg/math/big/), Go's bit integer library
  to do computations on such large numbers, or by computing it by hand.
  (For your answer, just show how you verified the modulus.  Including a
  snippet of code is fine.)

   </div>

## Generating a Key Pair

We have provided code in `keypair.go` that generates a bitcoin key pair.
You can try running this by running `go run keypair.go` (or you can
compile it with `go build keypair.go` and then run the resulting
executable `keypair`).  It will print out the generated private key and
corresponding public bitcoin address.  (Try running it a few times to
see that it produces a different key pair each time.)

Keys for ECDSC are generated by choosing a random private key, <span
class="math">_k_</span>, and finding the corresponding public key by
"multiplying" it by <span class="math">_G_</span>, the generator point.
(Multiplication here is not standard multiplication, but multiplication
on the elliptic curve, as discussed in [Class 3]({{< relref "classes/class3.md" >}}).)  The point <span
class="math">_G_</span> is defined by [lines
912-3](https://github.com/btcsuite/btcd/blob/BTCD_0_11_1_BETA/btcec/btcec.go#L912)):  
   ```
secp256k1.Gx = fromHex("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798")
secp256k1.Gy = fromHex("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8")
   ```

The resulting point is the public key.  It is easy to derive the public
key from the private key, but believed to be hard to learn anything
useful about the private key from the public key.  (The belief that it
is hard to reverse the elliptic curve multiplication is based on the
assumption that it is hard to compute discrete logarithms, which is not
proven, but underlies much of modern cryptography.)

The code for generating a new keypair is in [keypair.go](https://github.com/CryptocurrencyCabal/ps1/blob/master/keypair.go#L20):
````go
func generateKeyPair() (*btcec.PublicKey, *btcec.PrivateKey) {
   priv, err := btcec.NewPrivateKey(btcec.S256())
   if err != nil {
       // There was an error. Log it and bail out
       log.Fatal(err)
   }
   return priv.PubKey(), priv
}
````

The important work is the
[NewPrivateKey](https://github.com/btcsuite/btcd/blob/BTCD_0_11_1_BETA/btcec/privkey.go#L38)
function.  

   <div class="problem"> 
**Problem 5.** What are _all_ the things you need to trust if you are
going to send money to the key generated by running `keypair.go`?  You
should assume that you are an ultra-paranoid multi-billionaire who
intends to transfer her entire fortune to the generated address.  
    </div>

## Generating a Vanity Key

Anyone can have a bitcoin address like
`1H7tu2qUAyyr5aX1WA17eyvbetAGyqxfKZ` or
`1L3iGYBD5wbiki2SYUT5wmupy4TTgmEBg3`, but suppose you want a bitcoin
address that includes the name of your cat or your birthday.

  <div class="problem"> 
**Problem 6.** Define a function, `func
generateVanityAddress(pattern string) (*btcec.PublicKey,
*btcec.PrivateKey)` where `pattern` is a regular expression.  It should
return a valid key pair where the corresponding public address [matches
the pattern](http://golang.org/pkg/regexp/#MatchString).  
   </div>

You should be able to use your function to generate an address that
includes the digits of pi in sequence:
`generateVanityAddress("3.*1.*4.*1.*5.*9.*")` or contains `DAVE` without
any adjacent letters (`generateVanityAddress("[0-9]DAVE[0-9]")`) in its
public bitcoin address.  In deciding how vain you want to be for the
next exercise, think about how the running time scales with the number
of strings that match the target pattern.

  <div class="problem"> 
**Problem 7.** Use your `generateVanityAddress` function to create
your own vanity address containing your first or last name (or if that
is too long it could just be your initials).  If you are extra vain,
create a address where your name appears at the beginning (after the
initial `1`).  (Note that uppercase 'O' and 'I' and lowercase 'l' are
not used in any address, so if your name includes these letters you will
have to be creative.)
  </div>

   <div class="problem">
**Problem 8.**   Is your vanity address more or less secure than the first address you generated?
   </div>

There are on-line services that produce vanity addresses, like
   [http://bitcoinvanitygen.com/](http://bitcoinvanitygen.com/).  You
   should contemplate briefly whether using such a site is [more vain or
   stupid](http://www.reddit.com/r/Bitcoin/comments/21foj9/funds_currently_being_stolen_from_vanity/).

# Bitcoin Transactions

Having an address is not much fun without any funds!  

You should have received some money to the address you submitted in
[PS0](../../ps/ps0/ps0.md).  For these questions, you will need to have
actual money to transfer, so be careful to make small transfers for
these questions in case something goes wrong.  (In the event that you do
lose all of your bitcoin, you can get a new transfer to an adress of
your choosing by explaining to me what you have learned about software
development and or best practices.  It is not necessary to buy your own
bitcoin, even if you lose all of the original transfer.)

   <div class="problem">
<b>Problem 9.</b> Make a **small** (e.g., 1 mBTC) transfer from
your wallet address to your vanity address.  You can do this using
your MultiBit wallet.  Find the transaction in the blockchain (you can do this by
searching for your vanity address at
[insight.bitpay.com](https://insight.bitpay.com/) or
[blockchain.info](https://blockchain.info/)). You will need the
transaction ID for the next exercise.  (For your answer, just provide the transaction ID.)
   </div>

Once you have located the transaction that sends bitcoin to your vanity
address you should notice several things.

Your wallet most likely send bitcoins to your address _and_ back to a
new address.  We call this second address the 'change' address. Notice
that each output has an ordered position. This index (known as the `vout`)
along with the transaction ID lets us uniquely identify transaction
outputs. This is important if you want to use those outputs in a new
transaction.

   <div class="problem">
<b>Problem 10.</b> Transfer the coin from your vanity address back to your wallet. To do this you can run `spend.go` in  ps1. You can provide the parameters needed for the transaction at the command line (it is not necessary to modify the code).  
   </div>

If done correctly the script should look this when executed:

```bash
> go run spend.go \
    -txid="6bf98ac2e1a25ea9c2951bb6a40262e054514236f864a3414c16fe6b3a5f3f62" \
    -vout=0 \
    -privkey="8d66a9f85c4f737b231d1af0bd917c8e02f05d616f26c41f269a194a10c29029" \
    -address="1DTFjSGeij6woxyaJFaYLMzciKCYP83ZNB"

Here is your raw bitcoin transaction:
0100000001623f5f3a6bfe164c41a364f836425154e06202a4b61b95c2a95ea2e1c28af96b000000006a47304402202487724e5e42e60b5d8d98636797339b09db8c70452959cf9c48a0be1a45aba302203465c00d34a190dd1294c9c1561af7c0bcc2b0d452ae74b6d126ad06e0f03624012103c4a967e4e04f32f6b25b723c4d994bcc5f1f6db2cd7757d25643af0d1a1b4c1effffffff01905f0100000000001976a91488989f2c91337ad5241abedeafe8ed96fa28940488ac00000000
The sending api responded with:
{"txid":"8d1687cc05cd3264c8d41acf19efcd08ecd44b5b57d3105fe081fdfd4baa6b0b"}
```

Notice that when the command above is run with these parameters it only
works
[once](https://blockchain.info/tx/8d1687cc05cd3264c8d41acf19efcd08ecd44b5b57d3105fe081fdfd4baa6b0b). If
you tried to run it again or send the coins somewhere else the script
would fail.

   <div class="problem"> 
<b>Problem 11.</b> Try to double spend the same bitcoin by sending it
twice to your wallet address.  Figure out as much as you can about what
happens when the double spend transactions are attempted.  See how close
you can get to obtaining two verified transactions spending the same
coin (e.g., can you achieve two transactions with at least one
confirmation each?)
   </div>

Please don't try to actually rip anyone off; only attempt double
spending with your own addresses (or those of willing classmates).  You
are not expected to actually be able to double-spend successfully (if
this were easy, bitcoin would not work well as a currency!), but should
be able to learn something by attemting to do this.

### Submission

Follow the submission instructions at the beginning of this page by
8:29pm on Tuesday, 15 September.

