package main

import (
	"errors"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/google/shlex"
	"golang.org/x/term"
)

func (app *application) shellMode() error {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer tty.Close()

	w, h, err := term.GetSize(int(tty.Fd()))
	if err != nil {
		return err
	}

	oldState, err := term.MakeRaw(int(tty.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(tty.Fd()), oldState)

	terminal := term.NewTerminal(tty, "> ")

	if err := terminal.SetSize(w, h); err != nil {
		return err
	}

	app.terminal = terminal
	app.writer = tabwriter.NewWriter(
		terminal,
		0,
		0,
		2,
		' ',
		tabwriter.DiscardEmptyColumns,
	)
	for {
		line, err := terminal.ReadLine()
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				return nil
			default:
				return err
			}
		}
		if trimedLine := strings.TrimSpace(line); trimedLine != "" {
			if trimedLine == "exit" {
				terminal.Write([]byte("bye..\n"))
				break
			}
			app.args, err = shlex.Split(trimedLine)
			if err != nil {
				terminal.Write(append([]byte(err.Error()), '\n'))
			}
			msg, err := app.handleArgs()
			if err != nil {
				terminal.Write(append([]byte(err.Error()), '\n'))
			}
			if len(msg) != 0 {
				terminal.Write(append([]byte(msg), '\n'))
			}
		}
	}
	return nil
}
