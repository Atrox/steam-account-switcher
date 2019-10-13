package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

const registryName = "steam-account-switcher"

func isAutoStartActive() (bool, error) {
	reg, err := getAutoStartRegistry()
	if err != nil {
		return false, fmt.Errorf("could not get registry: %w", err)
	}

	val, _, err := reg.GetStringValue(registryName)
	if err != nil {
		if err == registry.ErrNotExist {
			return false, nil
		}

		return false, fmt.Errorf("could not get value: %w", err)
	}

	if val != executablePath {
		// path changed to executable, update it
		err = reg.SetStringValue(registryName, executablePath)
		if err != nil {
			return false, fmt.Errorf("could not update autostart entry: %w", err)
		}
	}

	return true, nil
}

func enableAutoStart() error {
	reg, err := getAutoStartRegistry()
	if err != nil {
		return fmt.Errorf("could not get registry: %w", err)
	}

	err = reg.SetStringValue(registryName, executablePath)
	if err != nil {
		return fmt.Errorf("could not set autostart entry: %w", err)
	}

	return nil
}

func disableAutoStart() error {
	reg, err := getAutoStartRegistry()
	if err != nil {
		return fmt.Errorf("could not get registry: %w", err)
	}

	err = reg.DeleteValue(registryName)
	if err != nil {
		return fmt.Errorf("could not remove autostart entry: %w", err)
	}

	return nil
}

func getAutoStartRegistry() (registry.Key, error) {
	return registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.READ|registry.WRITE)
}
