package aruba_os

type ShowApRadioDatabase struct {
	Group      string `json:"GROUP"`
	Ap_model   string `json:"AP_MODEL"`
	Status     string `json:"STATUS"`
	Flags      string `json:"FLAGS"`
	Primary    string `json:"PRIMARY"`
	Standby    string `json:"STANDBY"`
	Radio0     string `json:"RADIO0"`
	Ap_name    string `json:"AP_NAME"`
	Radio1     string `json:"RADIO1"`
	Radio2     string `json:"RADIO2"`
	Ip_address string `json:"IP_ADDRESS"`
}

var ShowApRadioDatabase_Template string = `Value AP_NAME (\S+)
Value GROUP (\S+)
Value AP_MODEL (\S+)
Value IP_ADDRESS ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value STATUS (.+?)
Value FLAGS (\S+)
Value PRIMARY ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value STANDBY ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value RADIO0 (\S+)
Value RADIO1 (\S+)
Value RADIO2 (\S+)

Start
  ^${AP_NAME}\s+${GROUP}\s+${AP_MODEL}\s+${IP_ADDRESS}\s+${STATUS}\s+${FLAGS}\s+${PRIMARY}\s+${STANDBY}\s+${RADIO0}\s+${RADIO1}\s+${RADIO2}\s*$$ -> Record
  ^AP Radio Database
  ^.+\.+
  ^Name\s+Group\s+AP\s+Type\s+IP Address\s+Status\s+Flags\s+Switch IP\s+Standby IP\s+Radio 0 Band/Chan/HT-Type/EIRP\s+Radio 1 Band/Chan/HT-Type/EIRP\s+Radio 2 Band/Chan/HT-Type/EIRP\s*$$
  ^Flags:
  ^\s*$$
  ^\s+
  ^"Spectrum"
  ^Total
  ^-+
  ^. -> Error

`
