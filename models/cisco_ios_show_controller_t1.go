package models

type CiscoIosShowControllerT1 struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Ais_state	string	`json:"AIS_STATE"`
	Los_state	string	`json:"LOS_STATE"`
	Lof_state	string	`json:"LOF_STATE"`
	Framing	string	`json:"FRAMING"`
	Line_code	string	`json:"LINE_CODE"`
	Clock_source	string	`json:"CLOCK_SOURCE"`
	Curr_interval_elapsed_seconds	string	`json:"CURR_INTERVAL_ELAPSED_SECONDS"`
	Curr_line_code_violations	string	`json:"CURR_LINE_CODE_VIOLATIONS"`
	Curr_path_code_violations	string	`json:"CURR_PATH_CODE_VIOLATIONS"`
	Curr_slip_secs	string	`json:"CURR_SLIP_SECS"`
	Curr_fr_loss_secs	string	`json:"CURR_FR_LOSS_SECS"`
	Curr_line_err_secs	string	`json:"CURR_LINE_ERR_SECS"`
	Curr_degraded_mins	string	`json:"CURR_DEGRADED_MINS"`
	Curr_err_secs	string	`json:"CURR_ERR_SECS"`
	Curr_bursty_err_secs	string	`json:"CURR_BURSTY_ERR_SECS"`
	Curr_sev_err_secs	string	`json:"CURR_SEV_ERR_SECS"`
	Curr_unavail_secs	string	`json:"CURR_UNAVAIL_SECS"`
	Total_line_code_violations	string	`json:"TOTAL_LINE_CODE_VIOLATIONS"`
	Total_path_code_violations	string	`json:"TOTAL_PATH_CODE_VIOLATIONS"`
	Total_slip_secs	string	`json:"TOTAL_SLIP_SECS"`
	Total_fr_loss_secs	string	`json:"TOTAL_FR_LOSS_SECS"`
	Total_line_err_secs	string	`json:"TOTAL_LINE_ERR_SECS"`
	Total_degraded_mins	string	`json:"TOTAL_DEGRADED_MINS"`
	Total_err_secs	string	`json:"TOTAL_ERR_SECS"`
	Total_bursty_err_secs	string	`json:"TOTAL_BURSTY_ERR_SECS"`
	Total_sev_err_secs	string	`json:"TOTAL_SEV_ERR_SECS"`
	Total_unavail_secs	string	`json:"TOTAL_UNAVAIL_SECS"`
}

var CiscoIosShowControllerT1_Template = `Value Required INTERFACE (\d/\d\/\d)
Value LINK_STATUS (down|up)
Value AIS_STATE (\w+)
Value LOS_STATE (\w+)
Value LOF_STATE (\w+)
Value FRAMING (\w+)
Value LINE_CODE (\w+)
Value CLOCK_SOURCE ([a-zA-Z0-9\s]+)
Value CURR_INTERVAL_ELAPSED_SECONDS (\d+)
Value CURR_LINE_CODE_VIOLATIONS (\d+)
Value CURR_PATH_CODE_VIOLATIONS (\d+)
Value CURR_SLIP_SECS (\d+)
Value CURR_FR_LOSS_SECS (\d+)
Value CURR_LINE_ERR_SECS (\d+)
Value CURR_DEGRADED_MINS (\d+)
Value CURR_ERR_SECS (\d+)
Value CURR_BURSTY_ERR_SECS (\d+)
Value CURR_SEV_ERR_SECS (\d+)
Value CURR_UNAVAIL_SECS (\d+)
Value TOTAL_LINE_CODE_VIOLATIONS (\d+)
Value TOTAL_PATH_CODE_VIOLATIONS (\d+)
Value TOTAL_SLIP_SECS (\d+)
Value TOTAL_FR_LOSS_SECS (\d+)
Value TOTAL_LINE_ERR_SECS (\d+)
Value TOTAL_DEGRADED_MINS (\d+)
Value TOTAL_ERR_SECS (\d+)
Value TOTAL_BURSTY_ERR_SECS (\d+)
Value TOTAL_SEV_ERR_SECS (\d+)
Value TOTAL_UNAVAIL_SECS (\d+)

Start 
  ^T1\s${INTERFACE}\sis\s${LINK_STATUS}\.
  ^\s+AIS\sState:${AIS_STATE}\s+LOS\sState:${LOS_STATE}\s+LOF\sState:${LOF_STATE}
  ^\s+Framing\sis\s${FRAMING},\sLine\sCode\sis\s${LINE_CODE},\sClock\sSource\sis\s${CLOCK_SOURCE}. 
  ^\s+Data\sin\scurrent\sinterval\s\(${CURR_INTERVAL_ELAPSED_SECONDS}\sseconds\selapsed\): -> ParseCurrentCounters
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

ParseCurrentCounters
  ^\s+${CURR_LINE_CODE_VIOLATIONS}\sLine\sCode\sViolations,\s${CURR_PATH_CODE_VIOLATIONS}\sPath\sCode\sViolations
  ^\s+${CURR_SLIP_SECS}\sSlip\sSecs,\s${CURR_FR_LOSS_SECS}\sFr\sLoss\sSecs,\s${CURR_LINE_ERR_SECS}\sLine\sErr\sSecs,\s${CURR_DEGRADED_MINS}\sDegraded\sMins
  ^\s+${CURR_ERR_SECS}\sErrored\sSecs,\s${CURR_BURSTY_ERR_SECS}\sBursty\sErr\sSecs,\s${CURR_SEV_ERR_SECS}\sSeverely\sErr\sSecs,\s${CURR_UNAVAIL_SECS}\sUnavail\sSecs
  ^\s+Total Data.+ -> ParseTotalCounters

ParseTotalCounters
  ^\s+${TOTAL_LINE_CODE_VIOLATIONS}\sLine\sCode\sViolations,\s${TOTAL_PATH_CODE_VIOLATIONS}\sPath\sCode\sViolations,
  ^\s+${TOTAL_SLIP_SECS}\sSlip\sSecs,\s${TOTAL_FR_LOSS_SECS}\sFr\sLoss\sSecs,\s${TOTAL_LINE_ERR_SECS}\sLine\sErr\sSecs,\s${TOTAL_DEGRADED_MINS}\sDegraded\sMins,
  ^\s+${TOTAL_ERR_SECS}\sErrored\sSecs,\s${TOTAL_BURSTY_ERR_SECS}\sBursty\sErr\sSecs,\s${TOTAL_SEV_ERR_SECS}\sSeverely\sErr\sSecs,\s${TOTAL_UNAVAIL_SECS}\sUnavail\sSecs -> Record Start

`