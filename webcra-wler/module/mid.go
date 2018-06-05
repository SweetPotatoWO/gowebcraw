package module

import (
	"net"
	"fmt"
	"gowebcraw/webcra-wler/errors"
)

//默认的序列号
var DefaultSNGen = NewSNGenertor(1,0)

//midTemplate D代表组件的ID的魔板
var midTemplate = "%s%d|%s"

//MID 代表组件的ID
type MID string

//GenMID 会根据参数生产对应的MID
func GenMID(metype Type,sn uint64,maddr net.Addr) (MID , error) {
	if !LegalType(metype) {
		errMsg := fmt.Sprintf("illegal module type: %s", mtype)
		return "", errors.NewIllegalParameterError(errMsg)
	}
}





