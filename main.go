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
			&prob,
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			start := time.Now()
			infile := cmd.String("inFile")
			prob := cmd.Float("prob")
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
			corpus.ngrams = ToNgrams(corpus.texts, t, corpus.transitions, prob)
			corpus.texts = SplitTriplets(corpus.texts, corpus.ngrams)
			rules := ToRules(corpus.texts)

			var ssd SSDMerger
			rules = ssd.apply(rules)

			var sds SDSMerger
			rules = sds.apply(rules)

			var dss DSSMerger
			rules = dss.apply(rules)

			var sss SSSMerger
			rules = sss.apply(rules)

			rules = ApplyFactor(rules)
			rules = SetIDs(rules)
			g := Grammar{Rules: rules}

			err = os.WriteFile("./outputs/test.jsgf", []byte(g.print()), 0644)
			if err != nil {
				return err
			}
			fmt.Println(time.Since(start))
			return nil
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
