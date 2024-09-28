package models

type AlcatelSrosShowServiceIdBase struct {
	Service_id	string	`json:"SERVICE_ID"`
	Customer_id	string	`json:"CUSTOMER_ID"`
	Description	string	`json:"DESCRIPTION"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Oper_state	string	`json:"OPER_STATE"`
	Sap_count	string	`json:"SAP_COUNT"`
	Sdp_count	string	`json:"SDP_COUNT"`
	Mtu	string	`json:"MTU"`
}