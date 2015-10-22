% Class 12: Script
% '2015-10-07'

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

# Hash Collisions!

**Computing a bitcoin address:** ([bitcoinwiki](https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses))

<!--more-->

**Private Key:** pick a random number, _k_.  
**Public Key:** compute (_Ux_, _Uy_) = _Gk_  (elliptic curve multiplication, _G_ is specified generator point)  
_Ux_ and _Uy_ are 32 bytes each.

The bitcoin address is (|| is bitstring concatenation):  
    raw = `1` || RIPEMD160(SHA256(_Ux_ || _Uy_))  _RIPEMD output is 160 bits (20 bytes) + one byte for `1`_  
    checksum = first 4 bytes of SHA256(SHA256(raw))  _Compute a checksum using SHA256 double hash_  
    address = Base58Check(raw || checksum)           _convert to printable, unambiguous characters_  


How important is pre-image resistance for the security of bitcoin addresses?

#


How important is collision resistance for the security of bitcoin addresses?

#


How important is pre-image resistance for the integrity of bitcoin's proof-of-work?  What about collision resistance?

#


Xiaoyun Wang and Hongbo Yu, [_How to Break MD5 and Other Hash Functions_](http://www.bitcoin-class.org/docs/hashcollisions.pdf), EuroCrypt 2005.

<!--page-->

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

|Opcode|Input|Output|Description|
|--|--|--|--|
|`OP_1`| - | **1** | Pushes a **1** (True) on the stack |
|`OP_DUP`| _a_ | _a_ _a_ | Duplicates the top element of the stack |
|`OP_ADD`| _a_ _b_ | (_a_+_b_) | Pushes the sum of the top two elements. |
|`OP_EQUAL`| _a_ _b_ | **0** or **1** | Pushes **1** if the top two elements are exactly equal, otherwise **0**. |
|`OP_VERIFY` | _a_ | - | If _a_ is not **True** (**1**), terminates as Invalid. |
|`OP_RETURN`| - | - | Terminates as Invalid. |
|`OP_EQUALVERIFY` | _a_ _b_ | - | If _a_ and _b_ are not equal, terminates as Invalid. |
|`OP_HASH160` | _a_ | H(_a_) | Pushes bitcoin address, RIPEMD160(SHA256(_a_)). |

Some more complex instructions:

`OP_IF` _statements_ `OP_ENDIF` - If the top of the stack is **1**, executes _statements_.  Otherwise does nothing.

`OP_CHECKSIG` - Pops two items from the stack, _publickey_ and _sig_.
Verifies the entire transaction (known from node state, not the stack)
using the _publickey_ and _sig_.  If the signature is valid, push **1**;
otherwise, **0**.

The most common locking script (send to public address):  
`OP_DUP` `OP_HASH160` `OP_DATA20` _(bitcoin address)_  `OP_EQUALVERIFY` `OP_CHECKSIG`  

What must be on the stack to unlock this locking script (end with **1** on top of stack)?

#

What must be on the stack to unlock `OP_HASH160` _20-byte hash_ `OP_EQUAL` ("Pay-to-Script-Hash")?

#

Is the bitcoin scripting language Turing-complete?  What would be risky about a crytocurrency scripting language being Turing-complete?
<div class="gap">
</div>

_See web notes for links to BTCD and Bitcoin Core Code for interpreting script._
