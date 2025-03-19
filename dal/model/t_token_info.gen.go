// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/shopspring/decimal"
)

const TableNameTTokenInfo = "t_token_info"

// TTokenInfo mapped from table <t_token_info>
type TTokenInfo struct {
	ID                int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UpdateTimestamp   time.Time       `gorm:"column:update_timestamp;not null;default:CURRENT_TIMESTAMP" json:"update_timestamp"`
	InsertTimestamp   time.Time       `gorm:"column:insert_timestamp;not null;default:CURRENT_TIMESTAMP" json:"insert_timestamp"`
	TokenName         string          `gorm:"column:token_name;not null" json:"token_name"`
	ChainName         string          `gorm:"column:chain_name;not null" json:"chain_name"`
	TokenAddress      string          `gorm:"column:token_address;not null" json:"token_address"`
	Decimals          int32           `gorm:"column:decimals;not null" json:"decimals"`
	FullName          string          `gorm:"column:full_name;not null" json:"full_name"`
	TotalSupply       decimal.Decimal `gorm:"column:total_supply;not null;default:0" json:"total_supply"`
	DiscoverTimestamp time.Time       `gorm:"column:discover_timestamp;not null;default:CURRENT_TIMESTAMP" json:"discover_timestamp"`
	Icon              string          `gorm:"column:icon;not null" json:"icon"`
	Mcap              float64         `gorm:"column:mcap;not null" json:"mcap"`
	Fdv               float64         `gorm:"column:fdv;not null" json:"fdv"`
	Volume24h         float64         `gorm:"column:volume24h;not null" json:"volume24h"`
	Price24h          float64         `gorm:"column:price24h;not null" json:"price24h"`
	Price6h           float64         `gorm:"column:price6h;not null" json:"price6h"`
}

// TableName TTokenInfo's table name
func (*TTokenInfo) TableName() string {
	return TableNameTTokenInfo
}
