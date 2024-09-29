package models

type CienaSaosPortShowNetworkConfig struct {
	Name	string	`json:"NAME"`
	Private_vlan_id	string	`json:"PRIVATE_VLAN_ID"`
	Accept_frame_type	string	`json:"ACCEPT_FRAME_TYPE"`
	Untag_ingress_vlan_id	string	`json:"UNTAG_INGRESS_VLAN_ID"`
	Vlan_ingress_filter	string	`json:"VLAN_INGRESS_FILTER"`
	Vs_ingress_filter	string	`json:"VS_INGRESS_FILTER"`
	Egress_untag_vlan	string	`json:"EGRESS_UNTAG_VLAN"`
	Vs_egress_filter	string	`json:"VS_EGRESS_FILTER"`
	Ingress_cos_policy	string	`json:"INGRESS_COS_POLICY"`
	Vpls_port_type	string	`json:"VPLS_PORT_TYPE"`
	Eth_vc_ethertype	string	`json:"ETH_VC_ETHERTYPE"`
}

var CienaSaosPortShowNetworkConfig_Template = `Value NAME (\S+)
Value PRIVATE_VLAN_ID (\S+)
Value ACCEPT_FRAME_TYPE (\S+)
Value UNTAG_INGRESS_VLAN_ID (\S+)
Value VLAN_INGRESS_FILTER (\S+)
Value VS_INGRESS_FILTER (\S+)
Value EGRESS_UNTAG_VLAN (\S+)
Value VS_EGRESS_FILTER (\S+)
Value INGRESS_COS_POLICY (\S+)
Value VPLS_PORT_TYPE (\S+)
Value ETH_VC_ETHERTYPE (\d+)

Start
  ^\+-+\s*PORT\s*ETHERNET\s*CONFIGURATION\s*-+\+ 
  ^\+-+
  ^\|\s*\s*\|\s*\s*\|\s*Accpt\s*\|\s*Untag\s*\|\s*VLAN\s*\|\s*VS\s*\|\s*Egress\s*\|\s*VS\s*\|\s*Ingress\s*\|\s*VPLS\s*\|\s*Eth\s*VC\s*\|\s*$$
  ^\|\s*Port\s*\|\s*\s*\|\s*Frame\s*\|\s*Ingress\s*\|\s*Ingress\s*\|\s*Ingress\s*\|\s*Untag\s*\|\s*Egress\s*\|\s*CoS\s*\|\s*Port\s*\|\s*Ether-\s*\|\s*$$
  ^\|\s*Name\s*\|\s*PVID\s*\|\s*Type\s*\|\s*Vid\s*\|\s*Filter\s*\|\s*Filter\s*\|\s*VLAN\s*\|\s*Filter\s*\|\s*Policy\s*\|\s*Type\s*\|\s*Type\s*\|\s*$$
  ^\|\s*${NAME}\s*\|\s*${PRIVATE_VLAN_ID}\s*\|\s*${ACCEPT_FRAME_TYPE}\s*\|\s*${UNTAG_INGRESS_VLAN_ID}\s*\|\s*(${VLAN_INGRESS_FILTER}|\s+)\s*\|\s*${VS_INGRESS_FILTER}\s*\|\s*${EGRESS_UNTAG_VLAN}\s*\|\s*${VS_EGRESS_FILTER}\s*\|\s*${INGRESS_COS_POLICY}\s*\|\s*${VPLS_PORT_TYPE}\s*\|\s*${ETH_VC_ETHERTYPE}\s*\| -> Record
  ^\s*$$
  ^. -> Error

`