package cisco_wlc

type SshShowMobilitySum struct {
	Mac_address  string `json:"MAC_ADDRESS"`
	Ip_address   string `json:"IP_ADDRESS"`
	Group_name   string `json:"GROUP_NAME"`
	Multicast_ip string `json:"MULTICAST_IP"`
	Status       string `json:"STATUS"`
}

var SshShowMobilitySum_Template string = `Value MAC_ADDRESS (([\d1-9a-f]{2}\:?){6})
Value IP_ADDRESS (([\d1-9]+\.?){4})
Value GROUP_NAME (\S+)
Value MULTICAST_IP (([\d1-9]+\.?){4})
Value STATUS (.+?)

Start
  ^Mobility\s+Protocol\s+Port
  ^Default\s+Mobility\s+Domain
  ^Multicast\s+Mode\s+
  ^Mobility\s+Domain\s+ID\s+for\s+802.11r
  ^Mobility\s+Keepalive\s+Interval
  ^Mobility\s+Keepalive\s+Count
  ^Mobility\s+Group\s+Members\s+Configured
  ^Mobility\s+Control\s+Message\s+DSCP\s+Value
  ^Controllers\s+configured\s+in\s+the\s+Mobility\s+Group 
  ^\s+MAC\s+Address\s+IP\s+Address\s+Group\s+Name\s+Multicast\s+IP\s+Status -> Mobility_Controllers


Mobility_Controllers
  ^\s+${MAC_ADDRESS}\s+${IP_ADDRESS}\s+${GROUP_NAME}\s+${MULTICAST_IP}\s+${STATUS}s*$$ -> Record
  ^\s*$$
  ^. -> Error`
