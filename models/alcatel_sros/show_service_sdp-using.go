package alcatel_sros

type ShowServiceSdpUsing struct {
	Sdp_id        string `json:"SDP_ID"`
	Sdp_svc_id    string `json:"SDP_SVC_ID"`
	Type          string `json:"TYPE"`
	Far_end       string `json:"FAR_END"`
	Oper_state    string `json:"OPER_STATE"`
	Ingress_label string `json:"INGRESS_LABEL"`
	Egress_label  string `json:"EGRESS_LABEL"`
	Service_id    string `json:"SERVICE_ID"`
}

var ShowServiceSdpUsing_Template string = `Value Required SERVICE_ID (\d+)
Value Required SDP_ID (\d+)
Value Required SDP_SVC_ID (\d+)
Value Required TYPE (\S+)
Value Required FAR_END (\S+)
Value Required OPER_STATE (\S+)
Value Required INGRESS_LABEL (\S+)
Value Required EGRESS_LABEL (\S+)

Start
  ^=+
  ^SDP\s+Using
  ^SvcId\s+SdpId\s+Type\s+Far\s+End\s+Opr\s+I.Label\s+E.Label\s*$$ -> SDP
  ^\s*$$
  ^. -> Error

SDP
  ^\s+State         
  ^-+
  ^${SERVICE_ID}\s+${SDP_ID}:${SDP_SVC_ID}\s+${TYPE}\s+${FAR_END}\s+${OPER_STATE}\s+${INGRESS_LABEL}\s+${EGRESS_LABEL} -> Record
  ^-+ -> Done
  ^Number\s+of
  ^=+
  ^\s*$$
  ^. -> Error

Done
`
