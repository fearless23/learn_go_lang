package utils

import (
	"learn-go-lang/basics"

	"github.com/google/uuid"
)

func Uuid() {
	uuid := uuid.NewString()
	basics.Log("uuid", uuid)
}
