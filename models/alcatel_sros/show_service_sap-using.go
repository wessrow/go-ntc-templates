package alcatel_sros

type ShowServiceSapUsing struct {
	Ingress_filter string `json:"INGRESS_FILTER"`
	Egress_qos     string `json:"EGRESS_QOS"`
	Egress_filter  string `json:"EGRESS_FILTER"`
	Admin_state    string `json:"ADMIN_STATE"`
	Oper_state     string `json:"OPER_STATE"`
	Port_id        string `json:"PORT_ID"`
	Service_id     string `json:"SERVICE_ID"`
	Ingress_qos    string `json:"INGRESS_QOS"`
}

var ShowServiceSapUsing_Template string = `Value Required PORT_ID (\S+)
Value Required SERVICE_ID (\d+)
Value Required INGRESS_QOS (\S+)
Value Required INGRESS_FILTER (\S+)
Value Required EGRESS_QOS (\d+)
Value Required EGRESS_FILTER (\S+)
Value Required ADMIN_STATE (Up|Down)
Value Required OPER_STATE (Up|Down)

Start
  ^=+
  ^Service\s+Access\s+Points
  ^PortId\s+SvcId\s+Ing.\s+Ing.\s+Egr.\s+Egr.\s+Adm\s+Opr\s*$$ -> SAPS
  ^\s*$$
  ^. -> Error

SAPS
  ^\s+QoS\s+Fltr\s+QoS\s+Fltr        
  ^-+
  ^${PORT_ID}\s+${SERVICE_ID}\s+${INGRESS_QOS}\s+${INGRESS_FILTER}\s+${EGRESS_QOS}\s+${EGRESS_FILTER}\s+${ADMIN_STATE}\s+${OPER_STATE} -> Record
  ^=+
  ^Number\s+of
  ^\*\s+indicates
  ^\s*$$
  ^. -> Error
`
