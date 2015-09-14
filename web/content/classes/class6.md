---
title: "Class 6: Hash Functions"
date: '2015-09-14'
nocomment: false
weight: 3
menu: "classes"
---

## Schedule

   <div class="todo">
**Tuesday, September 15** (8:29pm, tomorrow): [Problem Set 1](http://www.bitcoin-class.org/ps/ps1) due.

**Wednesday, September 23**: Check 2 (was originally scheduled for Monday, September 21)

**Readings for next week** (should be completed by Monday, September 21 at the latest, but earlier is better):

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

<!--more-->
<center>
<iframe src="//www.slideshare.net/slideshow/embed_code/key/JWtz4G4wUYKWU5"
width="625" height="400" frameborder="2" marginwidth="0"
marginheight="0" scrolling="no"></iframe>

<!--
   <div class="caption">
Note: ink markings may not appear in the
embedded viewer.  To see them, [download the slides](/classes/class5-post.pptx).
   </div>
-->

</center>

## Notes
Why do we typically hash a message before signing it? What's wrong if we always signed the full message?
<div class="gap"></div>

What are the properties we want in a cryptographically secure hash function?
<div class="gap"></div>

What is the "birthday attack" in the context of a hash function?
<div class="gap"></div>

<!--page-->
Say you have 3000 distinct files in the "Documents" folder of our laptop. If
you have SHA-256 hashes for each of them, do you expect any repeats? If we
truncated the hash output to just 20 bits, how many repeats do you expect to
see?
<div class="gap"></div>

What is the advantage of using a Merkle tree as opposed to directly hashing
the full string?
<div class="gap"></div>

