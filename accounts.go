package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/getlantern/systray"
	"github.com/pelletier/go-toml"
)

type Account struct {
	Username    string
	Description string

	menuItem *systray.MenuItem
}

func getAccounts() ([]*Account, error) {
	path := filepath.Join(applicationDir, "accounts.toml")

	file, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load file: %w", err)
		}

		file, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}

		_, err = file.WriteString(`[accounts]
username1 = "description..."
username2 = "another description"
username3 = "one more"`)
		if err != nil {
			return nil, fmt.Errorf("failed writing to file: %w", err)
		}

		_, err := file.Seek(0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed seeking back: %w", err)
		}
	}
	defer file.Close()

	tree, err := toml.LoadReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed parsing file: %w", err)
	}

	accountsTree, ok := tree.Get("accounts").(*toml.Tree)
	if !ok {
		return nil, fmt.Errorf("failed parsing file, missing accounts tree")
	}

	var accounts []*Account
	for key, value := range accountsTree.ToMap() {
		account := &Account{Username: key}

		switch v := value.(type) {
		case string:
			account.Description = v
		case int64:
			account.Description = strconv.FormatInt(v, 10)
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
