package models

type CiscoAsaShowVersion struct {
	Version	string	`json:"VERSION"`
	Device_mgr_version	string	`json:"DEVICE_MGR_VERSION"`
	Compile_date	string	`json:"COMPILE_DATE"`
	Image	string	`json:"IMAGE"`
	Hostname	string	`json:"HOSTNAME"`
	Uptime	string	`json:"UPTIME"`
	Hardware	string	`json:"HARDWARE"`
	Model	string	`json:"MODEL"`
	Flash	string	`json:"FLASH"`
	Interfaces	[]string	`json:"INTERFACES"`
	License_mode	string	`json:"LICENSE_MODE"`
	License_state	string	`json:"LICENSE_STATE"`
	Max_intf	string	`json:"MAX_INTF"`
	Max_vlans	string	`json:"MAX_VLANS"`
	Failover	string	`json:"FAILOVER"`
	Cluster	string	`json:"CLUSTER"`
	Serial	[]string	`json:"SERIAL"`
	Last_mod	string	`json:"LAST_MOD"`
}