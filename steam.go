package main

import (
	"fmt"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func switchAccount(account *Account) error {
	_ = stopSteam()

	reg, err := getRegistry()
	if err != nil {
		return fmt.Errorf("could not get registry: %w", err)
	}

	if err := reg.SetStringValue("AutoLoginUser", account.Username); err != nil {
		return fmt.Errorf("failed to set value in registry: %w", err)
	}
	if err := reg.SetDWordValue("RememberPassword", 1); err != nil {
		return fmt.Errorf("failed to set value in registry: %w", err)
	}

	return startSteam()
}

func getActiveUsername() (string, error) {
	reg, err := getRegistry()
	if err != nil {
		return "", fmt.Errorf("could not get registry: %w", err)
	}

	username, _, err := reg.GetStringValue("AutoLoginUser")
	if err != nil {
		return "", fmt.Errorf("could not get value in registry: %w", err)
	}

	return username, nil
}

func getRegistry() (registry.Key, error) {
	return registry.OpenKey(registry.CURRENT_USER, `Software\Valve\Steam`, registry.READ|registry.WRITE)
}

func startSteam() error {
	return exec.Command("cmd", "/C", "start", "steam://open/main").Start()
}

func stopSteam() error {
	return exec.Command("taskkill", "/F", "/IM", "steam.exe").Run()
}
