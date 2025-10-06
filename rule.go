// -*- coding: utf-8 -*-

// Created on Sun Jul 27 07:52:09 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"slices"
	"strings"
)

// Corresponds to one rule within a JSGF grammar
// includes the pre, root, suf structure of corpus texts
type Rule struct {
	pre      []string
	root     []string
	suf      []string
	isPublic bool
	id       int
}

// Checks if pre, root, and suf are empty slices or contain at least one non-empty string element
func (r *Rule) isEmpty() bool {
	var exp []string

	if len(r.pre) == 0 && len(r.root) == 0 && len(r.suf) == 0 {
		return true
	}
	if len(r.pre)+len(r.root)+len(r.suf) <= 3 {
		exp = append(exp, r.pre...)
		exp = append(exp, r.root...)
		exp = append(exp, r.suf...)
		for i := range exp {
			if exp[i] != "" {
				return false
			}
		}
		return true
	}
	return false
}

// Constructs the string representation of a rule
// removes spaces from in front of punctuation/other boundary characters
// returns each of pre, root, and suf wrapped in parentheses (brackets if there is an empty string present) and each element split by |
func (r *Rule) print(n string) string {
	joinBoundaries := func(s string) string {
		for i := range boundaryChars {
			s = strings.ReplaceAll(s, fmt.Sprint(" ", boundaryChars[i]), boundaryChars[i])
		}
		return s
	}

	fmtExpression := func(g []string) string {
		var opt bool

		slices.Sort(g)
		for j := range g {
			g[j] = joinBoundaries(g[j])
		}
		if slices.Contains(g, "") {
			opt = true
			ind := slices.Index(g, "")
			g = slices.Delete(g, ind, ind+1)
		}
		if len(g) == 0 {
			return ""
		}
		if opt {
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

	pre = fmtExpression(r.pre)
	root = fmtExpression(r.root)
	suf = fmtExpression(r.suf)

	if r.isPublic {
		b.WriteString("public ")
	}
	b.WriteString(fmt.Sprintf("<%s> =", n))

	for _, s := range []string{pre, root, suf} {
		if slices.Contains([]string{"()", "[]", ""}, s) {
			continue
		}
		b.WriteString(fmt.Sprintf(" %s", s))
	}
	b.WriteString(";")

	return b.String()
}

// Derives a rule name using the rule content and id
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

func (r *Rule) sort() Rule {
	slices.Sort(r.pre)
	slices.Sort(r.root)
	slices.Sort(r.suf)
	return *r
}

// Sorts rules and sets integer ids
func SetIDs(rules []Rule) []Rule {
	slices.SortStableFunc(rules, func(i, j Rule) int { return strings.Compare(i.name(), j.name()) })

	for i := range rules {
		rules[i].id = i
	}

	return rules
}
