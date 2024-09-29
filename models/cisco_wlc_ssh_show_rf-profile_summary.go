package models

type CiscoWlcSshShowRfProfileSummary struct {
	Profile	string	`json:"PROFILE"`
	Band	string	`json:"BAND"`
	Description	string	`json:"DESCRIPTION"`
	N_only	string	`json:"N_ONLY"`
	Applied	string	`json:"APPLIED"`
}

var CiscoWlcSshShowRfProfileSummary_Template = `Value PROFILE (\S+)
#Value PROFILE (\S+)
Value BAND (\d\.?4? GHz)
Value DESCRIPTION ((\S+(.?\S+)+))
Value N_ONLY (\w+)
Value APPLIED (\w+)


Start
  ^.+\.+
  ^RF\s+Profile\s+Name\s+Band\s+Description\s+11n-client-only\s+Applied -> Profile
  ^\s*$$
  ^. -> Error

Profile
  ^(?:High|Low|Typical)-Client-Density-802.11
  ^${PROFILE}\s+${BAND}\s+${DESCRIPTION}\s+${N_ONLY}\s+${APPLIED} -> Record
  ^-+[\s|-]+$$
  ^\s*$$
  ^. -> Error

`