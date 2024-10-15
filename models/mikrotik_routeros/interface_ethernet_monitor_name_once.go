package mikrotik_routeros

type InterfaceEthernetMonitorNameOnce struct {
	Sfp_link_length_9um      string   `json:"SFP_LINK_LENGTH_9UM"`
	Sfp_link_length_50um     string   `json:"SFP_LINK_LENGTH_50UM"`
	Sfp_supply_voltage       string   `json:"SFP_SUPPLY_VOLTAGE"`
	Name                     string   `json:"NAME"`
	Sfp_rx_loss              string   `json:"SFP_RX_LOSS"`
	Sfp_vendor_part_number   string   `json:"SFP_VENDOR_PART_NUMBER"`
	Sfp_vendor_serial        string   `json:"SFP_VENDOR_SERIAL"`
	Tx_flow_control          string   `json:"TX_FLOW_CONTROL"`
	Sfp_tx_fault             string   `json:"SFP_TX_FAULT"`
	Sfp_connector_type       string   `json:"SFP_CONNECTOR_TYPE"`
	Auto_negotiation         string   `json:"AUTO_NEGOTIATION"`
	Full_duplex              string   `json:"FULL_DUPLEX"`
	Sfp_tx_power             string   `json:"SFP_TX_POWER"`
	Sfp_rx_power             string   `json:"SFP_RX_POWER"`
	Eeprom                   []string `json:"EEPROM"`
	Link_partner_advertising string   `json:"LINK_PARTNER_ADVERTISING"`
	Sfp_wavelength           string   `json:"SFP_WAVELENGTH"`
	Sfp_temperature          string   `json:"SFP_TEMPERATURE"`
	Sfp_module_present       string   `json:"SFP_MODULE_PRESENT"`
	Sfp_type                 string   `json:"SFP_TYPE"`
	Sfp_link_length_62um     string   `json:"SFP_LINK_LENGTH_62UM"`
	Sfp_link_length_copper   string   `json:"SFP_LINK_LENGTH_COPPER"`
	Sfp_vendor_revision      string   `json:"SFP_VENDOR_REVISION"`
	Sfp_manufacturing_date   string   `json:"SFP_MANUFACTURING_DATE"`
	Status                   string   `json:"STATUS"`
	Rx_flow_control          string   `json:"RX_FLOW_CONTROL"`
	Default_cable_settings   string   `json:"DEFAULT_CABLE_SETTINGS"`
	Advertising              string   `json:"ADVERTISING"`
	Sfp_vendor_name          string   `json:"SFP_VENDOR_NAME"`
	Sfp_tx_bias_current      string   `json:"SFP_TX_BIAS_CURRENT"`
	Eeprom_checksum          string   `json:"EEPROM_CHECKSUM"`
	Rate                     string   `json:"RATE"`
	Combo_state              string   `json:"COMBO_STATE"`
}

var InterfaceEthernetMonitorNameOnce_Template string = `Value NAME (\S+)
Value STATUS (\S+)
Value AUTO_NEGOTIATION (\S+)
Value DEFAULT_CABLE_SETTINGS (\S+)
Value RATE (\S+)
Value FULL_DUPLEX (\S+)
Value TX_FLOW_CONTROL (\S+)
Value RX_FLOW_CONTROL (\S+)
Value COMBO_STATE (\S+)
Value SFP_MODULE_PRESENT (\S+)
Value SFP_RX_LOSS (\S+)
Value SFP_TX_FAULT (\S+)
Value SFP_TYPE (\S+)
Value SFP_CONNECTOR_TYPE (\S+)
Value SFP_LINK_LENGTH_9UM (\S+m)
Value SFP_LINK_LENGTH_50UM (\S+m)
Value SFP_LINK_LENGTH_62UM (\S+m)
Value SFP_LINK_LENGTH_COPPER (\S+m)
Value SFP_VENDOR_NAME (\S+)
Value SFP_VENDOR_PART_NUMBER (\S+)
Value SFP_VENDOR_REVISION (\S+)
Value SFP_VENDOR_SERIAL (\S+)
Value SFP_MANUFACTURING_DATE (\S+)
Value SFP_WAVELENGTH (\S+nm)
Value SFP_TEMPERATURE (\S+C)
Value SFP_SUPPLY_VOLTAGE (\S+V)
Value SFP_TX_BIAS_CURRENT (\S+mA)
Value SFP_TX_POWER (\S+dBm)
Value SFP_RX_POWER (\S+dBm)
Value EEPROM_CHECKSUM (\S+)
Value List EEPROM ([0-9a-f]+:(?:\s+[0-9a-f]{2})+.*)
Value ADVERTISING (\S+)
Value LINK_PARTNER_ADVERTISING (\S+)

Start
  ^\s*name:\s${NAME}
  ^\s*status:\s${STATUS}
  ^\s*auto-negotiation:\s${AUTO_NEGOTIATION}
  ^\s*default-cable-settings:\s${DEFAULT_CABLE_SETTINGS}
  ^\s*rate:\s${RATE}
  ^\s*full-duplex:\s${FULL_DUPLEX}
  ^\s*tx-flow-control:\s${TX_FLOW_CONTROL}
  ^\s*rx-flow-control:\s${RX_FLOW_CONTROL}
  ^\s*combo-state:\s${COMBO_STATE}
  ^\s*sfp-module-present:\s${SFP_MODULE_PRESENT}
  ^\s*sfp-rx-loss:\s${SFP_RX_LOSS}
  ^\s*sfp-tx-fault:\s${SFP_TX_FAULT}
  ^\s*sfp-type:\s${SFP_TYPE}
  ^\s*sfp-connector-type:\s${SFP_CONNECTOR_TYPE}
  ^\s*sfp-link-length-9um:\s${SFP_LINK_LENGTH_9UM}
  ^\s*sfp-link-length-50um:\s${SFP_LINK_LENGTH_50UM}
  ^\s*sfp-link-length-62um:\s${SFP_LINK_LENGTH_62UM}
  ^\s*sfp-link-length-copper:\s${SFP_LINK_LENGTH_COPPER}
  ^\s*sfp-vendor-name:\s${SFP_VENDOR_NAME}
  ^\s*sfp-vendor-part-number:\s${SFP_VENDOR_PART_NUMBER}
  ^\s*sfp-vendor-revision:\s${SFP_VENDOR_REVISION}
  ^\s*sfp-vendor-serial:\s${SFP_VENDOR_SERIAL}
  ^\s*sfp-manufacturing-date:\s${SFP_MANUFACTURING_DATE}
  ^\s*sfp-wavelength:\s${SFP_WAVELENGTH}
  ^\s*sfp-temperature:\s${SFP_TEMPERATURE}
  ^\s*sfp-supply-voltage:\s${SFP_SUPPLY_VOLTAGE}
  ^\s*sfp-tx-bias-current:\s${SFP_TX_BIAS_CURRENT}
  ^\s*sfp-tx-power:\s${SFP_TX_POWER}
  ^\s*sfp-rx-power:\s${SFP_RX_POWER}
  ^\s*eeprom-checksum:\s${EEPROM_CHECKSUM}
  ^\s*eeprom:\s${EEPROM}
  ^\s*${EEPROM}
  ^\s*advertising:\s*$$
  ^\s*advertising:\s${ADVERTISING}
  ^\s*link-partner-advertising:\s*$$
  ^\s*link-partner-advertising:\s${LINK_PARTNER_ADVERTISING}$$
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^\s*$$
  ^. -> Error
`
