package models

type ArubaAoscxShowVsfDetail struct {
	Member_id	string	`json:"MEMBER_ID"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Status	string	`json:"STATUS"`
	Cpu	string	`json:"CPU"`
	Memory	string	`json:"MEMORY"`
	Vsf_link_1	string	`json:"VSF_LINK_1"`
	Vsf_link_2	string	`json:"VSF_LINK_2"`
}