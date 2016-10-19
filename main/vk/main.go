package main

import (
	"fmt"
	"github.com/itimofeev/hustlesa/vk"
)

func main() {
	client := vk.NewVkClient()

	userResponse, err := client.GetGroupMembers(10, 0, "dwizhenie")
	if err != nil {
		fmt.Printf("!!!%+v\n", err) //TODO remove
	} else {
		fmt.Printf("!!!%+v\n", userResponse) //TODO remove
	}
}
