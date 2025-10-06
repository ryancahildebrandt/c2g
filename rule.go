// -*- coding: utf-8 -*-

// Created on Sun Jul 27 07:52:09 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"slices"
	"strings"
)

type Rule struct {
	pre      []string
	root     []string
	suf      []string
	isPublic bool
	id       int
}

func (r *Rule) isEmpty() bool {
	var e []string

	if len(r.pre) == 0 && len(r.root) == 0 && len(r.suf) == 0 {
		return true
	}
	if len(r.pre)+len(r.root)+len(r.suf) <= 3 {
		e = append(e, r.pre...)
		e = append(e, r.root...)
		e = append(e, r.suf...)
		for _, e := range e {
			if e != "" {
				return false
			}
		}
		return true
	}
	return false
}

func (r *Rule) print(n string) string {
	joinBoundaries := func(s string) string {
		for _, b := range boundaryChars {
			s = strings.ReplaceAll(s, fmt.Sprint(" ", b), b)
		}
		return s
	}

	fmtGroup := func(g []string) string {
		var o bool
		var i int

		for j := range g {
			g[j] = joinBoundaries(g[j])
		}
		if slices.Contains(g, "") {
			o = true
			i = slices.Index(g, "")
			g = slices.Delete(g, i, i+1)
		}
		if len(g) == 0 {
			return ""
		}
		if o {
			return fmt.Sprintf("[%s]", strings.Join(g, "|"))
		}
		return fmt.Sprintf("(%s)", strings.Join(g, "|"))
	}

	var (
		b    strings.Builder
		pre  string
		root string
		suf  string
	)

	if r.isEmpty() {
		return b.String()
	}

	pre = fmtGroup(r.pre)
	root = fmtGroup(r.root)
	suf = fmtGroup(r.suf)

	if r.isPublic {
		b.WriteString("public ")
	}
	b.WriteString(fmt.Sprintf("<%s> =", n))

	for _, s := range []string{pre, root, suf} {
		if s == "()" || s == "[]" || s == "" {
			continue
		}
		b.WriteString(fmt.Sprintf(" %s", s))
	}
	b.WriteString(";")

	return b.String()
}

func (r *Rule) name() string {
	var b string

	b = strings.Join(r.root, "_")
	b = strings.ReplaceAll(b, " ", "_")
	b = strings.ReplaceAll(b, "<", "")
	b = strings.ReplaceAll(b, ">", "")
	if r.id != 0 {
		return fmt.Sprintf("%.*s_%v", 20, b, r.id)
	}

	return fmt.Sprintf("%.*s", 20, b)
}

func SetIDs(rules []Rule) []Rule {
	slices.SortStableFunc(rules, func(i, j Rule) int { return strings.Compare(i.name(), j.name()) })

	for i := range rules {
		rules[i].id = i
	}

	return rules
}
