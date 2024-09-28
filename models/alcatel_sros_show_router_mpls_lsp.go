package models

type AlcatelSrosShowRouterMplsLsp struct {
	Lsp_name	string	`json:"LSP_NAME"`
	To	string	`json:"TO"`
	Tunnel_id	string	`json:"TUNNEL_ID"`
	Fastfail_config	string	`json:"FASTFAIL_CONFIG"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Oper_state	string	`json:"OPER_STATE"`
}