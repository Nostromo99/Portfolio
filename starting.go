package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://www.google.com/search?q=nvda+stock&safe=strict&rlz=1C1CHBF_enIE784IE784&sxsrf=ALeKk00R7a6xf6UTpwxc_R23lq_m9yQx0A:1600457550056&gbv=1&sei=TgtlX_eLA9uM1fAP7MacsAc")
	if err == nil {
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
		fmt.Println(value)

		// if err == nil {
		// 	help := string(outputstring[0])
		// 	fmt.Println(help)
		// 	fmt.Println("HI")
		// 	// var store int
		// 	for i := 0; i+6 != len(string(outputstring)); i++ {
		// 		// store = i
		// 		if string(outputstring[i:i+5]) == "iBp4i" {
		// 			fmt.Println(string(outputstring[i-20 : i+20]))
		// 		}
		// 	}
		// 	// answer := string(outputstring[store : store+20])
		// 	// fmt.Println(answer)

		// }

	} else {
		println(err)
		println("ho")
	}
}
