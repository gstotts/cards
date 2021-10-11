# Cards / Deck Module for Go

### Examples

#### Create a Deck
```golang
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