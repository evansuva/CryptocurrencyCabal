import requests
import blockcypher
from blockcypher import get_address_details 
from blockcypher import get_transaction_details
from sets import Set

API_KEY = '3b3561f9593e9a26f57c94e7e866efa3'

#read in addresses without 
#suspect_addresses = [line.rstrip('\n') for line in open('suspects.txt')]
suspect_addresses = [line.rstrip('\n') for line in open('infamous-addresses.txt')]

output_addresses = set()


def address_display(adr):
    try:
        r = requests.get('https://www.walletexplorer.com/api/1/address-lookup?address=' 
		         + adr + '&caller=virginia.edu').json()
        return r["label"]
        #return adr + " (" + r["found"] + ")"
    except:
        return adr



#get transactions with these addresses
#in these txs, find the transactions where the suspects are sending
#get the addresses of the outputs of these transactions
#check if these addresses are 

for a in suspect_addresses :
	print a
	#get_address_full(address=a)
	r = get_address_details(a.strip(), api_key=API_KEY)
	for b in r['txrefs']:
		input_of_tx = b['tx_input_n']
		hash_of_tx = b['tx_hash']
		if(input_of_tx > -1):
			#print input_of_tx
			#get the outputs of
			transaction = blockcypher.get_transaction_details(hash_of_tx, api_key=API_KEY)
			for output in transaction['outputs']:
				for addr in output['addresses']:
					output_addresses.add(addr)

for a in output_addresses:
	disp = address_display(a)
	if(len(disp) < 34):
		print disp
