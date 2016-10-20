package vk

import "github.com/labstack/gommon/log"

func NewVkFetcher(groups, users []string) *VkFetcher {
	return &VkFetcher{
		Groups:       groups,
		Users:        users,
		vkClient:     NewVkClient(),
		FetchedUsers: make(map[int64]VkUser),
	}
}

type VkFetcher struct {
	Groups   []string
	Users    []string
	vkClient *VkClient

	FetchedUsers map[int64]VkUser
}

func (f *VkFetcher) FetchAllUsers() []VkUser {
	for _, groupId := range f.Groups {
		users := f.FetchGroupMembers(groupId)
		fetchedSize := len(f.FetchedUsers)
		f.AddUsers(users)
		newFetchedSize := len(f.FetchedUsers)

		log.Infof("Fetched %d users from group %s, unique %d", len(users), groupId, newFetchedSize-fetchedSize)
	}

	for _, userId := range f.Users {
		users := f.FetchUserFriends(userId)
		fetchedSize := len(f.FetchedUsers)
		f.AddUsers(users)
		newFetchedSize := len(f.FetchedUsers)

		log.Infof("Fetched %d friends from user %s, unique %d", len(users), userId, newFetchedSize-fetchedSize)
	}

	var users []VkUser
	for _, value := range f.FetchedUsers {
		users = append(users, value)
	}

	return users
}

func (f *VkFetcher) AddUsers(users []VkUser) {
	for _, user := range users {
		user.Uid = user.GetId()
		user.Id = user.GetId()
		f.FetchedUsers[user.GetId()] = user
	}
}

func (f *VkFetcher) FetchGroupMembers(groupId string) (users []VkUser) {
	log.Infof("Fetching members of group %s", groupId)
	for i := 0; i < 100; i++ {
		response, err := f.vkClient.GetGroupMembers(1000, i*1000, groupId)
		if err != nil {
			log.Fatalf("Err: %s", err.Error())
			return
		}
		if len(response.Users) == 0 {
			return
		}
		users = append(users, response.Users...)
	}
	log.Fatalf("WTF?")
	return
}

func (f *VkFetcher) FetchUserFriends(userId string) (users []VkUser) {
	log.Infof("Fetching friends of user %s", userId)

	for i := 0; i < 100; i++ {
		response, err := f.vkClient.GetUserFriends(1000, i*1000, userId)
		if err != nil {
			log.Fatalf("Err: %s", err.Error())
			return
		}
		if len(response.Users) == 0 {
			return
		}
		users = append(users, response.Users...)
	}
	log.Fatalf("WTF?")
	return
}
