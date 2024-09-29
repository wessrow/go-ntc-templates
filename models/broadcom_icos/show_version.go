package broadcom_icos 

type ShowVersion struct {
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

var ShowVersion_Template = `Value DESCRIPTION (.+)
Value SWITCH_TYPE (.+)
Value SWITCH_MODEL (.+)
Value SERIAL (.+)
Value FRU_NUMBER (.+)
Value PART_NUMBER (.+)
Value MAINTENANCE_LEVEL (.+)
Value MANUFACTURER (.+)
Value MAC_ADDRESS (.+)
Value VERSION (.+)
Value OS_VERSION (.+)
Value NETWORK_PROCESSING_DEVICE (.+)
Value CPLD_VERSION (.+)
Value BOARD_REVISION (.+)
Value List ADDITIONAL_PACKAGES (.+)

Start
  # Captures show version for:
  # Accton AS4610-54P, Accton AS5610-52X, Quanta LY2R, Quanta LB9, DNI AG3448P-R   
  # The following can be an empty value as it doesnt exist in all the models:
  # FruNumber, PartNumber, CPLDversion, BoardRevision
  ^\s*Switch\s*:\s*\d+\s*$$
  ^\s*$$
  ^\s*System\s*Description\s*\.+\s*${DESCRIPTION}$$
  ^\s*Machine\s*Type\s*\.+\s*${SWITCH_TYPE}$$
  ^\s*Machine\s*Model\s*\.+\s*${SWITCH_MODEL}$$
  ^\s*Serial\s*Number\s*\.+\s*${SERIAL}$$
  ^\s*FRU\s*Number\s*\.+\s*${FRU_NUMBER}$$
  ^\s*Part\s*Number\s*\.+\s*${PART_NUMBER}$$
  ^\s*Maintenance\s*Level\s*\.+\s*${MAINTENANCE_LEVEL}$$
  ^\s*Manufacturer\s*\.+\s*${MANUFACTURER}$$
  ^\s*Burned\s*In\s*MAC\s*Address\s*\.+\s*${MAC_ADDRESS}$$
  ^\s*Software\s*Version\s*\.+\s*${VERSION}$$
  ^\s*Operating\s*System\s*\.+\s*${OS_VERSION}$$
  ^\s*Network\s*Processing\s*Device\s*\.+\s*${NETWORK_PROCESSING_DEVICE}$$
  ^\s*CPLD\s*version\s*\.+\s*${CPLD_VERSION}$$
  ^\s*Board\s*Revision\s*\.+\s*${BOARD_REVISION}$$
  ^\s*Additional\s*Packages\s*\.+\s*${ADDITIONAL_PACKAGES}$$
  ^\s+${ADDITIONAL_PACKAGES}$$
  ^. -> Error`