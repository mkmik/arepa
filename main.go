package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"

	"9fans.net/go/acme"
	"github.com/bitnami-labs/flagenv"
	"github.com/golang/glog"
)

var (
	triggerCmd = flag.String("t", "", "when this command terminates, trigger an execution")
)

func mainE(triggerCmd string, args []string) error {
	glog.Infof("args %q", args)

	needrun := make(chan bool, 1)

	win, err := acme.New()
	if err != nil {
		return err
	}
	pwd, _ := os.Getwd()
	win.Name(pwd + "/+arepa")
	win.Ctl("clean")
	win.Fprintf("tag", "Get ")

	go events(win, needrun)
	go runner(win, needrun, args)

	// if no trigger command is specified, then only react to manual "Get" commands.
	if triggerCmd == "" {
		// trigger once automatically
		go func() {
			needrun <- true
		}()
		select {}
	}

	for {
		needrun <- true
		runTrigger(triggerCmd)
	}
	return nil
}

func runTrigger(triggerCmd string) error {
	cmd := exec.Command("sh", "-c", triggerCmd)
	cmd.Stdout = os.Stderr // trigger output goes on stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func events(win *acme.Win, needrun chan<- bool) {
	for e := range win.EventChan() {
		switch e.C2 {
		case 'x', 'X': // execute
			if string(e.Text) == "Get" {
				select {
				case needrun <- true:
				default:
				}
				continue
			}
			if string(e.Text) == "Del" {
				win.Ctl("delete")
			}
		}
		win.WriteEvent(e)
	}
	os.Exit(0)
}

func runner(win *acme.Win, needrun <-chan bool, args []string) {
	for range needrun {
		runCmd(win, args)
	}
}

func runCmd(win *acme.Win, args []string) error {
	win.Ctl("dirty")

	glog.Infof("running cmd: %q", args)
	b, cerr := exec.Command(args[0], args[1:]...).CombinedOutput()
	glog.Infof("cmd completed with: %v", cerr)
	glog.V(2).Infof("out: %s", b)

	win.Addr(",")
	win.Write("data", nil)
	win.Ctl("clean")
	win.Fprintf("body", "$ %s\n", strings.Join(args, " "))

	win.Write("body", b)

	if cerr != nil {
		win.Fprintf("body", "error: %v\n", cerr)
	} else {
		win.Fprintf("body", "$\n")
	}
	win.Fprintf("addr", "#0")
	win.Ctl("dot=addr")
	win.Ctl("show")
	win.Ctl("clean")

	return nil
}

func main() {
	flagenv.SetFlagsFromEnv("AREPA", flag.CommandLine)
	flag.Parse()
	if err := mainE(*triggerCmd, flag.Args()); err != nil {
		glog.Fatal(err)
	}
}
