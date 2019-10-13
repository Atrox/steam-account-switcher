package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func logError(logErr error) {
	file, err := os.OpenFile(filepath.Join(applicationDir, "error.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	_, _ = file.WriteString(fmt.Sprintf("[%s][%s-%s-%s] %s\n", time.Now().Format("02.01.2006 15:04:05"), version, commit, date, logErr.Error()))
	log.Fatal(logErr)
}
