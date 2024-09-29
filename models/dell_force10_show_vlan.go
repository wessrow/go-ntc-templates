package models

type DellForce10ShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Status	string	`json:"STATUS"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Ports_u_gi	string	`json:"PORTS_U_GI"`
	Ports_t_gi	string	`json:"PORTS_T_GI"`
	Ports_u_te	string	`json:"PORTS_U_TE"`
	Ports_t_te	string	`json:"PORTS_T_TE"`
	Ports_u_fo	string	`json:"PORTS_U_FO"`
	Ports_t_fo	string	`json:"PORTS_T_FO"`
	Ports_u_ma	string	`json:"PORTS_U_MA"`
	Ports_t_ma	string	`json:"PORTS_T_MA"`
}

var DellForce10ShowVlan_Template = `Value Required VLAN_ID (\d+)
# Name is really the description and it currently only supports a single word (no spaces)
Value Required STATUS (\w+)
Value VLAN_NAME ([^TU][^\s]\S+(\s\S+)*)
Value PORTS_U_GI (\S+)
Value PORTS_T_GI (\S+)
Value PORTS_U_TE (\S+)
Value PORTS_T_TE (\S+)
Value PORTS_U_FO (\S+)
Value PORTS_T_FO (\S+)
Value PORTS_U_MA (\S+)
Value PORTS_T_MA (\S+)

Start
  ^(\*)?\s+\d+.* -> Continue.Record
  ^(\*)?\s+${VLAN_ID}\s+${STATUS}(\s+${VLAN_NAME})?(\s+(U Gi ${PORTS_U_GI}|T Gi ${PORTS_T_GI}|U Ma ${PORTS_U_MA}|T Ma ${PORTS_T_MA}|U Te ${PORTS_U_TE}|T Te ${PORTS_T_TE}|U Fo ${PORTS_U_FO}|T Fo ${PORTS_T_FO}))?
  ^\s+U Gi ${PORTS_U_GI}
  ^\s+T Gi ${PORTS_T_GI}
  ^\s+U Ma ${PORTS_U_MA}
  ^\s+T Ma ${PORTS_T_MA}
  ^\s+U Te ${PORTS_U_TE}
  ^\s+T Te ${PORTS_T_TE}
  ^\s+U Fo ${PORTS_U_FO}
  ^\s+T Fo ${PORTS_T_FO}

`