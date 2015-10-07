---
title: "Class 05: Becoming Paranoid"
date: '2015-09-09'
nocomment: false
weight: 3
menu: "classes"
---

## Schedule

   <div class="todo">
**Wednesday, 9 September** (now): Checkup 1 revisions (if desired).

**Tuesday, September 15** (8:29pm): [Problem Set 1](http://www.bitcoin-class.org/ps/ps1) due.

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
<iframe src="//www.slideshare.net/slideshow/embed_code/key/jXzZMV6Wpg9kCK"
width="625" height="400" frameborder="2" marginwidth="0"
marginheight="0" scrolling="no"></iframe>

   <div class="caption">
Note: ink markings may not appear in the
embedded viewer.  To see them, [download the slides](/classes/class5-post.pptx).
   </div>

</center>

## Notes

What does it mean for a problem to be _hard_?
<div class="gap"></div>

If you know an algorithm with running time in <span class="math">_O_(2<sup>_n_</sup>)</span> for problem _P_, what do you know about the hardness of problem _P_?
<div class="gap"></div>

What are the most common reasons for cryptosystems to fail in practice?

<!--page-->

   <div class="highlight">
We didn't get to this in class, but will cover it in a future class.
   </div>

## Bitcoin's Curve

Standards for Efficient Cryptography: [_SEC 2: Recommended Elliptic Curve Domain Parameters_](http://www.secg.org/sec2-v2.pdf) (Certicom Research, 27 January 2010)

> _Verifiably random parameters offer some additional conservative features. These parameters are
> chosen from a seed using SHA-1 as specified in ANSI X9.62 [X9.62]. This process ensures that
> the parameters cannot be predetermined. The parameters are therefore extremely unlikely to
> be susceptible to future special-purpose attacks, and no trapdoors can have been placed in the
> parameters during their generation. When elliptic curve domain parameters are chosen verifiably
> at random, the seed S used to generate the parameters may optionally be stored along with the
> parameters so that users can verify the parameters were chosen verifiably at random._

What does it mean for parameters to be "verifiably random"?
<div class="gap"></div>

# Randomness

**Kolmogorov Complexity:** <span class="math">_K_(_s_)</span> = the length of the shortest description of <span class="math">_s_</span>.

**Kolmogorov's Definition of Random:** A sequence <span
  class="math">_s_</span> is random, if <span class="math">_K_(_s_) =
  |_s_| + _C_</span> 

What is the Kolmogorov Complexity of the string `0001000010000011111111100111...`?
<div class="biggap"></div>

What is the Kolmogorov Complexity of the string: `1MRigEo5423vycLnUdSnA4C6Ts691fUiYu 18UikW89q9VgGDftQW3Gmuhe4sQDCFP5kD 19ZQwQmfAsgy47ErehfkW3SeSzNGFfH9iN 1AZCH1insc6JrT2Z9SiNvgtTugXg8sF8yd 15qYggRJvmyZfpchxvNqr6h3pNjw6bGBV9 1C943NwPPffUFY7VDzi3kt7KikXwc2vdkN 1JBhLLCgNYhR8f6AZcRS3mjfEAmMzPvwyf 1JvDrBSYm6o4ZTQUhwUE4FhPFxd2wuXWUR 1KcBM1RNhcp1oENycoD4AezA5Se4SrsZnA 16JZWC433XRxjWwR7X65uxRVFdLTmoPr4t 149LB8VYaT1BdMLyQUL92Kj6KrJfNwcp64 16zDGuzbwkHjW8dNYMw9stDjRbTzVSLZU1 1HfMaZn53ZDWKgmhWxk1UPZMjQ6QmpW6m 14gZWnuwKpRLTCUFCAgTZMciRaEdrkmEpr 1BZ2ateDPugmqLzYsXVy9EK5BguvXa2Bnj 1rCdRyMVcZHJaHA2LKUvRqYBcHqvAfQkc 1Ak8VwX6x4FPbA6aXTC3BQGQHnnhfaJuB8 129sBvF6Jternwdn5XcoA37LinQRcmAD5U1H2in 1HxEzSKHssPtog2krjFPiPfrKSiw4...` ?

<div class="gap"></div>


Daniel J. Bernstein, Tung Chou, Chitchanok Chuengsatiansup, Andreas
H&uuml;lsing, Tanja Lange, Ruben Niederhagen, and Christine van Vredendaal.
[_How to Manipulate Curve Standards: A White Paper for the Black  Hat_](https://eprint.iacr.org/2014/571.pdf), 2014.


How likely is it that the parameters for the secp256k1 curve used by bitcoin have a trapdoor?

<div class="gap"></div>

How should ECC parameters be generated for an important curve in a standard?
<div class="biggap"></div>

[_Root Zone DNSSEC KSK Ceremonies
Guide_](http://www.root-dnssec.org/wp-content/uploads/2010/02/draft-icann-dnssec-ceremonies-00.txt).
If you have a few hours to spare, you can watch a key signing for the
DNSSEC (Domain Name System): [DNSSEC KSK Ceremony
20](https://www.iana.org/dnssec/ceremonies/20)

## Dual-EC PRNG 

[NIST Special Publication 800-90A Recommendation for Random Number
Generation Using Deterministic Random Bit
Generators](http://csrc.nist.gov/publications/nistpubs/800-90A/SP800-90A.pdf)

<span class="math">_P_</span> and <span class="math">_Q_</span> are points on the curve, specified by the standard (but not
explained how <span class="math">_Q_</span> is choosen).  <span class="math">_P_</span> is a generator, so there exists some
<span class="math">_e_</span> such that <span class="math">_Q_<sup>_e_</sup> = _P_</span>.

<span class="math">_s_<sub>0</sub> = </span> initialize with seed randomness  
<span class="math">_s_<sub>_i_+1</sub> = &phiv;(_s_<sub>_i_</sub> &times; _P_)</span>  
<span class="math">_r_<sub>_i_</sub> = &phiv;(_s_<sub>_i_</sub> &times; _Q_)</span>  
<span class="math">_o_<sub>_i_</sub> =</span> the low-order 16 bits of the <span class="math">_x_</span>-coordinate of <span class="math">_r_<sub>_i_</sub></span>.

Given <span class="math">_o_<sub>i</sub></span>, how much work is it to find all the possible <span class="math">_r_<sub>i</sub> = (_x_, _y_)</span> values?
<div class="gap"></div>

Given <span class="math">_e_</span>, what is <span class="math">&phiv;(_e_ &times; _A_)</span> where <span class="math">_A_</span> is a possible <span class="math">_r_<sub>_i_</sub></span> value?
<div class="gap"></div>

Dan Shumow, Niels Ferguson.  [_On the Possibility of a Back Door in the
NIST SP800-90 Dual Ec Prng_](http://rump2007.cr.yp.to/15-shumow.pdf).
CRYPTO 2007 Rump Session.

Michael Wertheimer (NSA), [_Encryption and the NSA Role in International
Standards_](http://www.ams.org/notices/201502/rnoti-p165.pdf), Notices
of the American Mathematical Society, February 2015.

Wertheimer's letter is an attempt to respond to [_Mathematicians Discuss the
Snowden Revelations_](http://www.ams.org/notices/201406/rnoti-p623.pdf).

> _The recent revelations about the NSA’s spying
programs are both dismaying and encouraging.
What is encouraging is that they might lead not
just to a reform of the intelligence agencies but
also to a more serious look at what the ongoing
and inevitable erosion of privacy is doing to our
society. What is dismaying is less the intrusive data
collection itself and more what it reveals about the
decision-making processes inside the government._ (Andrew Odlyzko)


How satisfying is the NSA's response?  Are you more dismayed or encouraged?
<div class="gap"></div>
