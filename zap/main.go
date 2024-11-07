package main

import "go.uber.org/zap"

func main() {
	/*
		zap merupakan logging yang terstrukture yang didevelop oleh Uber dan di design dalam Go Application
		its offer "Blazing fast, structured, leveled logging in Go"
	*/
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	logger.Info("this is zap Logger")
	/*
		tidak seperti logging packages lainnya, zap tidak provide a preconfigured
		global logger ready for use. therefore, we need to create a zap.Logger instance
		to begin writting log, The NewProduction() method returns a Logger con
	*/

}
