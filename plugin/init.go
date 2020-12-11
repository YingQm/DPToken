package plugin

import (
	_ "github.com/YingQm/DPToken/plugin/consensus/init" //register consensus init package
	_ "github.com/YingQm/DPToken/plugin/crypto/init"
	_ "github.com/YingQm/DPToken/plugin/dapp/init"
	_ "github.com/YingQm/DPToken/plugin/mempool/init"
	_ "github.com/YingQm/DPToken/plugin/p2p/init"
	_ "github.com/YingQm/DPToken/plugin/store/init"
)
