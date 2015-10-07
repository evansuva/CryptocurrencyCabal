---
Title: 'Class 12: Script'
Date: '2015-10-07'
nocomment: false
print: "class12.pdf"
weight: 3
menu: "classes"

---

## Schedule

   <div class="todo"> 
**Read**: [_Chapter 3: Mechanics of Bitcoin_](http://bitcoin-class.org/docs/princeton-book/chapter_3.pdf), from
Arvind Narayanan, Joseph Bonneau, Edward Felten, Andrew Miller, Steven
Goldfeder. [_Bitcoin and Cryptocurrency
Technologies_](https://piazza.com/princeton/spring2015/btctech/resources).
Sections 3.2 and 3.3 are about bitcoin scripts, and should be read
carefully.  (You should read the whole chapter, but those sections are
most relevant to today's class.)

**Friday, October 9**: Problem Set 2 (8:29pm).

**Monday, October 19**: Midterm Exam</div>

<center>
<iframe src="//www.slideshare.net/slideshow/embed_code/key/aQYiH6zGCIgxbO" width="625" height="400" frameborder="2" marginwidth="0" marginheight="0" scrolling="no"> </iframe> 

   <div class="caption">
Sorry, ink markings were lost from today.  [Download the slides](/classes/class12-post.pptx)
   </div>

# Hash Collisions!

**Computing a bitcoin address:** ([bitcoinwiki](https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses))

<!--more-->

Private Key: pick a random number, _k_.  
Public Key: compute (_Ux_, _Uy_) = _Gk_
<div class="indented">(elliptic curve multiplication, _G_ is specified generator point)</div>
_Ux_ and _Uy_ are 32 bytes each.

The bitcoin address is (|| is bitstring concatenation):  
    <div class="indented">
    raw = `1` || RIPEMD160(SHA256(_Ux_ || _Uy_))  
         <div class="indented">_RIPEMD output is 160 bits (20 bytes) + one byte for `1`_</div>
    checksum = first 4 bytes of SHA256(SHA256(raw))  <div class="indented">_Compute a checksum using SHA256 double hash_</div>
    address = Base58Check(raw || checksum)  <div class="indented">_convert to printable, unambiguous characters_</div>
    </div>

How important is pre-image resistance for the security of bitcoin addresses?
<div class="gap"></div>

How important is collision resistance for the security of bitcoin addresses?
<div class="gap"></div>

How important is pre-image resistance for the integrity of bitcoin's proof-of-work?
<div class="gap"></div>

How important is collision resistance for the integrity of bitcoin's proof-of-work?
<div class="gap"></div>

Xiaoyun Wang and Hongbo Yu, [_How to Break MD5 and Other Hash Functions_](http://www.bitcoin-class.org/docs/hashcollisions.pdf), EuroCrypt 2005.

# Bitcoin Script

Transaction outputs in bitcoin are protected by _locking scripts_, and
must be unlocked by _unlocking scripts_.  The scripts are written in a
simple (compared to, say, the Java Virtual Machine language, but quite
complex and poorly specified for what one might expect would be needed
for bitcoin transactions) stack-based language.  A transaction output is
not unlocked unless an unlocking script is provided such that the result
of executing the unlocking script, followed by executing the locking
script, is a stack with value True on top (and no invalid transaction
results during the execution).

Some script instructions:


   <table cellpadding=5>
   <tr>
   <td align="center"><b>Opcode</b></td><td align="center"><b>Input</b></td><td align="center"><b>Output</b></td><td align="center"><b>Description</b></td></tr>
   <tr><td>`OP_1`</td><td> - </td><td> **1** </td><td> Pushes a **1** (True) on the stack </td></tr>
   <tr><td>`OP_DUP`</td><td> _a_ </td><td> _a_ _a_ </td><td> Duplicates the top element of the stack </td></tr>
   <tr><td>`OP_ADD`</td><td> _a_ _b_ </td><td> (_a_+_b_) </td><td> Pushes the sum of the top two elements. </td></tr>
   <tr><td>`OP_EQUAL`</td><td> _a_ _b_ </td><td> **0** or **1** </td><td> Pushes **1** if the top two elements are exactly equal, otherwise **0**. </td></tr>
   <tr><td>`OP_VERIFY` </td><td> _a_ </td><td> - </td><td> If _a_ is not **True** (**1**), terminates as Invalid. </td></tr>
   <tr><td>`OP_RETURN`</td><td> - </td><td> - </td><td> Terminates as Invalid. </td></tr>
   <tr><td>`OP_EQUALVERIFY` </td><td> _a_ _b_ </td><td> - </td><td> If _a_ and _b_ are not equal, terminates as Invalid. </td></tr>
   <tr><td>`OP_HASH160` </td><td> _a_ </td><td> H(_a_) </td><td> Pushes bitcoin address, RIPEMD160(SHA256(_a_)). </td></tr>
</table>

Some more complex instructions:

`OP_IF` _statements_ `OP_ENDIF` - If the top of the stack is **1**, executes _statements_.  Otherwise does nothing.

`OP_CHECKSIG` - Pops two items from the stack, _publickey_ and _sig_.
Verifies the entire transaction (known from node state, not the stack)
using the _publickey_ and _sig_.  If the signature is valid, push **1**;
otherwise, **0**.

`OP_1 OP_DUP OP_ADD OP_DUP OP_SUB OP_VERIFY`
<div class="gap">

</div>

The most common locking script (send to public address):  
`OP_DUP`  
`OP_HASH160`  
`OP_DATA20` _(bitcoin address)_  
`OP_EQUALVERIFY`  
`OP_CHECKSIG`  

What must be on the stack for the locking script to succeed (end with **1** on top of stack)?
<div class="gap">

</div>

`OP_HASH160` _20-byte hash_ `OP_EQUAL`  

What must be on the stack for the locking script above ("Pay-to-Script-Hash") to succeed?
<div class="gap">

</div>


According to [Most Popular Transaction
Scripts](http://www.quantabytes.com/articles/a-survey-of-bitcoin-transaction-types)
(analysis of all transactions in first 290,000 blocks), the ninth most popular script is:
`OP_RETURN OP_DATA_40`

What must be on the stack for the `OP_RETURN OP_DATA_40` locking script to succeed (end with
**1** on top of stack)? (Trick question: what happens to the coin
protected by this locking script?)
<div class="gap">

</div>

Is the bitcoin scripting language Turing-complete?
<div class="gap">

</div>

If you are not clear on what Turing-complete means, see [_Dori-Mic and the Universal Machine!_](http://www.dori-mic.org)

## BTCD Code

Type: [Script](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/script.go#L206) is the virtual machine the executes scripts (note that it has two Stacks) 

Execute a script: [Execute](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/script.go#L723)  
Execute one instruction: [Step](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/script.go#L782)  

[Opcodes](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L38)  
[exec](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L971) function executes one instruction  

Some interesting opcode implementations:
[OP_IF](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L1143)  
[OP_RETURN](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L1235)

## Bitcoin Core Code

[script/interpreter.cpp](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp)

[OP_DUP](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp#L524)  
[Crypto](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp#L752)
[OP_CHECKSIG](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp#L785)

## Links

[Script Playground](http://www.crmarsh.com/script-playground/)

Some interesting things you can do with bitcoin scripts:  
[Contracts](https://en.bitcoin.it/wiki/Contracts) (see also Nick Szabo's [_Formalizing and Securing Relationships on Public Networks_](http://szabo.best.vwh.net/formalize.html)  
[Secure Multiparty Computations](http://eprint.iacr.org/2013/784) (to implement lotteries)  

The `OP_RETURN`/pasted script execution bug doesn't even make this list of [_The 9 Biggest Screwups in Bitcoin History_](http://www.coindesk.com/9-biggest-screwups-bitcoin-history/).

[Block 71036](http://blockexplorer.com/rawblock/00000000000997f9fd2fe1ee376293ef8c42ad09193a5d2086dddf8e5c426b56)
