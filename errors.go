//go:build windows
// +build windows

// this file must be updated
// once the griServer updated its error codes constants

package griagent

// error codes
type Errors int

// err const
const (

	// FLAGS
	ERR_TARGETS_NOT_PROVIDED Errors = 1 + iota
	ERR_DOMAIN_NOT_PROVIDED
	ERR_USERNAME_NOT_PROVIDED
	ERR_F2C_NOT_PROVIDED
	ERR_F2R_NOT_PROVIDED
	ERR_READ_PASSWORD_STDIN

	// TARGETS
	ERR_CAN_T_READ_TARGETS
	ERR_TARGETS_IS_EMPTY

	// F2C
	ERR_CAN_T_READ_F2C
	ERR_F2C_IS_EMPTY

	// SMB
	ERR_COULD_NOT_DIAL_SMB

	A // TCP & SOCKET
	ERR_TCP_LISTEN
	ERR_TCP_ACCEPT
	ERR_UNIX_CLIENT_SOCKET
	ERR_UNIX_CLIENT_SOCKET_MARSHAL
	ERR_UNIX_CLIENT_SOCKET_INIT

	// AGENT ERR
	ERR_TCP_CLIENT_AGENT_ERROR
	ERR_AGENT_START_CMD_ERROR
)
