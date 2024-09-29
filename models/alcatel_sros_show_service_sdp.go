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

var AlcatelSrosShowServiceSdp_Template = `Value Required SDP_ID (\d+)
Value Required ADM_MTU (\d+)
Value Required OPR_MTU (\d+)
Value FAR_END (\d+.\d+.\d+.\d+)
Value Required ADM (Up|Down)
Value Required OPR (Up|Down)
Value Required DEL (\S+)
Value Required LSP (\S+)
Value Required SIG (\S+)

Start
  ^=+
  ^Services:
  ^SdpId\s+AdmMTU\s+OprMTU\s+Far\s+End\s+Adm\s+Opr\s+Del\s+LSP\s+Sig\s*$$
  ^-+ -> SDP
  ^\s*$$
  ^. -> Error

SDP
  ^${SDP_ID}\s+${ADM_MTU}\s+${OPR_MTU}\s+${FAR_END}\s+${ADM}\s+${OPR}\s+${DEL}\s+${LSP}\s+${SIG} -> Record
  ^${SDP_ID}\s+${ADM_MTU}\s+${OPR_MTU}\s+${ADM}\s+${OPR}\s+${DEL}\s+${LSP}\s+${SIG} -> Record
  ^-+ -> Done
  ^\s*$$
  ^=+
  ^. -> Error

Done

`