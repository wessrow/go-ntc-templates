package cisco_ios

type ShowAliases struct {
	Mode    string `json:"MODE"`
	Alias   string `json:"ALIAS"`
	Command string `json:"COMMAND"`
}

var ShowAliases_Template string = `Value Filldown MODE (((\w+)(\s)?){1,})
Value ALIAS (\w+)
Value COMMAND ((\w+(\s|\S))+)

Start
  ^${MODE}\saliases: -> ALIAS
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

ALIAS
  ^${MODE}\saliases: -> ALIAS  
  ^\s+${ALIAS}\s+${COMMAND} -> Record`
