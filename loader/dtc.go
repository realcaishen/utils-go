package loader

import (
	"database/sql"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/realcaishen/utils-go/alert"
	"github.com/realcaishen/utils-go/util"
)

type Dtc struct {
	TokenName     string
	FromChainName string
	ToChainName   string
	DtcLv1        float64
	DtcLv2        float64
	DtcLv3        float64
	DtcLv4        float64
	AmountLv1     float64
	AmountLv2     float64
	AmountLv3     float64
	AmountLv4     float64

	DtcLv1Str    string
	DtcLv2Str    string
	DtcLv3Str    string
	DtcLv4Str    string
	AmountLv1Str string
	AmountLv2Str string
	AmountLv3Str string
	AmountLv4Str string
}

type DtcManager struct {
	tokenFromToDtcs map[string]map[string]map[string]*Dtc

	db      *sql.DB
	alerter alert.Alerter
	mutex   *sync.RWMutex
}

func NewDtcManager(db *sql.DB, alerter alert.Alerter) *DtcManager {
	return &DtcManager{
		tokenFromToDtcs: make(map[string]map[string]map[string]*Dtc),

		db:      db,
		alerter: alerter,
		mutex:   &sync.RWMutex{},
	}
}

func (mgr *DtcManager) GetDtcs() map[string]map[string]map[string]*Dtc {
	mgr.mutex.RLock()
	defer mgr.mutex.RUnlock()
	return mgr.tokenFromToDtcs
}

func (mgr *DtcManager) GetDtc(token string, from string, to string) (*Dtc, bool) {
	mgr.mutex.RLock()
	defer mgr.mutex.RUnlock()
	ftInfos, ok := mgr.tokenFromToDtcs[strings.ToLower(strings.TrimSpace(token))]
	if ok {
		infos, ok := ftInfos[strings.ToLower(strings.TrimSpace(from))]
		if ok {
			info, ok := infos[strings.ToLower(strings.TrimSpace(to))]
			return info, ok
		}

	}
	return nil, false
}

func (mgr *DtcManager) LoadAllDtc() {
	// Query the database to select only id and name fields
	rows, err := mgr.db.Query("SELECT token_name, from_chain, to_chain, dtc_lv1, dtc_lv2, dtc_lv3, dtc_lv4, amount_lv1, amount_lv2, amount_lv3, amount_lv4 FROM t_dynamic_dtc")

	if err != nil || rows == nil {
		mgr.alerter.AlertText("select t_dynamic_dtc error", err)
		return
	}

	defer rows.Close()

	tokenFromToDtcs := make(map[string]map[string]map[string]*Dtc)
	counter := 0

	// Iterate over the result set
	for rows.Next() {
		var dtc Dtc

		if err := rows.Scan(&dtc.TokenName, &dtc.FromChainName, &dtc.ToChainName, &dtc.DtcLv1Str, &dtc.DtcLv2Str, &dtc.DtcLv3Str, &dtc.DtcLv4Str, &dtc.AmountLv1Str, &dtc.AmountLv2Str, &dtc.AmountLv3Str, &dtc.AmountLv4Str); err != nil {
			mgr.alerter.AlertText("scan t_dynamic_dtc row error", err)
		} else {
			dtc.FromChainName = strings.TrimSpace(dtc.FromChainName)
			dtc.ToChainName = strings.TrimSpace(dtc.ToChainName)
			dtc.TokenName = strings.TrimSpace(dtc.TokenName)

			dtc1, err := strconv.ParseFloat(dtc.DtcLv1Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc dtc1 not float", err)
				continue
			}
			dtc2, err := strconv.ParseFloat(dtc.DtcLv2Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc dtc2 not float", err)
				continue
			}
			dtc3, err := strconv.ParseFloat(dtc.DtcLv3Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc dtc3 not float", err)
				continue
			}
			dtc4, err := strconv.ParseFloat(dtc.DtcLv4Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc dtc4 not float", err)
				continue
			}

			amount1, err := strconv.ParseFloat(dtc.AmountLv1Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc amount1 not float", err)
				continue
			}
			amount2, err := strconv.ParseFloat(dtc.AmountLv2Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc amount2 not float", err)
				continue
			}
			amount3, err := strconv.ParseFloat(dtc.AmountLv3Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc amount3 not float", err)
				continue
			}
			amount4, err := strconv.ParseFloat(dtc.AmountLv4Str, 64)
			if err != nil {
				mgr.alerter.AlertText("t_dynamic_dtc amount4 not float", err)
				continue
			}

			dtc.DtcLv1 = dtc1
			dtc.DtcLv2 = dtc2
			dtc.DtcLv3 = dtc3
			dtc.DtcLv4 = dtc4
			dtc.AmountLv1 = amount1
			dtc.AmountLv2 = amount2
			dtc.AmountLv3 = amount3
			dtc.AmountLv4 = amount4

			ftInfos, ok := tokenFromToDtcs[strings.ToLower(dtc.TokenName)]
			if !ok {
				ftInfos = make(map[string]map[string]*Dtc)
				tokenFromToDtcs[strings.ToLower(dtc.TokenName)] = ftInfos
			}
			infos, ok := ftInfos[strings.ToLower(dtc.FromChainName)]
			if !ok {
				infos = make(map[string]*Dtc)
				ftInfos[strings.ToLower(dtc.FromChainName)] = infos
			}
			infos[strings.ToLower(dtc.ToChainName)] = &dtc

			counter++
		}
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		mgr.alerter.AlertText("get next t_dynamic_dtc row error", err)
		return
	}

	mgr.mutex.Lock()
	mgr.tokenFromToDtcs = tokenFromToDtcs
	mgr.mutex.Unlock()
}

func (mgr *DtcManager) GetIncludedDtc(tokenName string, fromChainName string, toChainName string, value float64) (float64, string, bool) {
	dtc, ok := mgr.GetDtc(tokenName, fromChainName, toChainName)
	if !ok {
		return 0, "", false
	}

	if value > (dtc.AmountLv3 + dtc.DtcLv3) {
		return dtc.DtcLv4, dtc.DtcLv4Str, true
	} else if value > (dtc.AmountLv2 + dtc.DtcLv2) {
		return dtc.DtcLv3, dtc.DtcLv3Str, true
	} else if value > (dtc.AmountLv1 + dtc.DtcLv1) {
		return dtc.DtcLv2, dtc.DtcLv2Str, true
	} else {
		return dtc.DtcLv1, dtc.DtcLv1Str, true
	}

}

func (mgr *DtcManager) GetDtcToInclude(tokenName string, fromChainName string, toChainName string, value float64) (float64, string, bool) {
	dtc, ok := mgr.GetDtc(tokenName, fromChainName, toChainName)
	if !ok {
		return 0, "", false
	}

	if value > dtc.AmountLv3 {
		return dtc.DtcLv4, dtc.DtcLv4Str, true
	} else if value > dtc.AmountLv2 {
		return dtc.DtcLv3, dtc.DtcLv3Str, true
	} else if value > dtc.AmountLv1 {
		return dtc.DtcLv2, dtc.DtcLv2Str, true
	} else {
		return dtc.DtcLv1, dtc.DtcLv1Str, true
	}

}

func (mgr *DtcManager) FromUiString(amount string, dtc string, decimals int32) *big.Int {
	value := big.NewInt(0)
	if amount != "" {
		amountValue, err := util.FromUiString(amount, decimals)
		if err == nil {
			value.Add(value, amountValue)
		}
	}

	if dtc != "" {
		dtcValue, err := util.FromUiString(dtc, decimals)
		if err == nil {
			value.Add(value, dtcValue)
		}
	}
	return value
}

func (mgr *DtcManager) GetIncludedDtcBigInt(tokenName string, fromChainName string, toChainName string, value *big.Int, decimals int32) (*big.Int, bool) {
	dtc, ok := mgr.GetDtc(tokenName, fromChainName, toChainName)
	if !ok {
		return nil, false
	}

	if value.Cmp(mgr.FromUiString(dtc.AmountLv1Str, dtc.DtcLv1Str, decimals)) <= 0 {
		return mgr.FromUiString("", dtc.DtcLv1Str, decimals), true
	} else if value.Cmp(mgr.FromUiString(dtc.AmountLv2Str, dtc.DtcLv2Str, decimals)) <= 0 {
		return mgr.FromUiString("", dtc.DtcLv2Str, decimals), true
	} else if value.Cmp(mgr.FromUiString(dtc.AmountLv3Str, dtc.DtcLv3Str, decimals)) <= 0 {
		return mgr.FromUiString("", dtc.DtcLv3Str, decimals), true
	} else {
		return mgr.FromUiString("", dtc.DtcLv4Str, decimals), true
	}
}

func (mgr *DtcManager) GetDtcToIncludeBigInt(tokenName string, fromChainName string, toChainName string, value *big.Int, decimals int32) (*big.Int, bool) {
	dtc, ok := mgr.GetDtc(tokenName, fromChainName, toChainName)
	if !ok {
		return nil, false
	}

	if value.Cmp(mgr.FromUiString(dtc.AmountLv1Str, "", decimals)) <= 0 {
		return mgr.FromUiString("", dtc.DtcLv1Str, decimals), true
	} else if value.Cmp(mgr.FromUiString(dtc.AmountLv2Str, "", decimals)) <= 0 {
		return mgr.FromUiString("", dtc.DtcLv2Str, decimals), true
	} else if value.Cmp(mgr.FromUiString(dtc.AmountLv3Str, "", decimals)) <= 0 {
		return mgr.FromUiString("", dtc.DtcLv3Str, decimals), true
	} else {
		return mgr.FromUiString("", dtc.DtcLv4Str, decimals), true
	}
}

func (mgr *DtcManager) GetMinValueIncludeGasFee(tokenName string, fromChainName string, toChainName string, decimals int32) (string, bool) {
	dtc, ok := mgr.GetDtc(tokenName, fromChainName, toChainName)
	if !ok {
		return "", false
	}

	vals := []string{dtc.DtcLv1Str, dtc.DtcLv2Str, dtc.DtcLv3Str, dtc.DtcLv4Str}
	for _, val := range vals {
		value := mgr.FromUiString("", val, decimals)
		includedDtc, ok := mgr.GetIncludedDtcBigInt(tokenName, fromChainName, toChainName, value, decimals)
		if ok {
			if includedDtc.Cmp(value) <= 0 {
				return val, true
			}
		}
	}
	return "", false
}
