package test_o_test

import (
	"testing"

	"github.com/meirongdev/go121expl/test_o"
)

// Generate test binary named test for pkg github.com/meirongdev/go121expl/test_o
// go test -c -o test github.com/meirongdev/go121expl/test_o
//  ./test -test.run TestHello

func TestHello(t *testing.T) {
	test_o.Hello()
}
