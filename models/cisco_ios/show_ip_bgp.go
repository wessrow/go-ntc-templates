package cisco_ios

type ShowIpBgp struct {
	Route_source   string `json:"ROUTE_SOURCE"`
	Next_hop       string `json:"NEXT_HOP"`
	Local_pref     string `json:"LOCAL_PREF"`
	Status         string `json:"STATUS"`
	Path_selection string `json:"PATH_SELECTION"`
	Network        string `json:"NETWORK"`
	Metric         string `json:"METRIC"`
	Weight         string `json:"WEIGHT"`
	As_path        string `json:"AS_PATH"`
	Origin         string `json:"ORIGIN"`
}

var ShowIpBgp_Template string = `Value Filldown STATUS ([bdhimrsSx*>])
Value Filldown PATH_SELECTION ([bdhimrsSx*> ])
Value Filldown ROUTE_SOURCE ([bdhimrsSx*> ])
Value Filldown NETWORK (\S{0,18})
Value Required NEXT_HOP (\S{0,15})
Value Filldown METRIC (\S{0,10})
Value LOCAL_PREF (\S{0,6})
Value WEIGHT (\S{0,6})
Value AS_PATH (.*?)
Value ORIGIN ([ie\?])

Start
  # Since using mostly position, play it safe and ensure we see header first
  ^\s+Network\s+Next Hop\s+Metric\s+LocPrf\s+Weight\s+Path -> Bgp_table
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

Bgp_table
  # Account for show ip bgp vpnv4 vrf command
  ^Route\s+Distinguisher
  #
  #
  # Match if subnet is 17,18 characters long, creates two lines
  # Example: *>i 10.104.192.208/29
  ^\s{0,1}${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}\s{0,2}(?=${NETWORK}).{17,18}$$ -> Next
  #
  #
  # Compliment to previous, status, path_selection, route_source, network is filldown.
  # Example:                     200.200.186.194          0    100  50000 64801 64808 64608 64601 64787 i
  ^\s{20,25}(?=${NEXT_HOP}).{15}\s(?=\s{0,10}${METRIC}).{10}\s(?=\s{0,6}${LOCAL_PREF}).{6}\s(?=\s{0,6}${WEIGHT}).{6}\s*${AS_PATH}\s*${ORIGIN}$$ -> Record
  #
  #
  # Match first when there is no network, since previous line had it already (compliment and filldown below)
  # Example: *>                  0.0.0.0                  0         32768 i
  ^\s{0,1}${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}\s{0,2}\s{16}\s(?=${NEXT_HOP}).{15}\s(?=\s{0,10}${METRIC}).{10}\s(?=\s{0,6}${LOCAL_PREF}).{6}\s(?=\s{0,6}${WEIGHT}).{6}\s*${AS_PATH}\s*${ORIGIN}$$ -> Record
  #
  #
  # Full normal example. metric, and as_path might not exist, regex defaults to blank line.
  # Example: * i172.16.1.0/24    172.16.1.2               0    100      0 i
  ^\s{0,1}${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}\s{0,2}(?=${NETWORK}).{16}\s(?=${NEXT_HOP}).{15}\s(?=\s{0,10}${METRIC}).{10}\s(?=\s{0,6}${LOCAL_PREF}).{6}\s(?=\s{0,6}${WEIGHT}).{6}\s*${AS_PATH}\s*${ORIGIN}$$ -> Record


EOF
`
