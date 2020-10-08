package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	update()
	var input string
	for {
		fmt.Print("-->")
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "exit" {
			break
		}
		if input == "update" {
			update()
		} else if len(input) > 4 && input[0:3] == "add" {
			add(input[4:])

		} else if len(input) > 7 && input[0:6] == "remove" {
			remove(input[7:])
		} else {
			request(input)
		}
	}

}
func request(stock string) bool {
	resp, err := http.Get("https://www.google.com/search?q=" + stock + "+stock&safe=strict&rlz=1C1CHBF_enIE784IE784&sxsrf=ALeKk00R7a6xf6UTpwxc_R23lq_m9yQx0A:1600457550056&gbv=1&sei=TgtlX_eLA9uM1fAP7MacsAc")
	if err != nil {
		fmt.Println("get request for ticker failed")
		return false
	}
	defer resp.Body.Close()
	output, _ := ioutil.ReadAll(resp.Body)

	outputstring := string(output)
	changepos := strings.LastIndex(outputstring, "lB8g7")
	changeval := ""
	if changepos == -1 {
		changepos = strings.LastIndex(outputstring, "AWuZUe")
		changeval = string(outputstring[changepos+8 : changepos+25])
	} else {

		changeval = string(outputstring[changepos+7 : changepos+25])
	}
	lastpos := strings.LastIndex(changeval, "<")
	changeval = changeval[0:lastpos]
	re:=regexp.MustCompile(`Bp4i AP7Wnd">[0-9]`)
	position:=re.FindStringIndex(outputstring)[1]
	value := string(outputstring[position-1 : position+25])
	endpoint := strings.Index(value, " ")
	value = value[0:endpoint]
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		fmt.Println("invalid input")
		return false
	}
	fmt.Println(stock + ": " + value + " " + changeval)

	return true
}
func update() {
	tickers, err := os.Open("tickers.txt")
	if err != nil {
		fmt.Println("could not access tickers file")
		return
	}
	defer tickers.Close()
	lines := bufio.NewScanner(tickers)
	name := []string{}
	count := 0
	for lines.Scan() {
		name = append(name, lines.Text())
		count++

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
func add(addition string) {
	file, err := os.OpenFile("tickers.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("could not open tickers file to add entry")
		return
	}
	if !(request(addition)) {
		return
	}
	defer file.Close()
	lines := bufio.NewScanner(file)
	match := false
	for lines.Scan() {
		if lines.Text() == addition {
			match = true
		}

	}
	if match == false {
		file.WriteString(addition + "\n")
		fmt.Println("added " + addition)

	} else {
		fmt.Println("ticker already in watchlist")
	}

}
func remove(removal string) {
	file, _ := ioutil.ReadFile("tickers.txt")
	filestring := string(file)
	replacement := strings.ReplaceAll(filestring, removal+"\n", "")
	if filestring == replacement {
		fmt.Println(removal + " not in watchlist")
		return

	}
	ioutil.WriteFile("tickers.txt", []byte(replacement), 0644)
	fmt.Println("removed " + removal)

}
