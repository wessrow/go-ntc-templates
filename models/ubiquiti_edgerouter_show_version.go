package models

type UbiquitiEdgerouterShowVersion struct {
	Version	string	`json:"VERSION"`
	Build_id	string	`json:"BUILD_ID"`
	Build_on	string	`json:"BUILD_ON"`
	Copyright	string	`json:"COPYRIGHT"`
	Hardware_model	string	`json:"HARDWARE_MODEL"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Uptime	string	`json:"UPTIME"`
}

var UbiquitiEdgerouterShowVersion_Template = `Value VERSION (\S+)
Value BUILD_ID (\S+)
Value BUILD_ON (.+)
Value COPYRIGHT (.+)
Value HARDWARE_MODEL (.+)
Value SERIAL_NUMBER (\S+)
Value UPTIME (.+)

Start
  ^Version:\s+${VERSION}
  ^Build\sID:\s+${BUILD_ID}
  ^Build\son:\s+${BUILD_ON}
  ^Copyright:\s+${COPYRIGHT}
  ^HW\smodel:\s+${HARDWARE_MODEL}
  ^HW\sS/N:\s+${SERIAL_NUMBER}
  ^Uptime:\s+${UPTIME}
  ^\s*$$
  ^. -> Error

`