import blockcypher
import requests

def get_receivers(sending_address):
    receiving_addresses = set()
    r = blockcypher.get_address_details(sending_address)
    for tx in r['txrefs']: # loop through all transactions
        hash = tx['tx_hash']
        transaction = blockcypher.get_transaction_details(hash, api_key="37901f0cd5d1f74e7a2ffb7df3795e89")
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

def address_display(adr):
    try:
        r = requests.get('https://www.walletexplorer.com/api/1/address-lookup?address='
		         + adr + '&caller=virginia.edu').json()
        return adr + " (" + r['label'] + ")"
    except:
        return adr

addrs = []
f = open('suspects.txt', 'r')
for line in f:
    addrs.append(line[:-1])

# -------------------------------- PROBLEM 7 ------------------------------------------------------------
#for sender in addrs:
#    for receiver in get_receivers(sender):
#        print("   " + address_display(receiver[0]) + " (" + str(satoshi_to_BTC(receiver[1])) + " BTC)")
#--------------------------------------------------------------------------------------------------------

#------------------------------------ PROBLEM 8 + 9 ------------------------------------------------------
recvd = {}
second = []
for sender in addrs:
    for receiver in get_receivers(sender):
        if(sender != receiver):
            if(receiver not in recvd):
             recvd[sender] = [receiver[0], receiver[1], sender, False]
        else:
            arr = recvd[sender]
            arr[1] = arr[1] + receiver[1]
            if sender == arr[2]: arr[3] = True;
            recvd[receiver] = arr
        second.append(receiver[0])

for sender2 in second:
    for receiver2 in get_receivers(sender2):
        if(sender != receiver):
            if(receiver not in recvd):
             recvd[sender] = [receiver[0], receiver[1], sender, False]
        else:
            arr = recvd[sender]
            arr[1] = arr[1] + receiver[1]
            if sender == arr[2]: arr[3] = True;
            recvd[receiver] = arr

for r in recvd.keys():
    if recvd[r][3] == True:
        print("   " + address_display(recvd[0]) + " (" + str(satoshi_to_BTC(recvd[1])) + " BTC)")