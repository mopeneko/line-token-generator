package main

import (
	"./android"
	"fmt"
)

func main() {
	var authKey string

	// Read authKey from stdin
	fmt.Scan(&authKey)

	accessToken := android.GenerateToken(authKey)
	fmt.Println(accessToken)
}
