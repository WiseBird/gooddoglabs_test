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
			fmt.Println("Sscanf " + err.Error())
			continue
		}

		err = client.Login(username, password)
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
			fmt.Printf("Enter new user's username, password, first name and last name with space between: ")
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}

			var username, password, firstname, lastname string
			_, err = fmt.Sscanf(line, "%s %s %s %s", &username, &password, &firstname, &lastname)
			if err != nil {
				fmt.Println("Sscanf " + err.Error())
				continue
			}

			err = client.CreateUser(username, password, firstname, lastname)
			if err != nil {
				fmt.Println(err)
				continue
			}

			break
		}
	}
}
