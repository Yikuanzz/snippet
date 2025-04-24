package ui

import "embed"

//go:embed html/*.tmpl
//go:embed html/partials/*.tmpl
//go:embed html/pages/*.tmpl
var Templates embed.FS

//go:embed static/*
var Files embed.FS
