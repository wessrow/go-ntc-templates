package cisco_ios

type ShowIpOspfDatabaseRouter struct {
	Area             string `json:"AREA"`
	Lsa_id           string `json:"LSA_ID"`
	Lsa_adv_router   string `json:"LSA_ADV_ROUTER"`
	Lsa_length       string `json:"LSA_LENGTH"`
	Lsa_abr          string `json:"LSA_ABR"`
	Lsa_asbr         string `json:"LSA_ASBR"`
	Ls_link_id       string `json:"LS_LINK_ID"`
	Router_id        string `json:"ROUTER_ID"`
	Ls_mtid_metrics  string `json:"LS_MTID_METRICS"`
	Ls_tos_0_metrics string `json:"LS_TOS_0_METRICS"`
	Ls_link_data     string `json:"LS_LINK_DATA"`
	Lsa_checksum     string `json:"LSA_CHECKSUM"`
	Lsa_options      string `json:"LSA_OPTIONS"`
	Lsa_seq_number   string `json:"LSA_SEQ_NUMBER"`
	Lsa_num_links    string `json:"LSA_NUM_LINKS"`
	Ls_link_type     string `json:"LS_LINK_TYPE"`
	Lsa_type         string `json:"LSA_TYPE"`
	Lsa_age          string `json:"LSA_AGE"`
	Process_id       string `json:"PROCESS_ID"`
}

var ShowIpOspfDatabaseRouter_Template string = `Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
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
Value Filldown LSA_NUM_LINKS (\d+)
Value Filldown LSA_ABR (Area\s+Border\s+Router)
Value Filldown LSA_ASBR (AS\s+Boundary\s+Router)
Value LS_LINK_TYPE (.+?)
Value LS_LINK_ID (\d+\.\d+\.\d+\.\d+)
Value LS_LINK_DATA (\d+\.\d+\.\d+\.\d+)
Value LS_MTID_METRICS (\d+)
Value LS_TOS_0_METRICS (\d+)

Start
  ^\s*OSPF\s+Router\s+with\s+ID\s+\(${ROUTER_ID}\)\s+\(Process\s+ID\s+${PROCESS_ID}\)
  ^\s+Router\s+Link\s+States\s+\(Area ${AREA}\) -> LSAInfo
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
  ^\s+Number\s+of\s+Links:\s+${LSA_NUM_LINKS} -> LSAs
  ^\s*$$
  ^. -> Error

LSAs
  ^\s+Link\s+connected\s+to:\s+${LS_LINK_TYPE}$$
  ^\s+\(Link ID\)\s+((Network\/subnet\s+number)|(Designated\s+Router\s+address)|(Neighboring\s+Router ID)):\s+${LS_LINK_ID}
  ^\s+\(Link Data\)\s+((Network\s+Mask)|(Router\s+Interface\s+address)):\s+${LS_LINK_DATA}
  ^\s+Number\s+of\s+MTID\s+metrics:\s+${LS_MTID_METRICS}
  ^\s+TOS\s+0\s+Metrics:\s+${LS_TOS_0_METRICS} -> Next.Record
  ^\s+LS\s+age:\s+${LSA_AGE} -> LSAInfo
  ^\s+Router\s+Link\s+States\s+\(Area ${AREA}\) -> LSAInfo
  ^\s*$$
  ^. -> Error

EOF
`
