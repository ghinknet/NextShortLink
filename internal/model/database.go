package model

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

// DatabaseLink table 'links'
type DatabaseLink struct {
	ID       int64  `xorm:"pk autoincr 'id'"`
	Link     string `xorm:"text notnull 'link'"`
	Validity *int64 `xorm:"bigint 'validity'"`
}

func (DatabaseLink) TableName() string {
	return "links"
}
