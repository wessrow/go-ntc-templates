package cisco_ios

type ShowApSummary struct {
	Slot              string `json:"SLOT"`
	Mac_address       string `json:"MAC_ADDRESS"`
	Location          string `json:"LOCATION"`
	Country           string `json:"COUNTRY"`
	Regulatory_domain string `json:"REGULATORY_DOMAIN"`
	Ip_address        string `json:"IP_ADDRESS"`
	State             string `json:"STATE"`
	Ap_name           string `json:"AP_NAME"`
	Ap_model          string `json:"AP_MODEL"`
	Radio_mac         string `json:"RADIO_MAC"`
}

var ShowApSummary_Template string = `Value AP_NAME (\S+)
Value SLOT (\d+)
Value AP_MODEL (\S+)
Value MAC_ADDRESS ([a-fA-F0-9:\.]+)
Value RADIO_MAC ([a-fA-F0-9:\.]+)
Value LOCATION (.+?)
Value COUNTRY (\S+)
Value REGULATORY_DOMAIN (\S+)
Value IP_ADDRESS ([a-fA-F0-9:\.]+)
Value STATE (\S+)

Start
  # AP count summary
  ^[a-zA-Z]+\s[a-z]+\s[a-zA-z]+:\s\d+$$
  # abbreviation legend
  ^[A-Z]+\s+=\s+.*$$
  # line of dashes
  ^-+\s*$$
  # legacy ios-xe wlc
  ^AP\s+Name\s+Slots\s+AP\s+Model\s+Ethernet\s+MAC\s+Radio\sMAC\s+Location\s+Country\s+IP\s+Address\s+State\s*$$
  ^${AP_NAME}\s+${SLOT}\s+${AP_MODEL}\s+${MAC_ADDRESS}\s+${RADIO_MAC}\s+${LOCATION}\s+${COUNTRY}\s+${IP_ADDRESS}\s+${STATE}\s*$$ -> Record
  # ios-xe wlc approximately 17.9.4 and later
  ^AP\s+Name\s+Slots\s+AP\s+Model\s+Ethernet\s+MAC\s+Radio\sMAC\s+CC\s+RD\s+IP\s+Address\s+State\s+Location\s*$$
  ^${AP_NAME}\s+${SLOT}\s+${AP_MODEL}\s+${MAC_ADDRESS}\s+${RADIO_MAC}\s+${COUNTRY}\s+${REGULATORY_DOMAIN}\s+${IP_ADDRESS}\s+${STATE}\s+${LOCATION}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
