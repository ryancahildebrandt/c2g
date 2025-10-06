// -*- coding: utf-8 -*-

// Created on Wed Sep  3 07:41:18 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
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
		Value:   0.1,
		Usage:   "transitional probability below which tokens will be split",
	}
	factor cli.IntFlag = cli.IntFlag{
		Name:    "factor",
		Aliases: []string{"f"},
		Value:   1,
		Usage:   "number of occurrences above which an expression group will be factored out to its own rule",
	}
	printMain cli.BoolFlag = cli.BoolFlag{
		Name:    "main",
		Aliases: []string{"m"},
		Value:   false,
		Usage:   "format output grammar with single public rule",
	}

	// loggers
	nilLogger    *log.Logger = log.New(io.Discard, "", log.LstdFlags)
	stdoutLogger *log.Logger = log.New(os.Stdout, "INFO:", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
)

func NewFileLogger(p string) (*log.Logger, error) {
	f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return &log.Logger{}, err
	}
	defer f.Close()
	return log.New(f, "INFO:", log.LstdFlags|log.Lmicroseconds|log.Llongfile), nil
}

// Checks that the provided path exists on disk and has extension txt/csv
func ValidateInFile(p string) error {
	_, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("in ValidateInFile(%v):\n%+w", p, err)
	}
	switch filepath.Ext(p) {
	case ".txt", ".csv":
		return nil
	default:
		return fmt.Errorf("in ValidateInFile(%v):\n%+w", p, errors.New("file extension is not one of .txt, .csv"))
	}
}

// Checks that the directory (if present) in the provided path exists and the file extension of the provided path is .jsgf
func ValidateOutFile(p string) error {
	_, err := os.Stat(filepath.Dir(p))
	if err != nil {
		return fmt.Errorf("in ValidateOutFile(%v):\n%+w", p, err)
	}
	_, err = os.Open(p)
	if err != nil {
		return fmt.Errorf("in ValidateOutFile(%v):\n%+w", p, err)
	}
	if filepath.Ext(p) != ".jsgf" {
		return fmt.Errorf("in ValidateOutFile(%v):\n%+w", p, errors.New("file extension is not .jsgf"))
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

func buildRules(cmd *cli.Command) ([]Rule, error) {
	var (
		err         error
		file        *os.File
		scanner     *bufio.Scanner
		texts       []Text
		transitions Transitions
		tokens      []string
		chunks      []string
		rules       []Rule
		tokenizer   = NewWordTokenizer()
	)

	file, err = os.Open(cmd.String("inFile"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)
	texts = ReadTexts(scanner)
	for i, t := range texts {
		texts[i].text = tokenizer.normalize(t.text)
	}

	transitions = CollectTransitions(texts, TokenSplit(tokenizer))
	for i, t := range texts {
		tokens = tokenizer.tokenize(t.text)
		texts[i].chunk = TransitionChunk(tokens, tokens, transitions, cmd.Float("prob"))
	}
	chunks = CollectChunks(texts)
	for i, t := range texts {
		texts[i] = ToTriplet(t, chunks)
	}
	for _, t := range texts {
		rules = append(rules, ToRule(t))
	}

	return rules, err
}
