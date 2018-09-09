package main

import (
	_ "github.com/emmaly/elephant/wasm/log"
	"github.com/emmaly/elephant/wasm/util"
)

func main() {
	util.WaitUntilUnload()
}
