package cisco_asa

type ShowCpuUsageDetailed struct {
	Cpu_5_sec    string `json:"CPU_5_SEC"`
	Cpu_5_sec_cp string `json:"CPU_5_SEC_CP"`
	Cpu_1_min    string `json:"CPU_1_MIN"`
	Cpu_5_min_cp string `json:"CPU_5_MIN_CP"`
	Core         string `json:"CORE"`
	Cpu_5_sec_dp string `json:"CPU_5_SEC_DP"`
	Cpu_1_min_dp string `json:"CPU_1_MIN_DP"`
	Cpu_1_min_cp string `json:"CPU_1_MIN_CP"`
	Cpu_5_min    string `json:"CPU_5_MIN"`
	Cpu_5_min_dp string `json:"CPU_5_MIN_DP"`
}

var ShowCpuUsageDetailed_Template string = `Value CORE (Core\s\d+)
Value CPU_5_SEC (\d+\.\d+)
Value CPU_5_SEC_DP (\d+\.\d+)
Value CPU_5_SEC_CP (\d+\.\d+)
Value CPU_1_MIN (\d+\.\d+)
Value CPU_1_MIN_DP (\d+\.\d+)
Value CPU_1_MIN_CP (\d+\.\d+)
Value CPU_5_MIN (\d+\.\d+)
Value CPU_5_MIN_DP (\d+\.\d+)
Value CPU_5_MIN_CP (\d+\.\d+)


Start
  ^Break\s+down\s+of\s+per-core\s+data\s+path\s+versus\s+control\s+point\s+cpu\s+usage:
  ^Core\s+5\s+sec\s+1\s+min\s+5\s+min
  ^${CORE}\s+${CPU_5_SEC}\s+\(${CPU_5_SEC_DP}\s+\+\s+${CPU_5_SEC_CP}\)\s+${CPU_1_MIN}\s+\(${CPU_1_MIN_DP}\s+\+\s+${CPU_1_MIN_CP}\)\s+${CPU_5_MIN}\s+\(${CPU_5_MIN_DP}\s+\+\s+${CPU_5_MIN_CP}.+ -> Record
  ^\s*
  ^Current\s+control\s+point\s+elapsed\s+versus\s+the\s+maximum\s+control\s+point\s+elapsed\s+for:
  ^CPU\s+utilization\s+of\s+external\s+processes\s+for:
  ^Total\s+CPU\s+utilization\s+for:
  ^5\s+seconds\s+=\s+\d+.\d+%;\s+1\s+minute:\s+\d+.\d+%;\s+5\s+minutes:\s+\d+.\d+%
  ^.*$$ -> Error
`
