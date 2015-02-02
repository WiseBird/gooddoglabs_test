package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	serviceUrl = flag.String("serviceurl", "http://0.0.0.0:9000", "Rest service url")
)

func main() {
	flag.Parse()

	client := NewClient(*serviceUrl)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Username and password with space between: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		var username, password string
		_, err = fmt.Sscanf(line, "%s %s", &username, &password)
		if err != nil {
			fmt.Println("Sscanf" + err.Error())
			continue
		}

		client.SetAuth(username, password)
		err = client.CheckAuth()
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	for {
		users, err := client.Users()
		if err != nil {
			fmt.Println(err)
			break
		}

		for _, user := range users {
			fmt.Println(user)
		}

		for {
			fmt.Printf("Enter new user first and last names with space between: ")
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}

			var firstname, lastname string
			_, err = fmt.Sscanf(line, "%s %s", &firstname, &lastname)
			if err != nil {
				fmt.Println("Sscanf" + err.Error())
				continue
			}

			err = client.CreateUser(firstname, lastname)
			if err != nil {
				fmt.Println(err)
				continue
			}

			break
		}
	}
}
