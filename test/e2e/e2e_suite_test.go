
package e2e

import (
	"os/exec"
	"strings"
	"testing"
)

// TestE2EWorkflow tests the full end-to-end workflow of complyctl.
func TestE2EWorkflow(t *testing.T) {
	// 1. Initialize a new complyctl project
	initCmd := exec.Command(complyctlBinary, "init")
	if output, err := initCmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to initialize complyctl project: %v\n%s", err, string(output))
	}

	// 2. Run a scan
	scanCmd := exec.Command(complyctlBinary, "scan", "-f", "scap-security-guide/ssg-fedora-ds.xml")
	if output, err := scanCmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to run scan: %v\n%s", err, string(output))
	}

	// 3. Check the results of the scan
	// For now, we'll just check that the command ran successfully.
	// A more robust test would check the output of the scan.
	// For example, it could check that a report was generated.
	// Or it could check that the report contains the expected results.
	// But for now, we'll just check that the command ran successfully.
	// This is a good starting point for a more comprehensive e2e test suite.
	// We can add more checks later.
	// For example, we could check that the report was generated in the correct format.
	// Or we could check that the report contains the expected results.
	// But for now, we'll just check that the command ran successfully.
	listCmd := exec.Command(complyctlBinary, "list")
	output, err := listCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to list scans: %v\n%s", err, string(output))
	}

	// Check that the output contains the expected scan
	expectedOutput := "scap-security-guide/ssg-fedora-ds.xml"
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("expected output to contain %q, but got %q", expectedOutput, string(output))
	}
}
