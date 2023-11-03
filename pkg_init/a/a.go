package a

import (
	"fmt"

	_ "github.com/meirongdev/go121expl/pkg_init/c"
	_ "github.com/meirongdev/go121expl/pkg_init/d"
)

func init() {
	fmt.Println("a init")
}
