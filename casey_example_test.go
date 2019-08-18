package casey_test

import (
	"fmt"

	"github.com/vporoshok/casey"
)

func ExampleCamel() {
	fmt.Println(casey.Camel("Some42TextWithABBR").SNAKE())
	// Output: SOME42_TEXT_WITH_ABBR
}
