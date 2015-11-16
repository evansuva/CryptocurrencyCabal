---
title: "Class 23: Before (and Beyond) Bitcoin"
date: '2015-11-16'
nocomment: false
print: "class23.pdf"
weight: 3
menu: "classes"
---

## Schedule

**Starting Wednesday and every following class**: Be prepared to give an
  elevator pitch for your project.  Your pitch should be no more than 2
  minutes long.  You may use visuals as long as you can obtain them by
  (quickly) entering a URL in a web browser.  Your pitch should get
  across in a convincing and engaging way:

- The purpose of your project (what problem are you solving)
- What you are actually doing
- Why we should care

**Monday, 23 November** (8:29pm): Project Progress Reports.  Send an
  email to
  [ccc-staff@cs.virginia.edu](mailto:ccc-staff@cs.virginia.edu), cc-ing
  all members of your project team.  The email should have a subject
  line, `Project: `_Title_, with your project title.  Its body should contain at least this information:

<!--more-->

1. A link to the website for your project (this could be a github page
if you want).  That site should have a front page that describes your
project, lists the team members, and provides more information about
your project.

2. A short paragraph explaining how your project has changed since the
preliminary proposal email.  This should explain if the goals of your
project have changed and why.

3. A description of what progress you have made on your project.  

4. A description of what you plan to do to complete your project, and
your plans for doing this.  If you have a multi-person team, this should
include an explanation of how your team is working together and who is
doing what.

5. (optional) Any questions you have for us.

## Notes

[High Trust Bank](https://www.fdic.gov/bank/individual/failed/hightrust.html) must be trusty!

David Chaum, Amos Fiat, and Moni Naor.  [_Untraceable Electronic
Cash_](./ecash.pdf).  CRYPTO 1988.

**Simple RSA Signatures**  
Public Key = <span class="math">(_e_, _n_)</span>
Private Key = <span class="math">_d_</span>

Identity: <span class="math">_M_<sup>_de_</sup> = _M_ mod _n_</span>

<span class="math">Sign(_m_) = _m_<sup>_d_</sup> mod _n_</span>

**Blind Signatures**
Alice picks random <span class="math">_k_</span> in <span class="math">[1, _n_)</span>  
<span class="math">_t_ = _mk_<sup>_e_</sup> mod _n_</span>  
Sends <span class="math">_t_</span> to signer.

Signer returns <span class="math">_t_<sup>_d_</sub></span>.


Signer returns <span class="math">_t_<sup>_d_</sup></span>.

<span class="math">_t_<sup>_d_</sup> = (_mk_<sup>_e_</sup> mod _n_)<sup>_d_</sup> mod _n_</span>  
&nbsp;&nbsp;&nbsp; <span class="math">= _m_<sup>_d_</sup>_k_<sup>_ed_</sup> mod _n_</span>  
&nbsp;&nbsp;&nbsp; <span class="math">= _m_<sup>_d_</sup>_k_ mod _n_</span>  

Dividing by <span class="math">_k_</span> gives <span class="math">Sign(_m_) = _m_<sup>_d_</sup> mod _n_</span>.


What should a signer know before signing a random-looking string?
<div class="gap"></div>

**Cut-and-Choose**

Suppose Alice sends 256 copies and the Bank checks 255 of them.  What is the probability Alice can cheat without getting caught?
<div class="gap"></div>

What should the maximimum bill size be to prevent cheating?
<div class="gap"></div>

### Identity Strings

_I_ = "alice@alice.org"  

<span class="math">_M_<sub>_i_</sub></span> = "Bill #[r<sub>i</sub>] : Bear’s Turns Bank owes the holder of this message $100."  
&nbsp;&nbsp;&nbsp; + identity strings:
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="math">_I_<sub>1</sub> = (_h_(_I_<sub>1L</sub>), _h_(_I_<sub>1R</sub>)), ..., _I_<sub>_n_</sub> = (_h_(_I_<sub>nL</sub>), _h_(_I_<sub>nR</sub>))</span>  
&nbsp;&nbsp;&nbsp; where <span class="math">_h_</span> is a one-way hash function and each <span class="math">_I_<sub>iL</sub> &oplus; _I_<sub>iR</sub> = _I_</span> (but <span class="math">_I_<sub>iL</sub> is choosen randomly).

To spend a bill, the reciever chooses either L or R for each pair for spender to open.

What is the probability Alice can spend a bill twice without revealing her identity?
<div class="gap"></div>


[Before Bitcoin: The Rise and Fall of DigiCash](http://globalcryptonews.com/before-bitcoin-the-rise-and-fall-of-digicash/)

> _By all accounts Chaum was a charismatic leader with an interesting management style, but he refused to compromise his artistic vision in any area against the best advice of his employees. He was suspicious of everyone and 'paranoid' with a habit of suddenly changing his mind without warning. At one time, Microsoft had offered DigiCash $180 million to allow them to preinstall Ecash software on Windows computers and the deal was on the verge of completion, but Chaum suddenly decided that his product was worth more and the deal collapsed. If the deal had gone through, cryptocurrency would now be as ubiquitous as Internet Explorer._

## Links

"Still think you’re anonymous on the Dark Web?", [NCA_UK's tweet](https://twitter.com/NCA_UK/status/530740716134490112), 7 November 2014

Sambuddho Chakravarty, Marco V. Barbera, Georgios Portokalidis, Michalis Polychronakis, and Angelos D. Keromytis. [_On the Effectiveness of Traffic Analysis Against Anonymity Networks Using Flow Records_](http://www3.cs.stonybrook.edu/~mikepo/papers/tor-netflow.pam14.pdf). _Passive and Active Measurement Conference_, March 2014.


[_Tor security advisory: "relay early" traffic confirmation attack_](https://blog.torproject.org/blog/tor-security-advisory-relay-early-traffic-confirmation-attack)

