package huawei_smartax

type DisplayOntInfoDescription struct {
	Ont_id      string `json:"ONT_ID"`
	Description string `json:"DESCRIPTION"`
	Fsp         string `json:"FSP"`
}

var DisplayOntInfoDescription_Template string = `Value Required FSP (\w+\/\s\w+\/\w+)
Value Key ONT_ID (\d+)
Value DESCRIPTION (.+?)


Start
  ^\s+F\/S\/P\s+(ONT|ONT-ID)\s+Description -> Descriptions

Descriptions
  ^\s+ID
  ^\s+-
  ^\s+${FSP}\s+${ONT_ID}\s+${DESCRIPTION}\s*$$ -> Record
  ^\s+In\s+port\s+\w+\/\s*\w+\/\w+\s*,\s+the\s+total(\s*number)?\s+of\s+ONTs\s+(is|are):\s+\w+(\,\s*online:\s*\w+)?(\,\s*offline:\s*\w+)? -> EOF
  ^. -> Error`
