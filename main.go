// -*- coding: utf-8 -*-

// Created on Fri Jun 27 03:42:54 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := &cli.Command{
		Name:                  "c2g",
		Usage:                 "Condense natural language expressions to a context free grammar",
		UsageText:             "c2g [COMMAND] [OPTIONS] example.txt",
		EnableShellCompletion: true,
		Suggest:               true,
		Before:                prepareContext,
		Flags: []cli.Flag{
			&inFile,
			&outFile,
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			start := time.Now()
			infile := cmd.String("inFile")
			file, err := os.Open(infile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			t := NewUnigramTokenizer()
			scanner := bufio.NewScanner(file)
			texts := ReadTexts(scanner)
			corpus := NewCorpus(texts)
			corpus.transitions = NewTransitions(corpus, t)
			corpus.ngrams = ToNgrams(corpus.texts, t, corpus.transitions)
			corpus.texts = SplitTriplets(corpus.texts, corpus.ngrams)
			rules := ToRules(corpus.texts)

			fmt.Println(len(rules))

			var ssd SSDMerger
			var sds SDSMerger
			var dss DSSMerger
			var sss SSSMerger

			rules = ssd.apply(rules)
			rules = sds.apply(rules)
			rules = dss.apply(rules)
			rules = sss.apply(rules)

			for _, rule := range rules {
				fmt.Println(rule.print("", true))
			}

			fmt.Println(len(rules))
			fmt.Println(time.Since(start))
			return nil
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
