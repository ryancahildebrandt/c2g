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
	Sep
	POSSep
)

var (
	boundaryChars   []string = []string{".", ",", "?", "!", ":", ";"}
	whiteSpaceChars []string = []string{" ", "\t", "\n", "\r"}
	preTokenizedSep []string = []string{"<SEP>"}
	posSep          []string = []string{"-"}
)

type Tokenizer interface {
	tokenize(s string) []string
	normalize(s string) string
}

type wordTokenizer struct{ *tokenizer.Tokenizer }

func NewWordTokenizer() wordTokenizer {
	var tok = wordTokenizer{tokenizer.New()}

	tok.SetWhiteSpaces([]byte{})
	tok.DefineTokens(WhiteSpace, whiteSpaceChars)
	tok.DefineTokens(Boundary, boundaryChars)

	return tok
}

func (t wordTokenizer) tokenize(s string) []string {
	var (
		res     string
		builder strings.Builder
		out                       = []string{}
		stream  *tokenizer.Stream = t.ParseString(s)
	)

	if s == "" {
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

func (t wordTokenizer) normalize(s string) string {
	return strings.Join(t.tokenize(s), " ")
}

type sepTokenizer struct{ *tokenizer.Tokenizer }

func NewSepTokenizer() sepTokenizer {
	var tok = sepTokenizer{tokenizer.New()}

	tok.SetWhiteSpaces([]byte{})
	tok.DefineTokens(Sep, preTokenizedSep)

	return tok
}

func (t sepTokenizer) tokenize(s string) []string {
	var (
		builder strings.Builder
		out                       = []string{}
		stream  *tokenizer.Stream = t.ParseString(s)
	)

	if s == "" {
		return out
	}
	if !strings.Contains(s, "<SEP>") {
		return []string{s}
	}

	for stream.IsValid() {
		switch {
		case stream.CurrentToken().Is(Sep):
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

func (t sepTokenizer) normalize(s string) string {
	return strings.Join(t.tokenize(s), " ")
}

type posTokenizer struct{ *tokenizer.Tokenizer }

func NewPOSTokenizer() posTokenizer {
	var tok = posTokenizer{tokenizer.New()}

	tok.SetWhiteSpaces([]byte{})
	tok.DefineTokens(POSSep, posSep)

	return tok
}

func (t posTokenizer) tokenize(s string) []string {
	var (
		builder strings.Builder
		out                       = []string{}
		stream  *tokenizer.Stream = t.ParseString(s)
	)

	if s == "" {
		return out
	}
	if !strings.Contains(s, "-") {
		return []string{s}
	}

	for stream.IsValid() {
		switch {
		case stream.CurrentToken().Is(POSSep):
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

func (t posTokenizer) normalize(s string) string {
	return strings.Join(t.tokenize(s), "-")
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
