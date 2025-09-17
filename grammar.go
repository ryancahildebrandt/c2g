// -*- coding: utf-8 -*-

// Created on Tue Jul 29 10:12:42 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"slices"
	"strings"
)

type Grammar struct {
	Rules []Rule
}

func (g *Grammar) print() string {
	var b strings.Builder
	b.WriteString("#JSGF V1.0 ISO8859-1 en;\n\ngrammar main;\n\n")

	slices.SortStableFunc(g.Rules, func(i, j Rule) int {
		return strings.Compare(i.print(""), j.print(""))
	})

	for _, v := range g.Rules {
		if v.isPublic && !v.isEmpty() {
			b.WriteString(v.print(v.name()))
			b.WriteString("\n")
		}
	}
	b.WriteString("\n")

	for _, v := range g.Rules {
		if !v.isPublic && !v.isEmpty() {
			b.WriteString(v.print(v.name()))
			b.WriteString("\n")
		}
	}

	return strings.TrimSpace(b.String())
}

func (g *Grammar) printMain() string {
	var b strings.Builder
	var main Rule

	for k, v := range g.Rules {
		if v.isPublic && !v.isEmpty() {
			main.root = append(main.root, fmt.Sprint("<", v.name(), ">"))
			v.isPublic = false
			g.Rules[k] = v
		}
	}
	main.isPublic = true
	b.WriteString("#JSGF V1.0 ISO8859-1 en;\n\ngrammar main;\n\n")
	b.WriteString(main.print("main"))
	b.WriteString("\n")

	for _, v := range g.Rules {
		if v.isPublic && !v.isEmpty() {
			b.WriteString(v.print(v.name()))
			b.WriteString("\n")
		}
	}
	b.WriteString("\n")

	for _, v := range g.Rules {
		if !v.isPublic && !v.isEmpty() {
			b.WriteString(v.print(v.name()))
			b.WriteString("\n")
		}
	}

	return strings.TrimSpace(b.String())
}
