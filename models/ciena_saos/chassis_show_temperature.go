package ciena_saos

type ChassisShowTemperature struct {
	Current string `json:"CURRENT"`
	Low     string `json:"LOW"`
	High    string `json:"HIGH"`
}

var ChassisShowTemperature_Template string = `Value CURRENT (\d+\s*\w+)
Value LOW (\d+\s*\w+)
Value HIGH (\d+\s*\w+)

Start
  ^\+-+\s*TEMPERATURE\s*THRESHOLD\s*-+\+
  ^\|\s*Low\s*\|\s*High\s*\|
  ^\+-+\+-+\+
  ^\|\s*-\d+\s*\w+\s*\|\s*\d+\s*\w+\s*\|
  ^\+-+\+-+\+
  ^\s*$$
  ^\+-+\s*TEMPERATURE\s*STATUS\s*-+\+
  ^\|\s*Current\s*\|\s*Low\s*\|\s*High\s*\|
  ^\+-+\+-+\+-+\+
  ^\|\s*${CURRENT}\s*\|\s*${LOW}\s*\|\s*${HIGH}\s*\|
  ^\+-+\+-+\+-+\+
  ^\s*$$
  ^. -> Error`
