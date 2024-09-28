package models

type CiscoNxosShowIpPimRpVrfAll struct {
	Vrf	string	`json:"VRF"`
	Bsr	string	`json:"BSR"`
	Auto_rp	string	`json:"AUTO_RP"`
	Bsr_candidate_policy	string	`json:"BSR_CANDIDATE_POLICY"`
	Bsr_policy	string	`json:"BSR_POLICY"`
	Auto_rp_announce_policy	string	`json:"AUTO_RP_ANNOUNCE_POLICY"`
	Auto_rp_discovery_policy	string	`json:"AUTO_RP_DISCOVERY_POLICY"`
}