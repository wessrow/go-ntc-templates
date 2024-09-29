package cisco_nxos 

type ShowVrfDetail struct {
	Name	string	`json:"NAME"`
	Id	string	`json:"ID"`
	Vpn_id	string	`json:"VPN_ID"`
	State	string	`json:"STATE"`
	Default_rd	string	`json:"DEFAULT_RD"`
	Max_routes	string	`json:"MAX_ROUTES"`
	Mid_threshold	string	`json:"MID_THRESHOLD"`
	Table_id	[]string	`json:"TABLE_ID"`
	Address_family	[]string	`json:"ADDRESS_FAMILY"`
	Fwd_id	[]string	`json:"FWD_ID"`
	Table_status	[]string	`json:"TABLE_STATUS"`
}

var ShowVrfDetail_Template = `Value Required NAME (\S+)
Value ID (\S+)
Value VPN_ID (\S+)
Value STATE (\S+)
Value DEFAULT_RD ((\d+|\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}):\d+|<not set>)
Value MAX_ROUTES (\S+)
Value MID_THRESHOLD (\S+)
Value List TABLE_ID (\S+)
Value List ADDRESS_FAMILY (\S+)
Value List FWD_ID (\S+)
Value List TABLE_STATUS (\S+)

Start
  ^VRF-Name:\s+${NAME},\s+VRF-ID:\s+${ID},\s+State:\s+${STATE} -> VRF
  ^. -> Error

VRF
  ^VRF-Name:\s+${NAME},\s+VRF-ID:\s+${ID},\s+State:\s+${STATE}
  ^\s+VPNID:\s+${VPN_ID}
  ^\s+RD:\s+${DEFAULT_RD}
  ^\s+Max\s+Routes:\s+${MAX_ROUTES}\s+Mid\-Threshold:\s+${MID_THRESHOLD}
  ^\s+Table\-ID:\s+${TABLE_ID},\s+AF:\s+${ADDRESS_FAMILY},\s+Fwd-ID:\s+${FWD_ID},\s+State:\s+${TABLE_STATUS}
  ^\s*$$ -> Record
  ^$$ -> Record
  ^. -> Error
`