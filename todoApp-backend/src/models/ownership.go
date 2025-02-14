package models

type ownership interface {
	User | Organization
}
