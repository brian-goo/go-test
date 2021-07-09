package gopkgs

import (
	"fmt"
	"gotest/gopkgs/lambdaio"
)

func Handler() {
	fmt.Println(lambdaio.GetPostData())
}
