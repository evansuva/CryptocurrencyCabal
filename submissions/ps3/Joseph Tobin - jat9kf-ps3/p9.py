import requests
import blockcypher
from blockcypher import get_address_details 
from blockcypher import get_transaction_details
from sets import Set

API_KEY = '3b3561f9593e9a26f57c94e7e866efa3'

#read in suspect and popular addresses 
suspect_addresses = [line.rstrip('\n') for line in open('infamous-addresses.txt')]
popular_addresses = [line.rstrip('\n') for line in open('popular-infamous-addresses.txt')]

#create a dictionary
addr_totals = {
	
}

#add popular addresses into the keys of the dictionary
#please note we're using satoshis so every address maps to 
#an int
for a in popular_addresses:
	addr_totals[a.strip()] = 0



'''
to get more info about the addresses found, I am going to
find the amount of bitcoin sent to each of the popular 
addresses 

first, get all of the suspect addresses (suspect_addresses)
second, get all of the popular addresses
//done

go through all of the transactions where suspect is input,
popular is output, 
-->find the location of the popular address
-->store the output value in map at the popular address
'''

all_rec_addresses = []

print "adding..."

for a in suspect_addresses :
	#get_address_full(address=a)
	print a
	rec_addresses = set()
	r = get_address_details(a.strip(), api_key=API_KEY)
	for b in r['txrefs']:
		input_of_tx = b['tx_input_n']
		hash_of_tx = b['tx_hash']
		if(input_of_tx > -1): 
			#if the address is an input to the tx
			transaction = blockcypher.get_transaction_details(hash_of_tx, api_key=API_KEY)
			
			for output in transaction['outputs']:
				for addr in output['addresses']:
					if (addr in addr_totals):
						addr_totals[addr] += output['value']


print "sorting..."
#sort map by value


unique_totals = set()
for v in addr_totals.values():
	unique_totals.add(v)

sorted_totals = []

for u in unique_totals:
	sorted_totals.append(u)

sorted_totals = sorted(sorted_totals)


for k in sorted_totals:
	print k
	for j in addr_totals.keys():
		if addr_totals[j] == k :
			print j

