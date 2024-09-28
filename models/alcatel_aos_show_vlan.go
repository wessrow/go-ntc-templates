package models

type AlcatelAosShowVlan struct {
	Vlan_id            string `json:"VLAN_ID"`
	Type               string `json:"TYPE"`
	Admin_state        string `json:"ADMIN_STATE"`
	Operational_state  string `json:"OPERATIONAL_STATE"`
	Spanning_tree_1x1  string `json:"SPANNING_TREE_1X1"`
	Spanning_tree_flat string `json:"SPANNING_TREE_FLAT"`
	Auth               string `json:"AUTH"`
	Ip_state           string `json:"IP_STATE"`
	Mbletag            string `json:"MBLETAG"`
	Source_learn       string `json:"SOURCE_LEARN"`
	Vlan_name          string `json:"VLAN_NAME"`
}
