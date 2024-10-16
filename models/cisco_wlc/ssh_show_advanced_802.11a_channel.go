package cisco_wlc

type SshShowAdvanced80211aChannel struct {
	Interfer            string   `json:"INTERFER"`
	Dca_min             string   `json:"DCA_MIN"`
	Unused_dca_channels []string `json:"UNUSED_DCA_CHANNELS"`
	Dca_sensitivity     string   `json:"DCA_SENSITIVITY"`
	Dca_channels        []string `json:"DCA_CHANNELS"`
	Assign_mode         string   `json:"ASSIGN_MODE"`
	Anchor_time         string   `json:"ANCHOR_TIME"`
	Load                string   `json:"LOAD"`
	Leader              string   `json:"LEADER"`
	Dca_width           string   `json:"DCA_WIDTH"`
	Update_time         string   `json:"UPDATE_TIME"`
	Noise               string   `json:"NOISE"`
	Device_aware        string   `json:"DEVICE_AWARE"`
	Cleanair_driven     string   `json:"CLEANAIR_DRIVEN"`
}

var SshShowAdvanced80211aChannel_Template string = `Value ASSIGN_MODE (\w+)
Value UPDATE_TIME (\d+ \w+)
Value ANCHOR_TIME (\d+)
Value NOISE (\w+)
Value INTERFER (\w+)
Value LOAD (\w+)
Value DEVICE_AWARE (\w+)
Value CLEANAIR_DRIVEN (\w+)
Value LEADER (\w+)
Value DCA_SENSITIVITY (\w+ \(\d+ dB\))
Value DCA_WIDTH (\d+ \w+)
Value DCA_MIN (-\d+ dBm)
Value List DCA_CHANNELS ((\d+,?)*?)
Value List UNUSED_DCA_CHANNELS ((\d+,?)*?)

Start
  ^\s*Leader\sAutomatic\sChannel\sAssignments*$$
  ^\s+Channel\sAssignment\sMode\.*\s${ASSIGN_MODE}s*$$
  ^\s+Channel\sUpdate\sInterval\.*\s${UPDATE_TIME}s*$$
  ^\s+Anchor\stime\s\(Hour\sof\sthe day\)\.*\s+${ANCHOR_TIME}\s+s*$$
  ^\s+Update\sContributions*$$
  ^\s+Noise\.*\s${NOISE}s*$$
  ^\s+Interference\.*\s${INTERFER}s*$$
  ^\s+Load\.*\s${LOAD}s*$$
  ^\s+Device\sAware\.*\s${DEVICE_AWARE}s*$$
  ^\s+CleanAir\sEvent-driven\sRRM option\.*\s${CLEANAIR_DRIVEN}s*$$
  ^\s+Channel\sAssignment\sLeader\.*\s+${LEADER}\s+\(([\d1-9]+\.?){4}\)\s+\(\:\:\)s*$$
  ^\s+Last\s+Run\.*\s+\d+\s+seconds\s+agos*$$
  ^\s+Last\s+Run\s+Time\.*\s+\d+\s+secondss*$$
  ^\s+DCA\sSensitivity\sLevel\:?\s?\.*\s${DCA_SENSITIVITY}s*$$
  ^\s+DCA\s\S+\sChannel\sWidth\.*\s${DCA_WIDTH}s*$$
  ^\s+DCA\sMinimum\sEnergy\sLimit\.*\s${DCA_MIN}s*$$
  ^\s+Channel\sEnergy\sLevels
  ^\s+Minimum
  ^\s+Average
  ^\s+Maximum
  ^\s+Channel\sDwell\sTimes
  ^\s+802\.11a\s5\sGHz\sAuto\-RF Channel Lists*$$ -> Channels
  ^.*Allowed\sChannel\sList
  ^\s+Unused\sChannel\sList
  ^\s+\d+,?
  ^\s+DCA\sOutdoor\sAP\soption
  ^\s*$$
  ^. -> Error

Channels
  ^.*Allowed\sChannel\sList\.*\s${DCA_CHANNELS}s*$$
  ^\s+${DCA_CHANNELS}s*$$
  ^\s+Unused Channel List.*\s${UNUSED_DCA_CHANNELS}s*$$
  ^\s+802.11a\s4.9\sGHz\sAuto-RF Channel Lists*$$ -> Start


`
