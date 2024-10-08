package oneaccess_oneos 

type ShowIsdnActive struct {
	App	string	`json:"APP"`
	Call_type	string	`json:"CALL_TYPE"`
	Calling_nbr	string	`json:"CALLING_NBR"`
	Called_nbr	string	`json:"CALLED_NBR"`
	Duration	string	`json:"DURATION"`
	Port	string	`json:"PORT"`
	Bchannel	string	`json:"BCHANNEL"`
	Call_ref	string	`json:"CALL_REF"`
	Call_id	string	`json:"CALL_ID"`
}

var ShowIsdnActive_Template = `Value Required APP (\S+)
Value CALL_TYPE (\S+)
Value CALLING_NBR (\S+)
Value CALLED_NBR (\S+)
Value DURATION (\d+:\d+:\d+)
Value PORT (\S+)
Value BCHANNEL (\d+)
Value CALL_REF (\d+)
Value CALL_ID (\d+)

Start
  ^${APP}\s+${CALL_TYPE}\s+${CALLING_NBR}\s+${CALLED_NBR}\s+${DURATION}\s+${PORT}\s+${BCHANNEL}\s+${CALL_REF}\s+${CALL_ID} -> Record
  ^\s*-+\s*$$
  ^\s+ISDN\s+ACTIVE\s+CALLS\s*$$
  ^\s*App\.\s+Call\s+Calling\s+Called\s+Call\s+Port\s+BChan\s+Call-ref\s+call-id\s*$$
  ^\s+Type\s+Number\s+Number\s+Duration\s*$$
  ^\s*no\s+isdn\s+active\s+calls                           
  ^\s*$$
  ^. -> Error
`