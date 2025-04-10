package content

import "embed"

//go:embed templates/*
var TemplateFs embed.FS

//go:embed web/*
var WebFs embed.FS
