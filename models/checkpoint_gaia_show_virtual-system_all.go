package models

type CheckpointGaiaShowVirtualSystemAll struct {
	Instance_id	string	`json:"INSTANCE_ID"`
	Instance_name	string	`json:"INSTANCE_NAME"`
}

var CheckpointGaiaShowVirtualSystemAll_Template = `Value INSTANCE_ID (\d+)
Value INSTANCE_NAME ([a-zA-Z0-9_-]+)


Start
  ^${INSTANCE_ID}\s+${INSTANCE_NAME} -> Record
`