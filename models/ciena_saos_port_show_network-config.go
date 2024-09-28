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