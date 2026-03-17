package installer

import (
	"fmt"
	"os"
	"os/exec"
)

func RunInstall() {
	fmt.Println("Installing atlas...")
	cmd := exec.Command("sh", "-c", "curl -sSf https://atlasgo.sh | sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Installed successfully")
}
