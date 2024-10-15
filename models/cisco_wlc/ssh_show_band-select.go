package cisco_wlc

type SshShowBandSelect struct {
	Band_select_client_mid_rssi string `json:"BAND_SELECT_CLIENT_MID_RSSI"`
	Band_select_probe_response  string `json:"BAND_SELECT_PROBE_RESPONSE"`
	Band_select_cycle_count     string `json:"BAND_SELECT_CYCLE_COUNT"`
	Band_select_cycle_thresh    string `json:"BAND_SELECT_CYCLE_THRESH"`
	Band_select_age_suppress    string `json:"BAND_SELECT_AGE_SUPPRESS"`
	Band_select_age_dual_band   string `json:"BAND_SELECT_AGE_DUAL_BAND"`
	Band_select_client_rssi     string `json:"BAND_SELECT_CLIENT_RSSI"`
}

var SshShowBandSelect_Template string = `Value BAND_SELECT_PROBE_RESPONSE (.+?)
Value BAND_SELECT_CYCLE_COUNT (\d+\s+\w+)
Value BAND_SELECT_CYCLE_THRESH (\d+\s+\w+)
Value BAND_SELECT_AGE_SUPPRESS (\d+\s+\w+)
Value BAND_SELECT_AGE_DUAL_BAND (\d+\s+\w+)
Value BAND_SELECT_CLIENT_RSSI (-\d+\s+dBm)
Value BAND_SELECT_CLIENT_MID_RSSI (-\d+\s+dBm)


Start
  ^Band\s+Select\s+Probe\s+Response\.*\s+${BAND_SELECT_PROBE_RESPONSE}s*$$
  ^\s+Cycle\s+Count\.*\s+${BAND_SELECT_CYCLE_COUNT}s*$$
  ^\s+Cycle\s+Threshold\.*\s+${BAND_SELECT_CYCLE_THRESH}s*$$
  ^\s+Age\s+Out\s+Suppression\.*\s+${BAND_SELECT_AGE_SUPPRESS}s*$$
  ^\s+Age\s+Out\s+Dual\s+Band\.*\s+${BAND_SELECT_AGE_DUAL_BAND}s*$$
  ^\s+Client\s+RSSI\.*\s+${BAND_SELECT_CLIENT_RSSI}s*$$
  ^\s+Client\s+Mid\s+RSSI\.*\s+${BAND_SELECT_CLIENT_MID_RSSI}s*$$
  ^\s*$$
  ^. -> Error`
