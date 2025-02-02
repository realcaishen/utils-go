package loader

import (
	"database/sql"
	"github.com/realcaishen/utils-go/log"
)

type MakerAddressGroupPO struct {
	Id        int64
	GroupName string
	Env       string
}

type MakerAddressPO struct {
	Id      int64
	GroupId int64
	Backend Backend
	Address string
}

type MakerAddress struct {
	GroupId           int64
	GroupName         string
	Env               string
	Addresses         []*MakerAddressPO
	SecurityAddresses []*MakerAddressPO
}

type MakerAddressManager struct {
	groupIdAddress        map[int64]*MakerAddress
	envGroup              map[string][]*MakerAddress
	backendAddressToGroup map[Backend]map[string]int64

	db *sql.DB
}

func NewMakerAddressManager(db *sql.DB) *MakerAddressManager {
	return &MakerAddressManager{
		groupIdAddress:        make(map[int64]*MakerAddress),
		envGroup:              make(map[string][]*MakerAddress),
		backendAddressToGroup: make(map[Backend]map[string]int64),
		db:                    db,
	}
}

func (mgr *MakerAddressManager) LoadAllMakerAddresses() {
	// Query the database for all maker address groups
	groupRows, err := mgr.db.Query("SELECT id, group_name, env FROM t_maker_address_groups")
	if err != nil || groupRows == nil {
		log.Errorf("select maker_address_groups error: %v", err)
		return
	}
	defer groupRows.Close()

	groups := make(map[int64]*MakerAddress)
	for groupRows.Next() {
		var group MakerAddressGroupPO
		if err = groupRows.Scan(&group.Id, &group.GroupName, &group.Env); err != nil {
			log.Errorf("scan maker_address_groups row error: %v", err)
			continue
		}

		makerAddress := &MakerAddress{
			GroupId:   group.Id,
			GroupName: group.GroupName,
			Env:       group.Env,
			Addresses: []*MakerAddressPO{},
		}
		groups[group.Id] = makerAddress
	}

	// Check for errors from iterating over rows
	if err = groupRows.Err(); err != nil {
		log.Errorf("get next maker_address_groups row error: %v", err)
		return
	}

	// Query the database for all maker addresses
	addressRows, err := mgr.db.Query("SELECT id, group_id, backend, address FROM t_maker_addresses")
	if err != nil || addressRows == nil {
		log.Errorf("select maker_addresses error: %v", err)
		return
	}
	defer addressRows.Close()

	backendAddressToGroup := make(map[Backend]map[string]int64)

	for addressRows.Next() {
		var address MakerAddressPO
		if err = addressRows.Scan(&address.Id, &address.GroupId, &address.Backend, &address.Address); err != nil {
			log.Errorf("scan maker_addresses row error: %v", err)
			continue
		}

		if group, ok := groups[address.GroupId]; ok {
			group.Addresses = append(group.Addresses, &address)
		}

		// Populate tempBackendAddressToGroup mapping
		if _, ok := backendAddressToGroup[address.Backend]; !ok {
			backendAddressToGroup[address.Backend] = make(map[string]int64)
		}
		backendAddressToGroup[address.Backend][address.Address] = address.GroupId
	}

	if err = addressRows.Err(); err != nil {
		log.Errorf("get next maker_addresses row error: %v", err)
		return
	}

	// Query the database for all security addresses
	securityAddressRows, err := mgr.db.Query("SELECT id, group_id, backend, address FROM t_security_addresses")
	if err != nil || securityAddressRows == nil {
		log.Errorf("select security_addresses error: %v", err)
		return
	}
	defer securityAddressRows.Close()

	for securityAddressRows.Next() {
		var securityAddress MakerAddressPO
		if err = securityAddressRows.Scan(&securityAddress.Id, &securityAddress.GroupId, &securityAddress.Backend, &securityAddress.Address); err != nil {
			log.Errorf("scan security_addresses row error: %v", err)
			continue
		}

		if group, ok := groups[securityAddress.GroupId]; ok {
			group.SecurityAddresses = append(group.SecurityAddresses, &securityAddress)
		}
	}

	if err = securityAddressRows.Err(); err != nil {
		log.Errorf("get next security_addresses row error: %v", err)
		return
	}

	mgr.groupIdAddress = groups
	envGroup := make(map[string][]*MakerAddress)
	for _, group := range groups {
		envGroup[group.Env] = append(envGroup[group.Env], group)
	}
	mgr.envGroup = envGroup
	mgr.backendAddressToGroup = backendAddressToGroup
}

func (mgr *MakerAddressManager) GetMakerAddressesByEnv(env string) []*MakerAddress {
	return mgr.envGroup[env]
}

func (mgr *MakerAddressManager) GetMakerAddressByGroupId(groupId int64) *MakerAddress {
	return mgr.groupIdAddress[groupId]
}

func (mgr *MakerAddressManager) GetGroupIDByBackendAndAddress(backend Backend, address string) int64 {
	if addressMap, ok := mgr.backendAddressToGroup[backend]; ok {
		if groupId, ok := addressMap[address]; ok {
			return groupId
		}
	}
	return 0
}
