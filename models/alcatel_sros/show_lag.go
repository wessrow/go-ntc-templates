package alcatel_sros

type ShowLag struct {
	Opr          string `json:"OPR"`
	Weighted     string `json:"WEIGHTED"`
	Threshold    string `json:"THRESHOLD"`
	Up_count     string `json:"UP_COUNT"`
	Mc_act_stdby string `json:"MC_ACT_STDBY"`
	Lag_id       string `json:"LAG_ID"`
	Adm          string `json:"ADM"`
}

var ShowLag_Template string = `Value LAG_ID (\d+)
Value ADM (up|down)
Value OPR (up|down)
Value WEIGHTED (Yes|No)
Value THRESHOLD (\d+)
Value UP_COUNT (\d+)
Value MC_ACT_STDBY (N/A|active|standby)

Start
  ^----------- -> Lag

Lag
  ^${LAG_ID}\s+${ADM}\s+${OPR}\s+${WEIGHTED}\s+${THRESHOLD}\s+${UP_COUNT}\s+${MC_ACT_STDBY}(\s|$$) -> Record
  ^\s*$$
  ^-----------
  ^===========
  ^Total
  ^. -> Error
`
