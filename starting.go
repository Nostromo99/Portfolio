package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	x := 5
	fmt.Println(x)
	resp, err := http.Get("https://www.google.com/search?q=nvda&safe=strict&rlz=1C1CHBF_enIE784IE784&sxsrf=ALeKk00R7a6xf6UTpwxc_R23lq_m9yQx0A:1600457550056&gbv=1&sei=TgtlX_eLA9uM1fAP7MacsAc")
	if err == nil {
		defer resp.Body.Close()
		output, err := ioutil.ReadAll(resp.Body)
		outputstring := string(output)
		if err == nil {
			help := string(outputstring[0])
			fmt.Println(help)
			fmt.Println("HI")
			// var store int
			for i := 0; i+6 != len(string(outputstring)); i++ {
				// store = i
				if string(outputstring[i:i+5]) == "iBp4i" {
					fmt.Println(string(outputstring[i-20 : i+20]))
				}
			}
			// answer := string(outputstring[store : store+20])
			// fmt.Println(answer)

		}

	} else {
		println(err)
		println("ho")
	}
}
