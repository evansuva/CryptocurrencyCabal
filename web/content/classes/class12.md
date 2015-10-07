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

Is the bitcoin scripting language Turing-complete?
<div class="gap">

</div>

If you are not clear on what Turing-complete means (and why real
machines are not ideal), see [_Dori-Mic and the Universal
Machine!_](http://www.dori-mic.org)

