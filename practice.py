import pickle
import re
import requests
class Portfolio():
    def __init__(self):
        self.stuff={"AMZN":0,"TSLA":0,"TTWO":0,"DELL":0,"NVDA":0,"AMD":0}
        try: pickle.load(open("save.p","rb"))
        except FileNotFoundError:
                self.update()
        self.stuff=pickle.load(open("save.p", "rb"))
    def update(self):

        for item in self.stuff.keys():
            x=0
            print(item)
            res=requests.get("https://finance.yahoo.com/quote/"+item+"?p="+item)
            x=re.search('"regularMarketPrice":{"raw":[0-9]*(\.[0-9]*)?',res.text)
            y=float(re.search('[0-9].*',x.group()).group())
            #z=re.search('[0-9].*',y)
            self.stuff[item]=y
            self.save()
        return self.stuff
    def save(self):
        pickle.dump(self.stuff,open("save.p","wb"))
test=Portfolio()
