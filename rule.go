// -*- coding: utf-8 -*-

// Created on Sun Jul 27 07:52:09 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"strings"
)

type Rule struct {
	root Ngram
	pre  []Expression
	suf  []Expression
}

func NewRule(n Ngram) Rule {
	var r Rule

	r.root = n

	return r
}

func (r *Rule) print(n string) string {
	// if len(r.pre)+len(r.suf) == 0 {
	// 	return fmt.Sprintf("<rule_name> = %s;", r.root.text)
	// }
	// if len(r.pre) == 0 {
	// 	return fmt.Sprintf("<rule_name> = (%s) %s;", strings.Join(r.pre, "|"), r.root.text)
	// }
	// if len(r.suf) == 0 {
	// 	return fmt.Sprintf("<rule_name> = %s (%s);", r.root.text, strings.Join(r.suf, "|"))
	// }

	return fmt.Sprintf("<(%s)> = (%s) (%s) (%s);", n, strings.Join(r.pre, "|"), r.root.text, strings.Join(r.suf, "|"))
}
