package browser

import (
	"fmt"
	"github.com/Arikael/shop-searcher/search_config"
	"os/exec"
	"runtime"
)

func OpenBrowser(config *search_config.Config) (err error) {
	for _, url := range config.Urls {
		var cmd *exec.Cmd
		if cmd, err = getOsCommand(url); err != nil {
			return err
		}

		if err = cmd.Start(); err != nil {
			return err
		}

		if err = cmd.Wait(); err != nil {
			return err
		}
	}

	return nil
}

func getOsCommand(url string) (cmd *exec.Cmd, err error) {
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = nil
		err = fmt.Errorf("unknown OS")
	}

	return cmd, err
}
