---
title: "Checkup 2 Revisions"
weight: 3
date: '2015-09-28'
---

If you didn't get a "gold star" on [Checkup 2](./checkup2.pdf), you can
improve your score by answering one or more of the questions below.  You
can submit your revisions in class Wednesday, or at the end of office
hours Wednesday.  To submit revisions, attach your answers to these
questions to your original submission.

**R1. Hashing** (Answer this question if you want to improve your score on questions 1-4)

There are several different places in bitcoin where cryptographic hashes are used:

A. Producing the public bitcoin address by hashing the public key.

B. Producing a transaction digest for use as the input in signing a transaction.

C. Producing the Merkle tree root for authenticating the transactions in
a block (using hashes all the way up the tree).

D. Producing the hash of the previous block to use in the block header.

E. Producing the double hash of the block (with nonces) to find a block
that satisfies the difficult needed in mining.

Suppose _H_ is a hash function that provides strong pre-image
resistance, but does not provide collision resistance. That is, an
adversary knows way to efficiently find pairs of values, _m_<sub>1</sub>
and _m_<sub>2</sub> such that _m_<sub>1</sub> is not equal to
_m_<sub>2</sub> and _H_(_m_<sub>1</sub>) = _H_(_m_<sub>2</sub>).

Explain clearly _one place_ where cryptographic hashing is used in
bitcoin that the adversary could gain an advantage by knowing how to
find collisions in _H_.  (Be careful to not overstate the advantage,
though.)

**R2. Blockchain** (Answer this question if you want to improve your score on questions 5-7)

Ishtosa suggests strengthening bitcoin by adding the hash of the
previous previous block to the block header.  So, the block header for
block 534 would now include the both the hash of the block 533, and
the hash of block 532.  Would this be a good idea?

<!--more-->
