package cisco_wlc

type SshShowApSummary struct {
	Radio_mac     string `json:"RADIO_MAC"`
	Dse_location  string `json:"DSE_LOCATION"`
	Number_of_aps string `json:"NUMBER_OF_APS"`
	Port          string `json:"PORT"`
	Ip_address    string `json:"IP_ADDRESS"`
	Slot          string `json:"SLOT"`
	Country       string `json:"COUNTRY"`
	Clients       string `json:"CLIENTS"`
	Ap_name       string `json:"AP_NAME"`
	Ap_model      string `json:"AP_MODEL"`
	Mac_address   string `json:"MAC_ADDRESS"`
	Location      string `json:"LOCATION"`
	Priority      string `json:"PRIORITY"`
	State         string `json:"STATE"`
}

var SshShowApSummary_Template string = `Value AP_NAME (\S+)
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
