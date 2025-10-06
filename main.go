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
		Usage:                 "Convert natural language expressions to a context free grammar",
		UsageText:             "c2g [COMMAND] [OPTIONS] example.txt",
		EnableShellCompletion: true,
		Suggest:               true,
		Commands: []*cli.Command{
			{
				Name:                  "clone",
				Usage:                 "Create a grammar such that each expression is contained in one rule, with no rule merging or factoring applied. This mode will not produce outputs not found in the source corpus.",
				UsageText:             "c2g clone [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&printMain},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						start = time.Now()
						rules []Rule
						g     Grammar
						err   error
					)

					rules, err = buildRules(cmd)
					if err != nil {
						return err
					}

					rules = SetIDs(rules)
					g = Grammar{Rules: rules}
					g.write(cmd)

					fmt.Println(time.Since(start))
					return nil
				},
			},
			{
				Name:                  "compress",
				Usage:                 "Create a grammar with rule merging and factoring applied to rules sharing 2 or more chunks. This mode will not produce outputs not found in the source corpus.",
				UsageText:             "c2g compress [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&prob,
					&factor,
					&printMain},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						start = time.Now()
						rules []Rule
						g     Grammar
						err   error
					)

					rules, err = buildRules(cmd)
					if err != nil {
						return err
					}

					rules = MergePR(rules)
					rules = MergePS(rules)
					rules = MergeRS(rules)
					rules = MergePRS(rules)
					rules = SetIDs(rules)
					rules = Factor(rules, cmd.Int("factor"))
					g = Grammar{Rules: rules}
					g.write(cmd)

					fmt.Println(time.Since(start))
					return nil
				},
			},
			{
				Name:                  "interpolate",
				Usage:                 "Create a grammar with rule merging and factoring applied to rules sharing 1 or more chunks. This mode will produce outputs not found in the source corpus.",
				UsageText:             "c2g interpolate [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&prob,
					&factor,
					&printMain},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						start = time.Now()
						rules []Rule
						g     Grammar
						err   error
					)

					rules, err = buildRules(cmd)
					if err != nil {
						return err
					}

					rules = MergePR(rules)
					rules = MergePS(rules)
					rules = MergeRS(rules)

					rules = MergeP(rules)
					rules = MergeR(rules)
					rules = MergeS(rules)

					rules = MergePRS(rules)

					rules = SetIDs(rules)
					rules = Factor(rules, cmd.Int("factor"))
					g = Grammar{Rules: rules}
					g.write(cmd)

					fmt.Println(time.Since(start))
					return nil
				},
			},
			{
				Name:                  "extrapolate",
				Usage:                 "Create a grammar with rule merging, factoring, and expansion applied to rules sharing 1 or more chunks. This mode will produce outputs not found in the source corpus.",
				UsageText:             "c2g extrapolate [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&prob,
					&factor,
					&printMain},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						start = time.Now()
						rules []Rule
						g     Grammar
						err   error
					)

					rules, err = buildRules(cmd)
					if err != nil {
						return err
					}

					rules = MergePR(rules)
					rules = MergePS(rules)
					rules = MergeRS(rules)
					rules = MergePRS(rules)
					rules = SetIDs(rules)
					rules = Factor(rules, cmd.Int("factor"))
					g = Grammar{Rules: rules}
					g.write(cmd)

					fmt.Println(time.Since(start))
					return nil
				},
			},
			{
				Name:                  "custom",
				Usage:                 "Create a grammar with a custom set of options.",
				UsageText:             "c2g custom [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&prob,
					&factor,
					&printMain},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						start = time.Now()
						rules []Rule
						g     Grammar
						err   error
					)

					rules, err = buildRules(cmd)
					if err != nil {
						return err
					}

					rules = MergePR(rules)
					rules = MergePS(rules)
					rules = MergeRS(rules)
					rules = MergePRS(rules)
					rules = SetIDs(rules)
					rules = Factor(rules, cmd.Int("factor"))
					g = Grammar{Rules: rules}
					g.write(cmd)

					fmt.Println(time.Since(start))
					return nil
				},
			},
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
