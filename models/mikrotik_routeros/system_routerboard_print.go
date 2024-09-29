package mikrotik_routeros 

type SystemRouterboardPrint struct {
	Routerboard	string	`json:"ROUTERBOARD"`
	Board_name	string	`json:"BOARD_NAME"`
	Hardware_model	string	`json:"HARDWARE_MODEL"`
	Revision	string	`json:"REVISION"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Firmware_type	string	`json:"FIRMWARE_TYPE"`
	Factory_firmware	string	`json:"FACTORY_FIRMWARE"`
	Current_firmware	string	`json:"CURRENT_FIRMWARE"`
	Upgrade_firmware	string	`json:"UPGRADE_FIRMWARE"`
}

var SystemRouterboardPrint_Template = `Value ROUTERBOARD (\S+)
Value BOARD_NAME (.+)
Value HARDWARE_MODEL (.+)
Value REVISION (.+)
Value SERIAL_NUMBER (\S+)
Value FIRMWARE_TYPE (\S+)
Value FACTORY_FIRMWARE ([\d.]+)
Value CURRENT_FIRMWARE ([\d.]+)
Value UPGRADE_FIRMWARE ([\d.]+)

Start
  ^\s*;;;.* -> Next
  ^\s*routerboard:\s${ROUTERBOARD}
  ^\s*board-name:\s${BOARD_NAME}
  ^\s*model:\s${HARDWARE_MODEL}
  ^\s*revision:\s${REVISION}
  ^\s*serial-number:\s${SERIAL_NUMBER}
  ^\s*firmware-type:\s${FIRMWARE_TYPE}
  ^\s*factory-firmware:\s${FACTORY_FIRMWARE}
  ^\s*current-firmware:\s${CURRENT_FIRMWARE}
  ^\s*upgrade-firmware:\s${UPGRADE_FIRMWARE}
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^. -> Error
`