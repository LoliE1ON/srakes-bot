package files

import (
	"fmt"
	"github.com/LoliE1ON/srakes-bot/config"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func Create() (err error) {

	cfg, err := config.Get()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed parse config"))
	}

	fmt.Println("file", cfg.AccountsPath)

	if _, err := os.Stat(cfg.AccountsPath); os.IsNotExist(err) {
		if err != nil {
			if err := ioutil.WriteFile(cfg.AccountsPath, nil, 0644); err != nil {
				return errors.Wrap(err, fmt.Sprintf("Create file error (%s)", cfg.AccountsPath))
			}
		}
	}

	if _, err := os.Stat(cfg.ProxyPath); os.IsNotExist(err) {
		if err != nil {
			if err := ioutil.WriteFile(cfg.ProxyPath, nil, 0644); err != nil {
				return errors.Wrap(err, fmt.Sprintf("Create file error (%s)", cfg.ProxyPath))
			}
		}
	}

	return

}
