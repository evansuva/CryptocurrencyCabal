---
title: "Class 08: Mining"
date: '2015-09-21'
print: "class8.pdf"
nocomment: false
weight: 3
menu: "classes"
---

## Schedule

   <div class="todo">
**Wednesday, September 23**: Checkup 2 (was originally scheduled for Monday, September 21).  See previous class notes for details on what is covered.  
**Friday, October 9**: Problem Set 2 (moved from original deadline of October 2).  Problem Set 2 will be posted later this week.  
**Monday, October 19**: Midterm Exam
   </div>

## Reminders

You can [subscribe to the course
calendar](https://www.google.com/calendar/ical/rmjagdrnmu3a9h2q5199lg4t28%40group.calendar.google.com/public/basic.ics).
This has updated information on deadlines and office hours.

If you want to receive course website updates by email, you can
subscribe to the RSS feed using a RSS reader or an emailing service like
[feedmyinbox.com](https://www.feedmyinbox.com/) or
[Blogtrottr](http://blogtrottr.com/).  The feed address is
[http://bitcoin-class.org/index.xml](http://bitcoin-class.org/index.xml).

<!--more-->

We generally will avoid sending out emails to the class, and will assume
you are observing the website closely.  

<center>
<iframe src="//www.slideshare.net/slideshow/embed_code/key/ir3gwqoaGyncK4"
width="625" height="400" frameborder="2" marginwidth="0"
marginheight="0" scrolling="no"></iframe>

   <div class="caption">
Note: ink markings may not appear in the
embedded viewer.  To see them, [download the slides](/classes/class8-post.pptx).
   </div>

</center>

# Exploring Blocks

<table>
<tr><td><b>Label</b></td><td><b>Bytes</b></td><td><b>Description</b></td></tr>
<tr><td>version</td><td>4</td><td>Block version information</td></tr>
<tr><td>prev_block</td><td>32</td><td>Hash of the previous block</td></tr>
<tr><td>merkle_root</td><td>32</td><td>Hash of Merkle tree of all transactions</td></tr>
<tr><td>timestamp</td><td>4</td><td>When block was created (overflows in 2106)</td></tr>
<tr><td>bits</td><td>4</td><td>Difficulty target used for this block</td></tr>
<tr><td>nonce</td><td>4</td><td>Nonce found to generate this block</td></tr>
</table>

# Merkle Trees

Ralph Merkle, [_Publishing a New Idea_](http://merkle.com/1974/).
Includes his [cs244 project
proposal](http://merkle.com/1974/FirstCS244projectProposal.pdf)
("Discussion: No, I am not joking.") and [ACM rejection
letter](http://merkle.com/1974/ExpertLetter.pdf) ("I am sorry to have to
inform you that the paper is not in the main stream of present
cryptography thinking and I would not recommend that it be
published...").

[https://github.com/btcsuite/btcd/blob/master/blockchain/merkle.go](https://github.com/btcsuite/btcd/blob/master/blockchain/merkle.go) (some comments removed)

```go
// HashMerkleBranches takes two hashes, treated as the left and right tree
// nodes, and returns the hash of their concatenation. 
func HashMerkleBranches(left *btcwire.ShaHash, right *btcwire.ShaHash) *btcwire.ShaHash {
   var sha [btcwire.HashSize * 2]byte
   copy(sha[:btcwire.HashSize], left.Bytes())
   copy(sha[btcwire.HashSize:], right.Bytes())
   newSha, _ := btcwire.NewShaHash(btcwire.DoubleSha256(sha[:]))
   return newSha
}

func BuildMerkleTreeStore(transactions []*btcutil.Tx) []*btcwire.ShaHash {
    nextPoT := nextPowerOfTwo(len(transactions))
    arraySize := nextPoT*2 - 1
    merkles := make([]*btcwire.ShaHash, arraySize)

    // Create the base transaction shas and populate the array with them.
    for i, tx := range transactions { merkles[i] = tx.Sha() }

    // Start the array offset after the last transaction and adjusted to the
    // next power of two.
    offset := nextPoT
    for i := 0; i < arraySize-1; i += 2 {
        switch {
           case merkles[i] == nil: 
              merkles[offset] = nil

           case merkles[i+1] == nil:
              newSha := HashMerkleBranches(merkles[i], merkles[i])
              merkles[offset] = newSha
       
           default:
              newSha := HashMerkleBranches(merkles[i], merkles[i+1])
              merkles[offset] = newSha
       }
       offset++
    }
    return merkles
}
```

<a href="/classes/merkle.png"><img src="/classes/merkle.png" width=600></a>

What is needed to verify <span class="math">T<sub>2</sub></span> in <span class="math">H<sub>root</sub></span>?
<div class="gap"></div>


What must be recomputed if <span class="math">T<sub>3</sub></span> is replaced?
<div class="gap"></div>

What must be computed if a new node, <span class="math">T<sub>5</sub></span>, is added?
<div class="gap"></div>

How many SHA-256 hashes must be computed to verify [Block 375474](https://blockexplorer.com/block/00000000000000000a02a2bf9d7b366e226160b8412830e6d72061ea7712d970)?
<div class="gap"></div>


# Mining Cost

The measured time to compute one SHA-256 hash on my EC2 node (2.5 GHz
processor) is 750 ns.  Approximately _how many instructions execute_ to
compute on SHA-256 hash?
<div class="gap"></div>

**Assumption.** SHA-256 produces a uniform random output.  (We know this
  is not really true, but it is a reasonable approximation, and
  necessary for the analysis.)  So, we can model SHA-256 on any (new)
  input <span class="math">_x_</span> as drawing randomly from 2<sup>256</sup> possible outputs:

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;SHA-256(<span class="math">_x_</span>) &larr; [0, 2<sup>256</sup>)

`$ Target = \frac{T_{max}}{Difficulty}$ `  
`$ T_{max} \approx 2^{224}$ `


[Current Bitcoin Difficulty](https://bitcoinwisdom.com/bitcoin/difficulty) = 59,335,351,234

# Energy

Why is energy/hash so much less for custom ASICs? 
<div class="gap"></div>

In an ASIC, it is possible to build an XOR using 4 transistors.  How
many transistors have to flip to do an XOR on a general purpose
processor like an Intel i7?
<div class="gap"></div>


