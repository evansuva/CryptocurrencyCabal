##
## Some estimates about bitcoin mining
##

from __future__ import division

difficulty = 44455415962 # from https://blockchain.info/stats, 15 Feb 2015

def find_target(difficulty):
    return 2 ** 224 / difficulty

success_probability = find_target(difficulty) / 2**256

ave_days_in_month = (365.25 / 12)
# difficulty increase per month
# this is a very low assumption - over past year, average rate was 0.35
rate_of_difficulty = 0.05

def guess_difficulty(month):
    # assume difficulty increases a rate_of_difficulty each month
    return difficulty * ((1 + rate_of_difficulty) ** month)

hash_rate = 1.7 * 10**12 # 1.7 Th/s
hashes_in_month = hash_rate * 60 * 60 * 24 * ave_days_in_month
block_value = 25 * 240 # assumes fixed BTC value 

kWh = 0.07
electricity_cost = (1.2 * kWh) * 24 * ave_days_in_month

def expected_blocks(month):
    success_probability = find_target(guess_difficulty(month)) / 2**256 
    return hashes_in_month * success_probability

def cumulative_expected_blocks(months):
    blocks = 0
    for month in range(months):
        blocks += expected_blocks(month)
        
    return blocks

