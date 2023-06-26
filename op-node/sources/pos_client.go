package sources

import (
	"context"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	//"net/http"
)

type PosClientConfig struct {
	PosURL                string
	PosBlockRefsCacheSize int
	DecSequencerHeight    int64
}

func PosClientDefaultConfig(config *rollup.Config) *PosClientConfig {
	// Cache 3/2 worth of sequencing window of receipts and txs
	// span := int(config.SeqWindowSize) * 3 / 2
	// fullSpan := span
	// if span > 1000 { // sanity cap. If a large sequencing window is configured, do not make the cache too large
	// 	span = 1000
	// }
	return &PosClientConfig{
		PosURL:             config.PosChainUrl,
		DecSequencerHeight: config.DecSequencerHeight,
	}
}

// L1Client provides typed bindings to retrieve L1 data from an RPC source,
// with optimized batch requests, cached results, and flag to not trust the RPC
// (i.e. to verify all returned contents against corresponding block hashes).
type PosClient struct {
	config *PosClientConfig
}

// NewL1Client wraps a RPC with bindings to fetch L1 data, while logging errors, tracking metrics (optional), and caching.
func NewPosClient(config *PosClientConfig) (*PosClient, error) {
	// ethClient, err := NewEthClient(client, log, metrics, &config.EthClientConfig)
	// if err != nil {
	// 	return nil, err
	// }

	return &PosClient{
		// EthClient:        ethClient,
		// l1BlockRefsCache: caching.NewLRUCache(metrics, "blockrefs", config.L1BlockRefsCacheSize),
		config: config,
	}, nil
}

type MetisSpanInfo struct {
	Height string         `json:"height"`
	Result MetisSpanResut `json:"result"`
}

type MetisSpanResut struct {
	EndBlock          int64         `json:"end_block"`
	MetisChainID      string        `json:"metis_chain_id"`
	SelectedProducers []Producer    `json:"selected_producers"`
	SpanID            int64         `json:"span_id"`
	StartBlock        int64         `json:"start_block"`
	ValidatorSet      ValidatorInfo `json:"validator_set"`
}

type Producer struct {
	ID          int64  `json:"ID"`
	Accum       int64  `json:"accum"`
	EndEpoch    int64  `json:"endEpoch"`
	Jailed      bool   `json:"jailed"`
	LastUpdated string `json:"last_updated"`
	Nonce       int64  `json:"nonce"`
	Power       int64  `json:"power"`
	PubKey      string `json:"pubKey"`
	Signer      string `json:"signer"`
	StartEpoch  int64  `json:"startEpoch"`
}

type ValidatorInfo struct {
	Proposer   Producer   `fmt:"proposer"`
	Validators []Producer `fmt:"validators"`
}

// L1BlockRefByLabel returns the [eth.L1BlockRef] for the given block label.
// Notice, we cannot cache a block reference by label because labels are not guaranteed to be unique.

//	curl -X 'GET' \
//	  'http://127.0.0.1:1317/metis/span/latest' \
//	  -H 'accept: application/json'
func (s *PosClient) GetSequencerByHeight(ctx context.Context, height int64) (common.Address, error) {

	switch height % 4 {
	case 0:
		return common.HexToAddress("0x690000000000000000000000000000000000000a"), nil
	case 1:
		return common.HexToAddress("0x690000000000000000000000000000000000000b"), nil
	case 2:
		return common.HexToAddress("0x690000000000000000000000000000000000000c"), nil
	default:
		return common.HexToAddress("0x690000000000000000000000000000000000000d"), nil
	}
	/*
		path := fmt.Sprintf("%v/metis/span/latest", s.config.PosURL)
		client := &http.Client{}
		resp, err := client.Get(path)
		if err != nil {
			return common.HexToAddress("0x0"), err
		}

		defer resp.Body.Close()
		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return common.HexToAddress("0x0"), err
		}

		var result MetisSpanResut

		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			return common.HexToAddress("0x0"), err
		}
		if result.StartBlock <= height && height <= result.EndBlock {
			return common.HexToAddress(result.SelectedProducers[0].Signer), nil
		}
		// current don't check height
		return common.HexToAddress(result.SelectedProducers[0].Signer), nil
	*/
}
