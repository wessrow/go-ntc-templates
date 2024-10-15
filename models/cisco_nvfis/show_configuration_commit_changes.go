package cisco_nvfis

type ShowConfigurationCommitChanges struct {
	Date_time  string `json:"DATE_TIME"`
	Client     string `json:"CLIENT"`
	Created_by string `json:"CREATED_BY"`
}

var ShowConfigurationCommitChanges_Template string = `Value CREATED_BY (\w+)
Value DATE_TIME (\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})
Value CLIENT (\w+)


Start
  ^!\sCreated\sby:\s${CREATED_BY}
  ^!\sDate:\s${DATE_TIME}
  ^!\sClient:\s${CLIENT} -> Record

EOF`
