package checkpoint_gaia 

type ShowVirtualSystemAll struct {
	Instance_id	string	`json:"INSTANCE_ID"`
	Instance_name	string	`json:"INSTANCE_NAME"`
}

var ShowVirtualSystemAll_Template = `Value INSTANCE_ID (\d+)
Value INSTANCE_NAME ([a-zA-Z0-9_-]+)


Start
  ^${INSTANCE_ID}\s+${INSTANCE_NAME} -> Record`