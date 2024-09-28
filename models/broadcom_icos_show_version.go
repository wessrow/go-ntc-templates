package models

type BroadcomIcosShowVersion struct {
	Description	string	`json:"DESCRIPTION"`
	Switch_type	string	`json:"SWITCH_TYPE"`
	Switch_model	string	`json:"SWITCH_MODEL"`
	Serial	string	`json:"SERIAL"`
	Fru_number	string	`json:"FRU_NUMBER"`
	Part_number	string	`json:"PART_NUMBER"`
	Maintenance_level	string	`json:"MAINTENANCE_LEVEL"`
	Manufacturer	string	`json:"MANUFACTURER"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Version	string	`json:"VERSION"`
	Os_version	string	`json:"OS_VERSION"`
	Network_processing_device	string	`json:"NETWORK_PROCESSING_DEVICE"`
	Cpld_version	string	`json:"CPLD_VERSION"`
	Board_revision	string	`json:"BOARD_REVISION"`
	Additional_packages	[]string	`json:"ADDITIONAL_PACKAGES"`
}