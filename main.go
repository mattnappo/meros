package main

import (
	"fmt"

	"github.com/xoreo/meros/core"
	"github.com/xoreo/meros/types"
)

func main() {
	bytes := []byte("Hello, world!")
	sizes, err := types.CalculateShardSizes(bytes, 5)
	if err != nil {
		panic(err)
	}

	splitBytes, err := core.SplitBytes(bytes, sizes)
	if err != nil {
		panic(err)
	}

	fmt.Print(splitBytes)

}
