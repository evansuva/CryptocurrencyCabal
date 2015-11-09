import requests
import blockcypher
from blockcypher import get_address_details 
from blockcypher import get_transaction_details
from sets import Set

API_KEY = '3b3561f9593e9a26f57c94e7e866efa3'

#read in addresses without 
#suspect_addresses = [line.rstrip('\n') for line in open('suspects.txt')]
suspect_addresses = [line.rstrip('\n') for line in open('infamous-addresses.txt')]

set_of_sets_of_output_addresses = set()


'''
first, get all of the suspect addresses (suspect_addresses)
then, for each of those suspect addresses, get all of the transactions where the suspect was input
for each address, store the output addresses of these transactions in a set (one set for each suspect)
take the set difference of every set
store the difference in a set
that set is all of the addresses
//increase range of for loop to reach multiple layers of transactions
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
		if(input_of_tx > -1): #if the address is an input to the tx
			#print input_of_tx
			#get the outputs of
			transaction = blockcypher.get_transaction_details(hash_of_tx, api_key=API_KEY)
			for output in transaction['outputs']:
				for addr in output['addresses']:
					rec_addresses.add(addr)
	all_rec_addresses.append(rec_addresses)


popular_addresses = set()

print "comparing..."

for b1 in all_rec_addresses:
	for b2 in all_rec_addresses:
		if not (b1.issuperset(b2) and b1.issubset(b2)):
			b3 = b1.intersection(b2)
			for elem in b3:
				popular_addresses.add(elem)


for pa in popular_addresses:
	print pa

def address_display(adr):
    try:
        r = requests.get('https://www.walletexplorer.com/api/1/address-lookup?address=' 
		         + adr + '&caller=virginia.edu').json()
        return r["label"]
        #return adr + " (" + r["found"] + ")"
    except:
        return adr