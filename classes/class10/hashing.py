##
## Some estimates about bitcoin mining
##

from __future__ import division

difficulty = 44455415962 # from https://blockchain.info/stats, 15 Feb 2015

def find_target(difficulty):
    return 2 ** 224 / difficulty

success_probability = find_target(difficulty) / 2**256
expected_hashes = 1.0 / success_probability # don't divide by 2, not looking for one key

def nanoseconds_to_years(ns):
    return ns / (10**9 * 60 * 60 * 24 * 365.25)

time_to_hash = 760 * 2 # ns, as measured on EC2 node

years_needed = nanoseconds_to_years(expected_hashes * time_to_hash)

# print "Years: " + str(years_needed)

ave_days_in_month = (365.25 / 12)
# difficulty increase per month
rate_of_difficulty = 0.05

def guess_difficulty(month):
    # assume difficulty increases a rate_of_difficulty each month
    return difficulty * ((1 + rate_of_difficulty) ** month)

hash_rate = 1.7 * 10**12 # 1.7 Th/s
hashes_in_month = hash_rate * 60 * 60 * 24 * ave_days_in_month
block_value = 25 * 240 # assumes fixed BTC value 

kWh = 0.07
electricity_cost = (1.2 * kWh) * 24 * ave_days_in_month

def expected_revenue(month):
    success_probability = find_target(guess_difficulty(month)) / 2**256 
    return block_value * hashes_in_month * success_probability

def expected_profit(month):
    return expected_revenue(month) - electricity_cost

def cummulative_profit(months):
    profit = 0.0
    month = 0
    while month < months:
        month_profit = expected_profit(month)
        if month_profit > 0:
            profit += month_profit
        else:
            print "Turn it off in month " + str(month)
            break
        month += 1
        
    return profit

