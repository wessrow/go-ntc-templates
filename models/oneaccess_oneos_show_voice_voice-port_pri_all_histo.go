package models

type OneaccessOneosShowVoiceVoicePortPriAllHisto struct {
	Port	string	`json:"PORT"`
	Physical_type	string	`json:"PHYSICAL_TYPE"`
	Proto_descriptor	string	`json:"PROTO_DESCRIPTOR"`
	Attached_voip_dial_peer	string	`json:"ATTACHED_VOIP_DIAL_PEER"`
	Date_last_reset	string	`json:"DATE_LAST_RESET"`
	Max_channels_used	string	`json:"MAX_CHANNELS_USED"`
	Max_channels_available	string	`json:"MAX_CHANNELS_AVAILABLE"`
}

var OneaccessOneosShowVoiceVoicePortPriAllHisto_Template = `Value Required PORT (\d+\/\d+)
Value PHYSICAL_TYPE (\S+)
Value PROTO_DESCRIPTOR (\S+)
Value ATTACHED_VOIP_DIAL_PEER (\d+)
Value DATE_LAST_RESET (.*)
Value MAX_CHANNELS_USED (\d+)
Value MAX_CHANNELS_AVAILABLE (\d+)

Start
  ^\s+voice\s+port -> Continue.Record
  ^\s*voice\sport\s+${PORT}
  ^\s*physical\stype\s+${PHYSICAL_TYPE}
  ^\s*protocol\sdescriptor\s+${PROTO_DESCRIPTOR}
  ^\s*attached\svoip\sdial\speer\s+${ATTACHED_VOIP_DIAL_PEER}
  ^\s*date\sof\slast\sreset\s+${DATE_LAST_RESET}
  ^\s*max\schannel\(s\)\sused\s+${MAX_CHANNELS_USED}\/${MAX_CHANNELS_AVAILABLE}
  ^\s+daily\s+occupancy
  ^\s+%\s*$$
  ^\s+\d+\s+(\||\+-)
  ^\s+\d+\s+\d+
  ^\s*$$
  ^. -> Error

`