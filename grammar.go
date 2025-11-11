// -*- coding: utf-8 -*-

// Created on Tue Jul 29 10:12:42 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/urfave/cli/v3"
)

// Struct to handle grammar export
type Grammar struct {
	Rules []Rule
}

// Constructs grammar headers including configuration and jsgf declarations
func (g *Grammar) frontMatter(c *cli.Command) string {
	var b strings.Builder
	var flags []cli.Flag

	b.WriteString("#JSGF V1.0 ISO8859-1 en;\n")
	b.WriteString("#created using c2g\n")
	b.WriteString("#cfg: {")
	b.WriteString(fmt.Sprintf("\"command\":%v, ", c.Name))
	b.WriteString(fmt.Sprintf("\"inFile\":%v, ", c.String("inFile")))

	for _, f := range c.Flags {
		if f.Names()[0] == "help" {
			continue
		}
		flags = append(flags, f)
	}

	slices.SortStableFunc(flags, func(i, j cli.Flag) int { return strings.Compare(i.Names()[0], j.Names()[0]) })
	for _, f := range flags {
		b.WriteString(fmt.Sprintf("\"%s\":%v, ", f.Names()[0], f.Get()))
	}

	b.WriteString("}\n\n")
	b.WriteString("grammar main;\n\n")

	return b.String()
}

// Constructs grammar body with all non-factored rules set as public
func (g *Grammar) body() string {
	var b strings.Builder
	slices.SortStableFunc(g.Rules, func(i, j Rule) int {
		return strings.Compare(i.print(""), j.print(""))
	})

	for _, rule := range g.Rules {
		if rule.isPublic && !rule.isEmpty() {
			b.WriteString(rule.print(rule.name()))
			b.WriteString("\n")
		}
	}
	b.WriteString("\n")

	for _, rule := range g.Rules {
		if !rule.isPublic && !rule.isEmpty() {
			b.WriteString(rule.print(rule.name()))
			b.WriteString("\n")
		}
	}

	return strings.TrimSpace(b.String())
}

// Constructs grammar body with one main public rule
func (g *Grammar) bodyMain() string {
	var b strings.Builder
	var main = Rule{isPublic: true}

	for _, rule := range g.Rules {
		if !rule.isPublic {
			continue
		}
		if rule.isEmpty() {
			continue
		}
		main.root = append(main.root, fmt.Sprint("<", rule.name(), ">"))
	}

	b.WriteString(main.print("main"))
	b.WriteString("\n\n")

	for _, rule := range g.Rules {
		if rule.isEmpty() {
			continue
		}
		rule.isPublic = false
		b.WriteString(rule.print(rule.name()))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	return strings.TrimSpace(b.String())
}

// Writes grammar to file or stdout
func (g *Grammar) write(c *cli.Command) error {
	var (
		err       error
		b         strings.Builder
		out       = c.String("outFile")
		printMain = c.Bool("main")
	)

	b.WriteString(g.frontMatter(c))

	switch {
	case out == "" && !printMain:
		b.WriteString(g.body())
		fmt.Println(b.String())
		return nil
	case out == "" && printMain:
		b.WriteString(g.bodyMain())
		fmt.Println(b.String())
		return nil
	case !printMain:
		b.WriteString(g.body())
		err = os.WriteFile(out, []byte(b.String()), 0644)
		if err != nil {
			return err
		}
	case printMain:
		b.WriteString(g.bodyMain())
		err = os.WriteFile(out, []byte(b.String()), 0644)
		if err != nil {
			return err
		}
	default:
	}
	return err
}
