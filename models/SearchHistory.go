package models

type SearchHistory struct {
	Id             uint   `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	SearchKeyWords string `json:"search_key_words"`
	UserId         User   `json:"user_id"`
	Status         string `json:"status"`
	Result         string `json:"result"`
}
