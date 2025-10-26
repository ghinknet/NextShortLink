package model

import "time"

// DatabaseConfig table 'config'
type DatabaseConfig struct {
	Key   int64        `xorm:"pk int notnull 'key'"`
	Value SystemConfig `xorm:"jsonb notnull 'value'"`
}

func (DatabaseConfig) TableName() string {
	return "config"
}

// DatabaseApplication table 'application'
type DatabaseApplication struct {
	ID        int64  `xorm:"pk autoincr 'id'"`
	SecretID  string `xorm:"text notnull unique 'secret_id'"`
	SecretKey string `xorm:"text notnull unique 'secret_key'"`
	Name      string `xorm:"text notnull unique 'name'"`
}

func (DatabaseApplication) TableName() string {
	return "application"
}

// DatabasePermission table 'permission'
type DatabasePermission struct {
	ID           int64    `xorm:"pk autoincr 'id'"`
	Application  int      `xorm:"int notnull 'application'"`
	Interface    string   `xorm:"text notnull 'interface'"`
	DisableKey   bool     `xorm:"bool notnull 'disable_key'"`
	DisableToken bool     `xorm:"bool notnull 'disable_token'"`
	Blacklist    []string `xorm:"jsonb 'blacklist'"`
	Whitelist    []string `xorm:"jsonb 'whitelist'"`
	QPS          int64    `xorm:"bigint 'qps'"`
	QPM          int64    `xorm:"bigint 'qpm'"`
}

func (DatabasePermission) TableName() string {
	return "permission"
}

// DatabasePackage table 'package'
type DatabasePackage struct {
	ID            int64  `xorm:"pk autoincr 'id'"`
	Application   int    `xorm:"int 'application'"`
	Interface     string `xorm:"text notnull 'interface'"`
	Total         int64  `xorm:"bigint 'total'"`
	Used          int64  `xorm:"bigint 'used'"`
	Unlimit       bool   `xorm:"boolean default false 'unlimit'"`
	Priority      int    `xorm:"int notnull 'priority'"`
	AvailableFrom int64  `xorm:"bigint 'available_from'"`
	AvailableTo   int64  `xorm:"bigint 'available_to'"`
}

func (DatabasePackage) TableName() string {
	return "package"
}

// DatabaseCNID table 'cnid'
type DatabaseCNID struct {
	IDEncrypted     string `xorm:"text notnull unique 'id_encrypted'"`
	NameEncrypted   string `xorm:"text notnull 'name_encrypted'"`
	IDFingerprint   string `xorm:"text notnull unique 'id_fingerprint'"`
	NameFingerprint string `xorm:"text notnull 'name_fingerprint'"`
}

func (DatabaseCNID) TableName() string {
	return "cnid"
}

// DatabaseMusic404 table 'music_404'
type DatabaseMusic404 struct {
	ID     int64  `xorm:"pk autoincr 'id'"`
	Name   string `xorm:"text notnull 'name'"`
	Artist string `xorm:"text notnull 'artist'"`
	Audio  string `xorm:"text notnull 'audio'"`
	Cover  string `xorm:"text notnull 'cover'"`
}

func (DatabaseMusic404) TableName() string {
	return "music404"
}

// DatabaseGreyFilter table 'grey_filter'
type DatabaseGreyFilter struct {
	ID          int64  `xorm:"pk autoincr 'id'"`
	Slot        string `xorm:"text notnull 'slot'"`
	Object      string `xorm:"text notnull 'object'"`
	Rule        string `xorm:"text notnull 'rule'"`
	Description string `xorm:"text notnull 'description'"`
}

func (DatabaseGreyFilter) TableName() string {
	return "grey_filter"
}

// DatabaseFurpassCase table 'furpass_case'
type DatabaseFurpassCase struct {
	ID        int64     `xorm:"pk autoincr 'id'"`
	Name      string    `xorm:"text notnull 'name'"`
	BeginDate time.Time `xorm:"date notnull 'begin_date'"`
	EndDate   time.Time `xorm:"date notnull 'end_date'"`
	Location  string    `xorm:"text notnull 'location'"`
	Cover     string    `xorm:"text notnull 'cover'"`
	Link      string    `xorm:"text 'link'"`
}

func (DatabaseFurpassCase) TableName() string {
	return "furpass_case"
}
