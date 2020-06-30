package model

type OptionModel struct {
	Id     uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Module string `json:"module" gorm:"column:module;not null"`
	Key    string `json:"key" gorm:"column:key;not null"`
	Value  string `json:"value" gorm:"column:value;"`
}

func (c *OptionModel) TableName() string {
	return "options"
}

// Create creates a new user account.
func (u *OptionModel) Create() error {
	return DB.Self.Create(&u).Error
}

func GetConfigs(module string, key string) (*OptionModel, error) {
	u := &OptionModel{}
	d := DB.Self.Where("`module` = ? AND `key` = ?", module, key).First(&u)

	return u, d.Error
}
