package arista_eos 

type ShowIpBgp struct {
	Status	string	`json:"STATUS"`
	Path_selection	string	`json:"PATH_SELECTION"`
	Route_source	string	`json:"ROUTE_SOURCE"`
	Network	string	`json:"NETWORK"`
	Next_hop	string	`json:"NEXT_HOP"`
	Metric	string	`json:"METRIC"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Weight	string	`json:"WEIGHT"`
	As_path	string	`json:"AS_PATH"`
	Origin	string	`json:"ORIGIN"`
	Vrf	string	`json:"VRF"`
	Local_as	string	`json:"LOCAL_AS"`
	Router_id	string	`json:"ROUTER_ID"`
}

var ShowIpBgp_Template = `Value STATUS ([bceEisS*>#?])
Value PATH_SELECTION ([bceEisS*>#? ])
Value ROUTE_SOURCE ([bceEisS*>#? ])
Value NETWORK (\S+)
Value NEXT_HOP (\S+)
Value METRIC (\S+)
Value LOCAL_PREF (\S+)
Value WEIGHT (\S+)
Value AS_PATH (.*?)
Value ORIGIN ([ie\?])
Value Filldown VRF (\S+)
Value Filldown LOCAL_AS (\d+\.\d+|\d+)
Value Filldown ROUTER_ID (\S+)

Start
  ^BGP\s+routing\s+table\s+information\s+for\s+VRF\s+${VRF}
  ^Router\s+identifier\s+${ROUTER_ID},\s+local\s+AS\s+number\s+${LOCAL_AS}
  ^AS Path Attributes.+
  ^\s${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}\s+${NETWORK}\s+${NEXT_HOP}\s+${METRIC}\s+${LOCAL_PREF}\s+${WEIGHT}\s+${AS_PATH}\s+${ORIGIN}$$ -> Record

EOF
`