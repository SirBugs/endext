package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func logo() {
	fmt.Printf("\n                  ______          ________     __ 	\n")
	fmt.Println("                 / ____/___  ____/ / ____/  __/ /_	")
	fmt.Println("                / __/ / __ \\/ __  / __/ | |/_/ __/	")
	fmt.Println("               / /___/ / / / /_/ / /____>  </ /_  	")
	fmt.Println("              /_____/_/ /_/\\__,_/_____/_/|_|\\__/  	")
	fmt.Println("")
	fmt.Println("            ( * ) EndpointsExtractor Tool By @SirBugs .go Version")
	fmt.Println("            ( * ) For Extracting all possilbe endpoints from Js files ")
	fmt.Println("            ( * ) Version: 1.0.5 (Updated 3.Vrs on 7/7/2023)")
	fmt.Println("            ( * ) Contact: Twitter@SirBagoza, GitHub@SirBugs, Medium@bag0zathev2")
	fmt.Printf("            ( * ) Command: go run main.go -l jsurls.txt\n\n")
	fmt.Println("            ( ! ) You can use only -u for single URL or -l for .JS file URLs, Not both")
	fmt.Printf("            ( ! ) This tool has been received the last 3 updates at once\n\n")
}

var EndPoints []string
var EPCount int

func IsinArray(element string, arr []string) bool {
	for _, value := range arr {
		if value == element {
			return true
		}
	}
	return false
}

func IsValid(stringo string) bool {
	if strings.Contains(stringo, "$") || strings.Contains(stringo, "#") || strings.Contains(stringo, "|") || strings.Contains(stringo, "\\") || strings.Contains(stringo, "?") || strings.Contains(stringo, "(") || strings.Contains(stringo, ")") || strings.Contains(stringo, "[") || strings.Contains(stringo, "]") || strings.Contains(stringo, "{") || strings.Contains(stringo, "}") || strings.Contains(stringo, ",") || strings.Contains(stringo, "<") || strings.Contains(stringo, ":") || strings.Contains(stringo, "*") || strings.Contains(stringo, ">") || strings.Contains(stringo, "\n") || strings.Contains(stringo, "$") || stringo == "/" || stringo == "./" || stringo == "\n" || stringo == "\n\n" || stringo == " " || stringo == "" || strings.Contains(stringo, ".svg") || strings.Contains(stringo, ".png") || strings.Contains(stringo, ".jpg") || strings.Contains(stringo, ".ico") || len(stringo) == 1 || stringo == "//" {
		return false
	} else {
		return true
	}
}

func gimmejslink(jsurl string, output string, activationFlag bool) {

	// #   Saving into the output file
	// # ###############################
	myOutPut, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer myOutPut.Close()

	// # #####################################################################

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
			// #########################################################################
			// Method 1 : Grepping all the values between " and "
			PossibleEndPoints := strings.Count(response, "\"/")
			for i := 1; i < PossibleEndPoints+1; i++ {
				// fmt.Println(i)
				CurrentEndPoint := strings.Split(response, "\"/")[i]
				CurrentEndPointt := strings.Split(CurrentEndPoint, "\"")[0]
				if IsValid(CurrentEndPointt) == false {
					fmt.Printf("")
				} else {
					if IsinArray(CurrentEndPointt, EndPoints) == true {
					} else {
						EPCount += 1
						EndPoints = append(EndPoints, CurrentEndPointt)
						if activationFlag == true {
							fmt.Println(" ( " + fmt.Sprint(EPCount) + " ) - " + jsurl + " :: (endpoint) " + CurrentEndPointt)
						} else {
							fmt.Println(" ( " + fmt.Sprint(EPCount) + " ) " + " ::  " + CurrentEndPointt)
						}
						_, err = fmt.Fprintln(myOutPut, " ( "+fmt.Sprint(EPCount)+" ) - "+jsurl+" :: (endpoint) "+CurrentEndPointt)
						if err != nil {
							panic(err)
						}
					}
				}
			}
			// #########################################################################
			// Method 2 : Grep using regex all the X Values in text for example, this.fetch(this.url("X")
			pattern := regexp.MustCompile(`this\.fetch\(this\.url\("([^"]+)"\)`)
			matches := pattern.FindAllStringSubmatch(response, -1)
			if matches != nil {
				for _, match := range matches {
					if IsValid(match[1]) == false {
					} else {
						if IsinArray(match[1], EndPoints) == true {
						} else {
							EPCount += 1
							EndPoints = append(EndPoints, match[1])
							if activationFlag == true {
								fmt.Println(" ( " + fmt.Sprint(EPCount) + " ) - " + jsurl + " ::  " + match[1])
							} else {
								fmt.Println(" ( " + fmt.Sprint(EPCount) + " ) " + " ::  " + match[1])
							}
							_, err = fmt.Fprintln(myOutPut, " ( "+fmt.Sprint(EPCount)+" ) - "+jsurl+" ::  "+match[1])
							if err != nil {
								panic(err)
							}
						}
					}
				}
			}
			// #########################################################################
		} else {
			fmt.Println("[ - ] Bad JS File Detected")
		}
	}
}

func main() {
	// Define three flags
	single_url := flag.String("u", "", "single URL to grep endpoints from")
	urls_list := flag.String("l", "", "this list would have more thana .js file URL to grep the endpoints from")
	output := flag.String("o", "js_endpoints.txt", "output file")
	activationFlag := flag.Bool("p", false, "public mode for showing the URLs of each endpoints & Showing the function (endpoints/fetch)")
	flag.Parse()

	// Check if the results file is already available or not, If not =>> Create it
	_, err := os.Stat("js_endpoints.txt")
	if os.IsNotExist(err) {
		file, createErr := os.Create("js_endpoints.txt")
		if createErr != nil {
			fmt.Println("Error creating file:", createErr)
			return
		}
		defer file.Close()
	}

	// Counter Variable
	EPCount = 0

	// Printing the logo function
	logo()

	// This part belongs to the old version before adding the flags of single_url use which is "-u"
	// And the list-URLs which is "-l"
	//arg := os.Args[1]
	//myFile, err := os.Open(arg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//myScanner := bufio.NewScanner(myFile)

	// Validating with the flags
	if *single_url != "" && *urls_list != "" || *single_url == "" && *urls_list == "" {
		panic("Please use single_url mode or URLs_list mode instead, You cannot use multiple or don't use one at least")
	}

	if *single_url != "" {
		gimmejslink(*single_url, *output, *activationFlag)
	} else if *urls_list != "" {
		arg := *urls_list
		myFile, err := os.Open(arg)
		if err != nil {
			panic(err)
		}
		myScanner := bufio.NewScanner(myFile)
		for myScanner.Scan() {
			go gimmejslink(myScanner.Text(), *output, *activationFlag)
			time.Sleep(100 * time.Millisecond)
		}
		time.Sleep(5 * time.Second)
	}

}

// Last Updates:
// # #############
// removing duplicates => 1.0.3

// add this.fetch(this.url("X") => 1.0.4
// short the urls filtering functionality => 1.0.4

// flag for single url -u or urls list -l => 1.0.5
// flag for public the urls -p => 1.0.5
// flag for output -o => 1.0.5

// Working on:
// burp extension => 1.0.6
