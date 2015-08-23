Title: Project 1 - Bitcoin Transactions
Date: 2015-01-19
Category: PS
Tags: Problem Sets, Kernel
Slug: project1

   <div class="due">
Due: Friday, 30 January at 11:59pm
   </div>

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


### Types of Questions

There are four types of questions on this assignment (and other projects in this class):

   <div class="exercise"> 
**Exercises** are things you should do and learn from, but there is nothing
that you are expected to turn in for these.  You will often need things
you should learn from the exercises or their actual results for later
questions.
   </div>

   <div class="discuss"> **Discussion Questions** are questions that we
hope you will think about and contribute to class discussions about in
the web site comments.  It is not compulsory to contribute to all
discussions, and better to make meaningful contributions when you have
something useful to add to a conversation than to make
"pseudo-contributions" because you feel you are required to.  You should
definitely read other students' contributions before making your own,
and avoid duplication. You should be worried, though, if you are not
contributing usefully to a significant fraction of the discussion
questions.  </div>
  
   <div class="problem"> **Problems** are questions that you will be
expected to submit something for in the assignment submission.
Sometimes these will involve writing programs (although it will often
not be necessary to submit the program itself, only the answer you
obtain by using the program). 
   </div>

   <div class="challenge"> **Challenges** are suggestions for things you
might do to particularly impress me and your classmates (and perhaps,
the rest of the world!)  It is not expected that most students complete
challenges (and some may be hard enough that it is not expected that
anyone completes them), and you should avoid spending lots of time on
these until you have finished the rest of the assignment (and any other
important responsibilities you have outside of this class).  (Unlike the
rest of the assignment, the deadline does not apply to challenges.)

</div>

In general, you should assume the purpose of everything is for you to
learn and do interesting things, not to produce something gradable.  If
you have ideas for doing something more interesting than what is
requested, by all means pursue them!


# Getting Going

You are free to use any programming language and open source bitcoin
libraries and openly-licensed code you want for this assignment, but
must follow the license requirements of any code you use and credit this
code in your submission.  

The directions we provide use the [BTC
Suite](https://github.com/btcsuite) library for bitcoin, implemented in
the [Go](https://golang.org/).  

Hardly any of you have experience using Go already, but it is not a
difficult language to learn coming from experience with Java (which all
of you have), and although its not [my favorite programming
language](http://rust-class.org) it is a language that nearly everyone
who learns enjoys programming in.  The main reason we are using it for
this, though, is because the [BTC library](https://github.com/btcsuite)
is the best bitcoin library we are aware of, and it is written in Go.

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
github account](|filename|../../tools/github.md) and follow the
directions there to set up your private repository containing the
starting code for project 1.  (It may seem like overkill to use git for
this assignment since you will not need to write much code or work with
teammates on this one.  But, it is good to get experience using git and
will become necessary to work effectively with teammates for later
projects.)

Once you have finished setting up your `project1` repository, it should
contain the files:
 
- [keypair.go](https://github.com/CryptoCurrencyCafe/project1/blob/master/keypair.go):
  code for generating a bitcoin key pair (including its public address).
- [spend.go](https://github.com/CryptoCurrencyCafe/project1/blob/master/spend.go): code for generating a bitcoin transaction.

# Elliptic Curve Cryptography

The btcsuite library includes,
[btcec](https://github.com/btcsuite/btcec), an implementation of the
ECDSC algorithm using the
[secp256k1](https://en.bitcoin.it/wiki/Secp256k1) elliptic curve used by
bitcoin.

Examine the
[btcec.go](https://github.com/btcsuite/btcec/blob/master/btcec.go) code.

   <div class="exercise">
<b>Exercise 1.</b> Find the <span class="math">_y_<sup>2</sup> = _x_<sup>3</sup> + 7</span> curve in this code.
   </div>

Elliptic curves for cryptography needs really big numbers.  The modulus for the `secp256k1` curve is found on [line 672](https://github.com/btcsuite/btcec/blob/master/btcec.go#L672):
```
 secp256k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
```
This should be the value <span class="math">2<sup>256</sup> - 2<sup>32</sup> - 2<sup>9</sup> - 2<sup>8</sup> - 2<sup>7</sup> - 2<sup>6</sup> - 2<sup>4</sup> - 1.  

   <div class="sidenote">
Note: this has been corrected, see comments below.  Thanks to Ankit Gupta for noticing the mistake. 
   </div>

   <div class="exercise">
<b>Exercise 2.</b> Verify that the modulus used as `secp256k1.P` in `btcec.go` is correct.  You can do this either using [math/big](https://golang.org/pkg/math/big/), Go's bit integer library to do computations on such large numbers, or by computing it by hand.
   </div>

   <div class="challenge">
**Challenge.** Make an improvement to the library code that is accepted by the btcec developers as a pull request.
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
on the elliptic curve, as discussed in [Class
3](../../../classes/class3/class3.md).)  The point <span
class="math">_G_</span> is defined by [lines
675-6](https://github.com/btcsuite/btcec/blob/master/btcec.go#L675): ```
secp256k1.Gx, _ =
new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798",
16) secp256k1.Gy, _ =
new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8",
16) ```

The resulting point is the public key.  It is easy to derive the public
key from the private key, but believed to be hard to learn anything
useful about the private key from the public key.  (The belief that it
is hard to reverse the elliptic curve multiplication is based on the
assumption that it is hard to compute discrete logarithms, which is not
proven, but underlies much of modern cryptography.)

The code for generating a new keypair is in [keypair.go](https://github.com/CryptoCurrencyCafe/project1/blob/master/keypair.go#L20):
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
[NewPrivateKey](https://github.com/btcsuite/btcec/blob/master/privkey.go#L38)
function.  

   <div class="discuss"> <b>Discussion Question.</b> What are all the things you
need to trust if you are going to send money to the key generated by
running `keypair.go`?  You should assume that you are an ultra-paranoid
multi-billionaire who intends to transfer her entire fortune to the
generated address.  (For this question, you should think about this
yourself first, or in discussions with other students, and then post
your answer to [the discussion page](|filename|project1-trust.md).)
   </div>

## Generating a Vanity Key

Anyone can have a bitcoin address like
`1H7tu2qUAyyr5aX1WA17eyvbetAGyqxfKZ` or
`1L3iGYBD5wbiki2SYUT5wmupy4TTgmEBg3`, but suppose you want a bitcoin
address that includes the name of your cat or your birthday.

  <div class="problem"> <b>Problem 1.</b> Define a function, `func
generateVanityAddress(pattern string) (*btcec.PublicKey,
*btcec.PrivateKey)` where `pattern` is a regular expression.  It should
return a valid key pair where the corresponding public address [matches
the pattern](http://golang.org/pkg/regexp/#MatchString).  </div>

You should be able to use your function to generate an address that
includes the digits of pi in sequence:
`generateVanityAddress("3.*1.*4.*1.*5.*9.*")` or contains `DAVE` without
any adjacent letters (`generateVanityAddress("[0-9]DAVE[0-9]")`) in its
public bitcoin address.  In deciding how vain you want to be for the
next exercise, think about how the running time scales with the number
of strings that match the target pattern.

  <div class="exercise"> 
<b>Exercise 3.</b> Use your `generateVanityAddress` function to create
your own vanity address containing your first or last name (or if that
is too long it could just be your initials).  If you are extra vain,
create a address where your name appears at the beginning (after the
initial `1`).  (Note that uppercase 'O' and 'I' and lowercase 'l' are
not used in any address, so if your name includes these letters you will
have to be creative.)
  </div>

   <div class="discuss"> <b>Discussion Question.</b> Is your vanity
   address more or less secure than the first address you generated?
   What about a service like
   [http://bitcoinvanitygen.com/](http://bitcoinvanitygen.com/)?  Would
   some who uses such a service to generate their bitcoin address be [too
   vain for their own good](http://www.reddit.com/r/Bitcoin/comments/21foj9/funds_currently_being_stolen_from_vanity/)?
Think about this
yourself first, or in discussions with other students, and then initiate or join the [discussion page](|filename|project1-vanity.md).
   </div>

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

   <div class="exercise">
<b>Exercise 4.</b> Make a **small** (e.g., 0.001 BTC) transfer from
your wallet address to your vanity address.  You can do this using
MultiBit.  Find the transaction in the blockchain (you can do this by
searching for your vanity address at
[insight.bitpay.com](https://insight.bitpay.com/) or
[blockchain.info](https://blockchain.info/)). You will need the
transaction ID for the next exercise.
   </div>

Hint: If you can't find your transaction look at this
[one](https://blockchain.info/tx/8b4b61f6c1a3be10c61edd449313920cc0666ce4a20d549674bd7c8b8a83356b).
Notice that the sum of the transaction inputs (the left side of the
arrow) is slightly less than the sum of the transaction outputs (the
right side of the arrow). The difference is being collected as a
processing fee.

Once you have located the transaction that sends bitcoin to your vanity
address you should notice several things.  

Your wallet most likely sent bitcoins to your address _and_ back to a
new address that only it controls.  We call this second address the
'change' address. Notice that each output has an ordered position. This
index (known as the vout) along with the transaction ID lets us uniquely
identify transaction outputs. This is important if you want to use those
outputs in a new transaction.

   <div class="exercise">
<b>Exercise 5.</b> Transfer the coin from your vanity address back to your wallet. To do this you can run `spend.go` in  project1. You can provide the parameters needed for the transaction at the command line (it is not necessary to modify the code).  
   </div>

If done correctly the script should look this when executed:

```shell
> go run spend.go \
    -txid="6bf98ac2e1a25ea9c2951bb6a40262e054514236f864a3414c16fe6b3a5f3f62"\
    -vout=0\
    -privkey="8d66a9f85c4f737b231d1af0bd917c8e02f05d616f26c41f269a194a10c29029"\
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

   <div class="discuss"> <b>Exercise and Discussion.</b> Try to double
spend the same bitcoin by sending it twice to your wallet address.
Figure out as much as you can about what happens when the double spend
transactions are attempted.  See how close you can get to obtaining two
verified transactions spending the same coin (e.g., can you achieve two
transactions with at least one confirmation each?)

Please don't try to actually rip anyone off; only attempt double
spending with your own addresses (or those of willing classmates).  

Post about your experiences on [the discussion
page](|filename|project1-double.md).  Include links to transactions in
the blockchain to demonstrate your results.

</div>

### Submission

Submit the [Project 1 Submission Form](http://goo.gl/forms/kdIbZ33ryo) (by 11:59pm
on **Friday, 30 January**):

<iframe src="https://docs.google.com/forms/d/1I2a2T9owqTvLx7GAT8EIVf-qAhR2NU2113cSvwVOOAE/viewform?embedded=true" width="760" height="800" frameborder="0" marginheight="0" marginwidth="0">Loading...</iframe>

<p><br></br></p>

<div class="disqus">
<div id="disqus_thread"></div>
</div>
