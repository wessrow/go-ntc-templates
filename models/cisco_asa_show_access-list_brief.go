package models

type CiscoAsaShowAccessListBrief struct {
	Acl_name	string	`json:"ACL_NAME"`
	Acl_total_elements	string	`json:"ACL_TOTAL_ELEMENTS"`
	Acl_name_hash	string	`json:"ACL_NAME_HASH"`
	Line_hash	string	`json:"LINE_HASH"`
	Group_hash	string	`json:"GROUP_HASH"`
	Counter	string	`json:"COUNTER"`
	Last_hit	string	`json:"LAST_HIT"`
}