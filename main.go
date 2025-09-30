// -*- coding: utf-8 -*-

// Created on Fri Jun 27 03:42:54 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
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
			&factor,
			&printMain,
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			start := time.Now()
			rules, err := buildRules(cmd)
			if err != nil {
				return err
			}

			rules = MergePR(rules)
			rules = MergePS(rules)
			rules = MergeRS(rules)
			rules = MergePRS(rules)
			rules = SetIDs(rules)
			rules = Factor(rules, cmd.Int("factor"))
			g := Grammar{Rules: rules}
			g.write(cmd)

			fmt.Println(time.Since(start))
			return nil
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
