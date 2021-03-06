---
title: "Class 3: Digital Signatures"
date: '2015-09-02'
print: "class3.pdf"
nocomment: false
weight: 3
menu: "classes"
---
## Schedule 

   <div class="todo">
**Monday, September 7**: Check-up 1.  This will be a short in-class quiz
  to test your understanding of the main concepts covered so far.  It
  will cover material from the readings (see Class 2 for details) and classes 1-3.

**Tuesday, September 15** (8:29pm): [Problem Set 1](http://www.bitcoin-class.org/ps/ps1) due.

   </div>
<!--more-->

<center>
<iframe src="//www.slideshare.net/slideshow/embed_code/key/aX0xcwqcvhABQg"
width="625" height="400" frameborder="2" marginwidth="0"
marginheight="0" scrolling="no"></iframe>

   <div class="caption">
Note: ink markings may not appear in the
embedded viewer.  To see them, [download the slides](/classes/class3-post.pptx).
   </div>

</center>


## Signatures

"Real-life" signatures. Properties:

  * Easy to verify
  * Forging unlikely
  * Hard to repudiate

Digital signatures. Should have same properties, in the absence of legal forces on the internet.

Topics:

  * Assymetric cryptography
  * Digital signatures
  * Elliptic curve cryptography
  * Implementation pitfalls

Ordinary (or symmetric) Crypto

  * Both parties have to agree on a key

Diffie-Hellman key exchange

  * Proposed in 1976
  * Establishes a private secret key, unknown to any evesdropper

Whitfield Diffie and Martin E. Hellman, [_New Directions in Cryptography_](http://www-ee.stanford.edu/~hellman/publications/24.pdf), IEEE Transactions on Information Theory, 1976.

The paper concludes with an interesting historical discussion (that gets
bonus points for mentioning Thomas Jefferson):

> _While at first the public key systems and one-way authentication systems suggested in this paper appear to be unportended by past cryptographic developments, it is possible to view them as the natural outgrowth of trends in cryptography stretching back hundreds of years. Secrecy is at the heart of cryptography. In early cryptography, however, there was a confusion about what was to be kept secret. ... After the invention of the telegraph, the distinction between a general system and a specific key allowed the general system to be compromised, for example by theft of a cryptographic device, without compromising future messages enciphered in new keys. ... The development of computers has led for the first time to a mathematical theory of algorithms which can begin to approach the difficult problem of estimating the computational difficulty of breaking a cryptographic system. The position of mathematical proof may thus come full circle and be reestablished as the best method of certification._

> _The last characteristic which we note in the history of cryptography is the division between amateur and professional cryptographers. Skill in production cryptanalysis has always been heavily on the side of the professionals, but innovation, particularly in the design of new types of cryptographic systems, has come primarily from the amateurs. Thomas Jefferson, a cryptographic amateur, invented a system which was still in use in World War II, while the most noted cryptographic system of the twentieth century, the rotor machine, was invented simultaneously by four separate people, all amateurs. We hope this will inspire others to work in this fascinating area in which participation has been discouraged in the recent past by a nearly total government monopoly._

Observe how the them of amateurs being the ones who design the important
cryptographic systems carries through with bitcoin. Although Satoshi's
actual identity is unknown, it seems pretty clear that they are not a
professionally-trained cryptographer, and (so far, at least!) all the
digital currencies developed by world renowned cryptographers have
failed.

### Discrete Logarithm Problem

Given <span class="math">_g_</span>, <span class="math">_y_</span>, and <span class="math">_p_</span>, find <span class="math">_x_</span> such that <span class="math">_g_<sup>_x_</sup> mod _p_ = _y_</span>.

~~Discrete~~ Logarithm Problem: easy to solve for real numbers.

What is the range of elements out of which we are randomly selecting?

<div class="gap"></div>

### Mod 5 exponentiation

**Multiplicative order** is the number of multiplication after which the result repeats. 

Multiplicative order of <span class="math">_g_</span> is <span class="math">_n_</span> if <span class="math">_n_</span> is the smallest positive integer
such that <span class="math">_g_<sup>_n_</sup> = 1</span>.

### Exponent Modulus

  * Multiplicative order is at most <span class="math">_p_ - 1</span>.
  * Pick random <span class="math">_x_</span> such that <span class="math">0 &le; _x_ < _p_ - 1</span>.
  * <span class="math">_g_<sup>_a_</sup> _g_<sup>_b_</sup> mod _p_ = _g_<sup>_a_+_b_</sup> mod _p_ = _g_<sup>(_a_+_b_) mod _n_</sup> mod _p_</span>

# Public-key Cryptography

  1. Google announces <span class="math">_g_<sup>_a_</sup></span>.
  2. Bob picks random secret <span class="math">_b_</span>, computes <span class="math">(_g_<sup>_a_</sup>)<sup>_b_</sup>=_g_<sup>_ab_</sup></span>.
  3. Bob encrypts message <span class="math">_m_</span> and sends: <span class="math">_g_<sup>_b_</sup>, _g_<sup>_ab_</sup>_m_</span>.


### Man-in-the-middle (MITM)

Active adversary can still read everything. We have to know messages are coming from the right person.


## Digital Signatures

Discrete-log based signature

### ElGamal Signature Scheme

  * Fixed global parameters: <span class="math">_g_</span>, <span class="math">_p_</span>
  * Private key: <span class="math">_a_</span>
  * Public key: <span class="math">_g_<sup>a</sup> mod _p_</span>
  * Signing:
    1. Input: message <span class="math">_m_</span>
    2. Pick random <span class="math">_k_</span>
    3. Compute <span class="math">_r_ = _g_<sup>_k_</sup> mod _p_</span>;
               <span class="math">_s_ = (_m_-_ar_)_k_<sup>-1</sup> mod _p_ - 1</span>
    4. Send <span class="math">(_r_, _s_)</span> with message <span class="math">_m_</span>
  * Verification:
    1. Input: message <span class="math">_m_, (_r_,_s_)</span>
    2. Verify that <span class="math">_r_<sup>_s_</sup>(_g_<sup>_a_</sup>)<sup>_r_</sup> = _g_<sup>_m_</sup> mod _p_.

## Avoiding (overly) long numbers

Real-life keys are long. We can use any group where discrete log is hard.

A group is a set of elements and an associated operation such that it satisfies the following:

  * **Closure**: <span class="math">_a_ * _b_</span> is also a group element.
  * **Associativity:** for all <span class="math">_a_, _b_, _c_: (_a_*_b_)*_c_ = _a_*(_b_*_c_)</span>.
  * **Identity element:** there exists an element <span class="math">_e_</span>, such that for every element <span class="math">_a_: _a_ * _e_ = _a_ = _e_ * _a_</span>
  * **Inverse:** for every element <span class="math">_a_</span>, there exists an element <span class="math">_b_</span> such that <span class="math">_a_ * _b_ = _e_ = _b_ * _a_</span>.

Additional Cryptographic Properties:

  * Discrete logarithm should be hard
  * Group operation should be efficient

### Elliptic Curve Cryptography (ECC)

  * Group elements are points on a curve, e.g., <span class="math">_y_<sup>2</sup> = _x_<sup>3</sup> + 7</sup>.
  * Point "addition" using "geometry"

### Elliptic Curve Digital Signature Algorithm

Follows the same structure as ElGamal signature, but only on <span class="math">_x_</span>-coordinate.

## Pitfalls

If we ever repeat signing nonce, we leak the private key.  Sony actually
did this with Playstation 3 consoles ([_Console Hacking 2010: PS3 Epic Fail_](https://events.ccc.de/congress/2010/Fahrplan/attachments/1780_27c3_console_hacking_2010.pdf), CCC 2010).

Poor randomness makes private keys predictable. Use `/dev/urandom`
(Linux) or `java.security.SecureRandom` in Java.  If you use a
pseudorandom number generator that is seeded with an easily guessed
value, you are in big trouble!  A common mistake was to use
`Math.random()` or `srand(time(0))`.

Logjam attack: downgrade security during handshake.


## Notes

  * How are digital signatures and real-life signatures different in terms of why we trust them? What stops each from being forged by others?

  * Assume somebody really clever has a way of solving the discrete logarithm problem easily. That is, for any given <span class="math">_g_</span>, <span class="math">_y_</span>, and <span class="math">_p_</span>, the adversary can compute <span class="math">_x_</span> such that <span class="math">_g_<sup>_x_ mod _p_ = _y_</span>. How can this algorithm be used to break security of Diffie-Hellman protocol?

  * What is a nonce? What breaks if we reuse it between encrypted messages?

  * In elliptic curve cryptography, why do we use <span class="math">mod _p_</span> integers? What would go wrong if we used real numbers?  What would go wrong if we used unbounded integers?
