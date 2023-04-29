package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const serverURL = "http://172.20.176.1"
const serverPort = 7861
const txt2img_req = "/sdapi/v1/txt2img"
const options_req = "/sdapi/v1/options"

type options struct {
	Number int `json:"number"`
}

func main() {

	fmt.Println("Select a choice")

	for {

		var choice int

		fmt.Println("1. View current options")
		fmt.Println("2. Create an image")

		fmt.Print("> ")
		fmt.Scanf("%d", &choice)

		if choice == 1 {
			URL := fmt.Sprintf("%s:%d%s", serverURL, serverPort, options_req)
			resp, err := http.Get(URL)
			if err != nil {
				log.Fatal(err)
			}

			if resp.StatusCode != http.StatusOK {
				log.Fatal("unexpected response status", resp.Status)
			}
			io.Copy(os.Stdout, resp.Body)
			continue

		} else if choice == 2 {
			fmt.Sprintf("placeholder: %s", txt2img_req)

		} else {
			fmt.Print("\033[H\033[2J") // Clear Screen
			fmt.Println("Invalid Choice!")
			fmt.Println("")
			fmt.Println("Select a choice from below:")
		}
	}
}
