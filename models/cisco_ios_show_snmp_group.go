package models

type CiscoIosShowSnmpGroup struct {
	Groupname	string	`json:"GROUPNAME"`
	Security_model	string	`json:"SECURITY_MODEL"`
	Contextname	string	`json:"CONTEXTNAME"`
	Storage_type	string	`json:"STORAGE_TYPE"`
	Readview	string	`json:"READVIEW"`
	Writeview	string	`json:"WRITEVIEW"`
	Notifyview	string	`json:"NOTIFYVIEW"`
	Row_status	string	`json:"ROW_STATUS"`
	Access_list	string	`json:"ACCESS_LIST"`
}