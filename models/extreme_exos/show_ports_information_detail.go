package extreme_exos

type ShowPortsInformationDetail struct {
	Random_early_drop string   `json:"RANDOM_EARLY_DROP"`
	Link_state        string   `json:"LINK_STATE"`
	Link_downs_last   string   `json:"LINK_DOWNS_LAST"`
	Vlan_id           []string `json:"VLAN_ID"`
	Interface         string   `json:"INTERFACE"`
	Type              string   `json:"TYPE"`
	Admin_state       string   `json:"ADMIN_STATE"`
	Link_ups_num      string   `json:"LINK_UPS_NUM"`
	Link_ups_last     string   `json:"LINK_UPS_LAST"`
	Link_downs_num    string   `json:"LINK_DOWNS_NUM"`
	Description       string   `json:"DESCRIPTION"`
	Virtual_router    string   `json:"VIRTUAL_ROUTER"`
}

var ShowPortsInformationDetail_Template string = `Value INTERFACE (\d+)
Value DESCRIPTION (.*)
Value VIRTUAL_ROUTER (.*)
Value TYPE (.*)
Value RANDOM_EARLY_DROP (.*)
Value ADMIN_STATE (.*)
Value LINK_STATE (\S+)
Value LINK_UPS_NUM (\d+)
Value LINK_UPS_LAST (.*)
Value LINK_DOWNS_NUM (\d+)
Value LINK_DOWNS_LAST (.*)
Value List VLAN_ID (\d+)

Start
  ^\s*Port:.*$$ -> Continue.Record
  ^\s*Port:\s*${INTERFACE}\s*$$
  ^\s*Description\s+String:\s*\"${DESCRIPTION}\"\s*$$
  ^\s*Virtual-router:\s*${VIRTUAL_ROUTER}\s*$$
  ^\s*Type:\s*${TYPE}\s*$$
  ^\s*Random\s+Early\s+drop:\s*${RANDOM_EARLY_DROP}\s*$$
  ^\s*Admin\s+state:\s*${ADMIN_STATE}
  ^\s*Link\s+State:\s*${LINK_STATE}(,)?(\s)?.*$$
  ^\s*Link\s+Ups:\s*${LINK_UPS_NUM}\s+Last:\s*${LINK_UPS_LAST}\s*$$
  ^\s*Link\s+Downs:\s*${LINK_DOWNS_NUM}\s+Last:\s*${LINK_DOWNS_LAST}\s*$$
	^\s*VLAN\s+cfg:\s*$$
  ^.*VLAN\s+ID:\s*${VLAN_ID}\s*$$
`
