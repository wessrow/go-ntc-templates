package oneaccess_oneos

type ShowRunningConfigIpDhcp struct {
	Vrf            string   `json:"VRF"`
	Default_router string   `json:"DEFAULT_ROUTER"`
	Domain_name    string   `json:"DOMAIN_NAME"`
	Dns            []string `json:"DNS"`
	Network        string   `json:"NETWORK"`
	Lease          string   `json:"LEASE"`
	Excluded       string   `json:"EXCLUDED"`
	Pool           string   `json:"POOL"`
}

var ShowRunningConfigIpDhcp_Template string = `Value POOL (\S+)
Value VRF (\S+)
Value DEFAULT_ROUTER (\d+\.\d+\.\d+\.\d+)
Value DOMAIN_NAME (\S+)
Value List DNS (\d+\.\d+\.\d+\.\d+)
Value NETWORK (\S+\s\S+)
Value LEASE (.*)
Value EXCLUDED (\d+\.\d+\.\d+\.\d+(\s\d+\.\d+\.\d+\.\d+)?)
 
Start
  # DHCP SCOPE records have a dhcp pool name
  ^ip\sdhcp\spool\s${POOL}
  # EXCLUDE records do not have a dhcp pool name
  ^ip\sdhcp\s(vrf ${VRF}\s)?excluded-address\s${EXCLUDED} -> Record
  ^\s*lease\s${LEASE}
  ^\s*vrf\s${VRF}
  ^\s*domain-name\s+${DOMAIN_NAME}
  ^\s*default-router\s${DEFAULT_ROUTER}
  ^\s*dns-server\s${DNS} -> Continue
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){1}${DNS}\s*$$
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){1}${DNS} -> Continue
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){2}${DNS}\s*$$
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){2}${DNS} -> Continue
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){3}${DNS}\s*$$
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){3}${DNS} -> Continue
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){4}${DNS}\s*$$
  ^\s*dns-server\s(\d+\.\d+\.\d+\.\d+\s){4}${DNS} -> Continue
  ^\s*network\s${NETWORK}
  ^\s*netmask
  ^(exit|!) -> Record Start
  ^\s$$
  ^. -> Error
`
