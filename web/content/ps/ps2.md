---
title: "PS2: The Blockchain"
date: '2015-09-20'
nocomment: false
menu: "assignments" 
---

# Problem Set 2:<br> The Blockchain

   <div class="due">
Due: Friday, 9 October at 8:29pm
   </div>

## Purpose

The goal of this assignment is for everyone in the class to understand
how bitcoin mining and the blockchain work, and evaluate risks to the blockchain.

<!--more-->

### Collaboration Policy

For this assignment, everyone should submit their own assignment and
should writeup their own answers to the questions.  You may, and are
encouraged to, discuss all of the problems with anyone else you want
(both on-line using the course web site or any other means you choose,
and in person).  

### Submission

Submit your answers as a single PDF file using [this
link](https://www.dropbox.com/request/oar0AQeg4SJRNtdHbcSS).  The name
of your file should be `<your email ID>-ps2.pdf`.

Your submission should include clearly marked answers for all the
problems (highlighted in yellow).  None of the questions require
submitting a program, although you may find it helpful to write programs
to develop your answers.  If the code you write is less than a page, it
is best to just include it in the PDF writeup.  If it is longer, you may
submit separate code files (and mention them in the PDF submission).

# Blockchain Consensus

These questions concern the original bitcoin paper:

- Satoshi Nakamoto, [_Bitcoin: A Peer-to-Peer Electronic Cash
System_](https://bitcoin.org/bitcoin.pdf), 2008. 

   <div class="problem">
**Problem 1.** In Section 6, Satoshi writes:
_"The incentive may help encourage nodes to stay honest. If a greedy
attacker is able to assemble more CPU power than all the honest nodes,
he would have to choose between using it to defraud people by stealing
back his payments, or using it to generate new coins. He ought to find
it more profitable to play by the rules, such rules that favour him with
more new coins than everyone else combined, than to undermine the system
and the validity of his own wealth."_  

What are the assumptions necessary to support Satoshi's claim that it is
more profitable for a greedy attacker with a majority of the mining
power "to play by the rules"? (other than the assumption that the greedy
attacker is a "he")</div>

   <div class="problem"> **Problem 2.** 
    At the end of Section 11, Satoshi presents a table for <span class="math">_p_</span> < 0.001,
    where it is listed "q=0.45 z=340". What does this mean in plain
    English, expressed in a short sentence? 
    </div>

   <div class="problem">**Problem 3.** For that same table, what would
the <span class="math">_z_</span> values be for <span class="math">_p_</span> < 0.05 (instead of _p_ < 0.001)?
   </div>


# Merkle Trees and Storage

Bitcoin uses Merkle Trees to record transactions in a way that enables a
single hash to be used to record a set of transactions, and a small
(logarithmic in the number of transactions) number of hashes to be
sufficient to verify a transaction.  In this question, you'll explore
another use of Merkle Trees to provide verifiable cloud storage.

Imagine you are storing a database using a cloud storage provider, since
you do not have local storage. Assume that you have just enough local
storage to store a small number of cryptographic hashes or keys. You
fear, however, that your storage provider may be skimping, and may not
actually be storing all of your data.

So, you want to verify each record as you receive it, to make sure the
data is the same as what you wrote.  In order to do so, you are willing
to write extra cryptographic information with each write, but need to
limit the amount of local storage.

Assume we are storing <span class="math">_n_</span> records, and we can
query the server for the <span class="math">_i_</span>-th record for any <span class="math">_i_</span> from 0 to <span class="math">_n_ - 1</span>. The
data content of the <span class="math">_i_</span>-th record will be denoted as <span
class="math">data[_i_]</span>.

   <div class="problem"> **Problem 4.** One naive approach would be to
just sign each record data with a private key, and verify the signature
on reading it. 

   (a) There is a problem with this scheme unless the bytes of <span class="math">data[_i_]</span> indicate that it is in fact from <span class="math">_i_</span>-th index. What is this problem?

   (b) Suppose we perform a write at position <span class="math">_i_</span>, both data and signature. Later on, when we read it back, if the signature matches the data, can we be sure that it is indeed the data item we wrote?
  Explain.
</div>

   <div class="problem">**Problem 5.** Another approach would be hashing
  the concatenation of all records in the database. This hash is a small
  item that can be stored locally. 

  (a) What is the write/read/verify procedure for this system?  

  (b) How does the cost of reading and writing to the database scale with <span class="math">_n_</span> (the number of records)?
  </div>

  <div class="problem">**Problem 6.** Instead of using the concatenating all the records linear, they were organized into a Merkle tree.

   (a) What is the write/read/verify procedure for this system?  

  (b) How does the cost of reading and writing to the database scale with _n_ (the number of records)?
    </div>

# Blockchain 

These questions are related to this paper (and what was covered in [Class 10](/classes/class10) and [Class 11](/classes/class11)):

- Ittay Eyal and Emin G&uuml;n Sirer, [_Majority is not Enough: Bitcoin
  Mining is Vulnerable_](http://arxiv.org/pdf/1311.0243v5.pdf), November
  2013.

[Orphaned Blocks](https://blockchain.info/orphaned-blocks) are blocks
that are submitted to the bitcoin network, and that are valid, but do
not become part of the longest (consensus) blockchain.

One way to detect selfish mining is to be on the lookout for "orphaned"
blocks, or blocks that were mined but never became part of the final
blockchain. 

Let's say you set up a node to monitor orphaned blocks, and you are in a
favorable position in the network that allows you to observe all
orphaned blocks.  Assume that the hashing difficulty stays constant and
the expected block rate is constant at 10 minutes per block.

   <div class="problem"> **Problem 7.**  
  (a) If a mining pool has 15% (<span class="math">&alpha;</span> = 0.15) of the
  total network hashing power, how many blocks is it expected to find in
  a day? 

  (b) Obtain a general formula for expected number of blocks a mining pool with <span class="math">&alpha;</span> fraction of the total hashing power should find in <span class="math">_t_</span> minutes.
  </div>

For the next questions, you may assume a very simplified network model
where you can view the network having two "supernodes": one represents a
particular mining pool, and the other represents all other nodes in the
network.  The latency between the two supernodes is <span class="math">_L_</span> seconds.

With this simple (but very unrealistic) model, we assume that when one
supernode announces a block, all of the rest of the network learns of
the new block <span class="math">_L_</span> seconds later (but no one
learns of it before that).

  <div class="problem"> **Problem 8.** Assuming all of the miners are
  honest, what is the expected number of orphaned blocks per day for an honest
  mining block with hashing power <span class="math">&alpha;</span> and latency <span class="math">_L_</span> (as simplified
  above).
  </div>

  <div class="problem"> **Problem 9.** How does this change if the
  mining pool is mining selfishly?  (For this question, assume that the
  selfish mining pool learns of a block announced by the rest of the
  network as soon as it is announced, so will immediately announce any
  withheld blocks at that time.  That is, you may still assume the
  simplified latency <span class="math">_L_</span> model, but that the selfish mining pool has a
  spy in the other supernode with a low-latency direct connection to the
  mining pool.)
  </div>

# Pool Hopping

Answer _either_ question 10A or 10B (your choice).

   <div class="problem"> **Problem 10A.** In class, we saw how pool-hopping
can be used to game a "proportional" reward scheme. Design a simple
reward scheme that eliminates pool-hopping incentive. In particular,
derive the expected reward for your scheme and show that it does not
depend on time since last block (or any variable controlled by the
miner).  </div>

   <div class="problem"> **Problem 10B.** In upcoming classes, we will
   have visits from a law professor (who also works for the State
   Department on international cyberlaw and promulgating the open
   Internet) and an FBI agent who works on criminals using bitcoin for
   ransom.  Come up with at least one good question that you would like
   one of them to answer.
   </div>


# Bonus Opportunities

The following questions are suggestions for further work, but it is not
expected that everyone will solve them.  (Note that some of these could
be starting ideas for larger projects, although a good answer to any of
them would be most impressive.)

**Challenge 1.** The network model in questions 8 and 9 is very
  unrealistic.  Answer these questions for a more realistic model of the
  bitcoin network, and compare your results with the actual rates of
  orphaned blocks.  To better detect misbehaving miners, you would want
  to also look into the contents of the dual blocks (note that the one
  that was orphaned is not necessarily the one from the selfish miner).
  There is an API for obtaining orphaned blocks provided by
  [https://api.biteasy.com](https://api.biteasy.com/blockchain/v1/blocks?type=ORPHANED).

**Challenge 2.** The analysis in the paper and in class assumed that if
   the selfish miner is ahead by 2 blocks, it will always win by
   releasing the blocks when it observes a new block on the blockchain.
   It could still lose if the rest of the network finds a second block
   before the selfish miner's blocks have propagated over the network.
   How does this possibility impact the analysis?

**Challenge 3.** The site that was demonstrated in Class 11 uses a very
  simple model to estimate the profitability of a particular bitcoin
  miner.  Produce a better model, and use it to evaluated the expected
  profit (or loss) for mining hardware such as [Antminer
  S7](https://bitmaintech.com/product.htm).  A better model would need
  to capture expected increases in the difficulty, cost of capital, and
  other costs of mining besides electricity.



### Submission

Follow the submission instructions at the beginning of this page by
8:29pm on Friday, 9 October.

