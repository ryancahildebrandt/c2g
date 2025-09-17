// -*- coding: utf-8 -*-

// Created on Wed Sep  3 07:41:18 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

var (
	inFile  cli.StringFlag = cli.StringFlag{Name: "inFile", Hidden: true}
	outFile cli.StringFlag = cli.StringFlag{
		Name:    "outFile",
		Aliases: []string{"o"},
		Usage:   "jsgf file to write grammar to. If blank, grammar is returned to stdout",
	}
	prob cli.FloatFlag = cli.FloatFlag{
		Name:    "prob",
		Aliases: []string{"p"},
		Usage:   "transitional probability below which tokens will be split",
	}
)

// Checks that the provided path exists on disk and has extension .jsgf/.jjsgf
func ValidateInFile(p string) error {
	_, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("in ValidateInFile(%v):\n%+w", p, err)
	}
	switch filepath.Ext(p) {
	case ".jsgf", ".jjsgf":
		return nil
	default:
		return fmt.Errorf("in ValidateInFile(%v):\n%+w", p, errors.New("file extension is not one of .jsgf, .jjsgf"))
	}
}

// Checks that the directory (if present) in the provided path exists
func ValidateOutFile(p string) error {
	_, err := os.Stat(filepath.Dir(p))
	if err != nil {
		return fmt.Errorf("in ValidateOutFile(%v):\n%+w", p, err)
	}

	return nil
}

// Sets additional cli args before each gsgf command is run
func prepareContext(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	if cmd.Args().Get(0) == "" {
		cli.ShowSubcommandHelpAndExit(cmd, 0)
	}

	cmd.Set("inFile", cmd.Args().Get(0))

	return ctx, nil
}
