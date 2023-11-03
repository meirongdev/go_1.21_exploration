package b

import (
	"fmt"

	_ "github.com/meirongdev/go121expl/pkg_init/c"
	_ "github.com/meirongdev/go121expl/pkg_init/e"
)

func init() {
	fmt.Println("b init")
}
