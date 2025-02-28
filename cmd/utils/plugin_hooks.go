package utils

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plugins"
)

func DefaultDataDir(pl *plugins.PluginLoader, path string) string {
	dataDirPath := ""
	fnList := pl.Lookup("SetDefaultDataDir", func(item interface{}) bool {
		_, ok := item.(func(string) string)
		return ok
	})
	for _, fni := range fnList {
		if fn, ok := fni.(func(string) string); ok {
			dataDirPath = fn(path)
		}
	}
	return dataDirPath
}

func pluginDefaultDataDir(path string) string {
	if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting DefaultDataDir, but default PluginLoader has not been initialized")
		return ""
	}
	return DefaultDataDir(plugins.DefaultPluginLoader, path)
}

func PluginSetBootStrapNodes(pl *plugins.PluginLoader) []string {
	var urls []string
	fnList := pl.Lookup("SetBootstrapNodes", func(item interface{}) bool {
		_, ok := item.(func() []string)
		return ok
	})
	for _, fni := range fnList {
		if fn, ok := fni.(func() []string); ok {
			urls = fn()
		}
	}
	return urls
}

func pluginSetBootstrapNodes() []string {
	if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting pluginSetBootStrapNodes, but default PluginLoader has not been initialized")
		return nil
	}
	return PluginSetBootStrapNodes(plugins.DefaultPluginLoader)
}

func PluginNetworkId(pl *plugins.PluginLoader) *uint64 {
	var networkId *uint64
	fnList := pl.Lookup("SetNetworkId", func(item interface{}) bool {
		_, ok := item.(func() *uint64)
		return ok
	})
	for _, fni := range fnList {
		if fn, ok := fni.(func() *uint64); ok {
			networkId = fn()
		}
	}
	return networkId
}

func pluginNetworkId() *uint64 {
	if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting pluginNetworkID, but default PluginLoader has not been initialized")
		return nil
	}
	return PluginNetworkId(plugins.DefaultPluginLoader)
}

func PluginETHDiscoveryURLs(pl *plugins.PluginLoader, mode bool) []string {
	var ethDiscoveryURLs []string
	fnList := pl.Lookup("SetETHDiscoveryURLs", func(item interface{}) bool {
		_, ok := item.(func(bool) []string)
		return ok
	})
	for _, fni := range fnList {
		if fn, ok := fni.(func(bool) []string); ok {
			ethDiscoveryURLs = fn(mode)
		}
	}
	return ethDiscoveryURLs
}

func pluginETHDiscoveryURLs(mode bool) []string {
	if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting pluginETHDiscoveryURLs, but default PluginLoader has not been initialized")
		return nil
	}
	return PluginETHDiscoveryURLs(plugins.DefaultPluginLoader, mode)
}

func PluginSnapDiscoveryURLs(pl *plugins.PluginLoader) []string {
	var snapDiscoveryURLs []string
	fnList := pl.Lookup("SetSnapDiscoveryURLs", func(item interface{}) bool {
		_, ok := item.(func() []string)
		return ok
	})
	for _, fni := range fnList {
		if fn, ok := fni.(func() []string); ok {
			snapDiscoveryURLs = fn()
		}
	}
	return snapDiscoveryURLs
}

func pluginSnapDiscoveryURLs() []string {
	if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting PluginSnapDiscoveryURLs, but default PluginLoader has not been initialized")
		return nil
	}
	return PluginSnapDiscoveryURLs(plugins.DefaultPluginLoader)
}

func PluginGenesisBlock(pl *plugins.PluginLoader) *core.Genesis {
	genesisJSON, ok := plugins.LookupOne[func() []byte](pl, "GenesisBlock")
	if !ok {
		return nil
	}
	var genesis core.Genesis
	if err := json.Unmarshal(genesisJSON(), &genesis); err != nil {
		log.Warn("Error unmarshalling genesis", "err", err)
		return nil
	}
	return &genesis
}

func pluginGenesisBlock() *core.Genesis {
	if plugins.DefaultPluginLoader == nil {
		log.Warn("Attempting PluginGenesisBlock, but default PluginLoader has not been initialized")
		return nil
	}
	return PluginGenesisBlock(plugins.DefaultPluginLoader)
}
