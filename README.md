# Usage

import go pkg as bellow:

```go
package main

import (
	"fmt"
	user "github.com/rexitorg/sample_go_pkg/domain/user"
)

func main() {
	u := user.NewUser("konchan")
	fmt.Println(u.Name)
}
```

