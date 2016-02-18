package fw

import (
	"github.com/Sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	//	log.Formatter = &logrus.TextFormatter{FullTimestamp:true, DisableColors:true}	// for file
	Log.Formatter = new(logrus.TextFormatter) // default
	Log.Level = logrus.DebugLevel

	// set file output
	//	file, err := os.Create("test.log")
	//	if err != nil {
	//		log.Fatalln("fail to create test.log file!")
	//	}
	//	log.Out = file
}

