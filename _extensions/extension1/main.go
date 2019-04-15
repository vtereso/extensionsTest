package main

import (
	"github.com/vtereso/extensionsTest/_extensions/extension1/hello"
	"go.uber.org/zap"
)

func main() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Info(hello.Hello())
}