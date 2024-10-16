package alcatel_sros

type ShowServiceIdBase struct {
	Oper_state  string `json:"OPER_STATE"`
	Sap_count   string `json:"SAP_COUNT"`
	Sdp_count   string `json:"SDP_COUNT"`
	Mtu         string `json:"MTU"`
	Service_id  string `json:"SERVICE_ID"`
	Customer_id string `json:"CUSTOMER_ID"`
	Description string `json:"DESCRIPTION"`
	Admin_state string `json:"ADMIN_STATE"`
}

var ShowServiceIdBase_Template string = `Value SERVICE_ID (\d+)
Value CUSTOMER_ID (\d+)
Value DESCRIPTION ([\w|\s]+|\(Not\sSpecified\))
Value ADMIN_STATE ([uU][pP]|[Dd][oO][wW][nN])
Value OPER_STATE ([uU][pP]|[Dd][oO][wW][nN])
Value SAP_COUNT ([0-9]{1,1500})
Value SDP_COUNT ([0-9]{1,150})
Value MTU (\d+)

Start
  ^Customer\sId\s+.\s${CUSTOMER_ID}
  ^MTU\s+:\s${MTU}
  ^Service\sId\s+:\s${SERVICE_ID}
  ^Description\s+:\s${DESCRIPTION}
  ^Admin\sState\s+:\s${ADMIN_STATE}\s+Oper\sState\s+:\s${OPER_STATE}
  ^SAP\sCount\s+:\s${SAP_COUNT}\s+SDP\sBind\sCount\s+:\s${SDP_COUNT} -> Record
`
