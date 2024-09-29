package cisco_nvfis 

type ShowConfigurationCommitChanges struct {
	Created_by	string	`json:"CREATED_BY"`
	Date_time	string	`json:"DATE_TIME"`
	Client	string	`json:"CLIENT"`
}

var ShowConfigurationCommitChanges_Template = `Value CREATED_BY (\w+)
Value DATE_TIME (\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})
Value CLIENT (\w+)


Start
  ^!\sCreated\sby:\s${CREATED_BY}
  ^!\sDate:\s${DATE_TIME}
  ^!\sClient:\s${CLIENT} -> Record

EOF`