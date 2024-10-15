package cisco_ios

type ShowLicense struct {
	Feature          string `json:"FEATURE"`
	Period_left      string `json:"PERIOD_LEFT"`
	Period_used      string `json:"PERIOD_USED"`
	License_type     string `json:"LICENSE_TYPE"`
	License_state    string `json:"LICENSE_STATE"`
	License_count    string `json:"LICENSE_COUNT"`
	License_priority string `json:"LICENSE_PRIORITY"`
}

var ShowLicense_Template string = `Value FEATURE (\S+)
Value PERIOD_LEFT (.*)
Value PERIOD_USED (.*)
Value LICENSE_TYPE (.*)
Value LICENSE_STATE (.*)
Value LICENSE_COUNT (.*)
Value LICENSE_PRIORITY (.*)

Start
  ^Index\s+\d+\s+Feature: -> Continue.Record
  ^Index\s+\d+\s+Feature:\s+${FEATURE}\s*$$
  ^\s+Period\s+left:\s+${PERIOD_LEFT}$$
  ^\s+Period\s+Used:\s+${PERIOD_USED}$$
  ^\s+License\s+Type:\s+${LICENSE_TYPE}$$
  ^\s+License\s+State:\s+${LICENSE_STATE}$$
  ^\s+License\s+Count:\s+${LICENSE_COUNT}$$
  ^\s+License\s+Priority:\s+${LICENSE_PRIORITY}$$
  ^\s*$$
  ^. -> Error
`
