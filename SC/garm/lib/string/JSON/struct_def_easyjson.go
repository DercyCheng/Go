// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package main

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

func (BasicInfo) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*BasicInfo) UnmarshalJSON([]byte) error        { return nil }
func (BasicInfo) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*BasicInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_BasicInfo *BasicInfo

func (Employee) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*Employee) UnmarshalJSON([]byte) error        { return nil }
func (Employee) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*Employee) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Employee *Employee

func (JobInfo) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*JobInfo) UnmarshalJSON([]byte) error        { return nil }
func (JobInfo) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*JobInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_JobInfo *JobInfo
