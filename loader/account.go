package loader

import (
	"database/sql"
	"strings"
	"sync"

	"github.com/realcaishen/utils-go/alert"
)

type Account struct {
	Id          int64
	ChainInfoId int64
	Address     string
}

type AccountManager struct {
	idAccounts         map[int64]*Account
	addressCidAccounts map[string]map[int64]*Account
	cidAddressAccounts map[int64]map[string]*Account
	db                 *sql.DB
	alerter            alert.Alerter
	mutex              *sync.RWMutex
}

func NewAccountManager(db *sql.DB, alerter alert.Alerter) *AccountManager {
	return &AccountManager{
		idAccounts:         make(map[int64]*Account),
		addressCidAccounts: make(map[string]map[int64]*Account),
		cidAddressAccounts: make(map[int64]map[string]*Account),
		db:                 db,
		alerter:            alerter,
		mutex:              &sync.RWMutex{},
	}
}

func (mgr *AccountManager) GetAccountById(id int64) (*Account, bool) {
	mgr.mutex.RLock()
	acc, ok := mgr.idAccounts[id]
	mgr.mutex.RUnlock()
	return acc, ok
}

func (mgr *AccountManager) HasAddress(address string) bool {
	mgr.mutex.RLock()
	_, ok := mgr.addressCidAccounts[strings.ToLower(strings.TrimSpace(address))]
	mgr.mutex.RUnlock()
	return ok
}

func (mgr *AccountManager) GetAddresses(cid int64) []string {
	addrs := make([]string, 0)
	mgr.mutex.RLock()
	accs, ok := mgr.cidAddressAccounts[cid]
	if ok {
		for _, acc := range accs {
			addrs = append(addrs, acc.Address)
		}
	}
	mgr.mutex.RUnlock()
	return addrs
}

func (mgr *AccountManager) GetAccountByAddressCid(address string, cid int64) (*Account, bool) {
	mgr.mutex.RLock()
	defer mgr.mutex.RUnlock()
	accs, ok := mgr.addressCidAccounts[strings.ToLower(strings.TrimSpace(address))]
	if ok {
		acc, ok := accs[cid]
		return acc, ok
	}
	return nil, false
}

func (mgr *AccountManager) LoadAllAccounts() {
	// Query the database to select only id and name fields
	rows, err := mgr.db.Query("SELECT id, chain_id, address FROM t_account")

	if err != nil || rows == nil {
		mgr.alerter.AlertText("select t_account error", err)
		return
	}

	defer rows.Close()

	idAccounts := make(map[int64]*Account)
	addressCidAccounts := make(map[string]map[int64]*Account)
	cidAddressAccounts := make(map[int64]map[string]*Account)
	counter := 0

	// Iterate over the result set
	for rows.Next() {
		var acc Account
		if err := rows.Scan(&acc.Id, &acc.ChainInfoId, &acc.Address); err != nil {
			mgr.alerter.AlertText("scan t_account row error", err)
		} else {
			acc.Address = strings.TrimSpace(acc.Address)

			idAccounts[acc.Id] = &acc
			lowerAddr := strings.ToLower(acc.Address)

			accs, ok := addressCidAccounts[lowerAddr]
			if !ok {
				accs = make(map[int64]*Account)
				addressCidAccounts[lowerAddr] = accs
			}
			accs[acc.ChainInfoId] = &acc

			addraccs, ok := cidAddressAccounts[acc.ChainInfoId]
			if !ok {
				addraccs = make(map[string]*Account)
				cidAddressAccounts[acc.ChainInfoId] = addraccs
			}
			addraccs[lowerAddr] = &acc

			counter++
		}
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		mgr.alerter.AlertText("get next t_account row error", err)
		return
	}

	mgr.mutex.Lock()
	mgr.idAccounts = idAccounts
	mgr.addressCidAccounts = addressCidAccounts
	mgr.cidAddressAccounts = cidAddressAccounts
	mgr.mutex.Unlock()
}
