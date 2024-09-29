package linux_ip 

type VrfShow struct {
	Name	string	`json:"NAME"`
	Table	string	`json:"TABLE"`
}

var VrfShow_Template = `Value Required NAME (\S+)
Value TABLE (\d+)

Start
  ^--- -> Start_record

Start_record
  ^${NAME}\s+${TABLE} -> Record
  ^. -> Error
`