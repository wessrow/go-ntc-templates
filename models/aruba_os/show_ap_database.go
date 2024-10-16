package aruba_os

type ShowApDatabase struct {
	Ap_name    string `json:"AP_NAME"`
	Group      string `json:"GROUP"`
	Ap_model   string `json:"AP_MODEL"`
	Ip_address string `json:"IP_ADDRESS"`
	Status     string `json:"STATUS"`
	Flags      string `json:"FLAGS"`
	Primary    string `json:"PRIMARY"`
	Standby    string `json:"STANDBY"`
}

var ShowApDatabase_Template string = `Value AP_NAME (\S+)
Value GROUP (\S+)
Value AP_MODEL (\S+)
Value IP_ADDRESS ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value STATUS (.+?)
Value FLAGS (\S+)
Value PRIMARY ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value STANDBY ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))

Start
  ^${AP_NAME}\s+${GROUP}\s+${AP_MODEL}\s+${IP_ADDRESS}\s+${STATUS}\s+${FLAGS}\s+${PRIMARY}\s+${STANDBY}\s*$$ -> Record
  ^AP Database
  ^.+\.+
  ^Name\s+Group\s+AP\s+Type\s+IP\s+Address\s+Status\s+Flags\s+Switch\s+IP\s+Standby IP\s*$$
  ^Flags:
  ^\s*$$
  ^\s+
  ^Total
  ^-+
  ^. -> Error
`
