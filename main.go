package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func logo() {
	fmt.Printf("\n               /$$$$$$$$                 /$$ /$$$$$$$$             /$$    	\n")
	fmt.Println("              | $$_____/                | $$| $$_____/            | $$    	")
	fmt.Println("              | $$       /$$$$$$$   /$$$$$$$| $$       /$$   /$$ /$$$$$$  	")
	fmt.Println("              | $$$$$   | $$__  $$ /$$__  $$| $$$$$   |  $$ /$$/|_  $$_/  	")
	fmt.Println("              | $$__/   | $$  \\ $$| $$  | $$| $$__/    \\  $$$$/   | $$    	")
	fmt.Println("              | $$      | $$  | $$| $$  | $$| $$        >$$  $$   | $$ /$$	")
	fmt.Println("              | $$$$$$$$| $$  | $$|  $$$$$$$| $$$$$$$$ /$$/\\  $$  |  $$$$/")
	fmt.Println("              |________/|__/  |__/ \\_______/|________/|__/  \\__/   \\___/  	")
	fmt.Println("")
	fmt.Println("                       JsValidator Tool By @SirBugs .go Version")
	fmt.Println("                             V: 1.0.1 Made With All Love")
	fmt.Println("                  For Extracting all possilbe endpoints from Js files ")
	fmt.Println("                         Twitter@SirBagoza -- GitHub@SirBugs")
	fmt.Printf("                           Run : go run main.go jsurls.txt\n\n")
}

func gimmejslink(jsurl string) {
	client := http.Client{
		Timeout: 7 * time.Second,
	}
	req, err := http.NewRequest("GET", jsurl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0") // Setting To Real User-Agent
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if resp == nil {
		fmt.Println("[ - ] Bad URL")
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		response := string(body)
		if resp.StatusCode == 200 {
			// fmt.Println(response + "\n\n")
			PossibleEndPoints := strings.Count(response, "\"/")
			for i := 1; i < PossibleEndPoints+1; i++ {
				// fmt.Println(i)
				CurrentEndPoint := strings.Split(response, "\"/")[i]
				CurrentEndPointt := strings.Split(CurrentEndPoint, "\"")[0]
				fmt.Println(CurrentEndPointt)
			}
		} else {
			fmt.Println("[ - ] Bad JS File Detected")
		}
	}
}

func main() {
	// gimmejslink("https://dashboard.nylas.com/_app/chunks/recurly-536e8947.js")

	logo()

	arg := os.Args[1]
	myFile, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}

	myScanner := bufio.NewScanner(myFile)

	for myScanner.Scan() {
		gimmejslink(myScanner.Text())
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)

}
