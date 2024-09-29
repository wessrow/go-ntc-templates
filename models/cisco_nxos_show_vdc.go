package models

type CiscoNxosShowVdc struct {
	Vdc_id	string	`json:"VDC_ID"`
	Vdc_name	string	`json:"VDC_NAME"`
	State	string	`json:"STATE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Lc	string	`json:"LC"`
}

var CiscoNxosShowVdc_Template = `Value VDC_ID (\d+)
Value VDC_NAME (\S+)
Value STATE (\S+)
Value MAC_ADDRESS (\S+)
Value TYPE (\S+)
Value LC (\S+)

Start
  ^[Ss]witchwide\s+mode
  ^vdc_id\s+vdc_name\s+state\s+mac\s+type\s+lc\s*$$
  ^-+
  ^\s*${VDC_ID}\s+${VDC_NAME}\s+${STATE}\s+${MAC_ADDRESS}\s+${TYPE}\s+${LC} -> Record
  ^\s*$$
  ^. -> Error

`