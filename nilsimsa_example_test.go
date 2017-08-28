package nilsimsa_test

import (
	"fmt"

	"github.com/chikamim/nilsimsa"
)

func ExampleHex() {
	fmt.Println(nilsimsa.HexSum([]byte("Hello nilsimsa!")))
	// Output: 436119240183882801210e002e1cb00122a20d11b4268ab8001a51190c08084b
}

func ExampleHexScore() {
	h1 := nilsimsa.HexSum([]byte("Hello nilsimsa!"))
	h2 := nilsimsa.HexSum([]byte("Hello world!"))
	h3 := nilsimsa.HexSum([]byte("Nobody Expects the Spanish Inquisition"))

	fmt.Println(nilsimsa.DiffHexScore(h1, h1))
	fmt.Println(nilsimsa.DiffHexScore(h1, h2))
	fmt.Println(nilsimsa.DiffHexScore(h1, h3))
	// Output:
	// 1
	// 0.4094488188976378
	// 0.14960629921259844
}
