package avaya_ers 

type ShowMltAllMembers struct {
	Mltid	string	`json:"MLTID"`
	Mltname	string	`json:"MLTNAME"`
	Mltactivemembers	string	`json:"MLTACTIVEMEMBERS"`
	Mltinactivemembers	string	`json:"MLTINACTIVEMEMBERS"`
	Mltbpdu	string	`json:"MLTBPDU"`
	Mltmode	string	`json:"MLTMODE"`
	Mltstatus	string	`json:"MLTSTATUS"`
	Mlttype	string	`json:"MLTTYPE"`
	Mltlacpkey	string	`json:"MLTLACPKEY"`
}

var ShowMltAllMembers_Template = `Value Required MLTID (\d+)
Value MLTNAME (.*?)
Value MLTACTIVEMEMBERS (\S+?)
Value MLTINACTIVEMEMBERS (\S+?)
Value MLTBPDU (\S+?)
Value MLTMODE (\S+?)
Value MLTSTATUS (Enabled|Disabled)
Value MLTTYPE (Access|Trunk|)
Value MLTLACPKEY (\S+|)

Start
  ^Id:\s+\d+\s*$$ -> Continue.Record
  ^Id:\s*${MLTID}\s*$$
  ^\s+Name:${MLTNAME}\s*$$
  ^\s+Active Members:\s${MLTACTIVEMEMBERS}\s*$$
  ^\s+Inactive Members:\s${MLTINACTIVEMEMBERS}\s*$$
  ^\s+Bpdu:\s${MLTBPDU}\s*$$
  ^\s+Mode:\s${MLTMODE}\s*$$
  ^\s+Status:\s${MLTSTATUS}\s*$$
  ^\s+Type:\s${MLTTYPE}\s*$$
  ^\s+LACP key:\s${MLTLACPKEY}\s*$$
  ^.*$$ -> Error
`