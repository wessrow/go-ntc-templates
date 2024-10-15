package cisco_nxos

type ShowIpPimRpVrfAll struct {
	Vrf                      string `json:"VRF"`
	Bsr                      string `json:"BSR"`
	Auto_rp                  string `json:"AUTO_RP"`
	Bsr_candidate_policy     string `json:"BSR_CANDIDATE_POLICY"`
	Bsr_policy               string `json:"BSR_POLICY"`
	Auto_rp_announce_policy  string `json:"AUTO_RP_ANNOUNCE_POLICY"`
	Auto_rp_discovery_policy string `json:"AUTO_RP_DISCOVERY_POLICY"`
}

var ShowIpPimRpVrfAll_Template string = `Value VRF (\S+)
Value BSR (\S+)
Value AUTO_RP (\S+)
Value BSR_CANDIDATE_POLICY (\S+)
Value BSR_POLICY (\S+)
Value AUTO_RP_ANNOUNCE_POLICY (\S+)
Value AUTO_RP_DISCOVERY_POLICY (\S+)

Start
  ^\s*PIM RP Status Information for VRF "${VRF}"\s*$$
  ^\s*BSR ${BSR}\s*$$
  ^\s*Auto-RP ${AUTO_RP}\s*$$
  ^\s*BSR RP Candidate policy: ${BSR_CANDIDATE_POLICY}\s*$$
  ^\s*BSR RP policy: ${BSR_POLICY}\s*$$
  ^\s*Auto-RP Announce policy: ${AUTO_RP_ANNOUNCE_POLICY}\s*$$
  ^\s*Auto-RP Discovery policy: ${AUTO_RP_DISCOVERY_POLICY}\s*$$ -> Record
  ^. -> Error
`
