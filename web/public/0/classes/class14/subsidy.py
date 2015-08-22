INTERVAL = 210000
COIN = 100000000

def subsidy(height):
    val = 50 * COIN
    halvings = height / INTERVAL
    val = val >> halvings
    return val

import datetime

def expect_block(block):
    minutes = block * (561.8 / 60) 
    start = datetime.datetime.strptime("2009-01-03", "%Y-%m-%d")
    return start + datetime.timedelta(minutes=minutes)
    
