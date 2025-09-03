// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitTriplets(t *testing.T) {
	var emp []Sentence

	type args struct {
		s []Sentence
		n []Ngram
	}
	tests := []struct {
		name string
		args args
		want []Sentence
	}{
		{name: "", args: args{s: []Sentence{}, n: []Ngram{}}, want: emp}, // TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SplitTriplets(tt.args.s, tt.args.n))
		})
	}
}
