// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAccount = "account"

// Account mapped from table <account>
type Account struct {
	UserID            string `gorm:"column:userId;not null" json:"userId"`
	Type              string `gorm:"column:type;not null" json:"type"`
	Provider          string `gorm:"column:provider;primaryKey" json:"provider"`
	ProviderAccountID string `gorm:"column:providerAccountId;primaryKey" json:"providerAccountId"`
	RefreshToken      string `gorm:"column:refresh_token" json:"refresh_token"`
	AccessToken       string `gorm:"column:access_token" json:"access_token"`
	ExpiresAt         int32  `gorm:"column:expires_at" json:"expires_at"`
	TokenType         string `gorm:"column:token_type" json:"token_type"`
	Scope             string `gorm:"column:scope" json:"scope"`
	IDToken           string `gorm:"column:id_token" json:"id_token"`
	SessionState      string `gorm:"column:session_state" json:"session_state"`
}

// TableName Account's table name
func (*Account) TableName() string {
	return TableNameAccount
}
