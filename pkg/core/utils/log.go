package utils

import "github.com/fatih/color"

var LogInfo = color.New(color.FgGreen).PrintfFunc()
var LogError = color.New(color.FgRed).PrintfFunc()
var Log = color.New(color.FgHiWhite).PrintfFunc()
