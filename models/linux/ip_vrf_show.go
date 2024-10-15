package linux

type IpVrfShow struct {
	Name  string `json:"NAME"`
	Table string `json:"TABLE"`
}

var IpVrfShow_Template string = `Value Required NAME (\S+)
Value TABLE (\d+)

Start
  ^--- -> Start_record

Start_record
  ^${NAME}\s+${TABLE} -> Record
  ^. -> Error
`
