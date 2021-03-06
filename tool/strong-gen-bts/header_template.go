package main

var _headerTemplate = `
// Code generated by strong tool genbts. DO NOT EDIT.

NEWLINE
/* 
  Package {{.PkgName}} is a generated cache proxy package.
  It is generated from:
  ARGS
*/
NEWLINE

package {{.PkgName}}

import (
	"context"
	{{if .EnableBatch }}"sync"{{end}}
NEWLINE
	"github.com/huangbosbos/go-hutool/cache"
	{{if .EnableBatch }}"github.com/huangbosbos/go-hutool/sync/errgroup.v2"{{end}}
	{{.ImportPackage}}
NEWLINE
	{{if .EnableSingleFlight}}	"golang.org/x/sync/singleflight" {{end}}
)

{{if .UseBTS}}
var _ _bts
{{end }}
{{if .EnableSingleFlight}}
var cacheSingleFlights = [SFCOUNT]*singleflight.Group{SFINIT} 
{{end }}
`
