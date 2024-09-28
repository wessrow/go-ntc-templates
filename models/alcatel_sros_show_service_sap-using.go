package models

type AlcatelSrosShowServiceSapUsing struct {
	Port_id	string	`json:"PORT_ID"`
	Service_id	string	`json:"SERVICE_ID"`
	Ingress_qos	string	`json:"INGRESS_QOS"`
	Ingress_filter	string	`json:"INGRESS_FILTER"`
	Egress_qos	string	`json:"EGRESS_QOS"`
	Egress_filter	string	`json:"EGRESS_FILTER"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Oper_state	string	`json:"OPER_STATE"`
}