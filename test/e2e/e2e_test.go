package e2e

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"strings"
	"testing"
)

func TestComplyctlHelp(t *testing.T) {
	// Run the "complyctl --help" command
	cmd := exec.Command("complyctl", "--help")
	output, err := cmd.CombinedOutput()

	// Ensure there is no error when running the command
	if err != nil {
		t.Fatalf("Error running complyctl --help: %v\nOutput: %s", err, string(output))
	}

	// Convert the output to a string and check if expected text is present
	outputStr := string(output)

	// Assert that "Usage" or the expected help message is part of the output
	assert.True(t, strings.Contains(outputStr, "Usage"), "Help output should contain 'Usage'")
	assert.True(t, strings.Contains(outputStr, "Aliases"), "Help output should contain 'Aliases'")
	assert.True(t, strings.Contains(outputStr, "Available Commands"), "Help output should contain 'Available Commands'")
	assert.True(t, strings.Contains(outputStr, "Flags"), "Help output should contain 'Flags'")
	assert.True(t, strings.Contains(outputStr, "complyctl [command]"), "Help output should contain 'complyctl [command]'")
}

func TestComplyctlList(t *testing.T) {
	// Run the "complyctl list" command
	cmd := exec.Command("complyctl", "list")
	output, err := cmd.CombinedOutput()

	// Ensure there is no error when running the command
	if err != nil {
		t.Fatalf("Error running complyctl list: %v\nOutput: %s", err, string(output))
	}

	// Convert the output to a string and check if expected content is returned
	outputStr := string(output)

	// Check if the output contains expected text; you can adjust this to match the actual output of "complyctl list"
	assert.True(t, strings.Contains(outputStr, "list"), "'list' should appear in the output of complyctl list")
	assert.True(t, len(outputStr) > 0, "Output from 'complyctl list' should not be empty")
}
