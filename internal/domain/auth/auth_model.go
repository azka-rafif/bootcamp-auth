package auth

import (
	"encoding/json"
	"time"

	"github.com/evermos/boilerplate-go/shared/encrypt"
	"github.com/evermos/boilerplate-go/shared/nuuid"
	"github.com/evermos/boilerplate-go/shared/roles"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type AuthPayload struct {
	UserName string `json:"userName" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type LoginPayload struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}


type NamePayload struct {
	Name string `json:"name" validate:"required"`
}

type JwtResponseFormat struct {
	AccessToken string `json:"access_token"`
}

type User struct {
	UserId     uuid.UUID   `db:"id"`
	UserName   string      `db:"username"`
	Name       string      `db:"name"`
	Password   string      `db:"password"`
	Role       string      `db:"role"`
	Created_at time.Time   `db:"created_at"`
	Updated_at time.Time   `db:"updated_at"`
	Deleted_at null.Time   `db:"deleted_at"`
	Created_by uuid.UUID   `db:"created_by"`
	Updated_by uuid.UUID   `db:"updated_by"`
	Deleted_by nuuid.NUUID `db:"deleted_by"`
}

func (u User) NewFromPayload(payload AuthPayload) (User, error) {
	userId, _ := uuid.NewV4()
	hashedPass, err := encrypt.HashPassword(payload.Password)
	if err != nil {
		return User{}, err
	}
	userRole := roles.GetStringFromRole(roles.GetRoleFromString(payload.Role))
	newUser := User{
		UserId:     userId,
		UserName:   payload.UserName,
		Name:       payload.Name,
		Password:   hashedPass,
		Role:       userRole,
		Created_at: time.Now().UTC(),
		Created_by: userId,
		Updated_at: time.Now().UTC(),
		Updated_by: userId,
	}
	return newUser, nil
}

func (j *JwtResponseFormat) MarshalJson() ([]byte, error) {
	return json.Marshal(j)
}

func (u *User) ValidatePassword(loginPass string) error {
	return encrypt.ComparePasswords(u.Password, loginPass)
}

func (u *User) UpdateName(payload NamePayload) {
	u.Name = payload.Name
}
	return encrypt.ComparePasswords(u.Password, loginPass)
}
