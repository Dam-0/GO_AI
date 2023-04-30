package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const serverURL = "http://172.20.176.1"
const serverPort = 7861
const txt2img_req = "/sdapi/v1/txt2img"
const options_req = "/sdapi/v1/options"

// json structure for all the options
type get_requests struct {
	model string `json:"sd_model_checkpoint"`
}

type post_requests struct {
}

// Function to obtain cetain hardcoded options
func url_get_request(url_get string) string {

	url_client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	required, err := http.NewRequest(http.MethodGet, url_get, nil)
	if err != nil {
		log.Fatal(err)
	}

	required.Header.Set("User-Agent", "go_webui_ai")

	response, getErr := url_client.Do(required)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	g := get_requests{}
	jsonErr := json.Unmarshal(body, &g)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	str, _ := json.MarshalIndent(g, "", "\t")
	return fmt.Sprintln(string(str))
	//return fmt.Sprintf("{Model_name:%s}", get_options.model)
	//return fmt.Sprintf("%+v\n", get_options)
}

func url_post_request() {

}

func main() {

	fmt.Println("Select a choice")

	for {

		var choice int

		fmt.Println("\n1. View current options")
		fmt.Println("2. Create an image")

		fmt.Print("\n> ")
		fmt.Scanf("%d", &choice)   // reads user choice
		fmt.Print("\033[H\033[2J") // Clear Screen

		if choice == 1 {
			url := fmt.Sprintf("%s:%d%s", serverURL, serverPort, options_req)

			fmt.Println(url_get_request(url))

		} else if choice == 2 {
			fmt.Printf("placeholder: %s \n", txt2img_req)

		} else {
			fmt.Print("\033[H\033[2J") // Clear Screen
			fmt.Println("Invalid Choice!")
			fmt.Println("")
			fmt.Println("Select a choice from below:")
		}
	}
}
