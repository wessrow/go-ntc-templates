package models

type OneaccessOneosShowCellularRadioNetwork struct {
	Operator	string	`json:"OPERATOR"`
	Signal_strength	string	`json:"SIGNAL_STRENGTH"`
	Rssi	string	`json:"RSSI"`
	Rsrq	string	`json:"RSRQ"`
	Rsrp	string	`json:"RSRP"`
	Snr	string	`json:"SNR"`
	Radio_technology	string	`json:"RADIO_TECHNOLOGY"`
	Register_state	string	`json:"REGISTER_STATE"`
	Attach_state	string	`json:"ATTACH_STATE"`
	Resets_loss_registrations	string	`json:"RESETS_LOSS_REGISTRATIONS"`
	Resets_failed_registrations	string	`json:"RESETS_FAILED_REGISTRATIONS"`
	Resets_modem	string	`json:"RESETS_MODEM"`
	Resets_unknown	string	`json:"RESETS_UNKNOWN"`
	Lac	string	`json:"LAC"`
	Cell_id	string	`json:"CELL_ID"`
	Tac	string	`json:"TAC"`
	Plmn	string	`json:"PLMN"`
}

var OneaccessOneosShowCellularRadioNetwork_Template = `Value Required OPERATOR (.*)
Value SIGNAL_STRENGTH (.*)
Value RSSI (\S+)
Value RSRQ (\S+)
Value RSRP (\S+)
Value SNR (\S+)
Value RADIO_TECHNOLOGY (\S+)
Value REGISTER_STATE (\S+)
Value ATTACH_STATE (\S+)
Value RESETS_LOSS_REGISTRATIONS (\d+)
Value RESETS_FAILED_REGISTRATIONS (\d+)
Value RESETS_MODEM (\d+)
Value RESETS_UNKNOWN (\d+)
Value LAC (.*)
Value CELL_ID (\d+)
Value TAC (\d+)
Value PLMN (\d+)

Start
  ^\s*Current\sselected\soperator\s+:\s+${OPERATOR}
  ^\s*Signal\sstrength\s+:\s+${SIGNAL_STRENGTH}
  ^\s*RSSI\s+:\s+${RSSI}
  ^\s*RSRQ\s+:\s+${RSRQ}
  ^\s*RSRP\s+:\s+${RSRP}
  ^\s*SNR\s+:\s+${SNR}
  ^\s*Current\sradio\saccess\stechnology\s+:\s+${RADIO_TECHNOLOGY}
  ^\s*Circuit-switched\sregister\sstate\s+:\s+${REGISTER_STATE}
  ^\s*Packet-switched\sattach\sstate\s+:\s+${ATTACH_STATE}
  ^\s*Reset\son\sloss\sof\s(?:GPRS\s)?registration\s+:\s+${RESETS_LOSS_REGISTRATIONS}
  ^\s*Reset\son\sfailed\sinitial\sregistration\s+:\s+${RESETS_FAILED_REGISTRATIONS}
  ^\s*Hardware\sreset\sof\smodem\s+:\s+${RESETS_MODEM}
  ^\s*Unknown\sreset\sof\smodem\s+:\s+${RESETS_UNKNOWN}
  ^\s*.*\(LAC\)\s+:\s+${LAC}
  ^\s*Cell\sID\s+:\s+${CELL_ID}
  ^\s*.*\(TAC\)\s+:\s+${TAC}
  ^.*Public Land Mobile Network.*:\s+${PLMN}
  ^\s*Total\s+Ec/Io\s+:
  ^\s*Statistics:
  ^\s*Details:
  ^\s*Toggle\s+w_disable\s+of\s+modem\s+:
  ^\s*$$
  ^. -> Error

`