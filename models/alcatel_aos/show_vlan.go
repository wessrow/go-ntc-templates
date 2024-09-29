package alcatel_aos 

type ShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Type	string	`json:"TYPE"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Operational_state	string	`json:"OPERATIONAL_STATE"`
	Spanning_tree_1x1	string	`json:"SPANNING_TREE_1X1"`
	Spanning_tree_flat	string	`json:"SPANNING_TREE_FLAT"`
	Auth	string	`json:"AUTH"`
	Ip_state	string	`json:"IP_STATE"`
	Mbletag	string	`json:"MBLETAG"`
	Source_learn	string	`json:"SOURCE_LEARN"`
	Vlan_name	string	`json:"VLAN_NAME"`
}

var ShowVlan_Template = `Value VLAN_ID (\d+)
Value TYPE (std|vstk|gvrp|ipmv)
Value ADMIN_STATE (on|off)
Value OPERATIONAL_STATE (on|off)
Value SPANNING_TREE_1X1 (on|off)
Value SPANNING_TREE_FLAT (on|off)
Value AUTH (on|off)
Value IP_STATE (on|off)
Value MBLETAG (on|off)
Value SOURCE_LEARN (on|off)
Value VLAN_NAME ((\S+\s*)+\S+)

Start
  ^\s*${VLAN_ID}\s+${TYPE}\s+${ADMIN_STATE}\s+${OPERATIONAL_STATE}\s+${SPANNING_TREE_1X1}\s+${SPANNING_TREE_FLAT}\s+${AUTH}\s+${IP_STATE}\s+${MBLETAG}\s+${SOURCE_LEARN}\s+${VLAN_NAME}\s*$$ -> Record
  ^\s*stree\s+mble\s+src\s*$$
  ^\s*vlan\s+type\s+admin\s+oper\s+1x1\s+flat\s+auth\s+ip\s+tag\s+lrn\s+name\s*$$
  ^-----\+-----\+------\+------\+------\+------\+----\+-----\+-----\+------\+----------
  ^.*$$ -> Error
`