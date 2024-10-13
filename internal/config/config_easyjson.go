// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package config

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6615c02eDecodeGithubComOlegsxmGoSseChatGitInternalConfig(in *jlexer.Lexer, out *AppConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Production":
			out.Production = bool(in.Bool())
		case "Server":
			easyjson6615c02eDecode(in, &out.Server)
		case "Swagger":
			easyjson6615c02eDecode1(in, &out.Swagger)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6615c02eEncodeGithubComOlegsxmGoSseChatGitInternalConfig(out *jwriter.Writer, in AppConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Production\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Production))
	}
	{
		const prefix string = ",\"Server\":"
		out.RawString(prefix)
		easyjson6615c02eEncode(out, in.Server)
	}
	{
		const prefix string = ",\"Swagger\":"
		out.RawString(prefix)
		easyjson6615c02eEncode1(out, in.Swagger)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AppConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6615c02eEncodeGithubComOlegsxmGoSseChatGitInternalConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AppConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6615c02eEncodeGithubComOlegsxmGoSseChatGitInternalConfig(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AppConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6615c02eDecodeGithubComOlegsxmGoSseChatGitInternalConfig(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AppConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6615c02eDecodeGithubComOlegsxmGoSseChatGitInternalConfig(l, v)
}
func easyjson6615c02eDecode1(in *jlexer.Lexer, out *struct {
	Url string `yaml:"url"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Url":
			out.Url = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6615c02eEncode1(out *jwriter.Writer, in struct {
	Url string `yaml:"url"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Url\":"
		out.RawString(prefix[1:])
		out.String(string(in.Url))
	}
	out.RawByte('}')
}
func easyjson6615c02eDecode(in *jlexer.Lexer, out *struct {
	Address string `yaml:"addr"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Address":
			out.Address = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6615c02eEncode(out *jwriter.Writer, in struct {
	Address string `yaml:"addr"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix[1:])
		out.String(string(in.Address))
	}
	out.RawByte('}')
}