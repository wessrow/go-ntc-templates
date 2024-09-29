package models

type AlcatelSrosShowRouterBgpRoutesVpnIpv4 struct {
	In_out_use	string	`json:"IN_OUT_USE"`
	Rd	string	`json:"RD"`
	Network	string	`json:"NETWORK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Med	string	`json:"MED"`
	Next_hop	string	`json:"NEXT_HOP"`
	Path_id	string	`json:"PATH_ID"`
	Label	string	`json:"LABEL"`
	As_path	string	`json:"AS_PATH"`
}

var AlcatelSrosShowRouterBgpRoutesVpnIpv4_Template = `Value IN_OUT_USE (\*>i|u\*>i|\*i)
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