package cisco_xr

type ShowBgp struct {
	Origin         string `json:"ORIGIN"`
	Network        string `json:"NETWORK"`
	Next_hop       string `json:"NEXT_HOP"`
	Metric         string `json:"METRIC"`
	Local_pref     string `json:"LOCAL_PREF"`
	Route_source   string `json:"ROUTE_SOURCE"`
	Netmask        string `json:"NETMASK"`
	Weight         string `json:"WEIGHT"`
	Path_selection string `json:"PATH_SELECTION"`
	Router_id      string `json:"ROUTER_ID"`
	Local_as       string `json:"LOCAL_AS"`
	Nsr            string `json:"NSR"`
	Dampening      string `json:"DAMPENING"`
	Bgp_state      string `json:"BGP_STATE"`
	Status         string `json:"STATUS"`
	As_path        string `json:"AS_PATH"`
}

var ShowBgp_Template string = `Value Filldown ROUTER_ID (\S+)
Value Filldown LOCAL_AS (\d+)
Value Filldown NSR (.+?)
Value Filldown BGP_STATE (.+?)
Value Filldown DAMPENING (.+?)
# Documenting STATUS, PATH_SELECTION, and ROUTE_SOURCE capture group options
# https://www.cisco.com/c/en/us/td/docs/ios_xr_sw/iosxr_r3-7/routing/command/reference/rr37bgp.pdf
# https://www.cisco.com/c/en/us/support/docs/routers/asr-9000-series-aggregation-services-routers/116386-configure-asr9000-00.html
Value STATUS ([NrSs*]?)
Value PATH_SELECTION ([>dh]?)
Value ROUTE_SOURCE ([i]?)
Value Filldown NETWORK (\S+?)
Value Filldown NETMASK (\d+)
Value Required NEXT_HOP (\S+)
Value METRIC (\d{0,10})
Value LOCAL_PREF (\d{0,10})
Value WEIGHT (\d+)
Value AS_PATH (.*?)
Value ORIGIN ([ie\?])

Start
  ^\S+\s+\S+\s+\d+\s+\d+:\d+:\d+\.\d+\s+\S+\s*$$
  ^BGP\s+router\s+identifier\s+${ROUTER_ID},\s+local\s+AS\s+number\s+${LOCAL_AS}\s*$$
  ^BGP\s+generic\s+scan
  ^Non-stop\s+routing\s+is\s+${NSR}\s*$$
  ^BGP\s+table\s+state:\s+${BGP_STATE}\s*$$
  ^BGP\s+(NSR|main|scan)
  ^Table\s+ID:
  ^Dampening\s+${DAMPENING}\s*$$
  ^Status\s+codes:
  ^\s+\S+\s+-\s+\S+
  ^Origin\s+codes
  # Checking for header
  ^\s*Network\s+Next(?:\s+|-)[Hh]op\s+Metric\s+LocPrf\s+Weight\s+Path\s*$$ -> BGPTable
  ^\s*$$
  ^. -> Error

BGPTable
  # Regex to match the complete line including network
  # *> 10.0.0.0/8          192.168.1.1          900    100      0 65135 65235 i
  ^${STATUS}\s*${PATH_SELECTION}\s*${ROUTE_SOURCE}\s*${NETWORK}/${NETMASK}\s+${NEXT_HOP}\s{1,19}${METRIC}\s{1,10}${LOCAL_PREF}\s{2,6}${WEIGHT}\s*${AS_PATH}\s*${ORIGIN}\s*$$ -> Record
  #
  # Regex to match the lines without network
  # * i                   192.168.1.2          900    100      0 65135 65235 i
  ^\s*${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}\s+${NEXT_HOP}\s{1,19}${METRIC}\s{1,10}${LOCAL_PREF}\s{2,6}${WEIGHT}\s*${AS_PATH}\s*${ORIGIN}\s*$$ -> Record
  ^Processed
  ^\s*$$
  ^. -> Error

EOF
`
