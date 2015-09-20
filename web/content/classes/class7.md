---
title: "Class 7: The Blockchain"
date: '2015-09-16'
print: "class7.pdf"
nocomment: false
weight: 3
menu: "classes"
---

## Schedule

   <div class="todo">
**Wednesday, September 23**: Checkup 2 (was originally scheduled for Monday, September 21)

**Readings** (should be completed by Monday, September 21 at the latest; covered by Checkup 2):

- Satoshi Nakamoto, [_Bitcoin: A Peer-to-Peer Electronic Cash
System_](https://bitcoin.org/bitcoin.pdf), 2008.  The is the original
bitcoin paper, which is quite readable and historically interesting.

- [_Chapter 6: The Bitcoin
Network_](https://github.com/aantonop/bitcoinbook/blob/develop/ch06.asciidoc)
and [_Chapter 7: The
Blockchain_](https://github.com/aantonop/bitcoinbook/blob/develop/ch07.asciidoc)
from Andreas Antonopoulos' book.  

- [_Chapter 2: How Bitcoin Achieves
Decentralization_](http://bitcoin-class.org/docs/princeton-book/chapter_2.pdf)
and [_Chapter 5: Bitcoin
Mining_](http://bitcoin-class.org/docs/princeton-book/chapter_5.pdf)
from Arvind Narayanan, Joseph Bonneau, Edward Felten, Andrew Miller,
Steven Goldfeder. [_Bitcoin and Cryptocurrency
Technologies_](https://piazza.com/princeton/spring2015/btctech/resources).
   </div>

# Blockchain in the News

[_Blockchain initiative backed by 9 large investment banks_](http://www.ft.com/cms/s/0/f358ed6c-5ae0-11e5-9846-de406ccb37f2.html), Financial Times, 15 Sept 2015.

[_Nine of the World’s Biggest Banks Form Blockchain Partnership_](http://recode.net/2015/09/15/nine-of-the-worlds-biggest-banks-form-blockchain-partnership/), Re/Code, 15 Sept 2015.

[_Bitcoin Is Only The Beginning For Blockchain Technology_](http://www.forbes.com/sites/mikemontgomery/2015/09/15/bitcoin-is-only-the-beginning-for-blockchain-technology/), Forbes, 15 Sept 2015.

[_Bitcoin's Shared Ledger Technology: Money's New Operating System_](http://www.forbes.com/sites/laurashin/2015/09/09/bitcoins-shared-ledger-technology-moneys-new-operating-system/), Forbes, 9 Sept 2015.

# Trust

What are valid sources of _trust_?
<div class="gap"></div>

<!--more-->

What are potentially misleading sources of _trust_?
<div class="gap"></div>

What mechanisms have humans evolved or constructed to enhance trust among strangers?
<div class="gap"></div>

# Distributed Consensus

How well does the 2-out-of-3 network consensus public ledger protocol work?
<div class="gap"></div>


# Proof-of-Work

Cynthia Dwork and Moni Naor.  [_Pricing via Processing or Combatting Junk Mail_](http://bitcoin-class.org/0/classes/class6/pvp.pdf), CRYPTO 1992.

**Pricing Function**: (<span class="math">_f_</span>)
- moderately easy to compute
- cannot be amortized 
- computing <span class="math">_f_(_m_<sub>1</sub>), ..., _f_(_m_<sub>l</sub>)</span> costs <span class="math">_l_</span> times as much as computing <span class="math">_f_(_m_<sub>i<sub>)</span>. 
- easily verified: given <span class="math">_x_</span>, <span class="math">_y_</span> easy to check <span class="math">_y_ = _f_(_x_)</span>.

Adam Back. [_Hash Cash Postage Implementation_](http://www.hashcash.org/papers/announce.txt)

**Interactive Hashcash**:  
1. Sender to Receiver: `Hello`  
2. Receiver to Sender: <span class="math">_r_</span> (random nonce)  
3. Sender to Receiver: <span class="math">_x_</span>, `Mail` where <span class="math">_x_ = _f_(_r_)</span>.  
4. Receiver verifies <span class="math">_x_ = _f_(_r_)</span>.  

How well does this protocol work for sending mail?
<div class="gap"></div>

How can we make this protocol non-interactive?
<div class="gap"></div>

# Bitcoin Mining

Proof-of-work: Find an <span class="math">_x_</span> such that: SHA-256(SHA-256(<span class="math">_r_</span> + <span class="math">_x_</span>)) < <span class="math">_T_/_d_</span>.

<span class="math">_d_</span> is the "difficulty" (varies).  
<span class="math">_T_</span> is a fixed target (256-bit number).  
<span class="math">_r_</span> depends on hash of previous block, transactions, and other information.

What does it mean for the bitcoin difficulty to go down?
<div class="gap"></div>



