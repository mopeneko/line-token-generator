line-token-generator
====

Generate X-Line-Access header

## Usage

### Android

```
package main

import (
	"fmt"
	"github.com/mopeneko/line-token-generator/android"
)

func main() {
	var authKey string

	// Read authKey from stdin
	fmt.Scan(&authKey)

	accessToken := android.GenerateToken(authKey)
	fmt.Println(accessToken)
}
```
