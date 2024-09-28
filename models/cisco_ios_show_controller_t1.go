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