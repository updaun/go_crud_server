package repository

import "example.com/m/types"

type UserRepository struct {
	UserMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{}
}
