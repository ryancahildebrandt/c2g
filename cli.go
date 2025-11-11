// -*- coding: utf-8 -*-

// Created on Wed Sep  3 07:41:18 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jdkato/prose/tag"
	"github.com/urfave/cli/v3"
)

var (
	inFile cli.StringFlag = cli.StringFlag{
		Name:   "inFile",
		Hidden: true,
		Validator: func(s string) error {
			_, err := os.Open(s)
			if err != nil {
				return fmt.Errorf("in ValidateInFile(%v):\n%+w", s, err)
			}
			switch filepath.Ext(s) {
			case ".txt", ".csv":
				return nil
			default:
				return fmt.Errorf("in ValidateInFile(%v):\n%+w", s, fmt.Errorf("file extension is not one of .txt, .csv"))
			}
		},
	}
	outFile cli.StringFlag = cli.StringFlag{
		Name: "outFile",
		Validator: func(s string) error {
			_, err := os.Stat(filepath.Dir(s))
			if err != nil {
				return fmt.Errorf("in ValidateOutFile(%v):\n%+w", s, err)
			}
			_, err = os.Open(s)
			if filepath.Ext(s) != ".jsgf" {
				return fmt.Errorf("in ValidateOutFile(%v):\n%+w", s, fmt.Errorf("file extension is not .jsgf"))
			}

			return nil
		},
		Usage: "jsgf file to write grammar to. If blank, grammar is returned to stdout",
	}
	printMain cli.BoolFlag = cli.BoolFlag{
		Name:  "main",
		Value: false,
		Usage: "format output grammar with single public rule",
	}
	filterQuantile cli.FloatFlag = cli.FloatFlag{
		Name:  "filter",
		Value: 0.0,
		Validator: func(f float64) error {
			if f < 0.0 || f > 1.0 {
				return fmt.Errorf("in ValidateFilterQuantile(%v):\n%+w", f, fmt.Errorf("quantile must be between 0 and 1"))
			}
			return nil
		},
		Usage: "quantile below which texts will be filtered out from the corpus, based on constituency tags",
	}
	preTokenized cli.BoolFlag = cli.BoolFlag{
		Name:  "preTokenized",
		Value: false,
		Usage: "assume corpus has been pre tokenized using delimter '<SEP>'",
	}

	chunk cli.StringFlag = cli.StringFlag{
		Name: "chunk",
		Validator: func(s string) error {
			switch s {
			case "token", "posTag", "conTag":
				return nil
			default:
				return fmt.Errorf("in ValidateChunk(%v):\n%+w", s, fmt.Errorf("chunk must be one of ['token', 'posTag', 'conTag']"))
			}
		},
		Usage: "strategy to use during expression chunking. one of ['token', 'posTag', 'conTag']",
	}
	prob cli.FloatFlag = cli.FloatFlag{
		Name:  "prob",
		Value: 0.1,
		Validator: func(f float64) error {
			if f < 0.0 || f > 1.0 {
				return fmt.Errorf("in ValidateProbability(%v):\n%+w", f, fmt.Errorf("probability must be between 0 and 1"))
			}
			return nil
		},
		Usage: "transitional probability below which consecutive tokens will be split into chunks",
	}

	merge cli.StringFlag = cli.StringFlag{
		Name: "merge",
		Validator: func(s string) error {
			switch s {
			case "literal", "charDistance", "tokenDistance", "tfidf", "posTag", "conTag":
				return nil
			default:
				return fmt.Errorf("in ValidateChunk(%v):\n%+w", s, fmt.Errorf("chunk must be one of ['literal', 'charDistance', 'tokenDistance', 'tfidf', 'posTag', 'conTag']"))
			}
		},
		Usage: "strategy to use during rule merging. one of ['literal', 'charDistance', 'tokenDistance', 'tfidf', 'posTag', 'conTag']",
	}
	similarity cli.FloatFlag = cli.FloatFlag{
		Name:  "sim",
		Value: 0.8,
		Validator: func(f float64) error {
			if f < 0.0 || f > 1.0 {
				return fmt.Errorf("in ValidateSimilarity(%v):\n%+w", f, fmt.Errorf("similarity must be between 0 and 1"))
			}
			return nil
		},
		Usage: "similarity threshold above which expression groups will be considered eqivalent",
	}

	factor cli.IntFlag = cli.IntFlag{
		Name:  "factor",
		Value: 1,
		Validator: func(i int) error {
			if i < 0 {
				return fmt.Errorf("in ValidateFactor(%v):\n%+w", i, fmt.Errorf("factor must be a positive number"))
			}
			return nil
		},
		Usage: "number of occurrences above which an expression group will be factored out to its own rule",
	}
	conFactor cli.BoolFlag = cli.BoolFlag{
		Name:  "conFactor",
		Value: false,
		Usage: "factor rules based on constituency tags. if unset, rules will be factored by expression group",
	}

	synFile cli.StringFlag = cli.StringFlag{
		Name: "synFile",
		Validator: func(s string) error {
			_, err := os.Open(s)
			if err != nil {
				return fmt.Errorf("in ValidateSynFile(%v):\n%+w", s, err)
			}
			switch filepath.Ext(s) {
			case ".json":
				return nil
			default:
				return fmt.Errorf("in ValidateSynFile(%v):\n%+w", s, fmt.Errorf("file extension is not .json"))
			}
		},
		Usage: "user provided json file containing synonyms. Overrides the value of expand flag if provided",
	}

	logging cli.BoolFlag = cli.BoolFlag{
		Name:  "log",
		Value: false,
		Usage: "log merge and factoring decisions",
	}
	logFile cli.StringFlag = cli.StringFlag{
		Name: "logFile",
		Validator: func(s string) error {
			_, err := os.Stat(filepath.Dir(s))
			if err != nil {
				return fmt.Errorf("in ValidateLogFile(%v):\n%+w", s, err)
			}
			_, err = os.Open(s)
			if err != nil {
				return fmt.Errorf("in ValidateLogFile(%v):\n%+w", s, err)
			}

			return nil
		},
		Usage: "log file. If -l is set and logFile is blank, logging is returned to stdout",
	}

	nilLogger    *log.Logger = log.New(io.Discard, "", log.LstdFlags)
	stdoutLogger *log.Logger = log.New(os.Stdout, "INFO:", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
)

func NewFileLogger(p string) (*log.Logger, error) {
	f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return &log.Logger{}, err
	}
	return log.New(f, "INFO:", log.LstdFlags|log.Lmicroseconds|log.Llongfile), nil
}

// Sets additional cli args before each gsgf command is run
func prepareContext(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	if cmd.Args().Get(0) == "" {
		cli.ShowSubcommandHelpAndExit(cmd, 0)
	}

	cmd.Set("inFile", cmd.Args().Get(0))

	return ctx, nil
}

// Helper func to read from corpus file, filter, and normalize
func readInfile(cmd *cli.Command) ([]Text, error) {
	var (
		err       error
		file      *os.File
		scanner   *bufio.Scanner
		texts     []Text
		tokenizer Tokenizer = setTokenizer(cmd)
		model     *tag.PerceptronTagger
		tagger    SyntacticTagger
	)

	file, err = os.Open(cmd.String("inFile"))
	if err != nil {
		return texts, err
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	texts = ReadTexts(scanner)
	if cmd.Float64("filterQuantile") != 0.0 {
		model = tag.NewPerceptronTagger()
		tagger = NewSyntacticTagger(model, tokenizer)
		texts = FilterTexts(texts, tagger, cmd.Float64("filterQuantile"))
	}

	for i := range texts {
		texts[i].text = tokenizer.normalize(texts[i].text)
	}

	return texts, err
}

// Helper function to apply chunking strategy to texts and convert to rules
func applyChunking(texts []Text, cmd *cli.Command) []Rule {
	var (
		chunks      []string
		rules       []Rule
		tokenizer   Tokenizer               = setTokenizer(cmd)
		chunkfunc   TransitionSplitFunction = setChunk(cmd)
		transitions                         = CollectTransitions(texts, chunkfunc)
	)

	for i := range texts {
		tokens := tokenizer.tokenize(texts[i].text)
		texts[i].chunk = TransitionChunk(tokens, tokens, transitions, cmd.Float("prob"))
	}
	chunks = CollectChunks(texts)
	for i := range texts {
		texts[i] = ToTriplet(texts[i], chunks)
	}
	for i := range texts {
		rules = append(rules, ToRule(texts[i]))
	}

	return rules
}

// Sets logging behavior based on cli flags
func setLogger(cmd *cli.Command) (*log.Logger, error) {
	switch {
	case cmd.String("logFile") != "":
		logger, err := NewFileLogger(cmd.String("logFile"))
		if err != nil {
			return nilLogger, err
		}
		return logger, nil
	case cmd.Bool("log"):
		return stdoutLogger, nil
	default:
		return nilLogger, nil
	}
}

// Sets tokenization behavior based on cli flags
func setTokenizer(cmd *cli.Command) Tokenizer {
	if cmd.Bool("preTokenized") {
		return NewSepTokenizer()
	}
	return NewWordTokenizer()
}

// Sets text chunking behavior based on cli flags
func setChunk(cmd *cli.Command) TransitionSplitFunction {
	tokenizer := setTokenizer(cmd)
	switch cmd.String("chunk") {
	case "posTag":
		model := tag.NewPerceptronTagger()
		tagger := NewSyntacticTagger(model, tokenizer)
		return POSSplit(tagger)
	case "conTag":
		model := tag.NewPerceptronTagger()
		tagger := NewSyntacticTagger(model, tokenizer)
		return ConstituencySplit(tagger)
	default:
		return TokenSplit(tokenizer)
	}
}

// Sets rule merging behavior based on cli flags
func setMerge(cmd *cli.Command) (EqualityFunction, error) {
	logger, err := setLogger(cmd)
	if err != nil {
		return func(e1, e2 []string) bool { return false }, err
	}
	switch cmd.String("merge") {
	case "charDistance":
		return CharacterLevenshteinThreshold(cmd.Float64("sim"), logger), nil
	case "tokenDistance":
		return TokenLevenshteinThreshold(cmd.Float64("sim"), logger), nil
	case "tfidf":
		tokenizer := setTokenizer(cmd)
		texts, err := readInfile(cmd)
		if err != nil {
			return func(e1, e2 []string) bool { return false }, err
		}
		v := CollectVocab(texts, tokenizer)
		idf := CollectIDF(texts, tokenizer)
		return TFIDFCosineThreshold(cmd.Float64("sim"), v, tokenizer, idf, logger), nil
	case "posTag":
		tokenizer := setTokenizer(cmd)
		model := tag.NewPerceptronTagger()
		tagger := NewSyntacticTagger(model, tokenizer)
		return POSTagEqual(tagger, logger), nil
	case "conTag":
		tokenizer := setTokenizer(cmd)
		model := tag.NewPerceptronTagger()
		tagger := NewSyntacticTagger(model, tokenizer)
		return ConstituencyTagEqual(tagger, logger), nil
	default:
		return LiteralEqual(logger), nil
	}
}

// Sets rule factoring behavior based on cli flags
func setFactor(cmd *cli.Command) (FactorFunction, error) {
	logger, err := setLogger(cmd)
	if err != nil {
		return func(r []Rule) []Rule { return r }, err
	}
	if cmd.Bool("conFactor") {
		tokenizer := setTokenizer(cmd)
		model := tag.NewPerceptronTagger()
		tagger := NewSyntacticTagger(model, tokenizer)
		return ConstituencyFactor(tagger, cmd.Int("factor"), logger), nil
	}
	return ExpressionFactor(cmd.Int("factor"), logger), nil
}

// Sets synonym expansion behavior based on cli flags
func setSynonyms(cmd *cli.Command) (FactorFunction, error) {
	logger, err := setLogger(cmd)
	if err != nil {
		return func(r []Rule) []Rule { return r }, err
	}
	syn, err := ReadSynonyms(cmd.String("synFile"))
	if err != nil {
		return func(r []Rule) []Rule { return r }, err
	}
	tokenizer := setTokenizer(cmd)

	return SynonymFactor(syn, tokenizer, logger), nil
}
