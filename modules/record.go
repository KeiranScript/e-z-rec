package modules

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// RecordScreen handles the screen recording logic
func RecordScreen(mode string) error {
	// Get current timestamp for filename
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("./%s.mp4", timestamp)

	// Determine the appropriate wf-recorder command based on mode
	var cmd *exec.Cmd

	switch mode {
	case "partial":
		fmt.Println("Select the region you want to record.")
		cmd = exec.Command("wf-recorder", "-g", "--file", filename)
	case "window":
		fmt.Println("Select the window you want to record.")
		cmd = exec.Command("wf-recorder", "--output", "window", "-f", filename)
	case "fullscreen":
		cmd = exec.Command("wf-recorder", "-f", filename)
	default:
		return fmt.Errorf("invalid mode: %s. Choose from: partial, window, fullscreen", mode)
	}

	// Start the recording command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting screen recording: %v", err)
	}

	fmt.Printf("Recording started in %s mode. Press Ctrl+C to stop...\n", mode)

	// Setup signal catching to handle interrupt (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-c

	// Stop the recording when the signal is caught
	if err := cmd.Process.Signal(os.Interrupt); err != nil {
		return fmt.Errorf("error stopping recording: %v", err)
	}

	fmt.Printf("\nRecording saved as %s\n", filename)
	return nil
}
