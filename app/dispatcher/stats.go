// +build !confonly

package dispatcher

import (
	"github.com/v2fly/v2ray-core/v4/common"
	"github.com/v2fly/v2ray-core/v4/common/buf"
	"github.com/v2fly/v2ray-core/v4/features/stats"
)

type SizeStatWriter struct {
	Counter stats.Counter
	Writer  buf.Writer
	Record  *int64
}

func (w *SizeStatWriter) WriteMultiBuffer(mb buf.MultiBuffer) error {
	bufLen := int64(mb.Len())
	if w.Record != nil {
		*w.Record += bufLen
	}
	w.Counter.Add(bufLen)
	return w.Writer.WriteMultiBuffer(mb)
}

func (w *SizeStatWriter) Close() error {
	return common.Close(w.Writer)
}

func (w *SizeStatWriter) Interrupt() {
	common.Interrupt(w.Writer)
}
