package ruckus_fastiron

type ShowVlan struct {
	Lag       string `json:"LAG"`
	Vlan_id   string `json:"VLAN_ID"`
	Vlan_name string `json:"VLAN_NAME"`
	Status    string `json:"STATUS"`
	Stack_id  string `json:"STACK_ID"`
	Slot      string `json:"SLOT"`
	Port      string `json:"PORT"`
}

var ShowVlan_Template string = `Value Filldown VLAN_ID (\d+)
Value Filldown VLAN_NAME (\S+)
Value Required STATUS ((Untagged|Tagged))
Value STACK_ID (\d+)
Value SLOT (\d+)
Value Required PORT (\d+)
Value LAG (\S+)

Start
  ^Total\sPORT-VLAN\sentries:\s\d+
  ^Maximum\sPORT-VLAN\sentries:\s\d+
  ^Legend:\s\[Stk=Stack-Id,\sS=Slot\]
  ^PORT-VLAN\s${VLAN_ID},\sName\s${VLAN_NAME},\sPriority\s+level\d+,\s(Off|On)\s*
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+${PORT} -> Continue.Record
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){1}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){2}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){3}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){4}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){5}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){6}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){7}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){8}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){9}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){10}${PORT} -> Continue.Record
  ^\s+${STATUS}\s+Ports:\s+\(U${STACK_ID}\/M${SLOT}\)\s+(?:\d+\s+){11}${PORT} -> Continue.Record
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+${PORT} -> Continue.Record
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){1}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){2}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){3}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){4}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){5}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){6}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){7}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){8}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){9}${PORT} -> Continue.Record 
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){10}${PORT} -> Continue.Record
  ^\s+${STATUS}\s+Ports:\s+\(${LAG}\)\s+(?:\d+\s+){11}${PORT} -> Continue.Record  
  ^\s+(Untagged|Tagged)\s+Ports:\s+\(\S+\)\s+(?:\d+\s+)*\d+\s*$$
  ^\s+(Untagged|Tagged)\s+Ports:\s+None
  ^\s+Mac-Vlan\s+Ports:\s+
  ^\s+Monitoring:\s+
  ^\s*$$
  ^.+ -> Error

Done
  ^.*`
