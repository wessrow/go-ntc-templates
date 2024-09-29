package models

type CiscoXrShowConfigurationCommitList struct {
	Number	string	`json:"NUMBER"`
	Commit	string	`json:"COMMIT"`
	User	string	`json:"USER"`
	Line	string	`json:"LINE"`
	Client	string	`json:"CLIENT"`
	Date_time	string	`json:"DATE_TIME"`
}

var CiscoXrShowConfigurationCommitList_Template = `Value NUMBER (\d+)
Value COMMIT (\S+)
Value USER (\S+)
Value LINE (\S+)
Value CLIENT (\S+)
Value DATE_TIME (\w+\s\w+\s\d+\s\S+\s\d+|\w+\s\w+\s\s\d+\s\S+\s\d+)

Start
  ^${NUMBER}\s+${COMMIT}\s+${USER}\s+${LINE}\s+${CLIENT}\s+${DATE_TIME} -> Record
  ^${NUMBER}\s+${COMMIT}\s+${USER}\s+${LINE}\s+${CLIENT}\s+${DATE_TIME} -> Record

#
EOF

`