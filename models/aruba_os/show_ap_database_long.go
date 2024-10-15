package aruba_os

type ShowApDatabaseLong struct {
	Ap_name     string `json:"AP_NAME"`
	Mac_address string `json:"MAC_ADDRESS"`
	Ip_address  string `json:"IP_ADDRESS"`
	Flags       string `json:"FLAGS"`
	Standby     string `json:"STANDBY"`
	Serial      string `json:"SERIAL"`
	Port        string `json:"PORT"`
	Group       string `json:"GROUP"`
	Ap_model    string `json:"AP_MODEL"`
	Status      string `json:"STATUS"`
	User        string `json:"USER"`
	Primary     string `json:"PRIMARY"`
	Fqln        string `json:"FQLN"`
	Outer_ip    string `json:"OUTER_IP"`
}

var ShowApDatabaseLong_Template string = `Value AP_NAME (\S+)
Value GROUP (\S+)
Value AP_MODEL (\S+)
Value IP_ADDRESS ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value STATUS (.+?)
Value FLAGS (\S+)
Value PRIMARY ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value STANDBY ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value MAC_ADDRESS ([a-fA-F0-9]{2}(:[a-fA-F0-9]{2}){5})
Value SERIAL (\S+)
Value PORT (.+?)
Value FQLN (.+?)
Value OUTER_IP (.+?)
Value USER (.+?)

Start
  ^${AP_NAME}\s+${GROUP}\s+${AP_MODEL}\s+${IP_ADDRESS}\s+${STATUS}\s+${FLAGS}\s+${PRIMARY}\s+${STANDBY}\s+${MAC_ADDRESS}\s+${SERIAL}\s+${PORT}\s+${FQLN}\s+${OUTER_IP}\s+${USER}\s*$$ -> Record
  ^AP Database
  ^.+\.+
  ^Name\s+Group\s+AP Type\s+IP Address\s+Status\s+Flags\s+Switch IP\s+Standby IP\s+Wired MAC Address\s+Serial #\s+Port\s+FQLN\s+Outer IP\s+User\s*$$
  ^Flags:
  ^\s*$$
  ^\s+
  ^Total
  ^-+
  ^. -> Error
`
