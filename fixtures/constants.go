package fixtures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	cid "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"

	th "github.com/filecoin-project/go-filecoin/testhelpers"
	"github.com/filecoin-project/go-filecoin/types"
)

// The file used to build these addresses can be found in:
// $GOPATH/src/github.com/filecoin-project/go-filecoin/fixtures/setup.json
//
// If said file is modified these addresses will need to change as well
// rebuild using
// TODO: move to build script
// https://github.com/filecoin-project/go-filecoin/issues/921
// cat ./fixtures/setup.json | ./gengen/gengen --json --keypath fixtures > fixtures/genesis.car 2> fixtures/gen.json

// TestAddresses is a list of pregenerated addresses.
var TestAddresses []string

// testKeys is a list of filenames, which contain the private keys of the pregenerated addresses.
var testKeys []string

// TestMiners is a list of pregenerated miner acccounts. They are owned by the matching TestAddress.
var TestMiners []string

type detailsStruct struct {
	Keys   map[string]*types.KeyInfo
	Miners []struct {
		Owner   string
		Address string
		Power   uint64
	}
	GenesisCid *cid.Cid
}

func init() {
	gopath, err := th.GetGoPath()
	if err != nil {
		panic(err)
	}

	detailspath := filepath.Join(gopath, "/src/github.com/filecoin-project/go-filecoin/fixtures/gen.json")
	detailsFile, err := os.Open(detailspath)
	if err != nil {
		fmt.Printf("Fixture data not found. Skipping fixture initialization: %s\n", err)
		return
	}
	defer func() {
		if err := detailsFile.Close(); err != nil {
			panic(err)
		}
	}()
	detailsFileBytes, err := ioutil.ReadAll(detailsFile)
	if err != nil {
		panic(err)
	}
	var details detailsStruct
	if err := json.Unmarshal(detailsFileBytes, &details); err != nil {
		panic(err)
	}

	var keys []string
	for key := range details.Keys {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	miners := details.Miners

	for _, key := range keys {
		info := details.Keys[key]
		addr, err := info.Address()
		if err != nil {
			panic(err)
		}
		TestAddresses = append(TestAddresses, addr.String())
		testKeys = append(testKeys, fmt.Sprintf("%s.key", key))
	}

	for _, miner := range miners {
		TestMiners = append(TestMiners, miner.Address)
	}
}

// KeyFilePaths returns the paths to the wallets of the testaddresses
func KeyFilePaths() []string {
	gopath, err := th.GetGoPath()
	if err != nil {
		panic(err)
	}
	folder := "/src/github.com/filecoin-project/go-filecoin/fixtures/"

	res := make([]string, len(testKeys))
	for i, k := range testKeys {
		res[i] = filepath.Join(gopath, folder, k)
	}

	return res
}

// Lab week cluster addrs
const (
	filecoinBootstrap0 string = "/dns4/cluster.kittyhawk.wtf/tcp/9000/ipfs/Qmd6xrWYHsxivfakYRy6MszTpuAiEoFbgE1LWw4EvwBpp4"
	filecoinBootstrap1 string = "/dns4/cluster.kittyhawk.wtf/tcp/9001/ipfs/QmXq6XEYeEmUzBFuuKbVEGgxEpVD4xbSkG2Rhek6zkFMp4"
	filecoinBootstrap2 string = "/dns4/cluster.kittyhawk.wtf/tcp/9002/ipfs/QmXhxqTKzBKHA5FcMuiKZv8YaMPwpbKGXHRVZcFB2DX9XY"
	filecoinBootstrap3 string = "/dns4/cluster.kittyhawk.wtf/tcp/9003/ipfs/QmZGDLdQLUTi7uYTNavKwCd7SBc5KMfxzWxAyvqRQvwuiV"
	filecoinBootstrap4 string = "/dns4/cluster.kittyhawk.wtf/tcp/9004/ipfs/QmZRnwmCjyNHgeNDiyT8mXRtGhP6uSzgHtrozc42crmVbg"
	filecoinBootstrap5 string = "/dns4/cluster.kittyhawk.wtf/tcp/9005/ipfs/QmQ8UoqmUUNQnCcYXNJPLbqT9aKpvfcoCWgxnV1aDGR3Ui"
	filecoinBootstrap6 string = "/dns4/cluster.kittyhawk.wtf/tcp/9006/ipfs/QmY7dK81dyant42kteuiKQNrtnZ5hRjEzdYn3NWwpBunXC"
	filecoinBootstrap7 string = "/dns4/cluster.kittyhawk.wtf/tcp/9007/ipfs/QmT3X2Z2ptrD3zpQc7wJe53ayAVr7sHA1wWMdT4T1d1u9J"
	filecoinBootstrap8 string = "/dns4/cluster.kittyhawk.wtf/tcp/9008/ipfs/QmVbkRxWyuEcWWtEiTj2NHWRt5NyFHB5KqL45agnHAXGNY"
	filecoinBootstrap9 string = "/dns4/cluster.kittyhawk.wtf/tcp/9009/ipfs/QmS5YAS5YsKSPvk63DkXjmhJMXXt7aSvtRfEtZEbhunLH2"
)

// LabWeekBootstrapAddrs are the dns multiaddrs for the nodes of the filecoin
// cluster running at lab week.
var LabWeekBootstrapAddrs = []string{
	filecoinBootstrap0, filecoinBootstrap1, filecoinBootstrap2,
	filecoinBootstrap3, filecoinBootstrap4, filecoinBootstrap5,
	filecoinBootstrap6, filecoinBootstrap7, filecoinBootstrap8,
	filecoinBootstrap9,
}