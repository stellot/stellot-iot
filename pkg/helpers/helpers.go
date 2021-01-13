package helpers

import (
	"math/rand"
	"strings"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellot/stellot-iot/pkg/utils"
)

var networkPassphrase = utils.MustGetenv("NETWORK_PASSPHRASE")
var stellarCoreUrls = utils.MustGetenv("STELLAR_CORE_URLS")
var stellarCoreUrlsSlice = strings.Split(stellarCoreUrls, " ")
var devicesSecrets = utils.MustGetenv("DEVICES_SECRETS")
var devicesSecretsSlice = strings.Split(devicesSecrets, " ")
var horizonServerUrls = utils.MustGetenv("HORIZON_SERVER_URLS")
var horizonServerUrlsSlice = strings.Split(horizonServerUrls, " ")
var horizonServers = CreateHorizonServers()
var masterKp, _ = keypair.FromRawSeed(network.ID(networkPassphrase))

func DevicesKeypairs() []*keypair.Full {
	keypairs := make([]*keypair.Full, len(devicesSecretsSlice))
	for i, v := range devicesSecretsSlice {
		key, err := keypair.ParseFull(v)
		if err != nil {
			panic(err)
		}
		keypairs[i] = key
	}
	return keypairs
}

func CreateHorizonServers() []*horizonclient.Client {
	horizons := make([]*horizonclient.Client, len(horizonServerUrlsSlice))
	for i, v := range horizonServerUrlsSlice {
		horizons[i] = &horizonclient.Client{
			HorizonURL: v,
		}
	}
	return horizons
}

func RandomHorizon() *horizonclient.Client {
	return horizonServers[rand.Intn(len(horizonServers))]
}

func RandomStellarCoreUrl() string {
	return stellarCoreUrlsSlice[rand.Intn(len(stellarCoreUrlsSlice))]
}

func LoadMasterAccount() (*horizon.Account, error) {
	return loadAccount(masterKp.Address())
}

func loadAccount(accountId string) (*horizon.Account, error) {
	accReq := horizonclient.AccountRequest{AccountID: accountId}
	masterAccount, err := RandomHorizon().AccountDetail(accReq)
	if err != nil {
		return nil, err
	}
	return &masterAccount, nil
}
