package cisco_nvfis

type ShowNic struct {
	Devid   string `json:"DEVID"`
	Mode    string `json:"MODE"`
	Devno   string `json:"DEVNO"`
	Pnics   string `json:"PNICS"`
	Slotid  string `json:"SLOTID"`
	Adapter string `json:"ADAPTER"`
	Vendor  string `json:"VENDOR"`
}

var ShowNic_Template string = `Value SLOTID (\d+)
Value ADAPTER (.+?)
Value VENDOR (\d+)
Value DEVID (\d+)
Value MODE (\w+)
Value DEVNO (\w+)
Value PNICS (\[.+\])

Start
  ^${SLOTID}\s+${ADAPTER}\s+${VENDOR}\s+${DEVID}\s+${MODE}\s+${DEVNO}\s+${PNICS} -> Record

EOF`
