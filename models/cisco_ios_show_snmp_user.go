package models

type CiscoIosShowSnmpUser struct {
	User_name	string	`json:"USER_NAME"`
	Engine_id	string	`json:"ENGINE_ID"`
	Storage_type	string	`json:"STORAGE_TYPE"`
	Access_list	string	`json:"ACCESS_LIST"`
	Authentication_protocol	string	`json:"AUTHENTICATION_PROTOCOL"`
	Privacy_protocol	string	`json:"PRIVACY_PROTOCOL"`
	Group_name	string	`json:"GROUP_NAME"`
}