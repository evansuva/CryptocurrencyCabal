---
Title: 'Class 15: Scripting Transactions'
Date: '2015-10-19'
nocomment: false
print: "class15.pdf"
weight: 3
menu: "classes"
---
  

## Midterm Exam

The midterm exam will be in class, **Wednesday, 21
  October**.  The midterm will be closed-resources, but is meant to test
  understanding, not memorization.  

  The midterm covers: Classes 1-15, Checkups 1-2, Problem Sets 1-2, and
  the assigned readings: Chapters 1, 2, 3, and 5 from [_Bitcoin and
  Cryptocurrency
  Technologies_](https://piazza.com/princeton/spring2015/btctech/resources);
  Chapters 1, 2, 3, 4, 6, and 7 from [_Mastering Bitcoin: Unlocking
  Digital Cryptocurrencies_](https://github.com/aantonop/bitcoinbook);
  Satoshi Nakamoto, [_Bitcoin: A Peer-to-Peer Electronic Cash
  System_](https://bitcoin.org/bitcoin.pdf); Ittay Eyal and Emin
  G&uuml;n Sirer, [_Majority is not Enough: Bitcoin Mining is
  Vulnerable_](http://arxiv.org/pdf/1311.0243v5.pdf).  

  The first two questions on the midterm will ask you to comment on the
  technical validity of some statements in the [Congressional Research
  Service report](https://www.fas.org/sgp/crs/misc/R43339.pdf) (in the
  first 8 pages, not the policy issues).  

**Extra Office Hours.** There will be some additional office hours this week:

- Monday, 5-6:30pm (Ori, Rice 442)
- Tuesday, 2-3:30pm (Dave, Rice 507) **(Added)**
- Tuesday, 3:30-4:30pm (Samee, Rice 442) **(Added)**
- Wednesday, 3:30-4:30pm (Samee, Rice 442)
- Thursday, 2:30-3:30 (Dave, Rice 507)


## Scripting Transactions

<center>
<iframe src="//www.slideshare.net/slideshow/embed_code/key/cYC6B2BCdrtoy6" width="625" height="400" frameborder="2" marginwidth="0" marginheight="0" scrolling="no"> </iframe> 

   <div class="caption">
[Download the slides](/classes/class15-post.pptx)
   </div>
</center>

Recall from class 12: Transaction outputs in bitcoin are protected by
_locking scripts_, and must be unlocked by _unlocking scripts_.  A
transaction output is not unlocked unless an unlocking script is
provided such that the result of executing the unlocking script,
followed by executing the locking script, is a stack with value True on
top.

<!--more-->

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
<div class="gap"></div>

<!-- page-->

`OP_HASH160`  
_20-byte hash_  
`OP_EQUAL`  

What must be on the stack for the locking script above ("Pay-to-Script-Hash") to succeed?
<div class="gap"></div>


According to [Most Popular Transaction
Scripts](http://www.quantabytes.com/articles/a-survey-of-bitcoin-transaction-types)
(analysis of all transactions in first 290,000 blocks), the ninth most popular script is:  

`OP_RETURN`  
`OP_DATA_40`  

What must be on the stack for the `OP_RETURN OP_DATA_40` locking script to succeed (end with
**1** on top of stack)? (Trick question: what happens to the coin
protected by this locking script?)

<div class="gap"></div>

## BTCD Code

Type: [Script](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/script.go#L206) is the virtual machine the executes scripts (note that it has two Stacks) 

Execute a script: [Execute](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/script.go#L723)  
Execute one instruction: [Step](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/script.go#L782)  

[Opcodes](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L38): [exec](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L971) function executes one instruction  

Some interesting opcode implementations: [OP_IF](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L1143), [OP_RETURN](https://github.com/btcsuite/btcd/blob/c153596542b3d87dd774c29aa5be5117ac01a234/txscript/opcode.go#L1235)

## Bitcoin Core Code

[script/interpreter.cpp](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp), [OP_DUP](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp#L524), [Crypto](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp#L752), [OP_CHECKSIG](https://github.com/bitcoin/bitcoin/blob/41e6e4caba9899ce7c165b0784461c55c867ee24/src/script/interpreter.cpp#L785)

## Links

[Script Playground](http://www.crmarsh.com/script-playground/)

Some interesting things you can do with bitcoin scripts:  
[Contracts](https://en.bitcoin.it/wiki/Contracts) (see also Nick Szabo's [_Formalizing and Securing Relationships on Public Networks_](http://szabo.best.vwh.net/formalize.html)  
[Secure Multiparty Computations](http://eprint.iacr.org/2013/784) (to implement lotteries)  

The `OP_RETURN`/pasted script execution bug doesn't even make this list of [_The 9 Biggest Screwups in Bitcoin History_](http://www.coindesk.com/9-biggest-screwups-bitcoin-history/).

<!-- [Block 71036](http://blockexplorer.com/rawblock/00000000000997f9fd2fe1ee376293ef8c42ad09193a5d2086dddf8e5c426b56)-->
