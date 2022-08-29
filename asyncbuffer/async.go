package asyncbuffer

import (
	"context"
	log "github.com/weecloudy/logger"
)

func SendOneCh(bufferName, chName string, data IBuffer) (err error) {
	err, sd := GetInstance(bufferName)
	if err != nil {
		log.ErrorfCtx(context.TODO(), "buffer.GetInstance: %s fail~!", bufferName)
		return
	}

	err = sd.AddBufferItem(&InputBufferItem{
		ChName: chName,
		Data:   data,
	})
	if err != nil {
		log.ErrorfCtx(context.TODO(), "sd.AddBufferItem: %s fail~!", bufferName)
		return
	}

	log.Debugf("GetInstance: %s instance.ISendObj: %v change", bufferName, sd.ISendObj)

	return
}
