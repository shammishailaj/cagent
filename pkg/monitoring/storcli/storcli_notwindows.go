//+build !windows

package storcli

import "fmt"

func (s *StorCLI) getCommandLine() []string {
	return []string{"sudo", s.binaryPath, fmt.Sprintf("/c%d", s.controller), "show", "all", "J"}
}
