package models

type LinuxIpVrfShow struct {
	Name	string	`json:"NAME"`
	Table	string	`json:"TABLE"`
}

var LinuxIpVrfShow_Template = `Value Required NAME (\S+)
Value TABLE (\d+)

Start
  ^--- -> Start_record

Start_record
  ^${NAME}\s+${TABLE} -> Record
  ^. -> Error

`