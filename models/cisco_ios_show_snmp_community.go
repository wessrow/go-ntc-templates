package models

type CiscoIosShowSnmpCommunity struct {
	Name	string	`json:"NAME"`
	Index	string	`json:"INDEX"`
	Security_name	string	`json:"SECURITY_NAME"`
	Storage_type	string	`json:"STORAGE_TYPE"`
	Access_list_number	string	`json:"ACCESS_LIST_NUMBER"`
}