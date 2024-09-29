package models

type CiscoIosShowIpOspfDatabaseNetwork struct {
	Router_id	string	`json:"ROUTER_ID"`
	Process_id	string	`json:"PROCESS_ID"`
	Area	string	`json:"AREA"`
	Lsa_age	string	`json:"LSA_AGE"`
	Lsa_options	string	`json:"LSA_OPTIONS"`
	Lsa_type	string	`json:"LSA_TYPE"`
	Lsa_id	string	`json:"LSA_ID"`
	Lsa_adv_router	string	`json:"LSA_ADV_ROUTER"`
	Lsa_seq_number	string	`json:"LSA_SEQ_NUMBER"`
	Lsa_checksum	string	`json:"LSA_CHECKSUM"`
	Lsa_length	string	`json:"LSA_LENGTH"`
	Lsa_network_mask	string	`json:"LSA_NETWORK_MASK"`
	Lsa_abr	string	`json:"LSA_ABR"`
	Lsa_asbr	string	`json:"LSA_ASBR"`
	Ls_att_router	string	`json:"LS_ATT_ROUTER"`
}

var CiscoIosShowIpOspfDatabaseNetwork_Template = `Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown PROCESS_ID (\d+)
Value Filldown AREA (\d+\.\d+\.\d+\.\d+|\d+)
Value Filldown LSA_AGE (\d+)
Value Filldown LSA_OPTIONS (.+)
Value Filldown LSA_TYPE (.+)
Value Filldown LSA_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown LSA_ADV_ROUTER (\d+\.\d+\.\d+\.\d+)
Value Filldown LSA_SEQ_NUMBER (\S+)
Value Filldown LSA_CHECKSUM (\S+)
Value Filldown LSA_LENGTH (\d+)
Value Filldown LSA_NETWORK_MASK (\S+)
Value Filldown LSA_ABR (Area\s+Border\s+Router)
Value Filldown LSA_ASBR (AS\s+Boundary\s+Router)
Value LS_ATT_ROUTER (\d+\.\d+\.\d+\.\d+)

Start
  ^\s*OSPF\s+Router\s+with\s+ID\s+\(${ROUTER_ID}\)\s+\(Process\s+ID\s+${PROCESS_ID}\)
  ^\s+Net\s+Link\s+States\s+\(Area ${AREA}\) -> LSAInfo
  ^\s*$$
  ^. -> Error

LSAInfo
  ^\s+LS\s+age:\s+${LSA_AGE}
  ^\s+Options:\s+\(${LSA_OPTIONS}\)
  ^\s+LS\s+Type:\s+${LSA_TYPE}
  ^\s+Link\s+State\s+ID:\s+${LSA_ID}
  ^\s+Advertising\s+Router:\s+${LSA_ADV_ROUTER}
  ^\s+LS\s+Seq\s+Number:\s+${LSA_SEQ_NUMBER}
  ^\s+Checksum:\s+${LSA_CHECKSUM}
  ^\s+Length:\s+${LSA_LENGTH}
  ^\s+${LSA_ABR}$$
  ^\s+${LSA_ASBR}$$
  ^\s+Network\s+Mask:\s+${LSA_NETWORK_MASK} -> LSAs
  ^\s*$$
  ^. -> Error

LSAs
  ^\s+Attached\s+Router:\s+${LS_ATT_ROUTER} -> Next.Record
  ^\s+LS\s+age:\s+${LSA_AGE} -> LSAInfo
  ^\s+Router\s+Link\s+States\s+\(Area ${AREA}\) -> LSAInfo
  ^\s*$$
  ^. -> Error

EOF

`