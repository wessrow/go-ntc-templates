package models

type AristaEosShowIpOspfSummary struct {
	Instance	string	`json:"INSTANCE"`
	Router_id	string	`json:"ROUTER_ID"`
	Vrf	string	`json:"VRF"`
	Area	string	`json:"AREA"`
	Type	string	`json:"TYPE"`
	Interfaces	string	`json:"INTERFACES"`
	Neighbors	string	`json:"NEIGHBORS"`
	Neighbors_full	string	`json:"NEIGHBORS_FULL"`
	Router_lsas	string	`json:"ROUTER_LSAS"`
	Network_lsas	string	`json:"NETWORK_LSAS"`
	Summary_lsas	string	`json:"SUMMARY_LSAS"`
	Asbr_lsas	string	`json:"ASBR_LSAS"`
	Nssa_lsas	string	`json:"NSSA_LSAS"`
}

var AristaEosShowIpOspfSummary_Template = `Value Filldown INSTANCE (\d+)
Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown VRF (\S+)
Value AREA (\d+\.\d+\.\d+\.\d+)
Value TYPE (\S+)
Value INTERFACES (\d+)
Value NEIGHBORS (\d+)
Value NEIGHBORS_FULL (\d+)
Value ROUTER_LSAS (\d+)
Value NETWORK_LSAS (\d+)
Value SUMMARY_LSAS (\d+)
Value ASBR_LSAS (\d+)
Value NSSA_LSAS (\d+)

Start
  ^OSPF instance ${INSTANCE} with ID ${ROUTER_ID}, VRF ${VRF},.*$$
  ^Time since.*$$
  ^Max LSAs.*$$
  ^Type-5.*$$
  ^ID\s+Type\s+Intf\s+Nbrs\s+\(full\)\s+RTR LSA\s+NW LSA\s+SUM LSA\s+ASBR LSA\s+TYPE-7 LSA\s*$$
  ^${AREA}\s+${TYPE}\s+${INTERFACES}\s+${NEIGHBORS}\s+\(${NEIGHBORS_FULL}\s*\)\s+${ROUTER_LSAS}\s+${NETWORK_LSAS}\s+${SUMMARY_LSAS}\s+${ASBR_LSAS}\s+${NSSA_LSAS}\s*$$ -> Record
`