package eltex

type ShowInterfacesSwitchport struct {
	Interface_mode               string   `json:"INTERFACE_MODE"`
	Added_by                     []string `json:"ADDED_BY"`
	Mac_based_vlan_group_id      []string `json:"MAC_BASED_VLAN_GROUP_ID"`
	Subnet_based_vlan_group_id   []string `json:"SUBNET_BASED_VLAN_GROUP_ID"`
	Subnet_based_vlan_id         []string `json:"SUBNET_BASED_VLAN_ID"`
	Acceptable_frame_type        string   `json:"ACCEPTABLE_FRAME_TYPE"`
	Protected                    string   `json:"PROTECTED"`
	Vlan                         []string `json:"VLAN"`
	Protocol_based_vlan_group_id []string `json:"PROTOCOL_BASED_VLAN_GROUP_ID"`
	Protocol_based_vlan_id       []string `json:"PROTOCOL_BASED_VLAN_ID"`
	Interface                    string   `json:"INTERFACE"`
	Gvrp_status                  string   `json:"GVRP_STATUS"`
	Name                         []string `json:"NAME"`
	Forbidden_vlan               []string `json:"FORBIDDEN_VLAN"`
	Forbidden_vlan_name          []string `json:"FORBIDDEN_VLAN_NAME"`
	Mac_based_vlan_id            []string `json:"MAC_BASED_VLAN_ID"`
	Ingress_filtering            string   `json:"INGRESS_FILTERING"`
	Native_vlan                  string   `json:"NATIVE_VLAN"`
	Egress_rule                  []string `json:"EGRESS_RULE"`
	Interface_membership_type    []string `json:"INTERFACE_MEMBERSHIP_TYPE"`
}

var ShowInterfacesSwitchport_Template string = `Value INTERFACE (.*?)
Value INTERFACE_MODE (.*?)
Value GVRP_STATUS (.*?)
Value INGRESS_FILTERING (.*?)
Value ACCEPTABLE_FRAME_TYPE (.*?)
Value NATIVE_VLAN (.*?)
Value PROTECTED (.*?)
Value List VLAN (\d+)
Value List NAME (.*?)
Value List EGRESS_RULE (Untagged|Tagged)
Value List INTERFACE_MEMBERSHIP_TYPE (\S+)
Value List ADDED_BY (\S+)
Value List FORBIDDEN_VLAN (\d+)
Value List FORBIDDEN_VLAN_NAME (.*?)
Value List PROTOCOL_BASED_VLAN_GROUP_ID (\d+)
Value List PROTOCOL_BASED_VLAN_ID (\d+)
Value List MAC_BASED_VLAN_GROUP_ID (\d+)
Value List MAC_BASED_VLAN_ID (\d+)
Value List SUBNET_BASED_VLAN_GROUP_ID (\d+)
Value List SUBNET_BASED_VLAN_ID (\d+)

Start
  ^\s*Added\s+by\s*:\s+D-Default,\s+S-Static,\s+G-GVRP,\s+R-Radius\s+Assigned\s+VLAN,\s+T-Guest\s+VLAN,\s+V-Voice\s+VLAN\s*$$
  ^\s*Port\s*:.*$$ -> Continue.Record
  ^\s*Port\s*:\s*${INTERFACE}\s*$$
  ^\s*Port\s+Mode\s*:\s*${INTERFACE_MODE}\s*$$
  ^\s*Gvrp\s+Status\s*:\s*${GVRP_STATUS}\s*$$
  ^\s*Ingress\s+Filtering\s*:\s*${INGRESS_FILTERING}\s*$$
  ^\s*Acceptable\s+Frame\s+Type\s*:\s*${ACCEPTABLE_FRAME_TYPE}\s*$$
  ^\s*Ingress\s+UnTagged\s+VLAN\s*\(\s*NATIVE\s*\)\s*:\s*${NATIVE_VLAN}\s*$$
  ^\s*Protected\s*:\s*${PROTECTED}\s*$$
  ^\s*Port\s+is\s+member\s+in\s*:\s*$$ -> VlanMemberTable
  ^\s*Forbidden\s+VLAN[Ss]\s*:\s*$$ -> ForbiddenVlansTable
  ^\s*Classification\s+rules\s*:\s*$$ -> ClassificationRules
  ^\s*Protocol\s+based\s+VLANs\s*:\s*$$ -> ProtocolBasedVlansTable
  ^\s*Mac\s+based\s+VLANs\s*:\s*$$ -> MacBasedVlansTable
  ^\s*Subnet\s+based\s+VLANs\s*:\s*$$ -> SubnetBasedVlansTable
  ^\s*$$
  ^. -> Error

VlanMemberTable
  ^\s*Vlan\s+Name\s+Egress\s+rule\s+Port\s+Membership\s+Type\s*$$ -> VlanMemberTableType1
  ^\s*Vlan\s+Name\s+Egress\s+rule\s+Added\s+by\s*$$ -> VlanMemberTableType2
  ^\s*$$
  ^. -> Error

VlanMemberTableType1
  ^(?:\s*-+)+\s*$$
  ^\s*${VLAN}\s+${NAME}\s+${EGRESS_RULE}\s+${INTERFACE_MEMBERSHIP_TYPE}\s*$$
  ^\s*$$ -> Start
  ^. -> Error

VlanMemberTableType2
  ^(?:\s*-+)+\s*$$
  ^\s*${VLAN}\s+${NAME}\s+${EGRESS_RULE}\s+${ADDED_BY}\s*$$
  ^\s*$$ -> Start
  ^. -> Error

ForbiddenVlansTable
  ^\s*Vlan\s+Name\s*$$
  ^(?:\s*-+)+\s*$$
  ^\s*${FORBIDDEN_VLAN}\s+${FORBIDDEN_VLAN_NAME}\s*$$
  ^\s*$$ -> Start
  ^. -> Error

ClassificationRules
  ^\s*$$ -> Start

ProtocolBasedVlansTable
  ^\s*Group\s+ID\s+Vlan\s+ID\s*$$
  ^(?:\s*-+)+\s*$$
  ^\s*${PROTOCOL_BASED_VLAN_GROUP_ID}\s+${PROTOCOL_BASED_VLAN_ID}\s*$$
  ^\s*$$ -> Start
  ^. -> Error

MacBasedVlansTable
  ^\s*Group\s+ID\s+Vlan\s+ID\s*$$
  ^(?:\s*-+)+\s*$$
  ^\s*${MAC_BASED_VLAN_GROUP_ID}\s+${MAC_BASED_VLAN_ID}\s*$$
  ^\s*$$ -> Start
  ^. -> Error

SubnetBasedVlansTable
  ^\s*Group\s+ID\s+Vlan\s+ID\s*$$
  ^(?:\s*-+)+\s*$$
  ^\s*${SUBNET_BASED_VLAN_GROUP_ID}\s+${SUBNET_BASED_VLAN_ID}\s*$$
  ^\s*$$ -> Start
  ^. -> Error
`
