package cisco_nxos

type ShowIpCommunityList struct {
	Community string   `json:"COMMUNITY"`
	Type      string   `json:"TYPE"`
	Name      string   `json:"NAME"`
	Seq       string   `json:"SEQ"`
	Action    string   `json:"ACTION"`
	As        []string `json:"AS"`
	Tag       []string `json:"TAG"`
}

var ShowIpCommunityList_Template string = `Value Required,Filldown TYPE (Standard|Expanded)
Value Required,Filldown NAME (\S+)
Value Required SEQ (\d+)
Value Required ACTION (permit|deny)
Value List AS (\d+)
Value List TAG (\d+)
Value COMMUNITY (\D+)

Start
  ^${TYPE}\s+Community\s+List\s+${NAME}\s*$$
  ^\s+${SEQ}\s+${ACTION}\s+\"*${AS}:${TAG}\"* -> Continue
  ^\s+\d+\s+(?:permit|deny)\s+\"*\d+:\d+\"*\s+\"*${AS}:${TAG}\"* -> Continue
  ^\s+\d+\s+(?:permit|deny)\s+(?:\"*\d+:\d+\"*\s+){2}\"*${AS}:${TAG}\"* -> Continue
  ^\s+\d+\s+(?:permit|deny)\s+(?:\"*\d+:\d+\"*\s+){3}\"*${AS}:${TAG}\"* -> Continue
  ^\s+.+\"*\d+:\d+\"*\s+${COMMUNITY}\s*$$ -> Continue
  ^\s+${SEQ}\s+${ACTION}\s+${COMMUNITY}\s*$$ -> Continue
  ^\s+ -> Record
  ^$$
  ^.*$$ -> Error
`
