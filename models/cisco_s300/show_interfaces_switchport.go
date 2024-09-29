package cisco_s300 

type ShowInterfacesSwitchport struct {
	Interface	string	`json:"INTERFACE"`
	Interface_mode	string	`json:"INTERFACE_MODE"`
	Gvrp_status	string	`json:"GVRP_STATUS"`
	Ingress_filtering_status	string	`json:"INGRESS_FILTERING_STATUS"`
	Acceptable_frame_type	string	`json:"ACCEPTABLE_FRAME_TYPE"`
	Native_vlan	string	`json:"NATIVE_VLAN"`
	Vlan	[]string	`json:"VLAN"`
	Name	[]string	`json:"NAME"`
	Egress_rule	[]string	`json:"EGRESS_RULE"`
	Interface_membership_type	[]string	`json:"INTERFACE_MEMBERSHIP_TYPE"`
	Forbidden_vlan	[]string	`json:"FORBIDDEN_VLAN"`
	Forbidden_vlan_name	[]string	`json:"FORBIDDEN_VLAN_NAME"`
	Mac_based_vlan_group_id	[]string	`json:"MAC_BASED_VLAN_GROUP_ID"`
	Mac_based_vlan_id	[]string	`json:"MAC_BASED_VLAN_ID"`
}

var ShowInterfacesSwitchport_Template = `Value INTERFACE (\S+)
Value INTERFACE_MODE (\S+)
Value GVRP_STATUS (\S+)
Value INGRESS_FILTERING_STATUS (\S+)
Value ACCEPTABLE_FRAME_TYPE (\S+)
Value NATIVE_VLAN (\S+)
Value List VLAN (\d+)
Value List NAME (.*?)
Value List EGRESS_RULE (Untagged|Tagged)
Value List INTERFACE_MEMBERSHIP_TYPE (\S+)
Value List FORBIDDEN_VLAN (\d+)
Value List FORBIDDEN_VLAN_NAME (.*)
Value List MAC_BASED_VLAN_GROUP_ID (\d+)
Value List MAC_BASED_VLAN_ID (\d+)

Start
  ^\s*Port\s*:\s*${INTERFACE}\s*$$
  ^\s*Port\s+Mode\s*:\s*${INTERFACE_MODE}\s*$$
  ^\s*Gvrp\s+Status\s*:\s*${GVRP_STATUS}\s*$$
  ^\s*Ingress\s+Filtering\s*:\s*${INGRESS_FILTERING_STATUS}\s*$$
  ^\s*Acceptable\s+Frame\s+Type\s*:\s*${ACCEPTABLE_FRAME_TYPE}\s*$$
  ^\s*Ingress\s+UnTagged\s+VLAN\s*\(\s*NATIVE\s*\)\s*:\s*${NATIVE_VLAN}\s*$$
  ^\s*Port\s+is\s+member\s+in:\s*$$ -> PartOfVlan
  ^\s*$$
  ^. -> Error

PartOfVlan
  ^\s*Vlan\s+Name\s+Egress\s+rule\s+Port\s+Membership\s+Type\s*$$
  ^(\s*-*)*\s*$$
  ^\s*${VLAN}\s+${NAME}\s+${EGRESS_RULE}\s+${INTERFACE_MEMBERSHIP_TYPE}\s*$$
  ^\s*Forbidden\s+VLANS:\s*$$ -> ForbiddenVlans
  ^\s*$$
  ^. -> Error

ForbiddenVlans
  ^\s*Vlan\s+Name\s*$$
  ^(\s*-*)*\s*$$
  ^\s*${FORBIDDEN_VLAN}\s+${FORBIDDEN_VLAN_NAME}\s*$$
  ^\s*Classification\s+rules:\s*$$ -> ClassificationRules
  ^\s*$$
  ^. -> Error

ClassificationRules
  ^\s*Mac\s+based\s+VLANs:\s*$$ -> MacBasedVlans
  ^\s*$$
  ^. -> Error

MacBasedVlans
  ^\s*Group\s+ID\s+Vlan\s+ID\s*$$
  ^(\s*-*)*\s*$$
  ^\s*${FORBIDDEN_VLAN}\s+${FORBIDDEN_VLAN_NAME}\s*$$
  ^\s*$$
  ^. -> Error
`