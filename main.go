// +build go1.8

package main

import (
	_ "github.com/33cn/dplatform/system"
	_ "github.com/YingQm/DPToken/plugin"

	"flag"
	"runtime/debug"

	"github.com/33cn/dplatform/util/cli"
)

var percent = flag.Int("p", 0, "SetGCPercent")

func main() {
	flag.Parse()
	if *percent < 0 || *percent > 100 {
		*percent = 0
	}
	if *percent > 0 {
		debug.SetGCPercent(*percent)
	}
	cli.RunDplatform("DPToken", DPToken)
}
