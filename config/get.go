package config

import "fmt"

var state stateType

func Get() (config ConfigType, err error) {

	config = state.config

	if !state.isLoaded {
		config, err = parse()
		if err != nil {
			fmt.Print("Get config error", err)
			return
		} else {
			state.config = config
			state.isLoaded = true
		}
	}

	return
}
