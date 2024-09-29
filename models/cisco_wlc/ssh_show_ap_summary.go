package cisco_wlc 

type SshShowApSummary struct {
	Ap_name	string	`json:"AP_NAME"`
	Slot	string	`json:"SLOT"`
	Ap_model	string	`json:"AP_MODEL"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Radio_mac	string	`json:"RADIO_MAC"`
	Location	string	`json:"LOCATION"`
	Port	string	`json:"PORT"`
	Country	string	`json:"COUNTRY"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Clients	string	`json:"CLIENTS"`
	Dse_location	string	`json:"DSE_LOCATION"`
	Priority	string	`json:"PRIORITY"`
	State	string	`json:"STATE"`
	Number_of_aps	string	`json:"NUMBER_OF_APS"`
}

var SshShowApSummary_Template = `Value AP_NAME (\S+)
Value SLOT (\d+)
Value AP_MODEL (\S+)
Value MAC_ADDRESS ([a-fA-F0-9:\.]+)
Value RADIO_MAC ([a-fA-F0-9:\.]+)
Value LOCATION (.+?)
Value PORT (.+?)
Value COUNTRY (\S+)
Value IP_ADDRESS ([a-fA-F0-9:\.]+)
Value CLIENTS (\d+)
Value DSE_LOCATION (.+?)
Value PRIORITY (\d+)
Value STATE (\S+)
Value NUMBER_OF_APS (\d+)

Start
  # Skip lines 
  ^Number\s+of\s+APs -> Record
  # Match the various record patterns
  ^${AP_NAME}\s+${SLOT}\s+${AP_MODEL}\s+${MAC_ADDRESS}\s+${LOCATION}\s+${COUNTRY}\s+${IP_ADDRESS}\s+${CLIENTS}\s+${DSE_LOCATION}\s*$$ -> Record
  ^${AP_NAME}\s+${SLOT}\s+${AP_MODEL}\s+${MAC_ADDRESS}\s+${LOCATION}\s+${PORT}\s+${COUNTRY}\s+${PRIORITY}\s*$$ -> Record
  ^${AP_NAME}\s+${SLOT}\s+${AP_MODEL}\s+${MAC_ADDRESS}\s+${RADIO_MAC}\s${LOCATION}\s+${COUNTRY}\s+${IP_ADDRESS}\s+${STATE}\s*$$ -> Record
  ^.+\.+
  ^\s*$$
  ^AP\s+Name\s+Slots\s+AP\s+Model\s+Ethernet\s+MAC\s+Location\s+Country\s+IP\s+Address\s+Clients\s+DSE\s+Location\s*$$
  ^AP\s+Name\s+Slots\s+AP\s+Model\s+Ethernet\s+MAC\s+Location\s+Port\s+Country\s+Priority\s*$$
  ^AP\s+Name\s+Slots\s+AP\s+Model\s+Ethernet\s+MAC\s+Radio\s+MAC\s+Location\s+Country\s+IP\s+Address\s+State\s*$$
  ^-+
  ^. -> Error
`