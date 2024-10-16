package cisco_ios

type ShowVrfDetail struct {
	Vrf         string   `json:"VRF"`
	Rd          string   `json:"RD"`
	Rt_import   []string `json:"RT_IMPORT"`
	Rt_export   []string `json:"RT_EXPORT"`
	Vpn_id      string   `json:"VPN_ID"`
	Description string   `json:"DESCRIPTION"`
	Interfaces  []string `json:"INTERFACES"`
}

var ShowVrfDetail_Template string = `Value VRF (\S+)
Value RD (\S+|<not set>)
Value List RT_IMPORT (\S+)
Value List RT_EXPORT (\S+)
Value VPN_ID (\S+|<not set>)
Value DESCRIPTION ([\S\s]+)
Value List INTERFACES ([\w\./-]+)

Start
  ^\s*VRF\s.+default\sRD\s.+;\sdefault\sVPNID\s.+$$ -> Continue.Record
  ^\s*Old\sCLI\sformat,\ssupports\s\S+\sonly
  ^\s*Flags:\s\S+
  ^\s*\s*VRF\s${VRF}\s.+;\sdefault\sRD\s${RD};\sdefault\sVPNID\s${VPN_ID}\s*$$
  ^\s*\s*Description:\s${DESCRIPTION}
  ^\s*No\sinterfaces.*\s*$$
  ^\s*Interfaces:+s*$$ -> Interfaces
  ^\s*Flags:\s\S+
  ^\s*Export\sVPN.*\s*$$ -> Rt_Export
  ^\s*Import\sVPN.*\s*$$ -> Rt_Import
  ^\s*No\sExport.*\s*$$
  ^\s*No\sImport.*\s*$$
  ^\s*No\simport.*\s*$$
  ^\s*No\sexport.*\s*$$
  ^\s*No\sglobal.*\s*$$
  ^\s*VRF\slabel.*\s*$$
  ^\s*Address\sfamily.*\s*$$
  ^\s*New\sCLI\sformat.*\s*$$
  ^\s*Old\sCLI\sformat.*\s*$$
  ^. -> Error

Interfaces
  ^\s*${INTERFACES}\s*$$
  ^\s*Address\sfamily.+$$ -> Start
  
Rt_Export
  ^\s*RT:${RT_EXPORT}\s*$$
  ^\s*(Import\sVPN.*|No\simport\sroute-map)+\s*$$ -> Rt_Import

Rt_Import
  ^\s*RT:${RT_IMPORT}\s*$$
  ^\s*No\simport\sroute-map\s*$$ -> Start
  
`
