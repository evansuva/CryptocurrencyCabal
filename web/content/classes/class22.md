---
title: "Class 22: Zero-knowledge proofs"
date: '2015-11-11'
nocomment: false
weight: 3
menu: "classes"
---

## Reading Materials

  * Ben Sasson et al. [Zerocash: Decentralized anonymous payments from Bitcoin](http://zerocash-project.org/media/pdf/zerocash-extended-20140518.pdf)

  * Miers, I.; Garman, C.; Green, M.; Rubin, A.D. [Zerocoin: Anonymous distributed e-cash from bitcoin](http://ieeexplore.ieee.org/iel7/6547086/6547088/06547123.pdf?arnumber=6547123)

<!--more-->

# Notes

In class, we saw how arbitrary statements in NP could be reduced to a questions in graph 3-coloring. This allowed us to use 3-coloring zero-knowledge procedure for seemingly unrelated problems. 

How would you represent a statement such as x=y+2 as a graph 3-coloring problem? Let's assume for simplicity that x is a 2-bit integer, while y can either be 0 or 1. You can represent the numbers in the graph in any way you want (as long as you specify the encoding). Design a graph that would be 3-colorable if and only if the numbers satisfy the relation x=y+2.

<div class="gap"></div>

The 3-coloring zero-knowledge protocol requires the prover to reveal the colors of any verifier-chosen pair of neighboring vertices. There is a risk that this may leak coloring information. Argue that no information is indeed being leaked, other than the fact that it is a valid 3-coloring. (Hint: You may want to rely on the fact that each time, the verifier already knows the probability with which each coloring should appear on the revealed nodes, so it sees no more information than what it already has.)

<div class="gap"></div>

