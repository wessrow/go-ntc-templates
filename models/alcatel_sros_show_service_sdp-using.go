package models

type AlcatelSrosShowServiceSdpUsing struct {
	Service_id	string	`json:"SERVICE_ID"`
	Sdp_id	string	`json:"SDP_ID"`
	Sdp_svc_id	string	`json:"SDP_SVC_ID"`
	Type	string	`json:"TYPE"`
	Far_end	string	`json:"FAR_END"`
	Oper_state	string	`json:"OPER_STATE"`
	Ingress_label	string	`json:"INGRESS_LABEL"`
	Egress_label	string	`json:"EGRESS_LABEL"`
}