package basics

import (
	"fmt"
)

// Roles for user
const (
	admin   = 1 << iota // 00001
	player              // 00010
	finance             // 00100
	editor              // 01000
	speaker             // 10000
)

func EnumCPart2() {
	var user1Roles = admin | editor
	var user2Roles = finance | speaker
	// above puts 1`s at right places
	fmt.Printf("%v, %b\n", user1Roles, user1Roles)
	fmt.Printf("%v, %b\n", user2Roles, user2Roles)

	// isUser1 admin
	var isUser1Admin = user1Roles&admin == admin
	// when we and with a role... only 1 left at role
	fmt.Printf("%v\n", isUser1Admin)

	var isUser2Admin = user2Roles&admin == admin
	fmt.Printf("%v\n", isUser2Admin)
}
