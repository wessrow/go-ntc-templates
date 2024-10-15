package cisco_asa

type ShowResourceUsage struct {
	Peak     string `json:"PEAK"`
	Limit    string `json:"LIMIT"`
	Denied   string `json:"DENIED"`
	Context  string `json:"CONTEXT"`
	Resource string `json:"RESOURCE"`
	Current  string `json:"CURRENT"`
}

var ShowResourceUsage_Template string = `Value Required RESOURCE (.+?)
Value CURRENT (\d+)
Value PEAK (\d+)
Value LIMIT (\S+)
Value DENIED (\d+)
Value CONTEXT (.+?)


Start
  ^Resource\s+Current\s+Peak\s+Limit\s+Denied\s+Context\s*$$ -> Start_record
  ^\s*$$
  ^. -> Error

Start_record
  ^${RESOURCE}\s+${CURRENT}\s+${PEAK}\s+${LIMIT}\s+${DENIED}\s+${CONTEXT}\s*$$ -> Record
  ^S\s+=\s+System:\s+Combined\s+context\s+limits
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
`
