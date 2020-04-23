package main

import (
	"esparse/logging"
	"esparse/parser"
	"fmt"  
)

var internalModules = []string{
	"assert",
	"async_hooks",
	"buffer",
	"child_process",
	"cluster",
	"console",
	"constants",
	"crypto",
	"dgram",
	"dns",
	"domain",
	"events",
	"fs",
	"http",
	"http2",
	"https",
	"inspector",
	"module",
	"net",
	"os",
	"path",
	"perf_hooks",
	"process",
	"punycode",
	"querystring",
	"readline",
	"repl",
	"stream",
	"string_decoder",
	"sys",
	"timers",
	"tls",
	"trace_events",
	"tty",
	"url",
	"util",
	"v8",
	"vm",
	"worker_threads",
	"zlib",
}

func main() {
	log, _ := logging.NewDeferLog()
	source := logging.Source{
		Index:        0,
		AbsolutePath: "/home/runner/TepidBewitchedAddin/index.js",
    PrettyPath:  "/home/runner/TepidBewitchedAddin/index.js",
		Contents:   "import def from 'underscore'\nrequire('react')\nrequire('fs')",
	}
  parseOptions := parser.ParseOptions{
    IsBundling: true,
  }
	ast, ok := parser.Parse(log, source, parseOptions)
  fmt.Println(ok)
  //fmt.Printf("%+v\n", ast)
  outer: for _, importPath := range ast.ImportPaths {
    for _, internal := range internalModules {      
      if internal == importPath.Path.Text { 
        continue outer
      }
    }

    fmt.Println(importPath.Path.Text)
  }
}