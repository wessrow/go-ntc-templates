package broadcom_icos

type ShowVersion struct {
	Switch_type               string   `json:"SWITCH_TYPE"`
	Switch_model              string   `json:"SWITCH_MODEL"`
	Version                   string   `json:"VERSION"`
	Cpld_version              string   `json:"CPLD_VERSION"`
	Board_revision            string   `json:"BOARD_REVISION"`
	Description               string   `json:"DESCRIPTION"`
	Fru_number                string   `json:"FRU_NUMBER"`
	Mac_address               string   `json:"MAC_ADDRESS"`
	Os_version                string   `json:"OS_VERSION"`
	Network_processing_device string   `json:"NETWORK_PROCESSING_DEVICE"`
	Additional_packages       []string `json:"ADDITIONAL_PACKAGES"`
	Serial                    string   `json:"SERIAL"`
	Part_number               string   `json:"PART_NUMBER"`
	Maintenance_level         string   `json:"MAINTENANCE_LEVEL"`
	Manufacturer              string   `json:"MANUFACTURER"`
}

var ShowVersion_Template string = `Value DESCRIPTION (.+)
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
