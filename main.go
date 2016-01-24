package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"os"
	"fmt"
	"log"
)

func list() ([]string, error) {
	resp, err := http.Get("https://www.gitignore.io/api/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(body), ","), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: gotignore [--list] language\n")
		return
	}

	if os.Args[1] == "--list" {
		langs, err := list()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Print(strings.Join(langs, "\n"))
		return
	}

	resp, err := http.Get("https://www.gitignore.io/api/" + os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(body))
}
