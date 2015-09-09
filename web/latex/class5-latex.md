% Class 5: Becoming More Paranoid
% '2015-09-09'

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

## Notes

What does it mean for a problem to be _hard_?

#


If you know an algorithm with running time in <span class="math">_O_($2^n$)</span> for problem _P_, what do you know about the hardness of problem _P_?

#


What are the most common reasons for cryptosystems to fail in practice?

\clearpage

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

#


# Randomness

**Kolmogorov Complexity:** <span class="math">_K_(_s_)</span> = the length of the shortest description of <span class="math">_s_</span>.

**Kolmogorov's Definition of Random:** A sequence <span
  class="math">_s_</span> is random, if <span class="math">_K_(_s_) =
  |_s_| + _C_</span> 

What is the Kolmogorov Complexity of the string `0001000010000011111111100111...`?

#
#


What is the Kolmogorov Complexity of the string: `1MRigEo5423vycLnUdSnA4C6Ts691fUiYu 18UikW89q9VgGDftQW3Gmuhe4sQDCFP5kD 19ZQwQmfAsgy47ErehfkW...' ?


#



Daniel J. Bernstein, Tung Chou, Chitchanok Chuengsatiansup, Andreas
H&uuml;lsing, Tanja Lange, Ruben Niederhagen, and Christine van Vredendaal.
[_How to Manipulate Curve Standards: A White Paper for the Black  Hat_](https://eprint.iacr.org/2014/571.pdf), 2014.


How likely is it that the parameters for the secp256k1 curve used by bitcoin have a trapdoor?


#


How should ECC parameters be generated for an important curve in a standard?
\clearpage

## Dual-EC PRNG 

\noindent
$P$ and $Q$ are points on the curve, specified by the standard (but not
explained how $Q$ is choosen).  $P$ is a generator, so there exists some
$e$ such that $Q^e = P$.

\noindent $s_0 =$ initalize with seed randomness  
\noindent $s_{i+1} = \varphi(s_i \times P)$

\noindent$r_i = \varphi(s_i \times Q)$
$o_i = $ the low-order 16 bits of the $x$-coordinate of $r_i$.

\noindent Given $o_i$, how much work is it to find all the possible $r_i = (x, y)$ values?

#
#


\noindent Given $e$, what is $\varphi(e \times A)$ where $A$ is a possible $r_i$ value?

#
#


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

