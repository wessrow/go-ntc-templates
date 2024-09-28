package models

type AlcatelSrosShowServiceSdp struct {
	Sdp_id	string	`json:"SDP_ID"`
	Adm_mtu	string	`json:"ADM_MTU"`
	Opr_mtu	string	`json:"OPR_MTU"`
	Far_end	string	`json:"FAR_END"`
	Adm	string	`json:"ADM"`
	Opr	string	`json:"OPR"`
	Del	string	`json:"DEL"`
	Lsp	string	`json:"LSP"`
	Sig	string	`json:"SIG"`
}