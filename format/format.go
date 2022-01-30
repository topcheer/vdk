package format

import (
	"github.com/topcheer/vdk/av/avutil"
	"github.com/topcheer/vdk/format/aac"
	"github.com/topcheer/vdk/format/flv"
	"github.com/topcheer/vdk/format/mp4"
	"github.com/topcheer/vdk/format/rtmp"
	"github.com/topcheer/vdk/format/rtsp"
	"github.com/topcheer/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
