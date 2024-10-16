package cisco_xr

type ShowVrfAllDetail struct {
	Description string   `json:"DESCRIPTION"`
	Interfaces  []string `json:"INTERFACES"`
	Vrf         string   `json:"VRF"`
	Rd          string   `json:"RD"`
	Rt_import   []string `json:"RT_IMPORT"`
	Rt_export   []string `json:"RT_EXPORT"`
	Vpn_id      string   `json:"VPN_ID"`
	Mode        string   `json:"MODE"`
}

var ShowVrfAllDetail_Template string = `Value VRF (\S+)
Value RD (\S+|not set)
Value List RT_IMPORT (\S+)
Value List RT_EXPORT (\S+)
Value VPN_ID (\S+|not set)
Value MODE (\S+)
Value DESCRIPTION ([\S\s]+|not set)
Value List INTERFACES ([\w\./-]+)

Start
  ^\s*VRF\s\S+;\sRD\s.+;\sVPN\sID\s.+\s*$$ -> Continue.Record
  ^\s*VRF ${VRF}; RD ${RD}; VPN ID ${VPN_ID}\s*$$
  ^\s*VRF mode: ${MODE}\s*$$
  ^\s*Description ${DESCRIPTION}\s*$$
  ^\s*Import\sVPN.*:+\s*$$ -> Rt_Import
  ^\s*Export\sVPN.*:+\s*$$ -> Rt_Export
  ^\s*Interfaces:+s*$$ -> Interfaces

Interfaces
  ^\s*${INTERFACES}\s*$$
  ^\s*Address family.+$$ -> Start

Rt_Import
  ^\s*${RT_IMPORT}\s*$$
  ^\s*(Export\sVPN.*:|No\simport\sroute\spolicy)+\s*$$ -> Rt_Export

Rt_Export
  ^\s*${RT_EXPORT}\s*$$
  ^\s*No\simport\sroute\spolicy+\s*$$ -> Start
`
