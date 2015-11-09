import blockcypher

from blockcypher import get_block_overview

API_KEY = '3b3561f9593e9a26f57c94e7e866efa3'


def get_close_outputs(block_id, input_amount):
    possible_addresses = set()
    lowerFee = BTC_to_satoshi(input_amount) * 0.005 + BTC_to_satoshi(0.0005)
    upperbound = BTC_to_satoshi(input_amount) - lowerFee
    upperFee = BTC_to_satoshi(input_amount) * 0.006 + BTC_to_satoshi(0.0005)
    lowerbound = BTC_to_satoshi(input_amount) - upperFee

    #give a margin of error
    #margin_of_error = 0.05
    #lowerbound = lowerbound - 0.05 * lowerbound
    #upperbound = upperbound + margin_of_error * upperbound

    print ("We are searching for values between " + str(lowerbound) + " and " + str(upperbound))

    #please note the block has a size of 2054
    b = blockcypher.get_block_overview(block_id, coin_symbol='btc', txn_limit=2053, api_key=API_KEY)
    for tx in b['txids']:
        #go through all of the hashed txs
        transaction = blockcypher.get_transaction_details(tx, api_key=API_KEY)
        for output in transaction['outputs']:
            #how to get outputs???
            output_val = output['value']

            if(output_val >= lowerbound and output_val <= upperbound):
                possible_addresses.add(output_val)
    return possible_addresses


def satoshi_to_BTC(sval):
    return sval / 100000000

def BTC_to_satoshi(bval):
    return bval * 100000000

block_number = 379817+1
print ("Possible output addresses:" )
for receiver in get_close_outputs(block_number, 0.01523):
    print(receiver)