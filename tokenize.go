// -*- coding: utf-8 -*-

// Created on Fri Jul 11 08:51:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/bzick/tokenizer"
)

const (
	Boundary = iota + 1
	WhiteSpace
)

var (
	boundaryChars   []string = []string{".", ",", "?", "!", ":", ";"}
	whiteSpaceChars []string = []string{" ", "\t", "\n", "\r"}
)

// Returns a whitespace and punctuation based tokenizer
func NewUnigramTokenizer() *tokenizer.Tokenizer {
	var lexer *tokenizer.Tokenizer = tokenizer.New()

	lexer.SetWhiteSpaces([]byte{})

	lexer.DefineTokens(WhiteSpace, whiteSpaceChars)
	lexer.DefineTokens(Boundary, boundaryChars)

	return lexer
}

func UnigramTokenize(e string, t *tokenizer.Tokenizer) []string {
	var (
		res     string
		builder strings.Builder
		out                       = []string{}
		stream  *tokenizer.Stream = t.ParseString(e)
	)

	if e == "" {
		return out
	}

	for stream.IsValid() {
		switch {
		case stream.CurrentToken().Is(Boundary):
			builder, out = flushBuilder(builder, out)
			res = stream.CurrentToken().ValueUnescapedString()
			out = append(out, res)
			stream.GoNext()
		case stream.CurrentToken().Is(WhiteSpace):
			builder, out = flushBuilder(builder, out)
			stream.GoNext()
		default:
			builder.WriteString(stream.CurrentToken().ValueUnescapedString())
			stream.GoNext()
		}
	}
	builder, out = flushBuilder(builder, out)

	return out
}

func UnigramNormalize(e string, t *tokenizer.Tokenizer) string {
	return strings.Join(UnigramTokenize(e, t), " ")
}

// Helper function to get current contents of strings.Builder and reset
func flushBuilder(b strings.Builder, o []string) (strings.Builder, []string) {
	var str string = b.String()

	b.Reset()
	if len(str) != 0 {
		o = append(o, str)
	}

	return b, o
}

// Checks that the provided string can be consumed by a tokenizer (is not empty and does not contain byte \x00)
func ValidateTokenizerString(s string) error {
	if s == "" {
		return fmt.Errorf("error when calling ValidateTokenizerString(%v):\n%+w", s, errors.New("cannot tokenize empty string"))
	}
	if bytes.Contains([]byte(s), []byte("\x00")) {
		return fmt.Errorf("error when calling ValidateTokenizerString(%v):\n%+w", s, errors.New("cannot tokenize string containing null char \x00"))
	}

	return nil
}
