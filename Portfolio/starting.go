package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

//add concurency
//read in lines and create function to send get requests
func main() {

	tickers, err := os.Open("portfolio/tickers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer tickers.Close()
	lines := bufio.NewScanner(tickers)
	name := []string{}
	count := 0
	for lines.Scan() {
		name = append(name, lines.Text())
		count++
		// stock := lines.Text()
		// resp, err := http.Get("https://www.google.com/search?q=" + stock + "+stock&safe=strict&rlz=1C1CHBF_enIE784IE784&sxsrf=ALeKk00R7a6xf6UTpwxc_R23lq_m9yQx0A:1600457550056&gbv=1&sei=TgtlX_eLA9uM1fAP7MacsAc")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer resp.Body.Close()
		// output, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// outputstring := string(output)
		// position := strings.LastIndex(outputstring, "iBp4i")
		// value := string(outputstring[position+14 : position+25])
		// for i := 0; i < len(value); i++ {
		// 	if string(value[i]) == " " {
		// 		value = value[0:i]
		// 		break
		// 	}
		// }
		// fmt.Println(stock + ": " + string(value))

	}
	var waiter sync.WaitGroup
	waiter.Add(count)
	for _, i := range name {
		go func(i string) {
			request(i)
			waiter.Done()
		}(i)
	}
	waiter.Wait()

}
func request(stock string) {
	resp, err := http.Get("https://www.google.com/search?q=" + stock + "+stock&safe=strict&rlz=1C1CHBF_enIE784IE784&sxsrf=ALeKk00R7a6xf6UTpwxc_R23lq_m9yQx0A:1600457550056&gbv=1&sei=TgtlX_eLA9uM1fAP7MacsAc")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	outputstring := string(output)
	position := strings.LastIndex(outputstring, "iBp4i")
	value := string(outputstring[position+14 : position+25])
	for i := 0; i < len(value); i++ {
		if string(value[i]) == " " {
			value = value[0:i]
			break
		}
	}
	fmt.Println(stock + ": " + string(value))
}
