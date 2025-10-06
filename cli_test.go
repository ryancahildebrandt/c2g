// -*- coding: utf-8 -*-

// Created on Wed Sep  3 07:41:18 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v3"
)

func TestValidateInFile(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{args: args{p: ""}, assertion: assert.Error},
		{args: args{p: "."}, assertion: assert.Error},
		{args: args{p: " "}, assertion: assert.Error},
		{args: args{p: ".csv"}, assertion: assert.Error},
		{args: args{p: "../a.csv"}, assertion: assert.Error},
		{args: args{p: "data/tests/test1.csv"}, assertion: assert.NoError},
		{args: args{p: "data/tests/test0.csv"}, assertion: assert.Error},
		{args: args{p: "data/tests/test1.text"}, assertion: assert.Error},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tt.assertion(t, ValidateInFile(tt.args.p))
		})
	}
}

func TestValidateOutFile(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{args: args{p: ""}, assertion: assert.Error},
		{args: args{p: "."}, assertion: assert.Error},
		{args: args{p: " "}, assertion: assert.Error},
		{args: args{p: ".jsgf"}, assertion: assert.Error},
		{args: args{p: "../a.jsgf"}, assertion: assert.Error},
		{args: args{p: "outputs/a.jsgf"}, assertion: assert.Error},
		{args: args{p: "outputs"}, assertion: assert.Error},
		{args: args{p: "data/tests/out/tmp"}, assertion: assert.Error},
		{args: args{p: "data/tests/out/tmp.jsgf"}, assertion: assert.NoError},
		{args: args{p: "data/tests/out/tmp.jjsgf"}, assertion: assert.Error},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tt.assertion(t, ValidateOutFile(tt.args.p))
		})
	}
}

