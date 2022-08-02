package logs

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func Init() {
	var formatter = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	l := logrus.WithFields(logrus.Fields{})
	l.Logger.SetReportCaller(true)
	l.Logger.SetFormatter(formatter)
	l.Logger.SetLevel(logrus.InfoLevel)
}
