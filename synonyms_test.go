// -*- coding: utf-8 -*-

// Created on Mon Oct 27 10:02:39 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSynonyms(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		args      args
		want      Synonyms
		assertion assert.ErrorAssertionFunc
	}{
		{args: args{p: ""}, want: Synonyms{}, assertion: assert.Error},
		{args: args{p: "./data/tests/syn1.json"}, want: Synonyms{"I want to know what": []string{"TEST1", "TEST2"}}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn2.json"}, want: Synonyms{"I want to know what": []string{"TEST5", "TEST6"}, "My bill": []string{"TEST7"}, "my bill": []string{}}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn3.json"}, want: Synonyms{"a": []string{"TEST3", "TEST4"}}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn4.json"}, want: Synonyms{}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn5.json"}, want: Synonyms{}, assertion: assert.Error},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := ReadSynonyms(tt.args.p)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
