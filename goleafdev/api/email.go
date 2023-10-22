package api

import (
	"fmt"
	"os/exec"
)

const (
	EMAIL_SCRIPT_PATH string = "api/gmail/email_script.py"
)

func CallEmailScript() (error) {
	cmd := exec.Command("/usr/bin/python3", EMAIL_SCRIPT_PATH)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("Python output:\t%s", string(output))
	return nil
}