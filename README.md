# Cards / Deck Module for Go

<details><summary>Create a Deck of Cards</summary>

```go
package main

import (
	"fmt"

	"github.com/gstotts/cards"
)

func main() {
	deck := cards.Create_Deck()
	fmt.Println(deck)
}
```
</details>