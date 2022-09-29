package main

import "github.com/sirupsen/logrus"

func main() {
	// Using default logger
	// you can set logger level
	logrus.SetLevel(logrus.DebugLevel)
	// then print
	logrus.Debugln("Debug level")
	logrus.Infoln("Info level")
	logrus.Warnln("Warning level")
	logrus.Errorln("Error level")
	// logrus.Fatalln("Fatal level") // Fatalln will exit the program immediately
	// there is also Debugf, Infof, Warnf, Errorf, and Fatalf
	// You can also instantiate new *logrus.Logger
	// this is useful if you need to pass a logger to a function or method
	// that accept logger interface
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	logrus.Debugln("Debug level")
	logrus.Infoln("Info level")
	logrus.Warnln("Warning level")
	logrus.Errorln("Error level")
	// with Error level, it would only print "Error level"
}
