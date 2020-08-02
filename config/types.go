package config

type ConfigType struct {
	Timeout           int
	AccountsPath      string
	AccountsSeparator string
	ProxyPath         string
}

type stateType struct {
	isLoaded bool
	config   ConfigType
}
