// -*- coding: utf-8 -*-

// Created on Tue Jul 29 10:12:42 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

type Grammar struct {
	Rules   map[string]Rule
}

func NewGrammar() Grammar {
	var grammar Grammar

	grammar.Rules = make(map[string]Rule)

	return grammar
}
