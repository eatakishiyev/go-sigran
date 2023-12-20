package primitives

import (
	"go-sigtran/tcap"
	"go-sigtran/tcap/operations"
	"go-sigtran/tcap/parameters"
)

type TCBegin struct {
	ApplicationContext *parameters.ApplicationContext
	TCAPDialog         *tcap.TcapDialog
	UserInformation    []byte
	Components         []operations.Operations
}
