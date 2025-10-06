// -*- coding: utf-8 -*-

// Created on Mon Oct 27 10:02:39 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"encoding/json"
	"os"
)

type Synonyms map[string][]string

func ReadSynonyms(p string) (Synonyms, error) {
	var syn = Synonyms{}
	var err error

	file, err := os.Open(p)
	if err != nil {
		return syn, err
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&syn)

	return syn, err
}
