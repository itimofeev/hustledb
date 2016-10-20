package vk

import "gopkg.in/mgutz/dat.v1/sqlx-runner"

func NewUserInserter(db *runner.DB) *UserInserter {
	return &UserInserter{db: db}
}

type UserInserter struct {
	db *runner.DB
}

func (i *UserInserter) InsertUsers(users []VkUser) {
	for _, user := range users {
		insertUser(i.db, &user)
	}
}

func insertUser(db *runner.DB, user *VkUser) (*VkUser, error) {
	err := db.
		InsertInto("vk_user").
		Columns("id", "first_name", "last_name").
		Record(user).
		Returning("id").
		QueryScalar(&user.Id)

	return user, err
}
