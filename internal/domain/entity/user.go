package entity

import "github.com/duosonic62/go-strings-history/internal/domain/valueobject"

type User struct {
	id    string
	name  string
	uid   string
	token valueobject.AuthorizationToken
}

func NewUser(id string, name string, uid string, token valueobject.AuthorizationToken) (*User, error) {
	if len(id) == 0 {
		return nil, NewApplicationError(400, "id must not be null", "id must not be null", nil)
	}

	if len(name) == 0 {
		return nil, NewApplicationError(400, "name must not be null", "name must not be null", nil)
	}

	if len(uid) == 0 {
		return nil, NewApplicationError(400, "uid must not be null", "uid must not be null", nil)
	}

	return &User{
		id: id,
		name: name,
		uid: uid,
		token: token,
	}, nil
}

func (user User) GetID() string {
	return user.id
}

func (user User) GetName() string {
	return user.name
}

func (user User) GetToken() valueobject.AuthorizationToken {
	return user.token
}

func (user User) GetUID() string {
	return user.uid
}