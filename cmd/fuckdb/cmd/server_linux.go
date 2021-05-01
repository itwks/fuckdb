package cmd

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

func openBrowser(url string) {
	openCmd := exec.Command("x-www-browser", url)
	if err := openCmd.Run(); err != nil {
		logrus.Errorln("open browser error, please open browser manually:", url)
	}
}
