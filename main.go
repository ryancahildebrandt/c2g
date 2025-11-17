// -*- coding: utf-8 -*-

// Created on Fri Jun 27 03:42:54 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:                  "c2g",
		Usage:                 "Convert natural language expressions to a context free grammar",
		UsageText:             "c2g [COMMAND] [OPTIONS] example.txt",
		EnableShellCompletion: true,
		Suggest:               true,
		Commands: []*cli.Command{
			{
				Name:                  "clone",
				Usage:                 "Create a grammar such that each expression corresponds to one rule, with no rule merging or factoring applied. This mode will produce a grammar covering all utterances in the orignal grammar and no more.",
				UsageText:             "c2g clone [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&printMain,
					&preTokenized,
					&chunk,
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						texts  []Text
						rules  []Rule
						g      Grammar
						err    error
						logger *log.Logger
					)

					texts, err = readInfile(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}

					rules = applyChunking(texts, cmd)
					rules = SetIDs(rules)
					g = Grammar{Rules: rules}
					g.write(cmd)

					return nil
				},
			},
			{
				Name:                  "compress",
				Usage:                 "Create a grammar with rule merging and factoring applied to rules sharing 2 or more chunks. This mode will produce a grammar covering all utterances in the orignal grammar and no more.",
				UsageText:             "c2g compress [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&printMain,
					&preTokenized,
					&chunk,
					&prob,
					&factorN,
					&logging,
					&logFile,
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						texts  []Text
						rules  []Rule
						g      Grammar
						err    error
						logger *log.Logger
					)

					logger, err = setLogger(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					texts, err = readInfile(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}

					rules = applyChunking(texts, cmd)
					rules = MergePR(rules, LiteralEqual(logger), logger)
					rules = MergePS(rules, LiteralEqual(logger), logger)
					rules = MergeRS(rules, LiteralEqual(logger), logger)
					rules = MergeMisc(rules, LiteralEqual(logger), logger)
					rules = SetIDs(rules)
					rules = ExpressionFactor(cmd.Int("factor"), logger)(rules)
					g = Grammar{Rules: rules}
					g.write(cmd)

					return nil
				},
			},
			{
				Name:                  "interpolate",
				Usage:                 "Create a grammar with rule merging and factoring applied to rules sharing 1 or more chunks. This mode may produce outputs not found in the source corpus.",
				UsageText:             "c2g interpolate [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&printMain,
					&preTokenized,
					&chunk,
					&prob,
					&factorN,
					&merge,
					&similarity,
					&conFactor,
					&filterQuantile,
					&logging,
					&logFile,
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						texts   []Text
						rules   []Rule
						g       Grammar
						err     error
						logger  *log.Logger
						eqfunc  EqualityFunction
						facfunc FactorFunction
					)
					logger, err = setLogger(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					eqfunc, err = setMerge(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					facfunc, err = setFactor(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					texts, err = readInfile(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}

					rules = applyChunking(texts, cmd)
					rules = MergePR(rules, eqfunc, logger)
					rules = MergePS(rules, eqfunc, logger)
					rules = MergeRS(rules, eqfunc, logger)
					rules = MergeP(rules, eqfunc, logger)
					rules = MergeR(rules, eqfunc, logger)
					rules = MergeS(rules, eqfunc, logger)
					rules = MergeMisc(rules, eqfunc, logger)
					rules = SetIDs(rules)
					rules = facfunc(rules)
					g = Grammar{Rules: rules}
					g.write(cmd)

					return nil
				},
			},
			{
				Name:                  "extrapolate",
				Usage:                 "Create a grammar with rule merging, factoring, and synonym expansion applied to rules sharing 1 or more chunks. This mode may produce outputs not found in the source corpus.",
				UsageText:             "c2g extrapolate [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&printMain,
					&preTokenized,
					&chunk,
					&prob,
					&factorN,
					&merge,
					&similarity,
					&conFactor,
					&filterQuantile,
					&synFile,
					&logging,
					&logFile,
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						texts   []Text
						rules   []Rule
						g       Grammar
						err     error
						logger  *log.Logger
						eqfunc  EqualityFunction
						facfunc FactorFunction
						synfunc FactorFunction
					)

					logger, err = setLogger(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					eqfunc, err = setMerge(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					facfunc, err = setFactor(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					synfunc, err = setSynonyms(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					texts, err = readInfile(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}

					rules = applyChunking(texts, cmd)
					rules = MergePR(rules, eqfunc, logger)
					rules = MergePS(rules, eqfunc, logger)
					rules = MergeRS(rules, eqfunc, logger)
					rules = MergeP(rules, eqfunc, logger)
					rules = MergeR(rules, eqfunc, logger)
					rules = MergeS(rules, eqfunc, logger)
					rules = MergeMisc(rules, eqfunc, logger)
					rules = SetIDs(rules)
					rules = facfunc(rules)
					rules = synfunc(rules)
					g = Grammar{Rules: rules}
					g.write(cmd)

					return nil
				},
			},
			{
				Name:                  "custom",
				Usage:                 "Create a grammar with a user-specified combination of merging, factoring, and expansion strategies. This mode may produce outputs not found in the source corpus.",
				UsageText:             "c2g custom [OPTIONS] example.txt",
				EnableShellCompletion: true,
				Suggest:               true,
				Before:                prepareContext,
				Flags: []cli.Flag{
					&inFile,
					&outFile,
					&printMain,
					&preTokenized,
					&chunk,
					&prob,
					&factorN,
					&merge,
					&similarity,
					&conFactor,
					&filterQuantile,
					&synFile,
					&merge1,
					&merge2,
					&mergeMisc,
					&factor,
					&logging,
					&logFile,
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					var (
						texts   []Text
						rules   []Rule
						g       Grammar
						err     error
						logger  *log.Logger
						eqfunc  EqualityFunction
						facfunc FactorFunction
						synfunc FactorFunction
					)

					logger, err = setLogger(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					eqfunc, err = setMerge(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					facfunc, err = setFactor(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					synfunc, err = setSynonyms(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}
					texts, err = readInfile(cmd)
					if err != nil {
						logger.Printf("Error: %v", err)
						return err
					}

					rules = applyChunking(texts, cmd)
					if cmd.Bool("merge2") {
						rules = MergePR(rules, eqfunc, logger)
						rules = MergePS(rules, eqfunc, logger)
						rules = MergeRS(rules, eqfunc, logger)
					}
					if cmd.Bool("merge1") {
						rules = MergeP(rules, eqfunc, logger)
						rules = MergeR(rules, eqfunc, logger)
						rules = MergeS(rules, eqfunc, logger)
					}
					if cmd.Bool("mergemisc") {
						rules = MergeMisc(rules, eqfunc, logger)
					}
					rules = SetIDs(rules)
					if cmd.Bool("factor") {
						rules = facfunc(rules)
					}
					rules = synfunc(rules)
					g = Grammar{Rules: rules}
					g.write(cmd)

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
