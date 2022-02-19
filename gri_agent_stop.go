//go:build windows
// +build windows

package griagent

// before stopping the server
// this will be called
func (s *GriAgent) Stop() error {
	close(s.exit)

	s.Conn.Close()

	s.wg.Wait()
	return nil
}
