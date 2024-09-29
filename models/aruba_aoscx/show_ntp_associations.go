package aruba_aoscx 

type ShowNtpAssociations struct {
	Id	string	`json:"ID"`
	Name	string	`json:"NAME"`
	Remote	string	`json:"REMOTE"`
	Ref_id	string	`json:"REF_ID"`
	Stratum	string	`json:"STRATUM"`
	Last	string	`json:"LAST"`
	Poll	string	`json:"POLL"`
	Reach	string	`json:"REACH"`
}

var ShowNtpAssociations_Template = `Value ID (\W?\s+\d+)
Value NAME (\d+\.\d+\.\d+\.\d+)
Value REMOTE (\d+\.\d+\.\d+\.\d+)
Value REF_ID (\d+\.\d+\.\d+\.\d+)
Value STRATUM (\d+)
Value LAST (\d+)
Value POLL (\d+)
Value REACH (\d+)

Start
  ^-+$$
  ^\s+ID\s+NAME\s+REMOTE\s+REF-ID\s+ST\s+LAST\s+POLL\s+REACH
  ^${ID}\s+${NAME}\s+${REMOTE}\s+${REF_ID}\s+${STRATUM}\s+${LAST}\s+${POLL}\s+${REACH} -> Record
  ^. -> Error
`