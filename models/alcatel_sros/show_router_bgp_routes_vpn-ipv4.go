package alcatel_sros

type ShowRouterBgpRoutesVpnIpv4 struct {
	Label         string `json:"LABEL"`
	As_path       string `json:"AS_PATH"`
	Network       string `json:"NETWORK"`
	Next_hop      string `json:"NEXT_HOP"`
	Prefix_length string `json:"PREFIX_LENGTH"`
	Local_pref    string `json:"LOCAL_PREF"`
	Med           string `json:"MED"`
	Path_id       string `json:"PATH_ID"`
	In_out_use    string `json:"IN_OUT_USE"`
	Rd            string `json:"RD"`
}

var ShowRouterBgpRoutesVpnIpv4_Template string = `Value IN_OUT_USE (\*>i|u\*>i|\*i)
Value RD (\d+\:\d+)
Value Required NETWORK (\d+(\.\d+){3})
Value Required PREFIX_LENGTH (\d{1,2})
Value LOCAL_PREF (\d+)
Value MED (None|d\+)
Value NEXT_HOP (\d+\.\d+\.\d+\.\d+)
Value PATH_ID (None|\d+)
Value LABEL (\d+)
Value AS_PATH (No As-Path|\d+)

Start
  ^----------- -> Prefix

Prefix
  ^${IN_OUT_USE}\s+${RD}\:${NETWORK}\/${PREFIX_LENGTH}\s+${LOCAL_PREF}\s+${MED}
  ^\s+${NEXT_HOP}\s+${PATH_ID}\s+${LABEL}
  ^\s+${AS_PATH} -> Record
`
