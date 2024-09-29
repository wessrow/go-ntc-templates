package cisco_wlc 

type SshShowWlanSum struct {
	Wlanid	string	`json:"WLANID"`
	Profile	string	`json:"PROFILE"`
	Ssid	string	`json:"SSID"`
	Status	string	`json:"STATUS"`
	Interface	string	`json:"INTERFACE"`
	Pmip_mobility	string	`json:"PMIP_MOBILITY"`
}

var SshShowWlanSum_Template = `Value WLANID (\d+)
Value PROFILE (.+?)
Value SSID (.+?)
Value STATUS (Enabled|Disabled)
Value INTERFACE (\w+(.?\w+)+)
Value PMIP_MOBILITY (\S+)


Start
  ^Number\s+of\s+WLANs
  ^WLAN\s+ID\s+WLAN\s+Profile\s+Name\s+\/\s+SSID\s+Status\s+Interface\s+Name(\s+PMIPv6\s+Mobility)?\s*$$
  ^-------\s+
  ^${WLANID}\s+${PROFILE}\s/\s${SSID}\s+${STATUS}\s+${INTERFACE}(\s+${PMIP_MOBILITY})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`