package main

//go:generate fileb0x filebox.json
//go:generate rsrc -ico icon\icon.ico

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/getlantern/systray"
	"go.atrox.dev/steam-account-switcher/icon"
)

var executablePath string
var applicationDir string

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	var err error
	executablePath, err = os.Executable()
	if err != nil {
		executablePath = "./"
	}
	applicationDir = filepath.Dir(executablePath)

	systray.SetIcon(icon.FileIconIco)
	systray.SetTooltip("steam account switcher")

	titleItem := systray.AddMenuItem("steam account switcher", "created by atrox.dev")

	systray.AddSeparator()

	accounts, err := getAccounts()
	if err != nil {
		logError(err)
	}

	activeUsername, err := getActiveUsername()
	if err != nil {
		logError(err)
	}

	cases := make([]reflect.SelectCase, len(accounts))
	for i, account := range accounts {
		var description string
		if account.Description != "" {
			description = fmt.Sprintf(" - %s", account.Description)
		}

		account.menuItem = systray.AddMenuItem(fmt.Sprintf("%s%s", account.Username, description), "")

		if account.Username == activeUsername {
			account.menuItem.Check()
		}

		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(account.menuItem.ClickedCh),
		}
	}
	go func() {
		for {
			chosen, _, ok := reflect.Select(cases)
			if !ok {
				return
			}

			activeAccount := accounts[chosen]
			if err := switchAccount(activeAccount); err != nil {
				logError(err)
			}

			for _, account := range accounts {
				if activeAccount.Username == account.Username && !account.menuItem.Checked() {
					account.menuItem.Check()
				} else if account.menuItem.Checked() {
					account.menuItem.Uncheck()
				}
			}
		}
	}()

	systray.AddSeparator()

	autostartItem := systray.AddMenuItem("automatically start on boot", "")
	autostartActive, err := isAutoStartActive()
	if err != nil {
		logError(err)
	}
	if autostartActive {
		autostartItem.Check()
	}

	quitItem := systray.AddMenuItem("quit", "")

	go func() {
		for {
			select {
			case <-autostartItem.ClickedCh:
				if autostartItem.Checked() {
					autostartItem.Uncheck()

					err := disableAutoStart()
					if err != nil {
						logError(err)
					}
				} else {
					autostartItem.Check()

					err := enableAutoStart()
					if err != nil {
						logError(err)
					}
				}
			case <-titleItem.ClickedCh:
				_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://atrox.dev").Start()
			case <-quitItem.ClickedCh:
				systray.Quit()
			}
		}
	}()
}
