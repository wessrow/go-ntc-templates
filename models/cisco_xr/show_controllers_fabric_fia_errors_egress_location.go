package cisco_xr

type ShowControllersFabricFiaErrorsEgressLocation struct {
	To_spaui_error_1       string `json:"TO_SPAUI_ERROR_1"`
	Rl_over_under_flow_cnt string `json:"RL_OVER_UNDER_FLOW_CNT"`
	Fia                    string `json:"FIA"`
	Category               string `json:"CATEGORY"`
	To_spaui_error_0       string `json:"TO_SPAUI_ERROR_0"`
}

var ShowControllersFabricFiaErrorsEgressLocation_Template string = `Value FIA (\S+)
Value CATEGORY (\S+)
Value TO_SPAUI_ERROR_0 (\d+)
Value TO_SPAUI_ERROR_1 (\d+)
Value RL_OVER_UNDER_FLOW_CNT (\d+)

Start
  ^\s*\*+\s+\S+ -> Continue.Record
  ^\s*\*+\s+${FIA}\s+\*+
  ^Category:\s+${CATEGORY}
  ^\s+To\s+Spaui\s+Error-0\s+${TO_SPAUI_ERROR_0}
  ^\s+To\s+Spaui\s+Error-1\s+${TO_SPAUI_ERROR_1}
  ^\s+RL\s+over/under\s+flow\s+cnt\s+${RL_OVER_UNDER_FLOW_CNT}
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
`
