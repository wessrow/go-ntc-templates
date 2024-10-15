package oneaccess_oneos

type ShowTrackAll struct {
	Vrrp_id       string `json:"VRRP_ID"`
	State         string `json:"STATE"`
	Last_change   string `json:"LAST_CHANGE"`
	Up_delay      string `json:"UP_DELAY"`
	Down_delay    string `json:"DOWN_DELAY"`
	Id            string `json:"ID"`
	Interface     string `json:"INTERFACE"`
	Vrf           string `json:"VRF"`
	State_changes string `json:"STATE_CHANGES"`
	Poll_interval string `json:"POLL_INTERVAL"`
	Category      string `json:"CATEGORY"`
}

var ShowTrackAll_Template string = `Value Required ID (\d+)
Value CATEGORY (\S+)
Value INTERFACE (\S+(?:\s\S+)?)
Value VRRP_ID (\d+)
Value VRF (\S+)
Value STATE (\S+)
Value STATE_CHANGES (\d+)
Value LAST_CHANGE (\S+)
Value UP_DELAY (\d+)
Value DOWN_DELAY (\d+)
Value POLL_INTERVAL (\d+)

Start
  ^\s*Track\s -> Continue.Record
  ^\s*Track\s${ID}
  ^\s*interface\s${INTERFACE}\s${CATEGORY}
  ^\s*${CATEGORY}\sId\s${VRRP_ID}(\svrf\s${VRF})
  ^\s*\S+\sis\s${STATE}
  ^\s*${STATE_CHANGES}\sChange,\sLast\sChange\s${LAST_CHANGE}
  ^\s*Up\sDelay\s${UP_DELAY},\sDown\sDelay\s${DOWN_DELAY}
  ^\s*Poll\sInterval\s\(in\smsec\)\s${POLL_INTERVAL}
  ^\s*$$
  ^. -> Error
`
