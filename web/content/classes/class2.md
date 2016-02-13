---
title: "Class 2: Cryptography"
date: '2015-08-31'
print: "class2.pdf"
nocomment: false
weight: 3
menu: "classes"
---
## Schedule 

   <div class="todo">
Before the next class (Wednesday, Sept 2):

- **Read:** [_Chapter 3: The Bitcoin
Client_](https://github.com/aantonop/bitcoinbook/blob/develop/ch03.asciidoc)
and [_Chapter 4: Keys, Addresses,
Wallets_](https://github.com/aantonop/bitcoinbook/blob/develop/ch04.asciidoc)
from Andreas M. Antonopoulos, [_Mastering Bitcoin: Unlocking Digital
Cryptocurrencies_](https://github.com/aantonop/bitcoinbook) book (also
available [in
print](http://www.amazon.com/Mastering-Bitcoin-Unlocking-Digital-Crypto-Currencies/dp/1449374042)).
You can skim most of Chapter 3, especially the parts about installing
bitcoin core, but should read Chapter 4.  (You should have already read
Chapters 1 and 2.)

- **Read:** _Chapter 1: Introduction to Cryptography and
Cryptocurrencies_,
from Arvind Narayanan, Joseph Bonneau, Edward Felten, Andrew Miller,
Steven Goldfeder.  This chapter starts with cryptographic hashing and authenticated data
structures (which we are deferring until until later, but is still worth
reading now), and Section 1.3 covers digital signatures.

**Monday, September 7**: Check-up 1.  This will be a short in-class quiz
  to test your understanding of the main concepts covered so far.  It
  will cover material from the readings and classes 1-3.

**Tuesday, September 15** (8:29pm): [Problem Set 1](http://www.bitcoin-class.org/ps/ps1) due.

   </div>
<!--more-->

<center>
<iframe src="//www.slideshare.net/slideshow/embed_code/key/ppf9BW9JlnIXC5"
width="476" height="400" frameborder="0" marginwidth="0"
marginheight="0" scrolling="no"></iframe>

   <div class="caption">
Note: ink markings may not appear in the
embedded viewer.  To see them, [download the slides](/classes/class2-post.pptx).
   </div>

</center>


<!--
how is it possible to own something digital?

- copyright!

England

1662 - Licensing of the Press Act

guild of printers, "Stationer's Company" (formed in 1403, royal charter in 1557)
granted monopoly on printing [cf. Chinese granting monopoly on salt production]
exclusive right to print - responsible for censoring

ended in 1694 - no restrictions	       


Act of Queen Anne
-->

## Cryptography

_kryptos_ is a Greek root meaning hidden ("secret")

_crypto_ + _graphy_ = "secret writing"

_Decryption_ is what the intended receiver does.  
_Cryptanalysis_ is what an attacker does.  

How are cryptography and security related?
<div class="gap">

</div>

### Simple Message Cryptosystem

Two functions:

- **Encrypt:** <span class="math">_E_(_m_: byte[]) &rarr; byte[]</span>.  The input is called the
    **plaintext**; the output is called the **ciphertext**.
- **Decrypt:** <span class="math">_D_(_c_: byte[]) &rarr; byte[]</span>.

Required properties:

- **Correctness:** for all possible messages, <span class="math">_m_, _D_(_E_(_m_)) = _m_</span>
- **Security:** given the output of <span class="math">_E_(_m_)</span>, it is "hard" to learn anything interesting about <span class="math">_m_</span>.  

[_Goldwasser and Micali win Turing Award: Team honored for
‘revolutionizing the science of
cryptography'_](http://web.mit.edu/newsoffice/2013/goldwasser-and-micali-win-turing-award-0313.html),
MIT News, 13 March 2013. 

Their paper that introduced semantic security notions is:
[_Probabilistic Encryption and How to Play Mental Poker Keeping Secret
All Partial
Information_](http://groups.csail.mit.edu/cis/pubs/shafi/1982-stoc.pdf),
ACM Symposium on Theory of Computing, 1982.  (We will not get into
formal security definitions or proofs in this class, but you should take
[Mohammad Mahmoody](http://www.cs.virginia.edu/~mohammad/)'s class to
learn them.)

### Keyed Symmetric Cryptosystem

Claude Shannon, [_Communication Theory of Secrecy Systems_](http://netlab.cs.ucla.edu/wiki/files/shannon1949.pdf), 1949 (work done during World War II, but declassified later).

Two functions:

- **Encrypt:** <span class="math">_E_(_<font color="red">k</font>_: byte[], _m_: byte[]) &rarr; byte[]</span>. 
- **Decrypt:** <span class="math">_D_(_<font color="red">k</font>_, _c_: byte[]) &rarr; byte[]</span>.

Required properties:

- **Correctness:** for all possible messages, <span class="math">_m_</span>, and keys, <span class="math">_k_</span>, <span class="math">_D_(_<font color="red">k</font>_, _E_(_<font color="red">k</font>_, _m_)) = _m_</span>.
- **Security:** given <span class="math">_E_</span>, <span class="math">_D_</span>, and the output of <span class="math">_E_(<font color="red">_k_</font>, _m_)</span> it is "hard" to learn anything interesting about <span class="math">_m_</span> (without knowing <span class="math"><font color="red">_k_</font></span>).

Are these properties enough to be secure against an active attacker?
<div class="gap">

</div>

**Keyspace:** set of all possible keys.  Assume (hopefully for
  user!) that key is drawn uniformly from this set.

**Brute Force Attack:** for all possible keys, <span
  class="math">_k_<sub>i</sub></span>, try computing <span
  class="math">_D_(_k_<sub>i</sub>)</span> to see if it looks like a
  reasonable plaintext.  

In order for a brute force attack to succeed, what properties are
necessary about (1) the keyspace and (2) the message space?

<div class="gap"></div>

Where is symmetric cryptography used in your bitcoin wallet?

<div class="biggap"></div>

<!--page-->

(This material was not actually covered in class 2, but will be covered in class 3.)

## Asymmetric Cryptosystems

**Asymmetric cryptosystems** use _different functions_ for encrypting
  and decrypting, with the property that revealing the encryption
  function does not reveal the decryption function.  With Kerckhoff's
  Principle, this means there are different keys for encryption and
  decryption.

- **Generate:** produce key pair, <span class="math">(_<font color="green">KU<sub>X</sub></font>_, _<font color="red">KR<sub>X</sub></font>_)</span>, and publish the public key, <span class="math">_<font color="green">KU<sub>X</sub></font>_</span>.

- **Encrypt:** <span class="math">_E_(_<font color="green">KU<sub>X</sub></font>_: byte[], _m_: byte[]) &rarr; byte[]</span>. 

- **Decrypt:** <span class="math">_D_(_<font color="red">KR<sub>X</sub></font>_, _c_: byte[]) &rarr; byte[]</span>.

**Messages:** Sender encrypts a message with the recipient's public key.
  Recipient decrypts the message using her private key.

**Signatures:** Signer encrypts a message with her own private key.
  Verifier checks the message using the signer's public key.

How can we use asymmetric cryptosystems to _prove_ ownership?
<div class="gap"></div>

How can we use asymmetric cryptosystems to _transfer_ ownership?
<div class="gap"></div>

Where is asymmetric cryptography used in your wallet?
<div class="gap"></div>

Assuming we have a strong asymmetric cryptosystem, what hard problems are left
to solve to make a cryptocurrency?

<div class="gap"></div>
