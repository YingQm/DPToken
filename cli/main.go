// +build go1.8

package main

import (
	_ "github.com/33cn/dplatform/system"
	_ "github.com/YingQm/DPToken/plugin"

	"github.com/33cn/dplatform/util/cli"
	"github.com/YingQm/DPToken/cli/buildflags"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8801"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "DPToken")
}
