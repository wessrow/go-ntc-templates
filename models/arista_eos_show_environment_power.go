package models

type AristaEosShowEnvironmentPower struct {
	Power_supply_location	string	`json:"POWER_SUPPLY_LOCATION"`
	Power_supply_pid	string	`json:"POWER_SUPPLY_PID"`
	Power_supply_watts	string	`json:"POWER_SUPPLY_WATTS"`
	Power_supply_input	string	`json:"POWER_SUPPLY_INPUT"`
	Power_supply_output	string	`json:"POWER_SUPPLY_OUTPUT"`
	Power_supply_watts_using	string	`json:"POWER_SUPPLY_WATTS_USING"`
	Power_supply_status	string	`json:"POWER_SUPPLY_STATUS"`
	Power_supply_uptime	string	`json:"POWER_SUPPLY_UPTIME"`
	System_total_power_watts	string	`json:"SYSTEM_TOTAL_POWER_WATTS"`
	System_total_power_used_watts	string	`json:"SYSTEM_TOTAL_POWER_USED_WATTS"`
}

var AristaEosShowEnvironmentPower_Template = `# Tested against EOS version 4.14.15M and newer
# When run against vEOS returns a single list item with status
# power_supply_status: "no power supplies"
Value POWER_SUPPLY_LOCATION (\d+)
Value POWER_SUPPLY_PID (\S+)
Value POWER_SUPPLY_WATTS (\d+)
Value POWER_SUPPLY_INPUT (\S+)
Value POWER_SUPPLY_OUTPUT (\S+)
Value POWER_SUPPLY_WATTS_USING (\d+[.]?\d+?)
# The expression below allows for up to three words in a status i.e. "Ok",
# "Not Installed" or "Some Other Status"
Value Required POWER_SUPPLY_STATUS (\S+( [a-zA-Z]+)?( [a-zA-Z]+)?)
# The Uptime and total values below are not present in earlier versions of EOS
Value POWER_SUPPLY_UPTIME (.*)
Value Fillup SYSTEM_TOTAL_POWER_WATTS (\d+)
Value Fillup SYSTEM_TOTAL_POWER_USED_WATTS (\d+[.]?\d+?)

Start
  # Virtual EOS devices with no power supplies
  ^%\sThere\sseem\sto\sbe\s${POWER_SUPPLY_STATUS}\sconnected.$$
  # Physical EOS devices with power supplies
  ^Power\s+Input\s+Output\s+Output.*$$
  ^Supply\s+Model\s+Capacity\s+Current\s+Current\s+Power\s+Status\s+Uptime$$
  ^Supply\s+Model\s+Capacity\s+Current\s+Current\s+Power\s+Status$$
  ^-.*$$
  ^${POWER_SUPPLY_LOCATION}\s+${POWER_SUPPLY_PID}\s+${POWER_SUPPLY_WATTS}W\s+${POWER_SUPPLY_INPUT}\s+${POWER_SUPPLY_OUTPUT}\s+${POWER_SUPPLY_WATTS_USING}W\s+${POWER_SUPPLY_STATUS} *${POWER_SUPPLY_UPTIME}\s*$$ -> Record
  ^${POWER_SUPPLY_LOCATION}\s+${POWER_SUPPLY_PID}\s+${POWER_SUPPLY_WATTS}W\s+${POWER_SUPPLY_INPUT}\s+${POWER_SUPPLY_OUTPUT}\s+${POWER_SUPPLY_WATTS_USING}W\s+${POWER_SUPPLY_STATUS}$$ -> Record
  ^Total\s+-+\s+${SYSTEM_TOTAL_POWER_WATTS}W\s+-+\s+-+\s+${SYSTEM_TOTAL_POWER_USED_WATTS}W\s+.*$$
  ^. -> Error

`