package ethconfig

import (
	"github.com/ethereum/go-ethereum/log"
	// "github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/plugins"
	"github.com/ethereum/go-ethereum/plugins/wrappers"
	wengine "github.com/ethereum/go-ethereum/plugins/wrappers/engine"
	"github.com/ethereum/go-ethereum/plugins/wrappers/backendwrapper"
	// "github.com/ethereum/go-ethereum/rpc"
	"github.com/openrelayxyz/plugeth-utils/core"
	"github.com/openrelayxyz/plugeth-utils/restricted"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	// "github.com/ethereum/go-ethereum/consensus/beacon"
	// "github.com/ethereum/go-ethereum/consensus/clique"
	// "github.com/ethereum/go-ethereum/consensus/ethash"
	// "github.com/ethereum/go-ethereum/core"
	// "github.com/ethereum/go-ethereum/core/txpool"
	// "github.com/ethereum/go-ethereum/eth/downloader"
	// "github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethdb"
	// "github.com/ethereum/go-ethereum/log"
	// "github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	// "github.com/ethereum/go-ethereum/params"

	pconsensus "github.com/openrelayxyz/plugeth-utils/restricted/consensus"
	// pparams "github.com/openrelayxyz/plugeth-utils/restricted/params"
)



func PluginGetEngine(pl *plugins.PluginLoader, stack *node.Node, notify []string, noverify bool, db ethdb.Database) consensus.Engine {
	fnList := pl.Lookup("CreateEngine", func(item interface{}) bool {
		_, ok := item.(func(core.Node, []string, bool, restricted.Database) pconsensus.Engine)
		return ok
	})
	for _, fni := range fnList {
		if fn, ok := fni.(func(core.Node, []string, bool, restricted.Database) pconsensus.Engine); ok { 
			if engine := fn(wrappers.NewNode(stack), notify, noverify, backendwrapper.NewDB(db)); engine != nil {
				wrappedEngine := wengine.NewWrappedEngine(engine)
				return wrappedEngine
			}
			
		}
	}
	return nil
}

func pluginGetEngine(stack *node.Node, notify []string, noverify bool, db ethdb.Database) consensus.Engine {
	if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting GetEngine, but default PluginLoader has not been initialized")
		return nil
	}
	return PluginGetEngine(plugins.DefaultPluginLoader, stack, notify, noverify, db)
}