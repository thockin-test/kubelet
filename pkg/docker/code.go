package docker

import (
	"fmt"

	"github.com/thockin-test/util/foobar"
)

func Docker() {
	fmt.Printf("This is very %s\n", foobar.FooBar())
}
