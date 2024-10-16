package cisco_xr

type ShowPlatformSummaryLocationAll struct {
	Platform_node      string   `json:"PLATFORM_NODE"`
	Last_reset_process string   `json:"LAST_RESET_PROCESS"`
	Last_reset_time    string   `json:"LAST_RESET_TIME"`
	Rommon_version     string   `json:"ROMMON_VERSION"`
	Faults             string   `json:"FAULTS"`
	Oper_state         string   `json:"OPER_STATE"`
	Ios_version        string   `json:"IOS_VERSION"`
	Main_power         string   `json:"MAIN_POWER"`
	Card_type          string   `json:"CARD_TYPE"`
	Part_id            string   `json:"PART_ID"`
	Hw_version         string   `json:"HW_VERSION"`
	Serial_number      string   `json:"SERIAL_NUMBER"`
	Last_reset         string   `json:"LAST_RESET"`
	Configuration      []string `json:"CONFIGURATION"`
}

var ShowPlatformSummaryLocationAll_Template string = `Value Required PLATFORM_NODE (\S+(\s\S+)*)
Value PART_ID (\S+)
Value CARD_TYPE (\S+(\s\S+)*)
Value HW_VERSION (\w+)
Value SERIAL_NUMBER (\w+)
Value OPER_STATE (\S+(\s\S+)*)
Value LAST_RESET (\S+(\s\S+)*)
Value LAST_RESET_PROCESS (\S+(\s\S+)*)
Value LAST_RESET_TIME (\S+(\s\S+)*)
Value List CONFIGURATION ([\w\.]+(\s[\w\.]+)*)
Value ROMMON_VERSION (\S+(\s\S+)*)
Value IOS_VERSION (\S+(\s\S+)*)
Value MAIN_POWER (\S+(\s\S+)*)
Value FAULTS (\S+(\s\S+)*)


Start
  ^\s*Platform Node : ${PLATFORM_NODE}\s*$$
  ^\s*PID : ${PART_ID}\s*$$
  ^\s*Card Type : ${CARD_TYPE}\s*$$
  ^\s*VID/SN : ${HW_VERSION} / ${SERIAL_NUMBER}\s*$$
  ^\s*Oper State : ${OPER_STATE}\s*$$ -> LastReset

LastReset
  ^\s*Last Reset : ${LAST_RESET}(\s+Process: ${LAST_RESET_PROCESS})?\s*$$
  ^\s*: ${LAST_RESET_TIME}\s*$$
  ^\s*Configuration : ${CONFIGURATION}\s*$$ -> Configuration

Configuration
  ^\s*${CONFIGURATION}\s*$$
  ^\s*Rommon Ver : ${ROMMON_VERSION}\s*$$ -> Finish

Finish
  ^\s*IOS SW Ver : ${IOS_VERSION}\s*$$
  ^\s*Main Power : ${MAIN_POWER}\s*$$
  ^\s*Faults : ${FAULTS}\s*$$ -> Record Start`
