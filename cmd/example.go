package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grootcz/ezlog"
)

func main() {
	ezlog.SetLevels(63)
	ezlog.SetModules([]string{"main"})

	ezlog.NoticefWithStack("my log=%d", 111)
	ezlog.CriticalfWithStack("my log=%d", 111)

	ezlog.Debugf("my log=%d", 111)
	ezlog.Infof("my log=%d", 111)
	ezlog.Warnf("my log=%d", 111)
	ezlog.Errorf("my log=%d", 111)
	ezlog.Errorf("my log=%d")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			{
				fmt.Printf("security exit by %s signal.\n", s)
				time.Sleep(time.Millisecond * time.Duration(1500))
				os.Exit(0)
			}
		default:
			{
				fmt.Printf("unknown exit by %s signal.\n", s)
				time.Sleep(time.Millisecond * time.Duration(1500))
				os.Exit(0)
			}
		}
	}
}
