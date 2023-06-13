// Workaround for plugin version mismatch issue: https://github.com/golang/go/issues/31354
// By wrapping geth as a dependency, its module path matches that in the plugin.
package main

import "github.com/ethereum/go-ethereum/cmd/geth"

func main() { geth.Main() }
