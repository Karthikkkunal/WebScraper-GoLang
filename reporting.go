package main

import (
	"encoding/json"
	"os"
)

func generateReportJSON(vulnerabilities []string, filename string) error {
	data, err := json.MarshalIndent(vulnerabilities, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
