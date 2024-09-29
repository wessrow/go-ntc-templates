package oneaccess_oneos 

type ShowVoiceMos struct {
	Curr_nbr_call	string	`json:"CURR_NBR_CALL"`
	Curr_avg_mos	string	`json:"CURR_AVG_MOS"`
	Curr_min_mos	string	`json:"CURR_MIN_MOS"`
	Curr_max_mos	string	`json:"CURR_MAX_MOS"`
	Curr_avg_erl	string	`json:"CURR_AVG_ERL"`
	Curr_avg_acom	string	`json:"CURR_AVG_ACOM"`
	Curr_avg_loss_rate	string	`json:"CURR_AVG_LOSS_RATE"`
	Curr_avg_jitter	string	`json:"CURR_AVG_JITTER"`
	Curr_avg_max_delay	string	`json:"CURR_AVG_MAX_DELAY"`
	Curr_avg_pdd	string	`json:"CURR_AVG_PDD"`
	Prev_nbr_call	string	`json:"PREV_NBR_CALL"`
	Prev_avg_mos	string	`json:"PREV_AVG_MOS"`
	Prev_min_mos	string	`json:"PREV_MIN_MOS"`
	Prev_max_mos	string	`json:"PREV_MAX_MOS"`
	Prev_avg_erl	string	`json:"PREV_AVG_ERL"`
	Prev_avg_acom	string	`json:"PREV_AVG_ACOM"`
	Prev_avg_loss_rate	string	`json:"PREV_AVG_LOSS_RATE"`
	Prev_avg_jitter	string	`json:"PREV_AVG_JITTER"`
	Prev_avg_max_delay	string	`json:"PREV_AVG_MAX_DELAY"`
	Prev_avg_pdd	string	`json:"PREV_AVG_PDD"`
}

var ShowVoiceMos_Template = `Value CURR_NBR_CALL (\d+\.\d+|\d+)
Value CURR_AVG_MOS (\d+\.\d+|\d+)
Value CURR_MIN_MOS (\d+\.\d+|\d+)
Value CURR_MAX_MOS (\d+\.\d+|\d+)
Value CURR_AVG_ERL (\d+\.\d+|\d+)
Value CURR_AVG_ACOM (\d+\.\d+|\d+)
Value CURR_AVG_LOSS_RATE (\d+\.\d+|\d+)
Value CURR_AVG_JITTER (\d+\.\d+|\d+)
Value CURR_AVG_MAX_DELAY (\d+\.\d+|\d+)
Value CURR_AVG_PDD (\d+\.\d+|\d+)
Value PREV_NBR_CALL (\d+\.\d+|\d+)
Value PREV_AVG_MOS (\d+\.\d+|\d+)
Value PREV_MIN_MOS (\d+\.\d+|\d+)
Value PREV_MAX_MOS (\d+\.\d+|\d+)
Value PREV_AVG_ERL (\d+\.\d+|\d+)
Value PREV_AVG_ACOM (\d+\.\d+|\d+)
Value PREV_AVG_LOSS_RATE (\d+\.\d+|\d+)
Value PREV_AVG_JITTER (\d+\.\d+|\d+)
Value PREV_AVG_MAX_DELAY (\d+\.\d+|\d+)
Value PREV_AVG_PDD (\d+\.\d+|\d+)

Start
  ^\s*-+\sCurrent\shour\s-+ -> CurrentHour
  ^\s*$$
  ^\s*-*
  ^. -> Error

CurrentHour
  ^\s*Number\sof\sCall\s+:\s+${CURR_NBR_CALL}
  ^\s*Average\sof\sMOS\s+:\s+${CURR_AVG_MOS}
  ^\s*Minimum\sMOS\s+:\s+${CURR_MIN_MOS}
  ^\s*Maximum\sMOS\s+:\s+${CURR_MAX_MOS}
  ^\s*Average\sof\sERL\s+:\s+${CURR_AVG_ERL}
  ^\s*Average\sof\sACOM\s+:\s+${CURR_AVG_ACOM}
  ^\s*Average\sof\sloss-rate\s+:\s+${CURR_AVG_LOSS_RATE}
  ^\s*Average\sof\sjitter\s+:\s+${CURR_AVG_JITTER}
  ^\s*Average\sof\sMax\sdelay\s+:\s+${CURR_AVG_MAX_DELAY}
  ^\s*Average\sPdd\s+:\s+${CURR_AVG_PDD}
  ^\s*-+\sPrevious\shour\s-+ -> PreviousHour
  ^\s*$$
  ^. -> Error

PreviousHour
  ^\s*Number\sof\sCall\s+:\s+${PREV_NBR_CALL}
  ^\s*Average\sof\sMOS\s+:\s+${PREV_AVG_MOS}
  ^\s*Minimum\sMOS\s+:\s+${PREV_MIN_MOS}
  ^\s*Maximum\sMOS\s+:\s+${PREV_MAX_MOS}
  ^\s*Average\sof\sERL\s+:\s+${PREV_AVG_ERL}
  ^\s*Average\sof\sACOM\s+:\s+${PREV_AVG_ACOM}
  ^\s*Average\sof\sloss-rate\s+:\s+${PREV_AVG_LOSS_RATE}
  ^\s*Average\sof\sjitter\s+:\s+${PREV_AVG_JITTER}
  ^\s*Average\sof\sMax\sdelay\s+:\s+${PREV_AVG_MAX_DELAY}
  ^\s*Average\sPdd\s+:\s+${PREV_AVG_PDD}
  ^\s*-+\s*$$
  ^\s*$$
  ^. -> Error
`