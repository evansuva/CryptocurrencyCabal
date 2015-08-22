import urllib.request, urllib.error
from decimal import Decimal


def remove_non_ascii_1(text):

    return ''.join(i for i in text if ord(i)<128) 


url = 'https://localbitcoins.com/sell-bitcoins-online/US/united-states/'
req = urllib.request.Request(url, headers={'User-Agent': 'Mozilla/5.0'})
html = urllib.request.urlopen(req).read()

chars = ''

for letter in html:
    chars = chars + chr(letter)

index = 0;
index2 = 0;
exchange = ''
evalue = ''
dict = {}


for x in range(0, 50):
    index = chars.find('<i class="icon icon-thumbs-up-alt safe-badge"></i>', index)
    index2 = chars.find('<td>', index - 200)
    exchange = chars[index2+24:index-1].strip()
    
    remove_non_ascii_1(exchange)
    exchange = exchange.encode(encoding='utf_8', errors='strict')
    
    index = index + 108
    evalue = chars[index:index+6]
    dict[float(evalue)] = exchange
 
 
lis = list(reversed(sorted(dict)))


url = 'https://btc-e.com/exchange/ltc_usd'
req = urllib.request.Request(url, headers={'User-Agent': 'Mozilla/5.0'})
html = urllib.request.urlopen(req).read()

chars = ''

for letter in html:
    chars = chars + chr(letter)    

newindex = chars.find("<a href='https://btc-e.com/exchange/btc_usd'>BTC/USD<span style='display:block' id='last1'>",0)+91 
bitcoinprice = chars[newindex:newindex+3]  
print("")
print("**************************************************************************************")
print("* Buy BTC from:     https://btc-e.com/exchange/btc_usd                               *")
print("* At a rate of:     " + bitcoinprice + " USD/BTC                                                    *")
print("**************************************************************************************")
print("")
print(" Sell BTC to:      https://localbitcoins.com/sell-bitcoins-online/US/united-states/ ")
print(" At a rate of:                                                                      ")
print("")
i = 1
for item in lis:
    print ("   " + str(i) + "). " + str(item) + " USD/BTC")
    print ("        Payment Method:     " + str(dict[item])  )
    print ("        Profit:             " + str(float(item) - float(bitcoinprice))[:5] + " USD/BTC" )
    print("")
    i = i + 1