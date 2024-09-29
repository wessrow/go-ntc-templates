package huawei_smartax 

type DisplayBoardSerialNumber struct {
	Slot_id	string	`json:"SLOT_ID"`
	Board_name	string	`json:"BOARD_NAME"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
}

var DisplayBoardSerialNumber_Template = `Value Key SLOT_ID (\d+)
Value BOARD_NAME (\w*)
Value SERIAL_NUMBER (\w*)

Start
  ^\s+Solt\s+ID\s+Board\s+Name\s+Serial\s+Number
  ^\s+${SLOT_ID}\s+${BOARD_NAME}\s+${SERIAL_NUMBER} -> Record
  ^\s+-
  ^. -> Error`