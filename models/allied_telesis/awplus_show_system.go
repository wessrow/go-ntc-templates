package allied_telesis 

type AwplusShowSystem struct {
	Hardware	[]string	`json:"HARDWARE"`
	Serial	[]string	`json:"SERIAL"`
	Version	string	`json:"VERSION"`
	Hostname	string	`json:"HOSTNAME"`
}

var AwplusShowSystem_Template = `Value List HARDWARE (\S+)
Value List SERIAL (\S+)
Value VERSION (\S+)
Value HOSTNAME (\S+)

Start
  ^\s*Base\s+\d+\s+Base\s+${HARDWARE}\s+\S+\s+${SERIAL}
  ^\s*Base\s+\d+\s+${HARDWARE}\s+\S+\s+${SERIAL}
  ^\s*PSU
  ^\s*Current\ssoftware\s+:\s${VERSION}
  ^\s*System\s+Name -> SystemName
  ^\s*System\s+Contact -> SystemContact
  ^\s*System\s+Location -> SystemLocation
  ^\s*System\s+Status
  ^\s*Switch\s+System\s+Status
  ^\s*Stack\s+member
  ^\s*Board\s+ID\s+Bay\s+Board\s+Name
  ^\s*---
  ^\s*RAM:\s+Total
  ^\s*Flash:
  ^\s*Environment\s+Status
  ^\s*User\s+Configured\s+Territory
  ^\s*Uptime
  ^\s*Bootloader\s+version
  ^\s*Software\s+version
  ^\s*Build\s+date
  ^\s*Current\s+boot\s+config
  ^. -> Error

SystemName
  ^\s+${HOSTNAME}
  ^. -> Start

SystemContact
  ^\s+
  ^. -> Start

SystemLocation
  ^\s+
  ^. -> Start
`