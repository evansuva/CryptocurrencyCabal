---
title: "Checkup 1 Comments"
weight: 3
date: '2015-09-20'
---

Here are answers and discussions for the Checkup 1 questions.  It would
not be at all surprising if there are one or two questions on Checkup 2
(scheduled for Wednesday, Sept 23) that are closely related to questions
on Checkup 1.

#### 1. What are the most essential properties something must have to be used as a currency?

   <div class="answer"> 
The two most essential properties are
_transferability_ and _scarcity_.  Recall that we defined a currency as
a "medium of exchange".  In order for something to be used in exchanges,
it must be possible to transfer it from one owner to another.  The
second essential property of scarcity is intrinsic to having value.  If
something can be created or duplicated freely, it cannot be used to
represent value.  
<!--more-->
Being _universally recognized as valuable_ would be a
good alternative answer to _scarcity_, but is a bit circular: nothing is
"universally" recognized as valuable for the entire universe; it is
universally recognized as valuable by all partcipants in an economy
using that currency, but this is a somewhat circular definition (of
universal).

Other properties that were often mentioned, but are not essential for
something to be a currency include _stable_ (it is good in most cases is
a currency holds it value over time, but note that prefect stability is
generally not considered a good thing, and inflation rates should not be
zero for a well-functioning economy); _anonymous_ (some people consider
this desirable, but others view it as a negative; it is certainly not
essential for an effective currency); _decentralized_ (this is what
distinguished bitcoin from other currencies, but essentially all other
currencies used in the history of humanity have been centralized); and
_divisible_ (important for many transactions).
   </div>

#### 2. What are the drawbacks of using a centralized bank to record transactions?

   <div class="answer">
The biggest problem is that it requires placing a large amount of trust
the bank.  The bank can decide what transactions happened, and can
decide who owns what (including, deciding that the bank owns everything).

Another problem with relying on a single, centralized entity is that
even if that entity is not malicious, it could still be "incompetent" -
if the bank is off-line, it is impossible to execute a transaction
(i.e., this is what happens when you try and use a credit card and the
merchant can't connect to the system to do the transaction), and if the
bank loses the transaction record everything is lost.
   </div>


#### 3. Where is asymmetric cryptography used in a bitcoin wallet?

   <div class="answer"> 
The most important place is to sign transactions.  It is necessary to
use the private key associated with public address A to sign a
transaction that transfers a coin owned by A to another address.

The wallet is also using asymmetric cryptography to verify a transaction
(checking that a transaction that trasfers coin to your wallet is signed
by the appropriate private key using the public key corresponding to the
bitcoin address of the sender).
   </div>


#### 4. Find <span class="math">_x_</span> such that <span class="math">2<sup>_x_</sup> mod 7 = 1</span>.

   <div class="answer">
The simplest answer is <span class="math">_x_ = 3</span> since <span class="math">2<sup>3</sup> = 8 mod 7 = 1</span>.   

Any value <span class="math">_x_ mod 6 = 3</span> satisfies the equation.  For example, <span class="math">_x_ = 9</span> also works since <span class="math">2<sup>9</sup> = 512 mod 7 = 1</span>.   
   </div>

#### 5. The problem in the previous question is an instance of the _discrete logarithm problem_.  Why do cryptographers consider discrete logarithm to be a _hard_ problem?

   <div class="answer"> 
Computer scientists consider a problem to be
_hard_ if there is no known algorithm that solves the problem with
resources that scale as a polynomial in the size of the problem
instance.  Discrete log is (currently) considered a hard problem because
the best known solutions are asymptotically not much better than trying
all possibilities until you find an input that works (that is, doing a
brute force search).  For a large enough input, the time expect to solve
a discrete problem can be made prohibative (given assumptions based on
what is known today).
   </div>

#### 6. Alice owns coin <em>X</em> and has public/private key pair (<em>KU<sub>A</sub></em>, <em>KR<sub>A</sub></em>); Bob has public/private key pair (<em>KU<sub>B</sub></em>, <em>KR<sub>B</sub></em>) for the strong asymmetric cryptosystem _E_ (the notation <em>E<sub>K</sub>(m)</em> denotes the encryption of input <em>m</em> with key <em>K</em>).  What message should Alice send to the public ledger to transfer <em>X</em> to Bob?
   
   <div class="answer"> 
The message needs to be something (1) only Alice could generate and (2)
everyone can verify as transferring the coin to Bob.

_msg_ = "I, KU<sub>A</sub> (Alice), transfer coin <em>X</em> to KU<sub>B</sub> (Bob)"

Send to ledger the message signed by Alice: (_msg_, E<sub>KR<sub>A</sub></sub>(_msg_))
   </div>
