package main

import (
	"fmt"
	"github.com/itimofeev/hustlesa/vk"
)

/*
// disco_dubrovka
// dwizhenie
// discoswing
// discofoxan
// club3688533 (Динамика)
// svitelev_hustle
*/
func main() {
	groups := []string{"dwizhenie", "disco_dubrovka"}
	userNames := []string{"3956992"}
	fetcher := vk.NewVkFetcher(groups, userNames)

	users := fetcher.FetchAllUsers()

	fmt.Printf("!!!Unique users: %d\n", len(users)) //TODO remove
}
