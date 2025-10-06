
// -*- coding: utf-8 -*-

// Created on Tue Oct 14 03:48:57 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
"os"
"log"
"io"
)

func NewFileLogger(p string) (*log.Logger, error) {
	f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return &log.Logger{}, err
	}
	defer f.Close()	
	return log.New(f, "INFO:", log.LstdFlags|log.Lmicroseconds|log.Llongfile), nil
}

func NewNilLogger() (*log.Logger) {
	return log.New(io.Discard, "", log.LstdFlags)
}

func NewStdoutLogger() (*log.Logger) {
	return log.New(os.Stdout, "INFO:", log.LstdFlags|log.Lmicroseconds|log.Llongfile)
}

