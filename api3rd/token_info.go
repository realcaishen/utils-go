package api3rd

import (
	"context"
	"fmt"

	"github.com/realcaishen/utils-go/loader"
	"github.com/realcaishen/utils-go/network"
	"github.com/realcaishen/utils-go/rpc"
	"github.com/shopspring/decimal"
)

type TokenDetail struct {
	ChainName         string      `json:"chain_name"`
	Address           string      `json:"address"`
	CirculatingSupply string      `json:"circulating_supply"`
	TotalSupply       string      `json:"total_supply"`
	Mcap              float64     `json:"mcap"`
	Fdv               float64     `json:"fdv"`
	Holders           int         `json:"holders"`
	Transactions      int         `json:"transactions"`
	Symbol            string      `json:"symbol"`
	TokenName         string      `json:"token_name"`
	Decimals          int         `json:"decimals"`
	Price             float64     `json:"price"`
	Liquidity         float64     `json:"liquidity"`
	Volume24h         float64     `json:"volume24h"`
	Pricechg24h       float64     `json:"pricechg24h"`
	Pricechg6h        float64     `json:"pricechg6h"`
	CreationTime      string      `json:"creation_time"`
	SocialInfos       SocialInfos `json:"social_infos"`
	TopHolders        TopHolders  `json:"top_holders"`
}

func (item *TokenDetail) FillMcapFdv() {
	price := decimal.NewFromFloat(item.Price)

	if item.Mcap <= 0 {
		if item.CirculatingSupply != "" && item.CirculatingSupply != "0" {
			supply, err := decimal.NewFromString(item.CirculatingSupply)
			if err == nil {
				item.Mcap = supply.Mul(price).InexactFloat64()
			}
		} else if item.Fdv > 0 {
			item.Mcap = item.Fdv
		} else if item.TotalSupply != "" && item.TotalSupply != "0" {
			supply, err := decimal.NewFromString(item.TotalSupply)
			if err == nil {
				item.Mcap = supply.Mul(price).InexactFloat64()
			}
		}
	}
	if item.Mcap <= 0 {
		item.Mcap = 0
	}

	if item.Fdv <= 0 {
		if item.TotalSupply != "" && item.TotalSupply != "0" {
			supply, err := decimal.NewFromString(item.TotalSupply)
			if err == nil {
				item.Fdv = supply.Mul(price).InexactFloat64()
			}
		} else if item.Mcap > 0 {
			item.Fdv = item.Mcap
		}
	}
	if item.Fdv <= 0 {
		item.Fdv = 0
	}
}

type SocialInfos struct {
	Email     string `json:"email"`
	Bitbucket string `json:"bitbucket"`
	Discord   string `json:"discord"`
	Facebook  string `json:"facebook"`
	Github    string `json:"github"`
	Instagram string `json:"instagram"`
	Linkedin  string `json:"linkedin"`
	Medium    string `json:"medium"`
	Reddit    string `json:"reddit"`
	Telegram  string `json:"telegram"`
	Tiktok    string `json:"tiktok"`
	Twitter   string `json:"twitter"`
	Website   string `json:"website"`
	Youtube   string `json:"youtube"`
}

type TopHolders struct {
	Items []*TopHolderItem `json:"items"`
}
type TopHolderItem struct {
	Amount       decimal.Decimal `json:"amount"`
	Decimals     int32           `json:"decimals"`
	Mint         string          `json:"mint"`
	Owner        string          `json:"owner"`
	TokenAccount string          `json:"token_account"`
	UiAmount     decimal.Decimal `json:"ui_amount"`
}

type TokenDetails struct {
	Infos map[string]TokenDetail `json:"infos"`
}

func GetTokenDetails(ctx context.Context, serverUrl, chainName string, addresses []string, audit bool, detail bool, price bool, pool bool) (map[string]TokenDetail, error) {
	param := map[string]interface{}{
		"ChainName": chainName,
		"Tokens":    addresses,
		"Audit":     audit,
		"Detail":    detail,
		"Price":     price,
		"Pool":      pool,
	}
	var resp TokenDetails
	err := network.Request(fmt.Sprintf("%s/info/tokens", serverUrl), param, &resp)
	if err != nil {
		return make(map[string]TokenDetail), err
	}

	for _, value := range resp.Infos {
		value.FillMcapFdv()
	}

	return resp.Infos, nil
}

func GetTokenInfoBy(mgr *loader.ChainInfoManager, chainName string, token string) (*loader.TokenInfo, error) {
	chainInfo, ok := mgr.GetChainInfoByName(chainName)
	if !ok {
		return nil, fmt.Errorf("no chain info for %s", chainName)
	}
	return GetTokenInfo(chainInfo, token)
}

func GetTokenInfo(chainInfo *loader.ChainInfo, token string) (*loader.TokenInfo, error) {
	r, err := rpc.GetRpc(chainInfo, nil)
	if err != nil {
		return nil, err
	}

	ti, err := r.GetTokenInfo(context.TODO(), token)
	if err != nil {
		return nil, err
	}

	return ti, nil
}
