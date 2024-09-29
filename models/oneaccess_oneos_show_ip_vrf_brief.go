package models

type OneaccessOneosShowIpVrfBrief struct {
	Vrf	string	`json:"VRF"`
	Vrf_id	string	`json:"VRF_ID"`
	Interfaces	[]string	`json:"INTERFACES"`
}

var OneaccessOneosShowIpVrfBrief_Template = `Value VRF (\w+)
Value VRF_ID (\d+)
Value List INTERFACES ([\w\-]+\s?[\w\/\.]+)

Start
  ^\s*VRF\s+Name\s+VRF\s+I[Dd]\s+Interfaces\s*$$
  ^\s?\w+\s+\d+ -> Continue.Record
  ^\s?${VRF}?\s+${VRF_ID}\s+${INTERFACES}$$
  ^\s?${VRF}?\s+${VRF_ID}\s*$$
  ^\s+${INTERFACES}$$
  ^\s*$$
  ^. -> Error

`