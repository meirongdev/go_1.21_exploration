package c

import (
	"fmt"

	_ "github.com/meirongdev/go121expl/pkg_init/e"
	_ "github.com/meirongdev/go121expl/pkg_init/f"
)

func init() {
	fmt.Println("c init")
}
