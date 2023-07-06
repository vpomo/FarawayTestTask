package checkchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
)

type CollectionItem struct {
	ColAddress common.Address
	Name       string
	Symbol     string
}

type CollectionItems struct {
	Items []CollectionItem
}

type TokenMintedItem struct {
	ColAddress common.Address
	Owner      common.Address
	TokenId    *big.Int
	TokenUri   string
}

type TokenMintedItems struct {
	Items []TokenMintedItem
}

const providerUrl = "https://polygon-mumbai.g.alchemy.com/v2/TEu1hzP2IRfBKZ27AYRy6nfgSK_7CQ1N"
const contractAddress = "0x652ea34de1926fc668625a4eb68a80848faa78ed"

func (box *CollectionItems) AddICollectionItem(item CollectionItem) []CollectionItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func (box *TokenMintedItems) AddITokenMintedItem(item TokenMintedItem) []TokenMintedItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func GetCollectionsCreated() []CollectionItem {
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address := common.HexToAddress(contractAddress)
	instance, err := NewCheckchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	filterOpts := &bind.FilterOpts{Context: ctx, Start: 9000000, End: nil}

	iter, err := instance.FilterCollectionCreated(filterOpts)
	if err != nil {
		log.Fatal("error making event request: %s\n", err)
		os.Exit(1)
	}
	items := CollectionItems{}

	for iter.Next() {
		event := iter.Event
		item := CollectionItem{ColAddress: event.Collection, Name: event.Name, Symbol: event.Symbol}
		items.AddICollectionItem(item)
	}
	fmt.Println(items)
	return items.Items
}

func GetTokenMinted(addr string) []TokenMintedItem {
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address := common.HexToAddress(contractAddress)
	instance, err := NewCheckchain(address, client)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	filterOpts := &bind.FilterOpts{Context: ctx, Start: 9000000, End: nil}

	colAddress := common.HexToAddress(addr)
	iter, err := instance.FilterTokenMinted(filterOpts, colAddress)
	if err != nil {
		log.Fatal("error making event request: %s\n", err)
		os.Exit(1)
	}
	items := TokenMintedItems{}

	for iter.Next() {
		event := iter.Event
		item := TokenMintedItem{ColAddress: event.Collection, Owner: event.Owner, TokenId: event.TokenId, TokenUri: event.TokenUri}
		items.AddITokenMintedItem(item)
	}
	fmt.Println(items)
	return items.Items
}
