//go:build windows
// +build windows

package griagent

// this function will be called after service is started
func (s *GriAgent) DoTheJob() {
	_ = s.ExecStarting()

}
