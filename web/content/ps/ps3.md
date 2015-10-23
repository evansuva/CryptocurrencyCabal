---
title: "PS3: CSI: Blockchain"
date: '2015-10-21'
nocomment: false
menu: "hidden"
---

# DRAFT - Not yet posted!


# Problem Set 3:<br> CSI: Blockchain

   <div class="due">
Due: Sunday, 8 November at 8:29pm
   </div>

## Purpose

The goal of this assignment is for everyone to gain experience analyzing
the blockchain and understanding of the anonymity of bitcoin (and
mechanisms that are used to increase or decrease it), and to hopefully
help the FBI catch some heinous criminals.

<!--more-->

### Collaboration Policy

For this assignment, you may either work alone or with one other person
of your choice.  If you work with a partner, you should together submit
one assignment with both of your names on it that reflects your combined
efforts.  You may, and are encouraged to, discuss all of the problems
with anyone else you want (both on-line using the course web site or any
other means you choose, and in person).

### Submission

Submit your answers as a single PDF file using [this
link](https://www.dropbox.com/request/oar0AQeg4SJRNtdHbcSS).  The name
of your file should be `<your email ID>-ps3.pdf` (if you worked alone)
or `<partner1 email ID>-<partner2 email ID>-ps3.pdf`

Your submission should include clearly marked answers for all the
problems (highlighted in yellow).  None of the questions require
submitting a program, although you may find it helpful to write programs
to develop your answers.  If the code you write is less than a page, it
is best to just include it in the PDF writeup.  If it is longer, you may
submit separate code files (and mention them in the PDF submission).

## Warm-Up: Manual Exploration

You should have received by email a list of bitcoin addresses that were
given to us by the FBI.  These are addresses that were used to collect
ransom payments from victims of bitcoin ransomware attacks.  In the
scheme perpetrators encrypt a victim’s files and then demand that ransom
payment be made to a specific bitcoin address in return for the key to
decrypt the files.  

One of the addresses are marked with your email ID.  The intent of this
is to ensure that the examination efforts are well distributed over the
provided addresses.  For the manual exploration question, you should
start with that address (but feel free to look at any of the addresses
in the file, and for the programmatic analysis questions, you'll want to
use all the addresses).

You can use any tool you want to answer these questions, but we suggest
using [blockchain.info](http://blockchain.info).  One useful thing you
can do there is look at the taint analysis to see which other addresses
a given addresses is interacting with (e.g., examine the [reverse taint
analysis](https://blockchain.info/taint/14TCv5MKcyYCM3qij36x8xrKvKHHY1NXmq?reversed=true)
to see the addresses that are receiving bitcoin from the given address).

   <div class="problem"> 
**Problem 1.** 
(a) What address are you investigating?  
(b) When was that address in use?
(c) How many victims appear to have paid into the address?   
(d) How much was collected from each victim (the most relevant value would be the value in US dollars at the time of the payment)?  Did all the victims pay the same amount, or did the requested amount vary by victim?  
(e) How did the ransomist (owner of this address) distribute the proceeds?  Is there a percentage split, or some other fee being paid by the ransomist to other addresses?  
(f) What else can you learn about the operation starting from your transaction?
   </div>

Since our goal is to learn as much as we can about the overall
operation, you should post your answers to Problem 1 (for your address)
as a comment on this page.

## Blockchain APIs

One way to analyze the blockchain is to run a full bitcoin node, and
operate on your own local copy of the blockchain.  This requires
downloading and processing a large amount of data (approximately [45
GB](https://blockchain.info/charts/blocks-size) today).  Instead, you
can use an external service that provides an API for interacting with
the blockchain.  Several such APIs exist including
[blockchain.info](https://blockchain.info/api), Bitpay's [Insight
API](https://insight.is/), and BlockCypher's [Blockchain
API](http://dev.blockcypher.com/).  

For this assignment, you are welcome to use any services and open source
bitcoin libraries and openly-licensed code you want, but must follow the
license requirements of any code you use and credit this code in your
submission.  

The starting examples we provide will use Python and the BlockCypher
[Blockchain API](http://dev.blockcypher.com/), and we provide some
directions to get started using this next (but feel free to use other
tools if you prefer).

### Setting up BlockCypher API

**Install Python.** Start by downloading and installing Python
  ([https://www.python.org/downloads/release/python-350/](https://www.python.org/downloads/release/python-350/).
  Note that Python V3 is incompatible with Python V2 (on MacOSX, the
  default `python` is Python 2.7.2; to use Python 3, use `python3`).
  Even if you are not familiar with Python, if you feel comfortable
  learning a new language on the fly you should be fine jumping right
  into this assignment. If you prefer a more structured intro to Python
  there are many tutorials available, including
  [http://www.learnpython.org/](http://www.learnpython.org/).

**Install the Blockcypher Python library.** The [BlockCypher Python
  library](https://github.com/blockcypher/blockcypher-python) provides a
  convenient way to use the APIs.  To install it, execute `pip install
  blockcypher` (in the command shell).

Here is some example code that uses the BlockCypher API to find all the
addresses which received bitcoin directly from a sending address:

```Python
import blockcypher

def get_receivers(sending_address):
    receiving_addresses = set()
    r = blockcypher.get_address_details(sending_address)
    for tx in r['txrefs']: # loop through all transactions
        hash = tx['tx_hash']
        transaction = blockcypher.get_transaction_details(hash)
        # was this address an input
        sender = False
        for input in transaction['inputs']:
            if sending_address in input['addresses']:
                sender = True
                break
        if sender:
            for output in transaction['outputs']:
                for adr in output['addresses']:
                    receiving_addresses.add((adr, output['value']))

    return receiving_addresses

def satoshi_to_BTC(sval):
    return sval / 100000000

sender = '1Ez69SnzzmePmZX3WpEzMKTrcBF2gpNQ55'
print ("Received from " + sender + ": ")
for receiver in get_receivers(sender):
    print("   " + receiver[0] + " (" + str(satoshi_to_BTC(receiver[1])) + " BTC)")
```




## Mixing Services






The directions we provide use blockchain.info’s Python API (https://github.com/blockchain/api-v1-client-python). We will be using its “blockexplorer” module. 


Installing our Tools
Install Python. 
Install Blockchain API library. Follow the directions to install Blockchain.info’s API (https://github.com/blockchain/api-v1-client-python). Familiarize yourself with the blockexplorer module (https://github.com/blockchain/api-v1-client-python/blob/master/docs/blockexplorer.md). You will mainly be interacting with the “Address” object.

**suspects.txt**

Problem 1. Determine the average number of transactions that a given address in the suspect set is used for. (Show how you determined the average, including a snippet of code is fine). What kind of risks, if any, does this scheme take on from using some of the same addresses multiple times? 

Problem 2. Looking manually through some of the transactions (i.e. on blockchain.info) describe any particular patterns you see in the spending behaviors of these addresses. Determine the cumulative amount of bitcoin received by all of the suspect addresses. Next, determine the largest single ransom payment in the set. Provide the tx_index and a justification for why you think this payment was made by a victim.

Problem 3. Now it’s time to follow the money. We want to figure out if there are any common addresses that eventually receive coin from more than one of the starting addresses. This could detect if the suspect addresses are running received coins through a mixnet, but outputting to the same address eventually, which could mean payments to the same “vendor”. A good approach to this problem could be writing a program that recursively follows successive outputs from a series of seed addresses down to a reasonable depth and then checks for overlap between sets. We are also interested in finding out the relative importance of the addresses found. You can determine this based on factors such as amount of “dirty” bitcoins they receive, number of seed addresses they are associated with, or how many hops away the addresses are on average from associated seed addresses.

Problem 4. Given the intel we have gathered from our analysis, what do you think some possible next steps could be to de-anonymize the addresses? What kind of additional outside data would be required to make blockchain based findings such as our own actionable intelligence for law enforcement?

Problem 4 (Bonus?). The following address (not included in suspects.txt) is suspected by the FBI to have been used directly by the attackers. See what interesting things you can find by following the money! 

18dwCxqqmya2ckWjCgTYReYyRL6dZF6pzL





# Anonymity and Bitcoin

Readings:



  
### Submission

Follow the submission instructions at the beginning of this page by
8:29pm on Sunday, 8 November.

