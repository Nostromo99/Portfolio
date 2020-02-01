import pickle
import re
import requests
stuff_2=pickle.load(open("save.p","rb"))
stuff={"AMZN":0,"TSLA":0,"TTWO":0,"DELL":0,"NVDA":0,"AMD":0}
def update(dict):

    for item in dict.keys():
        x=0
        print(item)
        res=requests.get("https://finance.yahoo.com/quote/"+item+"?p="+item)
        x=re.search('"regularMarketPrice":{"raw":[0-9]*(\.[0-9]*)?',res.text)
        y=float(re.search('[0-9].*',x.group()).group())
        #z=re.search('[0-9].*',y)
        dict[item]=y
    return dict
#stuff=update(stuff)
print(stuff_2)
pickle.dump(stuff,open("save.p","wb"))

