package alcatel_sros

type ShowRouterMplsLsp struct {
	Admin_state     string `json:"ADMIN_STATE"`
	Oper_state      string `json:"OPER_STATE"`
	Lsp_name        string `json:"LSP_NAME"`
	To              string `json:"TO"`
	Tunnel_id       string `json:"TUNNEL_ID"`
	Fastfail_config string `json:"FASTFAIL_CONFIG"`
}

var ShowRouterMplsLsp_Template string = `Value Required LSP_NAME (\S+)
Value Required TO (\d+.\d+.\d+.\d+)
Value Required TUNNEL_ID (\d+)
Value Required FASTFAIL_CONFIG (\S+)
Value Required ADMIN_STATE (Up|Dwn)
Value Required OPER_STATE (Up|Dwn)

Start
  ^=+
  ^MPLS\s+LSPs
  ^LSP\s+Name\s+Tun\s+Fastfail\s+Adm\s+Opr\s*$$ -> LSP
  ^\s*$$
  ^. -> Error

LSP
  ^\s+To\s+Id\s+Config
  ^-+
  ^${LSP_NAME}\s+${TUNNEL_ID}\s+${FASTFAIL_CONFIG}\s+${ADMIN_STATE}\s+${OPER_STATE}
  ^\s+${TO} -> Record
  ^-+ -> Done
  ^LSPs\s+:
  ^=+
  ^\s*$$
  ^. -> Error

Done
`
