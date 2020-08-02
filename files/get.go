package files

import (
	"fmt"
	"github.com/LoliE1ON/srakes-bot/config"
	"github.com/pkg/errors"
	"io/ioutil"
)

func Get() (files FilesType, err error) {

	cfg, err := config.Get()
	if err != nil {
		return files, errors.Wrap(err, fmt.Sprintf("Failed parse config from files/Get"))
	}

	accounts, err := ioutil.ReadFile(cfg.AccountsPath)
	if err != nil {
		return files, errors.Wrap(err, fmt.Sprintf("Failed read file from files/Get"))
	}
	files.accounts = string(accounts)

	proxy, err := ioutil.ReadFile(cfg.ProxyPath)
	if err != nil {
		return files, errors.Wrap(err, fmt.Sprintf("Failed read file from files/Get"))
	}
	files.proxy = string(proxy)

	return

}
