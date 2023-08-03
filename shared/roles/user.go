package roles

import "strings"

type Role int

const (
	Student Role = iota
	Teacher
)

func GetRoleFromString(s string) Role {
	switch strings.ToLower(s) {
	case "student":
		return Student
	case "teacher":
		return Teacher
	default:
		return -1
	}
}

func GetStringFromRole(r Role) string {
	switch r {
	case Student:
		return "student"
	case Teacher:
		return "teacher"
	default:
		return "student"
	}
}
