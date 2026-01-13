# Repository: example

**Source:** https://github.com/golang/example  
**Generated:** 2026-01-13T14:15:03Z  
**Files:** 64 | **Lines:** 11620 | **Size:** 389 KB

## Directory Structure

```text
├── README.md
├── appengine-hello
│   ├── static
│   │   ├── favicon.ico
│   │   ├── index.html
│   │   ├── script.js
│   │   └── style.css
│   ├── app.go
│   ├── app.yaml
│   └── README.md
├── gotypes
│   ├── defsuses
│   │   └── main.go
│   ├── doc
│   │   └── main.go
│   ├── hello
│   │   └── hello.go
│   ├── hugeparam
│   │   └── main.go
│   ├── implements
│   │   └── main.go
│   ├── lookup
│   │   └── lookup.go
│   ├── nilfunc
│   │   └── main.go
│   ├── pkginfo
│   │   └── main.go
│   ├── skeleton
│   │   └── main.go
│   ├── typeandvalue
│   │   └── main.go
│   ├── gen.go
│   ├── go-types.md
│   ├── Makefile
│   └── README.md
├── hello
│   ├── reverse
│   │   ├── example_test.go
│   │   ├── reverse.go
│   │   └── reverse_test.go
│   ├── go.mod
│   └── hello.go
├── helloserver
│   ├── go.mod
│   └── server.go
├── internal
│   └── cmd
│       └── weave
│           └── weave.go
├── outyet
│   ├── go.mod
│   ├── main.go
│   └── main_test.go
├── ragserver
│   ├── ragserver
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── json.go
│   │   ├── main.go
│   │   └── weaviate.go
│   ├── ragserver-genkit
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── json.go
│   │   └── main.go
│   ├── ragserver-langchaingo
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── json.go
│   │   └── main.go
│   ├── tests
│   │   ├── add-documents.sh
│   │   ├── docker-compose.yml
│   │   ├── query.sh
│   │   ├── weaviate-delete-objects.sh
│   │   └── weaviate-show-all.sh
│   └── README.md
├── slog-handler-guide
│   ├── indenthandler1
│   │   ├── indent_handler.go
│   │   └── indent_handler_test.go
│   ├── indenthandler2
│   │   ├── indent_handler.go
│   │   └── indent_handler_test.go
│   ├── indenthandler3
│   │   ├── indent_handler.go
│   │   └── indent_handler_test.go
│   ├── indenthandler4
│   │   ├── indent_handler.go
│   │   ├── indent_handler_norace_test.go
│   │   └── indent_handler_test.go
│   ├── guide.md
│   ├── Makefile
│   └── README.md
├── template
│   ├── image.tmpl
│   ├── index.tmpl
│   └── main.go
├── go.mod
├── go.sum
├── LICENSE
└── PATENTS
```

## Files

### README.md

```markdown
# Go example projects

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

## Clone the project

```
$ git clone https://go.googlesource.com/example
$ cd example
```
https://go.googlesource.com/example is the canonical Git repository.
It is mirrored at https://github.com/golang/example.

## [hello](hello/) and [hello/reverse](hello/reverse/)

```
$ cd hello
$ go build
$ ./hello -help
```
A trivial "Hello, world" program that uses a library package.

The [hello](hello/) command covers:

* The basic form of an executable command
* Importing packages (from the standard library and the local repository)
* Printing strings ([fmt](//golang.org/pkg/fmt/))
* Command-line flags ([flag](//golang.org/pkg/flag/))
* Logging ([log](//golang.org/pkg/log/))

The [reverse](hello/reverse/) reverse covers:

* The basic form of a library
* Conversion between string and []rune
* Table-driven unit tests ([testing](//golang.org/pkg/testing/))

## [helloserver](helloserver/)

```
$ cd helloserver
$ go run .
```

A trivial "Hello, world" web server.

Topics covered:

* Command-line flags ([flag](//golang.org/pkg/flag/))
* Logging ([log](//golang.org/pkg/log/))
* Web servers ([net/http](//golang.org/pkg/net/http/))

## [outyet](outyet/)

```
$ cd outyet
$ go run .
```
A web server that answers the question: "Is Go 1.x out yet?"

Topics covered:

* Command-line flags ([flag](//golang.org/pkg/flag/))
* Web servers ([net/http](//golang.org/pkg/net/http/))
* HTML Templates ([html/template](//golang.org/pkg/html/template/))
* Logging ([log](//golang.org/pkg/log/))
* Long-running background processes
* Synchronizing data access between goroutines ([sync](//golang.org/pkg/sync/))
* Exporting server state for monitoring ([expvar](//golang.org/pkg/expvar/))
* Unit and integration tests ([testing](//golang.org/pkg/testing/))
* Dependency injection
* Time ([time](//golang.org/pkg/time/))

## [appengine-hello](appengine-hello/)

A trivial "Hello, world" App Engine application intended to be used as the
starting point for your own code. Please see
[Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go)
and [Quickstart for Go in the App Engine Standard Environment](https://cloud.google.com/appengine/docs/standard/go/quickstart).

## [gotypes](gotypes/)

The `go/types` package is a type-checker for Go programs. It is one of the most
complex packages in Go's standard library, so we have provided this tutorial to
help you find your bearings. It comes with several example programs that you
can obtain using `go get` and play with as you learn to build tools that analyze
or manipulate Go programs.

## [template](template/)

A trivial web server that demonstrates the use of the
[`template` package](https://golang.org/pkg/text/template/)'s `block` feature.

## [slog-handler-guide](slog-handler-guide/)

The `log/slog` package supports structured logging.
It features a flexible backend in the form of a `Handler` interface.
This guide can help you write your own handler.
```

### appengine-hello/static/index.html

```html
<!DOCTYPE html>

<!--
Copyright 2023 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->

<html>

<head>
    <title>Hello, world</title>
    <link rel="stylesheet" href="/style.css">
    <script src="/script.js"></script>
</head>

<body>
    <h1>Hello, world</h1>

    <button onclick="fetchMessage()">Fetch Message</button>
    <p id="message">Click on the button to fetch the message.</p>
</body>

</html>
```

### appengine-hello/static/script.js

```javascript
/*
Copyright 2023 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

"use strict";

function fetchMessage() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/hello", false);
    xmlHttp.send(null);
    document.getElementById("message").innerHTML = xmlHttp.responseText;
}
```

### appengine-hello/static/style.css

```css
/*
Copyright 2023 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

h1 {
  text-align: center;
}
```

### appengine-hello/app.go

```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hello is a simple App Engine application that replies to requests
// on /hello with a welcoming message.
package hello

import (
	"fmt"
	"net/http"
)

// init is run before the application starts serving.
func init() {
	// Handle all requests with path /hello with the helloHandler function.
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Go app")
}
```

### appengine-hello/app.yaml

```yaml
# Copyright 2023 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

application: you-application-id
version: 1
runtime: go
api_version: go1

handlers:
- url: /hello
  script: _go_app

- url: /favicon.ico
  static_files: static/favicon.ico
  upload: static/favicon.ico

- url: /
  static_files: static/index.html
  upload: static/index.html

- url: /
  static_dir: static
```

### appengine-hello/README.md

```markdown
This code is a starting point for your Google App Engine application in
Go.

To run the application locally you need the install the [Go Cloud
SDK](https://cloud.google.com/appengine/docs/go/#Go_tools) and execute the next
command from the directory containing this file:

    $ goapp serve app.yaml

To deploy the application you have to first create an App Engine project
and use it as the application file in all the yaml files.
```

### gotypes/defsuses/main.go

```go
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const hello = `package main

import "fmt"

func main() {
        fmt.Println("Hello, world")
}
`

// !+
func PrintDefsUses(fset *token.FileSet, files ...*ast.File) error {
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}
	_, err := conf.Check("hello", fset, files, info)
	if err != nil {
		return err // type error
	}

	for id, obj := range info.Defs {
		fmt.Printf("%s: %q defines %v\n",
			fset.Position(id.Pos()), id.Name, obj)
	}
	for id, obj := range info.Uses {
		fmt.Printf("%s: %q uses %v\n",
			fset.Position(id.Pos()), id.Name, obj)
	}
	return nil
}

//!-

func main() {
	// Parse one file.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		log.Fatal(err) // parse error
	}
	if err := PrintDefsUses(fset, f); err != nil {
		log.Fatal(err) // type error
	}
}

/*
//!+output
$ go build golang.org/x/example/gotypes/defsuses
$ ./defsuses
hello.go:1:9: "main" defines <nil>
hello.go:5:6: "main" defines func hello.main()
hello.go:6:9: "fmt" uses package fmt
hello.go:6:13: "Println" uses func fmt.Println(a ...interface{}) (n int, err error)
//!-output
*/
```

### gotypes/doc/main.go

```go
// The doc command prints the doc comment of a package-level object.
package main

import (
	"fmt"
	"go/ast"
	"log"
	"os"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: doc <package> <object>")
	}
	//!+part1
	pkgpath, name := os.Args[1], os.Args[2]

	// Load complete type information for the specified packages,
	// along with type-annotated syntax.
	// Types for dependencies are loaded from export data.
	conf := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(conf, pkgpath)
	if err != nil {
		log.Fatal(err) // failed to load anything
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1) // some packages contained errors
	}

	// Find the package and package-level object.
	pkg := pkgs[0]
	obj := pkg.Types.Scope().Lookup(name)
	if obj == nil {
		log.Fatalf("%s.%s not found", pkg.Types.Path(), name)
	}
	//!-part1
	//!+part2

	// Print the object and its methods (incl. location of definition).
	fmt.Println(obj)
	for _, sel := range typeutil.IntuitiveMethodSet(obj.Type(), nil) {
		fmt.Printf("%s: %s\n", pkg.Fset.Position(sel.Obj().Pos()), sel)
	}

	// Find the path from the root of the AST to the object's position.
	// Walk up to the enclosing ast.Decl for the doc comment.
	for _, file := range pkg.Syntax {
		pos := obj.Pos()
		if !(file.FileStart <= pos && pos < file.FileEnd) {
			continue // not in this file
		}
		path, _ := astutil.PathEnclosingInterval(file, pos, pos)
		for _, n := range path {
			switch n := n.(type) {
			case *ast.GenDecl:
				fmt.Println("\n", n.Doc.Text())
				return
			case *ast.FuncDecl:
				fmt.Println("\n", n.Doc.Text())
				return
			}
		}
	}
	//!-part2
}

// (The $GOROOT below is the actual string that appears in file names
// loaded from export data for packages in the standard library.)

/*
//!+output
$ ./doc net/http File
type net/http.File interface{Readdir(count int) ([]os.FileInfo, error); Seek(offset int64, whence int) (int64, error); Stat() (os.FileInfo, error); io.Closer; io.Reader}
$GOROOT/src/io/io.go:92:2: method (net/http.File) Close() error
$GOROOT/src/io/io.go:71:2: method (net/http.File) Read(p []byte) (n int, err error)
/go/src/net/http/fs.go:65:2: method (net/http.File) Readdir(count int) ([]os.FileInfo, error)
$GOROOT/src/net/http/fs.go:66:2: method (net/http.File) Seek(offset int64, whence int) (int64, error)
/go/src/net/http/fs.go:67:2: method (net/http.File) Stat() (os.FileInfo, error)

 A File is returned by a FileSystem's Open method and can be
served by the FileServer implementation.

The methods should behave the same as those on an *os.File.
//!-output
*/
```

### gotypes/hello/hello.go

```go
// !+
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

//!-
```

### gotypes/hugeparam/main.go

```go
// The hugeparam command identifies by-value parameters that are larger than n bytes.
//
// Example:
//
//	$ ./hugeparams encoding/xml
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
)

// !+
var bytesFlag = flag.Int("bytes", 48, "maximum parameter size in bytes")

func PrintHugeParams(fset *token.FileSet, info *types.Info, sizes types.Sizes, files []*ast.File) {
	checkTuple := func(descr string, tuple *types.Tuple) {
		for i := 0; i < tuple.Len(); i++ {
			v := tuple.At(i)
			if sz := sizes.Sizeof(v.Type()); sz > int64(*bytesFlag) {
				fmt.Printf("%s: %q %s: %s = %d bytes\n",
					fset.Position(v.Pos()),
					v.Name(), descr, v.Type(), sz)
			}
		}
	}
	checkSig := func(sig *types.Signature) {
		checkTuple("parameter", sig.Params())
		checkTuple("result", sig.Results())
	}
	for _, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncDecl:
				checkSig(info.Defs[n.Name].Type().(*types.Signature))
			case *ast.FuncLit:
				checkSig(info.Types[n.Type].Type.(*types.Signature))
			}
			return true
		})
	}
}

//!-

func main() {
	flag.Parse()

	// Load complete type information for the specified packages,
	// along with type-annotated syntax and the "sizeof" function.
	// Types for dependencies are loaded from export data.
	conf := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(conf, flag.Args()...)
	if err != nil {
		log.Fatal(err) // failed to load anything
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1) // some packages contained errors
	}

	for _, pkg := range pkgs {
		PrintHugeParams(pkg.Fset, pkg.TypesInfo, pkg.TypesSizes, pkg.Syntax)
	}
}

/*
//!+output
% ./hugeparam encoding/xml
/go/src/encoding/xml/marshal.go:167:50: "start" parameter: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/marshal.go:734:97: "" result: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/marshal.go:761:51: "start" parameter: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/marshal.go:781:68: "start" parameter: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/xml.go:72:30: "" result: encoding/xml.StartElement = 56 bytes
//!-output
*/
```

### gotypes/implements/main.go

```go
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

// !+input
const input = `package main

type A struct{}
func (*A) f()

type B int
func (B) f()
func (*B) g()

type I interface { f() }
type J interface { g() }
`

//!-input

func main() {
	// Parse one file.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "input.go", input, 0)
	if err != nil {
		log.Fatal(err) // parse error
	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	//!+implements
	// Find all named types at package level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}

	// Test assignability of all distinct pairs of
	// named types (T, U) where U is an interface.
	for _, T := range allNamed {
		for _, U := range allNamed {
			if T == U || !types.IsInterface(U) {
				continue
			}
			if types.AssignableTo(T, U) {
				fmt.Printf("%s satisfies %s\n", T, U)
			} else if !types.IsInterface(T) &&
				types.AssignableTo(types.NewPointer(T), U) {
				fmt.Printf("%s satisfies %s\n", types.NewPointer(T), U)
			}
		}
	}
	//!-implements
}

/*
//!+output
$ go build golang.org/x/example/gotypes/implements
$ ./implements
*hello.A satisfies hello.I
hello.B satisfies hello.I
*hello.B satisfies hello.J
//!-output
*/
```

### gotypes/lookup/lookup.go

```go
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"strings"
)

// !+input
const hello = `
package main

import "fmt"

// append
func main() {
        // fmt
        fmt.Println("Hello, world")
        // main
        main, x := 1, 2
        // main
        print(main, x)
        // x
}
// x
`

//!-input

// !+main
func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.ParseComments)
	if err != nil {
		log.Fatal(err) // parse error
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	// Each comment contains a name.
	// Look up that name in the innermost scope enclosing the comment.
	for _, comment := range f.Comments {
		pos := comment.Pos()
		name := strings.TrimSpace(comment.Text())
		fmt.Printf("At %s,\t%q = ", fset.Position(pos), name)
		inner := pkg.Scope().Innermost(pos)
		if _, obj := inner.LookupParent(name, pos); obj != nil {
			fmt.Println(obj)
		} else {
			fmt.Println("not found")
		}
	}
}

//!-main

/*
//!+output
$ go build golang.org/x/example/gotypes/lookup
$ ./lookup
At hello.go:6:1,        "append" = builtin append
At hello.go:8:9,        "fmt" = package fmt
At hello.go:10:9,       "main" = func cmd/hello.main()
At hello.go:12:9,       "main" = var main int
At hello.go:14:9,       "x" = var x int
At hello.go:16:1,       "x" = not found
//!-output
*/
```

### gotypes/nilfunc/main.go

```go
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

// !+input
const input = `package main

import "bytes"

func main() {
	var buf bytes.Buffer
	if buf.Bytes == nil && bytes.Repeat != nil && main == nil {
		// ...
	}
}
`

//!-input

var fset = token.NewFileSet()

func main() {
	f, err := parser.ParseFile(fset, "input.go", input, 0)
	if err != nil {
		log.Fatal(err) // parse error
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	if _, err = conf.Check("cmd/hello", fset, []*ast.File{f}, info); err != nil {
		log.Fatal(err) // type error
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if n != nil {
			CheckNilFuncComparison(info, n)
		}
		return true
	})
}

// !+
// CheckNilFuncComparison reports unintended comparisons
// of functions against nil, e.g., "if x.Method == nil {".
func CheckNilFuncComparison(info *types.Info, n ast.Node) {
	e, ok := n.(*ast.BinaryExpr)
	if !ok {
		return // not a binary operation
	}
	if e.Op != token.EQL && e.Op != token.NEQ {
		return // not a comparison
	}

	// If this is a comparison against nil, find the other operand.
	var other ast.Expr
	if info.Types[e.X].IsNil() {
		other = e.Y
	} else if info.Types[e.Y].IsNil() {
		other = e.X
	} else {
		return // not a comparison against nil
	}

	// Find the object.
	var obj types.Object
	switch v := other.(type) {
	case *ast.Ident:
		obj = info.Uses[v]
	case *ast.SelectorExpr:
		obj = info.Uses[v.Sel]
	default:
		return // not an identifier or selection
	}

	if _, ok := obj.(*types.Func); !ok {
		return // not a function or method
	}

	fmt.Printf("%s: comparison of function %v %v nil is always %v\n",
		fset.Position(e.Pos()), obj.Name(), e.Op, e.Op == token.NEQ)
}

//!-

/*
//!+output
$ go build golang.org/x/example/gotypes/nilfunc
$ ./nilfunc
input.go:7:5: comparison of function Bytes == nil is always false
input.go:7:25: comparison of function Repeat != nil is always true
input.go:7:48: comparison of function main == nil is always false
//!-output
*/
```

### gotypes/pkginfo/main.go

```go
// !+
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const hello = `package main

import "fmt"

func main() {
        fmt.Println("Hello, world")
}`

func main() {
	fset := token.NewFileSet()

	// Parse the input string, []byte, or io.Reader,
	// recording position information in fset.
	// ParseFile returns an *ast.File, a syntax tree.
	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		log.Fatal(err) // parse error
	}

	// A Config controls various options of the type checker.
	// The defaults work fine except for one setting:
	// we must specify how to deal with imports.
	conf := types.Config{Importer: importer.Default()}

	// Type-check the package containing only file f.
	// Check returns a *types.Package.
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	fmt.Printf("Package  %q\n", pkg.Path())
	fmt.Printf("Name:    %s\n", pkg.Name())
	fmt.Printf("Imports: %s\n", pkg.Imports())
	fmt.Printf("Scope:   %s\n", pkg.Scope())
}

//!-

/*
//!+output
$ go build golang.org/x/example/gotypes/pkginfo
$ ./pkginfo
Package  "cmd/hello"
Name:    main
Imports: [package fmt ("fmt")]
Scope:   package "cmd/hello" scope 0x820533590 {
.  func cmd/hello.main()
}
//!-output
*/
```

### gotypes/skeleton/main.go

```go
// The skeleton command prints the boilerplate for a concrete type
// that implements the specified interface type.
//
// Example:
//
//	$ ./skeleton io ReadWriteCloser buffer
//	// *buffer implements io.ReadWriteCloser.
//	type buffer struct{ /* ... */ }
//	func (b *buffer) Close() error { panic("unimplemented") }
//	func (b *buffer) Read(p []byte) (n int, err error) { panic("unimplemented") }
//	func (b *buffer) Write(p []byte) (n int, err error) { panic("unimplemented") }
package main

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/packages"
)

const usage = "Usage: skeleton <package> <interface> <concrete>"

// !+
func PrintSkeleton(pkg *types.Package, ifacename, concname string) error {
	obj := pkg.Scope().Lookup(ifacename)
	if obj == nil {
		return fmt.Errorf("%s.%s not found", pkg.Path(), ifacename)
	}
	if _, ok := obj.(*types.TypeName); !ok {
		return fmt.Errorf("%v is not a named type", obj)
	}
	iface, ok := obj.Type().Underlying().(*types.Interface)
	if !ok {
		return fmt.Errorf("type %v is a %T, not an interface",
			obj, obj.Type().Underlying())
	}
	// Use first letter of type name as receiver parameter.
	if !isValidIdentifier(concname) {
		return fmt.Errorf("invalid concrete type name: %q", concname)
	}
	r, _ := utf8.DecodeRuneInString(concname)

	fmt.Printf("// *%s implements %s.%s.\n", concname, pkg.Path(), ifacename)
	fmt.Printf("type %s struct{}\n", concname)
	mset := types.NewMethodSet(iface)
	for i := 0; i < mset.Len(); i++ {
		meth := mset.At(i).Obj()
		sig := types.TypeString(meth.Type(), (*types.Package).Name)
		fmt.Printf("func (%c *%s) %s%s {\n\tpanic(\"unimplemented\")\n}\n",
			r, concname, meth.Name(),
			strings.TrimPrefix(sig, "func"))
	}
	return nil
}

//!-

func isValidIdentifier(id string) bool {
	for i, r := range id {
		if !unicode.IsLetter(r) &&
			!(i > 0 && unicode.IsDigit(r)) {
			return false
		}
	}
	return id != ""
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal(usage)
	}
	pkgpath, ifacename, concname := os.Args[1], os.Args[2], os.Args[3]

	// Load only exported type information for the specified package.
	conf := &packages.Config{Mode: packages.NeedTypes}
	pkgs, err := packages.Load(conf, pkgpath)
	if err != nil {
		log.Fatal(err) // failed to load anything
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1) // some packages contained errors
	}
	if err := PrintSkeleton(pkgs[0].Types, ifacename, concname); err != nil {
		log.Fatal(err)
	}
}

/*
//!+output1
$ ./skeleton io ReadWriteCloser buffer
// *buffer implements io.ReadWriteCloser.
type buffer struct{}
func (b *buffer) Close() error {
	panic("unimplemented")
}
func (b *buffer) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}
func (b *buffer) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}
//!-output1

//!+output2
$ ./skeleton net/http Handler myHandler
// *myHandler implements net/http.Handler.
type myHandler struct{}
func (m *myHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}
//!-output2
*/
```

### gotypes/typeandvalue/main.go

```go
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

// !+input
const input = `
package main

var m = make(map[string]int)

func main() {
	v, ok := m["hello, " + "world"]
	print(rune(v), ok)
}
`

//!-input

var fset = token.NewFileSet()

func main() {
	f, err := parser.ParseFile(fset, "hello.go", input, 0)
	if err != nil {
		log.Fatal(err) // parse error
	}

	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	if _, err := conf.Check("cmd/hello", fset, []*ast.File{f}, info); err != nil {
		log.Fatal(err) // type error
	}

	//!+inspect
	// f is a parsed, type-checked *ast.File.
	ast.Inspect(f, func(n ast.Node) bool {
		if expr, ok := n.(ast.Expr); ok {
			if tv, ok := info.Types[expr]; ok {
				fmt.Printf("%-24s\tmode:  %s\n", nodeString(expr), mode(tv))
				fmt.Printf("\t\t\t\ttype:  %v\n", tv.Type)
				if tv.Value != nil {
					fmt.Printf("\t\t\t\tvalue: %v\n", tv.Value)
				}
			}
		}
		return true
	})
	//!-inspect
}

// nodeString formats a syntax tree in the style of gofmt.
func nodeString(n ast.Node) string {
	var buf bytes.Buffer
	format.Node(&buf, fset, n)
	return buf.String()
}

// mode returns a string describing the mode of an expression.
func mode(tv types.TypeAndValue) string {
	s := ""
	if tv.IsVoid() {
		s += ",void"
	}
	if tv.IsType() {
		s += ",type"
	}
	if tv.IsBuiltin() {
		s += ",builtin"
	}
	if tv.IsValue() {
		s += ",value"
	}
	if tv.IsNil() {
		s += ",nil"
	}
	if tv.Addressable() {
		s += ",addressable"
	}
	if tv.Assignable() {
		s += ",assignable"
	}
	if tv.HasOk() {
		s += ",ok"
	}
	return s[1:]
}

/*
//!+output
$ go build golang.org/x/example/gotypes/typeandvalue
$ ./typeandvalue
make(map[string]int)            mode:  value
                                type:  map[string]int
make                            mode:  builtin
                                type:  func(map[string]int) map[string]int
map[string]int                  mode:  type
                                type:  map[string]int
string                          mode:  type
                                type:  string
int                             mode:  type
                                type:  int
m["hello, "+"world"]            mode:  value,assignable,ok
                                type:  (int, bool)
m                               mode:  value,addressable,assignable
                                type:  map[string]int
"hello, " + "world"             mode:  value
                                type:  string
                                value: "hello, world"
"hello, "                       mode:  value
                                type:  untyped string
                                value: "hello, "
"world"                         mode:  value
                                type:  untyped string
                                value: "world"
print(rune(v), ok)              mode:  void
                                type:  ()
print                           mode:  builtin
                                type:  func(rune, bool)
rune(v)                         mode:  value
                                type:  rune
rune                            mode:  type
                                type:  rune
...more not shown...
//!-output
*/
```

### gotypes/gen.go

```go
//go:generate bash -c "go run ../internal/cmd/weave/weave.go ./go-types.md > README.md"

package gotypes
```

### gotypes/go-types.md

```markdown

# `go/types`: The Go Type Checker

This document is maintained by Alan Donovan `adonovan@google.com`.

[October 2015 GothamGo talk on go/types](https://docs.google.com/presentation/d/13OvHYozAUBeISPRoLgG7kMBuja1NsU1D_mMlmbaYojk/view)


# Contents

%toc

# Changes in Go 1.18

Go 1.18 introduces generics, and several corresponding new APIs for `go/types`.
This document is not yet up-to-date for these changes, but a guide to the new
changes exists at
[`x/exp/typeparams/example`](https://github.com/golang/exp/tree/master/typeparams/example).

# Introduction


The [`go/types` package]('https://golang.org/pkg/go/types) is a
type-checker for Go programs, designed by Robert Griesemer.
It became part of Go's standard library in Go 1.5.
Measured by lines of code and by API surface area, it is one of the
most complex packages in Go's standard library, and using it requires
a firm grasp of the structure of Go programs.
This tutorial will help you find your bearings.
It comes with several example programs that you can obtain with `go get` and play with.
We assume you are a proficient Go programmer who wants to build tools
to analyze or manipulate Go programs and that you have some knowledge
of how a typical compiler works.

The type checker complements several existing
standard packages for analyzing Go programs.
We've listed them below.


	→   go/types
	    go/constant
	    go/parser
	    go/ast
	    go/scanner
	    go/token


Starting at the bottom, the
[`go/token` package](http://golang.org/pkg/go/token)
defines the lexical tokens of Go.
The [`go/scanner` package](http://golang.org/pkg/go/scanner) tokenizes an input stream and records
file position information for use in diagnostics
or for file surgery in a refactoring tool.
The [`go/ast` package](http://golang.org/pkg/go/ast)
defines the data types of the abstract syntax tree (AST).
The [`go/parser` package](http://golang.org/pkg/go/parser)
provides a robust recursive-descent parser that constructs the AST.
And [`go/constant`](http://golang.org/pkg/go/constant)
provides representations and arithmetic operations for the values of compile-time
constant expressions, as we'll see in
[Constants](#constants).



The [`golang.org/x/tools/go/packages` package](https://pkg.go.dev/golang.org/x/tools/go/packages)
from the `x/tools` repository is a client of the type
checker that loads, parses, and type-checks a complete Go program from
source code.
We use it in some of our examples and you may find it useful too.



The Go type checker does three main things.
First, for every name in the program, it determines which declaration
the name refers to; this is known as _identifier resolution_.
Second, for every expression in the program, it determines what type
that expression has, or reports an error if the expression has no
type, or has an inappropriate type for its context; this is known as
_type deduction_.
Third, for every constant expression in the program, it determines the
value of that constant; this is known as _constant evaluation_.
Superficially, it appears that these three processes could be done
sequentially, in the order above, but perhaps surprisingly, they must
be done together.
For example, the value of a constant may depend on the type of an
expression due to operators like `unsafe.Sizeof`.
Conversely, the type of an expression may depend on the value of a
constant, since array types contain constants.
As a result, type deduction and constant evaluation must be done
together.
As another example, we cannot resolve the identifier `k` in the composite
literal `T{k: 0}` until we know whether `T` is a struct type.
If it is, then `k` must be found among `T`'s fields.
If not, then `k` is an ordinary reference
to a constant or variable in the lexical environment.
Consequently, identifier resolution and type deduction are also
inseparable in the general case.



Nonetheless, the three processes of identifier resolution, type
deduction, and constant evaluation can be separated for the purpose of
explanation.


# An Example


The code below shows the most basic use of the type checker to check
the _hello, world_ program, supplied as a string.
Later examples will be variations on this one, and we'll often omit
boilerplate details such as parsing.
To check out and build the examples,
run `go get golang.org/x/example/gotypes/...`.


%include pkginfo/main.go


First, the program creates a
[`token.FileSet`](http://golang.org/pkg/go/token/#FileSet).
To avoid the need to store file names and line and column
numbers in every node of the syntax tree, the `go/token` package
provides `FileSet`, a data structure that stores this information
compactly for a sequence of files.
A `FileSet` records each file name only once, and records
only the byte offsets of each newline, allowing a position within
any file to be identified using a small integer called a
`token.Pos`.
Many tools create a single `FileSet` at startup.
Any part of the program that needs to convert a `token.Pos` into
an intelligible location---as part of an error message, for
instance---must have access to the `FileSet`.



Second, the program parses the input string.
More realistic packages contain several source files, so the parsing
step must be repeated for each one, or better, done in parallel.
Third, it creates a `Config` that specifies type-checking options.
Since the _hello, world_ program uses imports, we must indicate
how to locate the imported packages.
Here we use `importer.Default()`, which loads compiler-generated
export data, but we'll explore alternatives in [Imports](#imports).



Fourth, the program calls `Check`.
This creates a `Package` whose path is `"cmd/hello"`, and
type-checks each of the specified files---just one in this example.
The final (nil) argument is a pointer to an optional `Info`
struct that returns additional deductions from the type checker; more
on that later.
`Check` returns a `Package` even when it also returns an error.
The type checker is robust to ill-formed input,
and goes to great lengths to report accurate
partial information even in the vicinity of syntax or type errors.
`Package` has this definition:


	type Package struct{ ... }
	func (*Package) Path() string
	func (*Package) Name() string
	func (*Package) Scope() *Scope
	func (*Package) Imports() []*Package


Finally, the program prints the attributes of the package, shown below.
(The hexadecimal number may vary from one run to the next.)


%include pkginfo/main.go output -


A package's `Path`, such as `"encoding/json"`, is the string
by which import declarations identify it.
It is unique within a workspace,
and for published packages it must be globally unique.


A package's `Name` is the identifier in the `package`
declaration of each source file within the package, such as `json`.
The type checker reports an error if not all the package declarations in
the package agree.
The package name determines how the package is known when it is
imported into a file (unless a renaming import is used),
but is otherwise not visible to a program.


`Scope` returns the package's [_lexical block_](#scopes),
which provides access to all the named entities or
[_objects_](#objects) declared at package level.
`Imports` returns the set of packages directly imported by this
one, and  may be useful for computing dependencies
(see [Initialization Order](#initialization-order)).



# Objects


The task of identifier resolution is to map every identifier in the
syntax tree, that is, every `ast.Ident`, to an object.
For our purposes, an _object_ is a named entity created by a
declaration, such as a `var`, `type`, or `func`
declaration.
(This is different from the everyday meaning of object in
object-oriented programming.)



Objects are represented by the `Object` interface:


	type Object interface {
	    Name() string   // package-local object name
	    Exported() bool // reports whether the name starts with a capital letter
	    Type() Type     // object type
	    Pos() token.Pos // position of object identifier in declaration
	
	    Parent() *Scope // scope in which this object is declared
	    Pkg() *Package  // nil for objects in the Universe scope and labels
	    Id() string     // object id (see Ids section below)
	}


The first four methods are straightforward; we'll explain the other
three later.
`Name` returns the object's name---an identifier.
`Exported` is a convenience method that reports whether the first
letter of `Name` is a capital, indicating that the object may be
visible from outside the package.
It's a shorthand for `ast.IsExported(obj.Name())`.
`Type` returns the object's type; we'll come back to that in
[Types](#types).



`Pos` returns the source position of the object's declaring identifier.
To make sense of a `token.Pos`, we need to call the
`(*token.FileSet).Position` method, which returns a struct with
individual fields for the file name, line number, column, and byte
offset, though usually we just call its `String` method:


	fmt.Println(fset.Position(obj.Pos())) // "hello.go:10:6"


Objects for predeclared functions and types such as `len` and `int`
do not have a valid (non-zero) position: `!obj.Pos().IsValid()`.


There are eight kinds of objects in the Go type checker.
Most familiar are the kinds that can be declared at package level:
constants, variables, functions, and types.
Less familiar are statement labels, imported package names
(such as `json` in a file containing an `import "encoding/json"`
declaration), built-in functions (such as `append` and
`len`), and the pre-declared `nil`.
The eight types shown below are the only concrete types that satisfy
the `Object` interface.
In other words, `Object` is a _discriminated union_ of 8
possible types, and we commonly use a type switch to distinguish them.


	Object = *Func         // function, concrete method, or abstract method
	       | *Var          // variable, parameter, result, or struct field
	       | *Const        // constant
	       | *TypeName     // type name
	       | *Label        // statement label
	       | *PkgName      // package name, e.g. json after import "encoding/json"
	       | *Builtin      // predeclared function such as append or len
	       | *Nil          // predeclared nil



Objects are canonical.
That is, two Objects `x` and `y` denote the same entity if and only if `x==y`.
(This is generally true but beware that parameterized types complicate matters; see
https://github.com/golang/exp/tree/master/typeparams/example for details.)

Object identity is significant, and objects are routinely compared by
the addresses of the underlying pointers.
A package-level object (func/var/const/type) can be uniquely
identified by its name and enclosing package.
The [`golang.org/x/tools/go/types/objectpath`](https://pkg.go.dev/golang.org/x/tools/go/types/objectpath)
package defines a naming scheme for objects that are
[exported](#imports) from their package or are unexported but form part of the
type of an exported object.
But for most objects, including all function-local objects,
there is no simple way to obtain a string that uniquely identifies it.



The `Parent` method returns the `Scope` (lexical block) in
which the object was declared; we'll come back to this in
[Scopes](#scopes).
Fields and methods are not found in the lexical environment, so
their objects have no `Parent`.
<!-- TODO check this -->



The `Pkg` method returns the `Package` to which this object
belongs, even for objects not declared at package level.
Only predeclared objects have no package.
The `Id` method will be explained in [Ids](#ids).



Not all methods make sense for each kind of object.  For instance,
the last four kinds above have no meaningful `Type` method.
And some kinds of objects have methods in addition to those required by the
`Object` interface:


	func (*Func) Scope() *Scope
	func (*Var) Anonymous() bool
	func (*Var) IsField() bool
	func (*Const) Val() constant.Value
	func (*TypeName) IsAlias() bool
	func (*PkgName) Imported() *Package


`(*Func).Scope` returns the [lexical block](#scopes)
containing the function's type parameters, parameters, results,
and other local declarations.
`(*Var).IsField` distinguishes struct fields from ordinary
variables, and `(*Var).Anonymous` discriminates named fields like
the one in `struct{T T}` from anonymous fields like the one in `struct{T}`.
`(*Const).Val` returns the value of a named [constant](#constants).


`(*TypeName).IsAlias` reports whether the type name declares an alias
for an existing type (as in `type I = int`), as opposed to defining a new
[`Named`](#named-types) type, as in `type Celsius float64`.
(Most `TypeName`s for which `IsAlias()` is true have a `Type()` of
type `*types.Alias`, but `IsAlias()` is also true for the predeclared
`byte` and `rune` types, which are aliases for `uint8` and `int32`.)


`(*PkgName).Imported` returns the package (for instance,
`encoding/json`) denoted by a given import name such as `json`.
Each time a package is imported, a new `PkgName` object is
created, usually with the same name as the `Package` it
denotes, but not always, as in the case of a renaming import.
`PkgName`s are objects, but `Package`s are not.
We'll look more closely at this in [Imports](#imports).



All relationships between the syntax trees (`ast.Node`s) and type
checker data structures such as `Object`s and `Type`s are
stored in mappings outside the syntax tree itself.
Be aware that the `go/ast` package also defines an older deprecated
type called `Object` that resembles---and predates---the type
checker's `Object`, and that `ast.Object`s are held directly by
identifiers in the AST.
They are created by the parser, which has a necessarily limited view
of the package, so the information they represent is at best partial and
in some cases wrong, as in the `T{k: 0}` example mentioned above.
If you are using the type checker, there is no reason to use the older
`ast.Object` mechanism.



# Identifier Resolution


Identifier resolution computes the relationship between
identifiers and objects.
Its results are recorded in the `Info` struct optionally passed
to `Check`.
The fields related to identifier resolution are shown below.


	type Info struct {
		Defs       map[*ast.Ident]Object
		Uses       map[*ast.Ident]Object
		Implicits  map[ast.Node]Object
		Selections map[*ast.SelectorExpr]*Selection
		Scopes     map[ast.Node]*Scope
		...
	}


Since not all facts computed by the type checker are needed by every
client, the API lets clients control which components of the result
should be recorded and which discarded: only fields that hold a
non-nil map will be populated during the call to `Check`.



The two fields of type `map[*ast.Ident]Object` are the most important:
`Defs` records _declaring_ identifiers and
`Uses` records _referring_ identifiers.
In the example below, the comments indicate which identifiers are of
which kind.


	var x int        // def of x, use of int
	fmt.Println(x)   // uses of fmt, Println, and x
	type T struct{U} // def of T, use of U (type), def of U (field)


The final line above illustrates why we don't combine `Defs` and
`Uses` into one map.
In the anonymous field declaration `struct{U}`, the
identifier `U` is both a use of the type `U` (a
`TypeName`) and a definition of the anonymous field (a
`Var`).



The function below prints the location of each referring and defining
identifier in the input program, and the object it refers to.


%include defsuses/main.go


Let's use the _hello, world_ program again as the input:


%include hello/hello.go


This is what it prints:


%include defsuses/main.go output -


Notice that the `Defs` mapping may contain nil entries in a few
cases.
The first line of output reports that the package identifier
`main` is present in the `Defs` mapping, but has no
associated object.



The `Implicits` mapping handles two cases of the syntax in
which an `Object` is declared without an `ast.Ident`, namely type
switches and import declarations.
<!--
A third case: an anonymous function parameter or result variable.
These objects are returned by Signature.Param and Signature.Result.
-->
In the type switch below, which declares a local variable `y`,
the type of `y` is different in each case of the switch:


	switch y := x.(type) {
	case int:
		fmt.Printf("%d", y)
	case string:
		fmt.Printf("%q", y)
	default:
		fmt.Print(y)
	}


To represent this, for each single-type case, the type checker creates
a separate `Var` object for `y` with the appropriate type,
and `Implicits` maps each `ast.CaseClause` to the `Var`
for that case.
The `default` case, the `nil` case, and cases with more than one
type all use the regular `Var` object that is associated with the
identifier `y`, which is found in the `Defs` mapping.



The import declaration below defines the name `json` without an
`ast.Ident`:


	import "encoding/json"


`Implicits` maps this `ast.ImportSpec` to the `PkgName`
object named `json` that it implicitly declares.



The `Selections` mapping, of type
`map[*ast.SelectorExpr]*Selection`, records the meaning of each
expression of the form _`expr`_`.f`, where _`expr`_ is
an expression or type and `f` is the name of a field or method.
These expressions, called _selections_, are represented by
`ast.SelectorExpr` nodes in the AST.
We'll talk more about the `Selection` type in [Selections](#selections).



Not all `ast.SelectorExpr` nodes represent selections.
Expressions like `fmt.Println`, in which a package name precedes
the dot, are _qualified identifiers_.
They do not appear in the `Selections` mapping, but their
constituent identifiers (such as `fmt` and `Println`) both
appear in `Uses`.



Referring identifiers that are not part of an `ast.SelectorExpr`
are _lexical references_.
That is, they are resolved to an object by searching for the
innermost enclosing lexical declaration of that name.
We'll see how that search works in the next section.


# Scopes


The `Scope` type is a mapping from names to objects.


	type Scope struct{ ... }
	
	func (s *Scope) Names() []string
	func (s *Scope) Lookup(name string) Object


`Names` returns the set of names in the mapping, in sorted order.
(It is not a simple accessor though, so call it sparingly.)
The `Lookup ` method returns the object for a given name, so we
can print all the entries or _bindings_ in a scope like this:


	for _, name := range scope.Names() {
		fmt.Println(scope.Lookup(name))
	}


The _scope_ of a declaration of a name is the region of
program source in which a reference to the name resolves to that
declaration.  That is, scope is a property of a declaration.
However, in the `go/types` API, the `Scope` type represents
a _lexical block_, which is one component of the lexical
environment.
Consider the _hello, world_ program again:


	package main
	
	import "fmt"
	
	func main() {
		const message = "hello, world"
		fmt.Println(message)
	}


There are four lexical blocks in this program.
The outermost one is the _universe block_, which maps the
pre-declared names like `int`, `true`, and `append` to
their objects---a `TypeName`, a `Const`, and a
`Builtin`, respectively.
The universe block is represented by the global variable
`Universe`, of type `*Scope`, although it's logically a
constant so you shouldn't modify it.



Next is the _package block_, which maps `"main"` to the
`main` function.
Following that is the _file block_, which maps `"fmt"` to
the `PkgName` object for this import of the `fmt` package.
And finally, the innermost block is that of function `main`, a
local block, which contains the declaration of `message`, a `Const`.
The `main` function is trivial, but many functions contain
several blocks since each `if`, `for`, `switch`,
`case`, or `select` statement creates at least one
additional block.
Local blocks nest to arbitrary depths.



The structure of the lexical environment thus forms a tree, with the
universe block at the root, the package blocks beneath it, the file
blocks beneath them, and then any number of local blocks beneath the
files.
We can access and navigate this tree structure with the following
methods of `Scope`:


	func (s *Scope) Parent() *Scope
	func (s *Scope) NumChildren() int
	func (s *Scope) Child(i int) *Scope


`Parent` lets us walk up the tree, and `Child`
lets us walk down it.
Note that although the `Parent` of every package `Scope` is
`Universe`, `Universe` has no children.
This asymmetry is a consequence of using a global variable to hold
`Universe`.



To obtain the universe block, we use the `Universe` global variable.
To obtain the lexical block of a `Package`, we call its
`Scope` method.
To obtain the scope of a file (`*ast.File`), or any smaller piece
of syntax such as an `*ast.IfStmt`, we consult the `Scopes`
mapping in the `Info` struct, which maps each block-creating
syntax node to its block.
The lexical block of a named function or method can also be obtained
by calling its `(*Func).Scope` method.


<!--
TODO: explain explicit and implicit blocks
TODO: explain Dot imports.
TODO: explain that Func blocks are associated with FuncType (not FuncDecl or FuncLit)
-->


To look up a name in the lexical environment, we must search the tree
of lexical blocks, starting at a particular `Scope` and walking
up to the root until a declaration of the name is found.
For convenience, the `LookupParent` method does this, returning
not just the object, if found, but also the `Scope` in which it was
declared, which may be an ancestor of the initial one:


	func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)


The `pos` parameter determines the position in the source code at
which the name should be resolved.
The effective lexical environment is different at each point in the
block because it depends on which local declarations appear
before or after that point.
(We'll see an illustration in a moment.)



`Scope` has several other methods relating to source positions:


	func (s *Scope) Pos() token.Pos
	func (s *Scope) End() token.Pos
	func (s *Scope) Contains(pos token.Pos) bool
	func (s *Scope) Innermost(pos token.Pos) *Scope


`Pos` and `End` report the `Scope`'s start and end
position which, for explicit blocks, coincide with its curly
braces.
`Contains` is a convenience method that reports whether a
position lies in this interval.
`Innermost` returns the innermost scope containing the specified
position, which may be a child or other descendent of the initial
scope.



These features are useful for tools that wish to resolve names or
evaluate constant expressions as if they had appeared at a particular
point within the program.
The next example program finds all the comments in the input,
treating the contents of each one as a name.  It looks up each name in
the environment at the position of the comment, and prints what it
finds.
Observe that the `ParseComments` flag directs the parser to
preserve comments in the input.


%include lookup/lookup.go main


The expression `pkg.Scope().Innermost(pos)` finds the innermost
`Scope` that encloses the comment, and `LookupParent(name, pos)`
does a name lookup at a specific position in that lexical block.



A typical input is shown below.
The first comment causes a lookup of `"append"` in the file block.
The second comment looks up `"fmt"` in the `main` function's block,
and so on.


%include lookup/lookup.go input -


Here's the output:


%include lookup/lookup.go output -


Notice how the two lookups of `main` return different results,
even though they occur in the same block, because one precedes the
declaration of the local variable named `main` and the other
follows it.
Also notice that there are two lookups of the name `x` but only
the first one, in the function block, succeeds.



Download the program and modify both the input program and
the set of comments to get a better feel for how name resolution works.



The table below summarizes which kinds of objects may be declared at
each level of the tree of lexical blocks.


                Universe File Package Local
    Builtin        ✔
    Nil            ✔
    Const          ✔            ✔      ✔
    TypeName       ✔            ✔      ✔
    Func                        ✔
    Var                         ✔      ✔
    PkgName               ✔
    Label                              ✔


# Initialization Order


In the course of identifier resolution, the type checker constructs a
graph of references among declarations of package-level variables and
functions.
The type checker reports an error if the initializer expression for a
variable refers to that variable, whether directly or indirectly.



The reference graph determines the initialization order of the
package-level variables, as required by the Go spec, using a
breadth-first algorithm.
First, variables in the graph with no successors are removed, sorted
into the order in which they appear in the source code, then added
to a list.  This creates more variables that have no successors.
The process repeats until they have all been removed.



The result is available in the `InitOrder` field of the
`Info` struct, whose type is `[]Initializer`.


	type Info struct {
		...
		InitOrder []Initializer
		...
	}
	
	type Initializer struct {
		Lhs []*Var // var Lhs = Rhs
		Rhs ast.Expr
	}


Each element of the list represents a single initializer expression
that must be executed, and the variables to which it is assigned.
The variables may number zero, one, or more, as in these examples:


	var _ io.Writer = new(bytes.Buffer)
	var rx = regexp.MustCompile("^b(an)*a$")
	var cwd, cwdErr = os.Getwd()


This process governs the initialization order of variables within a
package.
Across packages, dependencies must be initialized first, although the
order among them is not specified.
That is, any topological order of the import graph will do.
The `(*Package).Imports` method returns the set of direct
dependencies of a package.


# Types


The main job of the type checker is, of course, to deduce the type
of each expression and to report type errors.
Like `Object`, `Type` is an interface type used as a
discriminated union of several concrete types but, unlike
`Object`, `Type` has very few methods because types have
little in common with each other.
Here is the interface:


	type Type interface {
		Underlying() Type
	}


And here are the 14 concrete types that satisfy it:


	Type = *Basic
	     | *Pointer
	     | *Array
	     | *Slice
	     | *Map
	     | *Chan
	     | *Struct
	     | *Tuple
	     | *Signature
	     | *Alias
	     | *Named
	     | *Interface
	     | *Union
	     | *TypeParam


With the exception of `Named` types, instances of `Type` are
not canonical.
(Even for `Named` types, parameterized types complicate matters; see
https://github.com/golang/exp/tree/master/typeparams/example.)
That is, it is usually a mistake to compare types using `t1==t2`
since this equivalence is not the same as the
[type identity relation](https://golang.org/ref/spec#Type_identity)
defined by the Go spec.
Use this function instead:


	func Identical(t1, t2 Type) bool


For the same reason, you should not use a `Type` as a key in a map.
The [`golang.org/x/tools/go/types/typeutil` package](https://pkg.go.dev/golang.org/x/tools/go/types/typeutil)
provides a map keyed by types that uses the correct
equivalence relation.


The Go spec defines three relations over types.
[_Assignability_](https://golang.org/ref/spec#Assignability)
governs which pairs of types may appear on the
left- and right-hand side of an assignment, including implicit
assignments such as function calls, map and channel operations, and so
on.
[_Comparability_](https://golang.org/ref/spec#Comparison_operators)
determines which types may appear in a comparison `x==y` or a
switch case or may be used as a map key.
[_Convertibility_](https://golang.org/ref/spec#Conversions)
governs which pairs of types are allowed in a conversion operation
`T(v)`.
You can query these relations with the following predicate functions:


	func AssignableTo(V, T Type) bool
	func Comparable(T Type) bool
	func ConvertibleTo(V, T Type) bool


Let's take a look at each kind of type.


## Basic types


`Basic` represents all types that are not composed from simpler
types.
This is essentially the set of underlying types that a constant expression is
permitted to have--strings, booleans, and numbers---but it also
includes `unsafe.Pointer` and untyped nil.


	type Basic struct{...}
	func (*Basic) Kind() BasicKind
	func (*Basic) Name() string
	func (*Basic) Info() BasicInfo


The `Kind` method returns an "enum" value that indicates which
basic type this is.
The kinds `Bool`, `String`, `Int16`, and so on,
represent the corresponding predeclared boolean, string, or numeric
types.
There are two aliases: `Byte` is an alias for `Uint8`
and `Rune` is an alias for `Int32`.
The kind `UnsafePointer` represents `unsafe.Pointer`.
The kinds `UntypedBool`, `UntypedInt` and so on represent
the six kinds of "untyped" constant types: boolean, integer, rune,
float, complex, and string.
The kind `UntypedNil` represents the type of the predeclared
`nil` value.
And the kind `Invalid` indicates the invalid type, which is used
for expressions containing errors, or for objects without types, like
`Label`, `Builtin`, or `PkgName`.



The `Name` method returns the name of the type, such as
`"float64"`, and the `Info` method returns a bitfield that
encodes information about the type, such as whether it is signed or
unsigned, integer or floating point, or real or complex.



`Typ` is a table of canonical basic types, indexed by
kind, so `Typ[String]` returns the `*Basic` that represents
`string`, for instance.
Like `Universe`, `Typ` is logically a constant, so don't
modify it.



A few minor subtleties:

- According to the Go spec, pre-declared types such as `int` are
  named types for the purposes of assignability, even though the type
  checker does not represent them using `Named`.

- `unsafe.Pointer` is a pointer type for the purpose of
  determining whether the receiver type of a method is legal, even
  though the type checker does not represent it using `Pointer`.

- The "untyped" types are usually only ascribed to constant expressions,
  but there is one exception.
  A comparison `x==y` has type "untyped bool", so the result of
  this expression may be assigned to a variable of type `bool` or
  any other named boolean type.


## Simple Composite Types


The types `Pointer`, `Array`, `Slice`, `Map`,
and `Chan` are pretty self-explanatory.
All have an `Elem` method that returns the element type `T`
for a pointer `*T`, an array `[n]T`, a slice `[]T`, a
map `map[K]T`, or a channel `chan T`.
This should feel familiar if you've used the `reflect.Value` API.



In addition, the `*Map`, `*Chan`, and `*Array` types
have accessor methods that return their key type, direction, and
length, respectively:


	func (*Map) Key() Type
	func (*Chan) Dir() ChanDir      // = Send | Recv | SendRecv
	func (*Array) Len() int64



## Struct Types


A struct type has an ordered list of fields and a corresponding
ordered list of field tags.


	type Struct struct{ ... } 
	func (*Struct) NumFields() int
	func (*Struct) Field(i int) *Var
	func (*Struct) Tag(i int) string


Each field is a `Var` object whose `IsField` method returns true.
Field objects have no `Parent` scope, because they are
resolved through selections, not through the lexical environment.
<!-- TODO check this -->



Thanks to embedding, the expression `new(S).f` may be a shorthand
for a longer expression such as `new(S).d.e.f`, but in the
representation of `Struct` types, these field selection
operations are explicit.
That is, the set of fields of struct type `S` does not include `f`.
An anonymous field is represented like a regular field, but its
`Anonymous` method returns true.



One subtlety is relevant to tools that generate documentation.
When analyzing a declaration such as this,


	type T struct{x int}


it may be tempting to consider the `Var` object for field `x` as if it
had the name `"T.x"`, but beware: field objects do not have
canonical names and there is no way to obtain the name `"T"`
from the `Var` for `x`.
That's because several types may have the same underlying struct type,
as in this code:


	type T struct{x int}
	type U T


Here, the `Var` for field `x` belongs equally to `T`
and to `U`, and short of inspecting source positions or walking
the AST---neither of which is possible for objects loaded from compiler
export data---it is not possible to ascertain that `x` was declared as
part of `T`.
The type checker builds the exact same data structures given this input:


	type T U
	type U struct{x int}


A similar issue applies to the methods of named interface types.


## Tuple Types


Like a struct, a tuple type has an ordered list of fields, and fields
may be named.


	type Tuple struct{ ... }
	func (*Tuple) Len() int
	func (*Tuple) At(i int) *Var


Although tuples are not the type of any variable in Go, they are
the type of some expressions, such as the right-hand sides of these
assignments:


	v, ok = m[key]
	v, ok = <-ch
	v, ok = x.(T)
	f, err = os.Open(filename)


Tuples also represent the types of the parameter list and the result
list of a function, as we will see.



Since empty tuples are common, the nil `*Tuple` pointer is a valid empty tuple.



## Function and Method Types


The types of functions and methods are represented by a `Signature`,
which has a tuple of parameter types and a tuple of result types.


	type Signature struct{ ... }
	func (*Signature) Recv() *Var
	func (*Signature) Params() *Tuple
	func (*Signature) Results() *Tuple
	func (*Signature) Variadic() bool


Variadic functions such as `fmt.Println` have the `Variadic`
flag set.
The final parameter of such functions is always a slice, or in the
special case of certain calls to `append`, a string.



A `Signature` for a method, whether concrete or abstract, has a
non-nil receiver parameter, `Recv`.
The type of the receiver is usually a named type or a pointer to a named type, 
but it may be an unnamed struct or interface type in some cases.
Method types are rather second-class: they are only used for the
`Func` objects created by method declarations, and no Go
expression has a method type.
When printing a method type, the receiver does not appear, and the
`Identical` predicate ignores the receiver.



The types of `Builtin` objects like `append` cannot be
expressed as a `Signature` since those types require parametric
polymorphism.
`Builtin` objects are thus ascribed the `Invalid` basic type.
However, the type of each _call_ to a built-in function has a specific
and expressible Go type.
These types are recorded during type checking for later use
([TypeAndValue](#typeandvalue)).



## Alias Types

Type declarations come in two forms, aliases and defined types.

Aliases, though introduced only in Go 1.9 and not very common, are
simplest, so we'll present them first and explain defined types in
the next section ("Named Types").

An alias type declaration declares an alternative name for an existing
type. For example, this declaration lets you use the type `Dictionary`
as a synonym for `map[string]string`:

	type Dictionary = map[string]string

The declaration creates a `TypeName` object for `Dictionary`.
The object's `IsAlias` method returns true,
and its `Type` method returns an `Alias`:

        type Alias struct{ ... }
	func (a *Alias) Obj() *TypeName
	func (a *Alias) Origin() *Alias
	func (a *Alias) Rhs() Type
	func (a *Alias) SetTypeParams(tparams []*TypeParam)
	func (a *Alias) TypeArgs() *TypeList
	func (a *Alias) TypeParams() *TypeParamList

The type on the right-hand side of an alias declaration,
such as `map[string]string` in the example above,
can be accessed using the `Rhs()` method.
The `types.Unalias(t)` helper function recursively applies `Rhs`,
removing all `Alias` types from the operand t and returning the
outermost non-alias type.

The `Obj` method returns the declaring `TypeName` object, such as
`Dictionary`; it provides the name, position, and other properties of
the declaration. Conversely, the `TypeName` object's `Type` method
returns the `Alias` type.

Starting with Go 1.24, alias types may have type parameters.
For example, this declaration creates an Alias type with
a type parameter:

	type Set[T comparable] = map[T]bool

Each instantiation such as `Set[string]` is identical to the
corresponding instantiation of the alias' right-hand side type, such
as `map[string]bool`.

The remaining methods--Origin, SetTypeParams, TypeArgs,
TypeParams--are all concerned with type parameters. For now, see
https://github.com/golang/exp/tree/master/typeparams/example.

Prior to Go 1.22, aliases were not materialized as `Alias` types:
each reference to an alias type such as `Dictionary` would be
immediately replaced by its right-hand side type, leaving no
indication in the output of the type checker that an alias was
present.
By materializing alias types, optionally in Go 1.22 and by default
starting in Go 1.23, we can more faithfully record the structure of
the program, which improves the quality of diagnostic messages and
enables certain analyses and code transformations. And, crucially, it
enabled the addition of parameterized aliases in Go 1.24.)



## Named Types



The second form of type declaration, and the only kind prior to Go
1.9, does not use an equals sign:

	type Celsius float64

This declaration does more than just give a name to a type.
It first defines a new `Named` type
whose underlying type is `float64`; this `Named` type is different
from any other type, including `float64`.  The declaration binds the
`TypeName` object to the `Named` type.

Since Go 1.9, the Go language specification has used the term _defined
types_ instead of named types;
the essential property of a defined type is not that it has a name
(aliases and type parameters also have names)
but that it is a distinct type with its own method set.
However, the type checker API predates that
change and instead calls defined types `Named` types.

	type Named struct{ ... }
	func (*Named) NumMethods() int
	func (*Named) Method(i int) *Func
	func (*Named) Obj() *TypeName
	func (*Named) Underlying() Type

The `Named` type's `Obj` method returns the `TypeName` object, which
provides the name, position, and other properties of the declaration.
Conversely, the `TypeName` object's `Type` method returns the `Named` type.

A `Named` type may appear as the receiver type in a method declaration.
Methods are associated with the `Named` type, not the name (the
`TypeName` object); it's possible---though cryptic---to declare a
method on a `Named` type using one of its aliases.
The `NumMethods` and `Method` methods enumerate the declared
methods associated with this `Named` type (or a pointer to it),
in the order they were declared.
However, due to the subtleties of anonymous fields and the difference
between value and pointer receivers, a named type may have more or fewer
methods than this list.  We'll return to this in [Method Sets](#method-sets).



Every `Type` has an `Underlying` method, but for all of them
except `*Named` and `*Alias`, it is simply the identity function.
For a named or alias type, `Underlying` returns its underlying type, which
is always an unnamed type.
Thus `Underlying` returns `int` for both `T` and
`U` below.


	type T int
	type U T


Clients of the type checker often use type assertions or type switches
with a `Type` operand.
When doing so, it is often necessary to switch on the type that
underlies the type of interest, and failure to do so may be a bug.

This is a common pattern:


	// handle types of composite literal
	switch u := t.Underlying().(type) {  // remove any *Named and *Alias types
	case *Struct:        // ...
	case *Map:           // ...
	case *Array, *Slice: // ...
	default:
		panic("impossible")
	}

## Interface Types


Interface types are represented by `Interface`.


	type Interface struct{ ... }
	func (*Interface) Empty() bool
	func (*Interface) NumMethods() int
	func (*Interface) Method(i int) *Func
	func (*Interface) NumEmbeddeds() int
	func (*Interface) Embedded(i int) *Named
	func (*Interface) NumExplicitMethods() int
	func (*Interface) ExplicitMethod(i int) *Func


Syntactically, an interface type has a list of explicitly declared
methods (`ExplicitMethod`), and a list of embedded named
interface types (`Embedded`), but many clients care only about
the complete set of methods, which can be enumerated via
`Method`.
All three lists are ordered by name.
Since the empty interface is an important special case, the
`Empty` predicate provides a shorthand for `NumMethods() ==
0`.



As with the fields of structs (see above), the methods of interfaces
may belong equally to more than one interface type.
The `Func` object for method `f` in the code below is shared
by `I` and `J`:


	type I interface { f() }
	type J I


Because the difference between interface (abstract) and
non-interface (concrete) types is so important in Go, the
`IsInterface` predicate is provided for convenience.


	func IsInterface(Type) bool


The type checker provides three utility methods relating to interface
satisfaction:


	func Implements(V Type, T *Interface) bool
	func AssertableTo(V *Interface, T Type) bool
	func MissingMethod(V Type, T *Interface, static bool) (method *Func, wrongType bool)


The `Implements` predicate reports whether a type satisfies an
interface type.
`MissingMethod` is like `Implements`, but instead of
returning false, it explains why a type does not satisfy the
interface, for use in diagnostics.



`AssertableTo` reports whether a type assertion `v.(T)` is legal.
If `T` is a concrete type that doesn't have all the methods of
interface `v`, then the type assertion is not legal, as in this example:


	// error: io.Writer is not assertible to int
	func f(w io.Writer) int { return w.(int) }




## TypeParam types


A `TypeParam` is the type of a type parameter.
For example, the type of the variable `x` in the `identity` function
below is a `TypeParam`:

    func identity[T any](x T) T { return x }

As with `Alias` and `Named` types, each `TypeParam` has an associated
`TypeName` object that provides its name, position, and other
properties of the declaration.

See https://github.com/golang/exp/tree/master/typeparams/example
for a more thorough treatment of parameterized types.




## Union types

A `Union` is the type of type-parameter constraint of the form `func
f[T int | string]`.

See https://github.com/golang/exp/tree/master/typeparams/example
for a more thorough treatment of parameterized types.



## TypeAndValue


The type checker records the type of each expression in another field
of the `Info` struct, namely `Types`:


	type Info struct {
		...
		Types map[ast.Expr]TypeAndValue
	}


No entries are recorded for identifiers since the `Defs` and
`Uses` maps provide more information about them.
Also, no entries are recorded for pseudo-expressions like
`*ast.KeyValuePair` or `*ast.Ellipsis`.



The value of the `Types` map is a `TypeAndValue`, which
(unsurprisingly) holds the type and value of the expression, and in
addition, its _mode_.
The mode is opaque, but has predicates to answer questions such as:
Does this expression denote a value or a type?  Does this value have an
address?  Does this expression appear on the left-hand side of an
assignment?  Does this expression appear in a context that expects two
results?
The comments in the code below give examples of expressions that
satisfy each predicate.


	type TypeAndValue struct {
		Type  Type
		Value constant.Value // for constant expressions only
		...
	}
	
	func (TypeAndValue) IsVoid() bool      // e.g. "main()"
	func (TypeAndValue) IsType() bool      // e.g. "*os.File"
	func (TypeAndValue) IsBuiltin() bool   // e.g. "len(x)"
	func (TypeAndValue) IsValue() bool     // e.g. "*os.Stdout"
	func (TypeAndValue) IsNil() bool       // e.g. "nil"
	func (TypeAndValue) Addressable() bool // e.g. "a[i]" but not "f()", "m[key]"
	func (TypeAndValue) Assignable() bool  // e.g. "a[i]", "m[key]"
	func (TypeAndValue) HasOk() bool       // e.g. "<-ch", "m[key]"


The statement below inspects every expression within the AST of a single
type-checked file and prints its type, value, and mode:


%include typeandvalue/main.go inspect


It makes use of these two helper functions, which are not shown:


	// nodeString formats a syntax tree in the style of gofmt.
	func nodeString(n ast.Node) string
	
	// mode returns a string describing the mode of an expression.
	func mode(tv types.TypeAndValue) string


Given this input:


%include typeandvalue/main.go input -


the program prints:


%include typeandvalue/main.go output -


Notice that the identifiers for the built-ins `make` and
`print` have types that are specific to the particular calls in
which they appear.
Also notice `m["hello"]` has a 2-tuple type `(int, bool)`
and that it is assignable, but unlike the variable `m`, it is not
addressable.



Download the example and vary the inputs and see what the program prints.



Here's another example, adapted from the `govet` static checking tool.
It checks for accidental uses of a method value `x.f` when a
call `x.f()` was intended;
comparing a method `x.f` against nil is a common mistake.


%include nilfunc/main.go


Given this input,


%include nilfunc/main.go input -


the program reports these errors:


%include nilfunc/main.go output -

# Selections


A _selection_ is an expression _`expr`_`.f` in which
`f` denotes either a struct field or a method.
A selection is resolved not by looking for a name in the lexical
environment, but by looking within a _type_.
The type checker ascertains the meaning of each selection in the
package---a surprisingly tricky business---and records it in the
`Selections` mapping of the `Info` struct, whose values are
of type `Selection`:


	type Selection struct{ ... }
	func (s *Selection) Kind() SelectionKind // = FieldVal | MethodVal | MethodExpr
	func (s *Selection) Recv() Type
	func (s *Selection) Obj() Object
	func (s *Selection) Type() Type
	func (s *Selection) Index() []int
	func (s *Selection) Indirect() bool


The `Kind` method discriminates between the three (legal) kinds
of selections, as indicated by the comments below.


	type T struct{Field int}
	func (T) Method() {}
	var v T
	
                         // Kind            Type
        var _ = v.Field  // FieldVal        int
        var _ = v.Method // MethodVal       func()
        var _ = T.Method // MethodExpr      func(T)


Because of embedding, a selection may denote more than one field or
method, in which case it is ambiguous, and no `Selection` is
recorded for it.



The `Obj` method returns the `Object` for the selected field
(`*Var`) or method (`*Func`).
Due to embedding, the object may belong to a different type than that
of the receiver expression _`expr`_.
The `Type` method returns the type of the selection.  For a field
selection, this is the type of the field, but for method selections,
the result is a function type that is not the same as the type of the
method.
For a `MethodVal`, the receiver parameter is dropped, and
for a `MethodExpr`, the receiver parameter becomes a regular
parameter, as shown in the example above.



The `Index` and `Indirect` methods report information about
implicit operations occurring during the selection that a compiler
would need to know about.
Because of embedding, a selection _`expr`_`.f` may be
shorthand for a sequence containing several implicit field selections,
_`expr`_`.d.e.f`, and `Index` reports the complete
sequence.
And because of automatic pointer dereferencing during struct field
accesses and method calls, a selection may imply one or more indirect
loads from memory; `Indirect` reports whether this occurs.



Clients of the type checker can call `LookupFieldOrMethod` to
look up a name within a type, as if by a selection.
This function has an intimidating signature, but conceptually it
accepts just a `Type` and a name, and returns a `Selection`:


	func LookupFieldOrMethod(T Type, addressable bool, pkg *Package, name string) \
	    (obj Object, index []int, indirect bool)


The result is not actually a `Selection`, but it contains the
three main components of one: `Obj`, `Index`,
and `Indirect`.



The `addressable` flag should be set if the receiver is a
_variable_ of type `T`, since in a method selection on a
variable, an implicit address-of operation (`&`) may occur.
The flag indicates whether the methods of type `*T` should be
considered during the lookup.
(You may wonder why this parameter is necessary.  Couldn't clients
instead call `LookupFieldOrMethod` on the pointer type `*T`
if the receiver is a `T` variable?  The answer is that if
`T` is an interface type, the type `*T` has no methods at
all.)



The final two parameters of `LookupFieldOrMethod` are `(pkg
*Package, name string)`.
Together they specify the name of the field or method to look up.
This brings us to `Id`s.



# Ids


`LookupFieldOrMethod`'s need for a `Package` parameter
is a subtle consequence of the
[_Uniqueness of identifiers_](https://golang.org/ref/spec#Uniqueness_of_identifiers)
section in the Go spec: "Two
identifiers are different if they are spelled differently, or if they
appear in different packages and are not exported."
In practical terms, this means that a type may have two methods
(or two fields, or one of each) both named `f` so long as those
methods are defined in different packages, as in this example:


	package a
	type A int
	func (A) f()
	
	package b
	type B int
	func (B) f()
	
	package c
	import ( "a"; "b" )
	type C struct{a.A; b.B} // C has two methods called f


The type `c.C` has two methods named `f`, but there is
no ambiguity because the two `f`s are distinct
identifiers---think of them as `fᵃ` and `fᵇ`.
For an exported method, this situation _would_ be ambiguous
because there is no distinction between `Fᵃ` and `Fᵇ`; there
is only `F`.



Despite having two methods called `f`, neither of them can be
called from within package `c` because `c` has no way to
identify them.
Within `c`,  `f` is the identifier `fᶜ`, and
type `C` has no method of that name.
But if we pass an instance of `C` to code in package `a`
and call its `f` method via an interface, `fᵃ` is called.



The practical consequence for tool builders is that any time you need
to look up a field or method by name, or construct a map of fields and/or
methods keyed by name, it is not sufficient to use the object's name
as a key.
Instead, you must call the `Object.Id` method, which returns
a string that incorporates the object name, and for unexported
objects, the package path too.
There is also a standalone function `Id` that combines a name and
the package path in the same way:


	func Id(pkg *Package, name string) string


This distinction applies to selections _`expr`_`.f`, but not
to lexical references `x` because for unexported identifiers,
declarations and references always appear in the same package.



Fun fact: the `reflect.StructField` type records both the
`Name` and the `PkgPath` strings for the same reason.
The `FieldByName` methods of `reflect.Value` and
`reflect.Type` match field names without regard to the package.
If there is more than one match, they return an invalid value.




# Method Sets


The _method set_ of a type is the set of methods that can be
called on any value of that type.
(A variable of type `T` has access to all the methods of type
`*T` as well, due to the implicit address-of operation during
method calls, but those extra methods are not part of the method set
of `T`.)



Clients can request the method set of a type `T` by calling
`NewMethodSet(T)`:


	type MethodSet struct{ ... }
	func NewMethodSet(T Type) *MethodSet
	func (s *MethodSet) Len() int
	func (s *MethodSet) At(i int) *Selection
	func (s *MethodSet) Lookup(pkg *Package, name string) *Selection


The `Len` and `At` methods access a list of
`Selections`, all of kind `MethodVal`, ordered by `Id`.
The `Lookup` function allows lookup of a single method by
name (and package path, as explained in the previous section).



`NewMethodSet` can be expensive, so for applications that compute
method sets repeatedly, `golang.org/x/tools/go/types/typeutil`
provides a `MethodSetCache` type that records previous results.
If you only need a single method, don't construct the
`MethodSet` at all; it's cheaper to use
`LookupFieldOrMethod`.



The next program generates a boilerplate
declaration of a new concrete type that satisfies an existing
interface.
Here's an example:


%include skeleton/main.go output1 -


The three arguments are the package and the name of the existing
interface type, and the name of the new concrete type.
The `main` function (not shown) loads the specified package and
calls `PrintSkeleton` with the remaining two arguments:


%include skeleton/main.go


First, `PrintSkeleton` locates the package-level named interface
type, handling various error cases.
Then it chooses the name for the receiver of the new methods: the
first letter of the concrete type.
Finally, it iterates over the method set of the interface, printing
the corresponding concrete method declarations.



There's a subtlety in the declaration of `sig`, which is the
string form of the method signature.
We could have obtained this string from `meth.Type().String()`,
but this would cause any named types within it to be formatted with
the complete package path, for instance
`net/http.ResponseWriter`, which is informative in diagnostics
but not legal Go syntax.
The `TypeString` function (explained in [Formatting Values](#formatting-values)) allows the
caller to control how packages are printed.
Passing `(*types.Package).Name` causes only the package name
`http` to be printed, not the complete path.
Here's another example that illustrates it:


%include skeleton/main.go output2 -


The following program inspects all pairs of package-level named types
in `pkg`, and reports the types that satisfy each interface type.


%include implements/main.go implements


Given this input,


%include implements/main.go input


the program prints:


%include implements/main.go output -


Notice that the method set of `B` does not include `g`, but
the method set of `*B` does.
That's why we needed the second assignability check, using the pointer
type `types.NewPointer(T)`.



# Constants


A constant expression is one whose value is guaranteed to be computed at
compile time.
Constant expressions may appear in types, specifically as the length
of an array type such as `[16]byte`, so one of the jobs of the
type checker is to compute the value of each constant expression.



As we saw in the `typeandvalue` example, the type checker records
the value of each constant expression like `"Hello, " + "world"`,
storing it in the `Value` field of the `TypeAndValue` struct.
Constants are represented using the `Value` interface from the
`go/constant` package.


	package constant // go/constant
	
	type Value interface {
		Kind() Kind
	}
	
	type Kind int // one of Unknown, Bool, String, Int, Float, Complex


The interface has only one method, for discriminating the various
kinds of constants, but the package provides many functions for
inspecting a value of a known kind,


	// Accessors
	func BoolVal(x Value) bool
	func Float32Val(x Value) (float32, bool)
	func Float64Val(x Value) (float64, bool)
	func Int64Val(x Value) (int64, bool)
	func StringVal(x Value) string
	func Uint64Val(x Value) (uint64, bool)
	func Bytes(x Value) []byte
	func BitLen(x Value) int
	func Sign(x Value) int


for performing arithmetic on values,


	// Operations
	func Compare(x Value, op token.Token, y Value) bool
	func UnaryOp(op token.Token, y Value, prec uint) Value
	func BinaryOp(x Value, op token.Token, y Value) Value
	func Shift(x Value, op token.Token, s uint) Value
	func Denom(x Value) Value
	func Num(x Value) Value
	func Real(x Value) Value
	func Imag(x Value) Value


and for constructing new values:


	// Constructors
	func MakeBool(b bool) Value
	func MakeFloat64(x float64) Value
	func MakeFromBytes(bytes []byte) Value
	func MakeFromLiteral(lit string, tok token.Token, prec uint) Value
	func MakeImag(x Value) Value
	func MakeInt64(x int64) Value
	func MakeString(s string) Value
	func MakeUint64(x uint64) Value
	func MakeUnknown() Value


All numeric `Value`s, whether integer or floating-point, signed or
unsigned, or real or complex, are represented more precisely than
ordinary Go types like `int64` and `float64`.
Internally, the `go/constant` package uses multi-precision data types
like `Int`, `Rat`, and `Float` from the `math/big` package so that
`Values` and their arithmetic operations are accurate to at least 256
bits, as required by the Go specification.


<!-- TODO example -->


# Size and Alignment


Because the calls `unsafe.Sizeof(v)`, `unsafe.Alignof(v)`,
and `unsafe.Offsetof(v.f)` are all constant expressions, the type
checker must be able to compute the memory layout of any value
`v`.



By default, the type checker uses the same layout algorithm as the Go
1.5 `gc` compiler targeting `amd64`.
Clients can configure the type checker to use a different algorithm by
providing an instance of the `types.Sizes` interface in the
`types.Config` struct:


	package types
	
	type Sizes interface {
		Alignof(T Type) int64
		Offsetsof(fields []*Var) []int64
		Sizeof(T Type) int64
	}



For common changes, like reducing the word size to 32 bits, clients
can use an instance of `StdSizes`:


	type StdSizes struct {
		WordSize int64
		MaxAlign int64
	}


This type has two basic size and alignment parameters from which it
derives all the other values using common assumptions.
For example, pointers, functions, maps, and channels fit in one word,
strings and interfaces require two words, and slices need three.
The default behaviour is equivalent to `StdSizes{8, 8}`.
For more esoteric layout changes, you'll need to write a new
implementation of the `Sizes` interface.



The `hugeparam` program below prints all function parameters and
results whose size exceeds a threshold.
By default, the threshold is 48 bytes, but you can set it via the
`-bytes` command-line flag.
Such a tool could help identify inefficient parameter passing in your
programs.


%include hugeparam/main.go


As before, `Inspect` applies a function to every node in the AST.
The function cares about two kinds of nodes: declarations of named
functions and methods (`*ast.FuncDecl`) and function literals
(`*ast.FuncLit`).
Observe the two cases' different logic to obtain the type of each
function.



Here's a typical invocation on the standard `encoding/xml` package.
It reports a number of places where the 7-word
[`StartElement` type](https://pkg.go.dev/encoding/xml#StartElement)
is copied.


%include hugeparam/main.go output -

# Imports


The type checker's `Check` function processes a slice of parsed
files (`[]*ast.File`) that make up one package.
When the type checker encounters an import declaration, it needs the
type information for the objects in the imported package.
It gets it by calling the `Import` method of the `Importer`
interface shown below, an instance of which must be provided by the
`Config`.
This separation of concerns relieves the type checker from having to
know any of the details of Go workspace organization, `GOPATH`,
compiler file formats, and so on.


	type Importer interface {
		Import(path string) (*Package, error)
	}


Most of our examples used the trivial `Importer` implementation,
`importer.Default()`, provided by the `go/importer` package.
This importer looks for `.a` files written by the compiler
that was used to build the program.
In addition to object code, these files contain _export data_,
that is, a description of all the objects declared by the package, and
also of any objects from other packages that were referred to indirectly.
Because export data includes information about dependencies, the type
checker need load at most one file per import, instead of one per
transitive dependency.



Compiler export data is compact and efficient to locate, load, and
parse, but it has some shortcomings.
First, it does not contain complete syntax trees nor semantic
information about the bodies of all functions, so it is not
suitable for interprocedural analyses.
Second, compiler object data may be stale.  Nothing detects or ensures
that the object files are more recent than the source files from which
they were derived.
Generally, object data for standard packages is likely to be
up-to-date, but for user packages, it depends on how recently the user
ran a `go install` or `go build -i` command.



The [`golang.org/tools/x/go/packages`
package](https://pkg.go.dev/golang.org/x/tools/go/packages) provides
a comprehensive means of loading packages from source.
It runs `go list` to query the project metadata,
performs [`cgo`](https://golang.org/cmd/cgo/cgo) preprocessing if necessary,
reads and parses the source files,
and optionally type-checks each package.
It can load a whole program from source, or load just the initial
packages from source and load all their dependencies from export data.
It loads independent packages in parallel to hide I/O latency, and
detects and reports import cycles.
For each package, it provides the `types.Package` containing the
package's lexical environment, the list of `ast.File` syntax
trees for each file in the package, the `types.Info` containing
type information for each syntax node, a list of type errors
associated with that package, and other information too.
Since some of this information is more costly to compute,
the API allows you to select which parts you need,
but since this is a tutorial we'll generally request complete
information so that it is easier to explore.

The `doc` program below demonstrates a simple use of `go/packages`.
It is a rudimentary implementation of `go doc` that prints the type,
methods, and documentation of the package-level object specified on
the command line.
Here's an example:


%include doc/main.go output -


Observe that it prints the correct location of each method
declaration, even though, due to embedding, some of
`http.File`'s methods were declared in another package.
Here's the first part of the program, showing how to load
complete type information including typed syntax,
for a single package `pkgpath`,
plus exported type information for its dependencies.


%include doc/main.go part1


By default, `go/packages`, instructs the parser to retain comments during parsing.
The rest of the program prints the output:


%include doc/main.go part2


We used `IntuitiveMethodSet` to compute the method set, instead
of `NewMethodSet`.
The result of this convenience function, which is intended for use in
user interfaces, includes methods of `*T` as well as those of
`T`, since that matches most users' intuition about the method
set of a type.
(Our example, `http.File`, didn't illustrate the difference, but  try
running it on a type with both value and pointer methods.)



Also notice `PathEnclosingInterval`, which finds the set of AST
nodes that enclose a particular point, in this case, the object's
declaring identifier.
By walking up with path, we find the enclosing declaration, to which
the documentation is attached.



# Formatting support


All types that satisfy `Type` or `Object` define a
`String` method that formats the type or object in a readable
notation.   `Selection` also provides a `String` method.
All package-level objects within these data structures are
printed with the complete package path, as in these examples:


	[]encoding/json.Marshaler                                     // a *Slice type
	encoding/json.Marshal                                         // a *Func object
	(*encoding/json.Encoder).Encode                               // a *Func object (method)
	func (enc *encoding/json.Encoder) Encode(v interface{}) error // a method *Signature
	func NewEncoder(w io.Writer) *encoding/json.Encoder           // a function *Signature


This notation is unambiguous, but it is not legal Go syntax.
Also, package paths may be long, and the same package path may appear
many times in a single string, for instance, when formatting a
function of several parameters.
Because these strings often form part of a tool's user interface---as
with the diagnostic messages of `hugeparam` or the code generated
by `skeleton`---many clients want more control over the
formatting of package names.



The `go/types` package provides these alternatives to the
`String` methods:


	func ObjectString(obj Object, qf Qualifier) string
	func TypeString(typ Type, qf Qualifier) string
	func SelectionString(s *Selection, qf Qualifier) string
	
	type Qualifier func(*Package) string


The `TypeString`, `ObjectString`, and `SelectionString`
functions are like the `String` methods of the respective types,
but they accept an additional argument, a `Qualifier`.



A `Qualifier` is a client-provided function that determines how a
package name is rendered as a string.
If it is nil, the default behavior is to print the package's
path, just like the `String` methods do.
If a caller passes `(*Package).Name` as the qualifier, that is, a
function that accepts a package and returns its `Name`, then
objects are qualified only by the package name.
The above examples would look like this:


	[]json.Marshaler
	json.Marshal
	(*json.Encoder).Encode
	func (enc *json.Encoder) Encode(v interface{}) error
	func NewEncoder(w io.Writer) *json.Encoder


Often when a tool prints some output, it is implicitly in the
context of a particular package, perhaps one specified by the
command line or HTTP request.
In that case, it is more natural to omit the package qualification
altogether for objects belonging to that package, but to qualify all
other objects by their package's path.
That's what the `RelativeTo(pkg)` qualifier does:


	func RelativeTo(pkg *Package) Qualifier


The examples below show how `json.NewEncoder` would be printed
using three qualifiers, each relative to a different package:


	// RelativeTo "encoding/json":
	func NewEncoder(w io.Writer) *Encoder
	
	// RelativeTo "io":
	func NewEncoder(w Writer) *encoding/json.Encoder
	
	// RelativeTo any other package:
	func NewEncoder(w io.Writer) *encoding/json.Encoder


Another qualifier that may be relevant to refactoring tools (but is
not currently provided by the type checker) is one that renders each
package name using the locally appropriate name within a given source
file.
Its behavior would depend on the set of import declarations, including
renaming imports, within that source file.



# Getting from A to B


The type checker and its related packages represent many aspects of a
Go program in many different ways, and analysis tools must often map
between them.
For instance, a named entity may be identified by its `Object`;
by its declaring identifier (`ast.Ident`) or by any referring
identifier; by its declaring `ast.Node`; by the position
(`token.Pos`) of any those nodes; or by the filename and
line/column number (or byte offset) of those `token.Pos` values.



In this section, we'll list solutions to a number of common problems
of the form "I have an A; I need the corresponding B".



To map **from a `token.Pos` to an `ast.Node`**, call the
helper function
[`astutil.PathEnclosingInterval`](https://pkg.go.dev/golang.org/x/tools/go/ast/astutil#PathEnclosingInterval).
It returns the enclosing `ast.Node`, and all its ancestors up to
the root of the file.
If you don't know which file `*ast.File` the `token.Pos` belongs to,
you can iterate over the parsed files of the package and quickly test
whether its position falls within the file's range,
from `File.FileStart` to `File.FileEnd`.


To map **from an `Object` to its declaring syntax**, call
`Pos` to get its position, then use `PathEnclosingInterval` as before.
This approach is suitable for a one-off query.  For repeated use, it
may be more efficient to visit the syntax tree and construct the
mapping between declarations and objects.



To map **from an `ast.Ident` to the `Object`** it refers to (or
declares), consult the `Uses` or `Defs` map for the
package, as shown in [Identifier Resolution](#identifier-resolution).



To map **from an `Object` to its documentation**, find the
object's declaration, and look at the attached `Doc` field.
You must have set the parser's `ParseComments` flag.
See the `doc` example in [Imports](#imports).
```

### gotypes/Makefile

```

all: README.md
	go build ./...

README.md: go-types.md ../internal/cmd/weave/*.go $(wildcard */*.go)
	go run ../internal/cmd/weave go-types.md >README.md

# This is for previewing using github.
# $HOME/markdown must be a github client.
test: README.md
	cp README.md $$HOME/markdown/
	(cd $$HOME/markdown/ && git commit -m . README.md && git push)

```

### gotypes/README.md

```markdown
<!-- Autogenerated by weave; DO NOT EDIT -->

# `go/types`: The Go Type Checker

This document is maintained by Alan Donovan `adonovan@google.com`.

[October 2015 GothamGo talk on go/types](https://docs.google.com/presentation/d/13OvHYozAUBeISPRoLgG7kMBuja1NsU1D_mMlmbaYojk/view)


# Contents

1. [Changes in Go 1.18](#changes-in-go-1.18)
1. [Introduction](#introduction)
1. [An Example](#an-example)
1. [Objects](#objects)
1. [Identifier Resolution](#identifier-resolution)
1. [Scopes](#scopes)
1. [Initialization Order](#initialization-order)
1. [Types](#types)
	1. [Basic types](#basic-types)
	1. [Simple Composite Types](#simple-composite-types)
	1. [Struct Types](#struct-types)
	1. [Tuple Types](#tuple-types)
	1. [Function and Method Types](#function-and-method-types)
	1. [Alias Types](#alias-types)
	1. [Named Types](#named-types)
	1. [Interface Types](#interface-types)
	1. [TypeParam types](#typeparam-types)
	1. [Union types](#union-types)
	1. [TypeAndValue](#typeandvalue)
1. [Selections](#selections)
1. [Ids](#ids)
1. [Method Sets](#method-sets)
1. [Constants](#constants)
1. [Size and Alignment](#size-and-alignment)
1. [Imports](#imports)
1. [Formatting support](#formatting-support)
1. [Getting from A to B](#getting-from-a-to-b)

# Changes in Go 1.18

Go 1.18 introduces generics, and several corresponding new APIs for `go/types`.
This document is not yet up-to-date for these changes, but a guide to the new
changes exists at
[`x/exp/typeparams/example`](https://github.com/golang/exp/tree/master/typeparams/example).

# Introduction


The [`go/types` package]('https://golang.org/pkg/go/types) is a
type-checker for Go programs, designed by Robert Griesemer.
It became part of Go's standard library in Go 1.5.
Measured by lines of code and by API surface area, it is one of the
most complex packages in Go's standard library, and using it requires
a firm grasp of the structure of Go programs.
This tutorial will help you find your bearings.
It comes with several example programs that you can obtain with `go get` and play with.
We assume you are a proficient Go programmer who wants to build tools
to analyze or manipulate Go programs and that you have some knowledge
of how a typical compiler works.

The type checker complements several existing
standard packages for analyzing Go programs.
We've listed them below.


	→   go/types
	    go/constant
	    go/parser
	    go/ast
	    go/scanner
	    go/token


Starting at the bottom, the
[`go/token` package](http://golang.org/pkg/go/token)
defines the lexical tokens of Go.
The [`go/scanner` package](http://golang.org/pkg/go/scanner) tokenizes an input stream and records
file position information for use in diagnostics
or for file surgery in a refactoring tool.
The [`go/ast` package](http://golang.org/pkg/go/ast)
defines the data types of the abstract syntax tree (AST).
The [`go/parser` package](http://golang.org/pkg/go/parser)
provides a robust recursive-descent parser that constructs the AST.
And [`go/constant`](http://golang.org/pkg/go/constant)
provides representations and arithmetic operations for the values of compile-time
constant expressions, as we'll see in
[Constants](#constants).



The [`golang.org/x/tools/go/packages` package](https://pkg.go.dev/golang.org/x/tools/go/packages)
from the `x/tools` repository is a client of the type
checker that loads, parses, and type-checks a complete Go program from
source code.
We use it in some of our examples and you may find it useful too.



The Go type checker does three main things.
First, for every name in the program, it determines which declaration
the name refers to; this is known as _identifier resolution_.
Second, for every expression in the program, it determines what type
that expression has, or reports an error if the expression has no
type, or has an inappropriate type for its context; this is known as
_type deduction_.
Third, for every constant expression in the program, it determines the
value of that constant; this is known as _constant evaluation_.
Superficially, it appears that these three processes could be done
sequentially, in the order above, but perhaps surprisingly, they must
be done together.
For example, the value of a constant may depend on the type of an
expression due to operators like `unsafe.Sizeof`.
Conversely, the type of an expression may depend on the value of a
constant, since array types contain constants.
As a result, type deduction and constant evaluation must be done
together.
As another example, we cannot resolve the identifier `k` in the composite
literal `T{k: 0}` until we know whether `T` is a struct type.
If it is, then `k` must be found among `T`'s fields.
If not, then `k` is an ordinary reference
to a constant or variable in the lexical environment.
Consequently, identifier resolution and type deduction are also
inseparable in the general case.



Nonetheless, the three processes of identifier resolution, type
deduction, and constant evaluation can be separated for the purpose of
explanation.


# An Example


The code below shows the most basic use of the type checker to check
the _hello, world_ program, supplied as a string.
Later examples will be variations on this one, and we'll often omit
boilerplate details such as parsing.
To check out and build the examples,
run `go get golang.org/x/example/gotypes/...`.


	// go get golang.org/x/example/gotypes/pkginfo

```go
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const hello = `package main

import "fmt"

func main() {
        fmt.Println("Hello, world")
}`

func main() {
	fset := token.NewFileSet()

	// Parse the input string, []byte, or io.Reader,
	// recording position information in fset.
	// ParseFile returns an *ast.File, a syntax tree.
	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		log.Fatal(err) // parse error
	}

	// A Config controls various options of the type checker.
	// The defaults work fine except for one setting:
	// we must specify how to deal with imports.
	conf := types.Config{Importer: importer.Default()}

	// Type-check the package containing only file f.
	// Check returns a *types.Package.
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	fmt.Printf("Package  %q\n", pkg.Path())
	fmt.Printf("Name:    %s\n", pkg.Name())
	fmt.Printf("Imports: %s\n", pkg.Imports())
	fmt.Printf("Scope:   %s\n", pkg.Scope())
}
```


First, the program creates a
[`token.FileSet`](http://golang.org/pkg/go/token/#FileSet).
To avoid the need to store file names and line and column
numbers in every node of the syntax tree, the `go/token` package
provides `FileSet`, a data structure that stores this information
compactly for a sequence of files.
A `FileSet` records each file name only once, and records
only the byte offsets of each newline, allowing a position within
any file to be identified using a small integer called a
`token.Pos`.
Many tools create a single `FileSet` at startup.
Any part of the program that needs to convert a `token.Pos` into
an intelligible location---as part of an error message, for
instance---must have access to the `FileSet`.



Second, the program parses the input string.
More realistic packages contain several source files, so the parsing
step must be repeated for each one, or better, done in parallel.
Third, it creates a `Config` that specifies type-checking options.
Since the _hello, world_ program uses imports, we must indicate
how to locate the imported packages.
Here we use `importer.Default()`, which loads compiler-generated
export data, but we'll explore alternatives in [Imports](#imports).



Fourth, the program calls `Check`.
This creates a `Package` whose path is `"cmd/hello"`, and
type-checks each of the specified files---just one in this example.
The final (nil) argument is a pointer to an optional `Info`
struct that returns additional deductions from the type checker; more
on that later.
`Check` returns a `Package` even when it also returns an error.
The type checker is robust to ill-formed input,
and goes to great lengths to report accurate
partial information even in the vicinity of syntax or type errors.
`Package` has this definition:


	type Package struct{ ... }
	func (*Package) Path() string
	func (*Package) Name() string
	func (*Package) Scope() *Scope
	func (*Package) Imports() []*Package


Finally, the program prints the attributes of the package, shown below.
(The hexadecimal number may vary from one run to the next.)


```go
$ go build golang.org/x/example/gotypes/pkginfo
$ ./pkginfo
Package  "cmd/hello"
Name:    main
Imports: [package fmt ("fmt")]
Scope:   package "cmd/hello" scope 0x820533590 {
.  func cmd/hello.main()
}
```


A package's `Path`, such as `"encoding/json"`, is the string
by which import declarations identify it.
It is unique within a workspace,
and for published packages it must be globally unique.


A package's `Name` is the identifier in the `package`
declaration of each source file within the package, such as `json`.
The type checker reports an error if not all the package declarations in
the package agree.
The package name determines how the package is known when it is
imported into a file (unless a renaming import is used),
but is otherwise not visible to a program.


`Scope` returns the package's [_lexical block_](#scopes),
which provides access to all the named entities or
[_objects_](#objects) declared at package level.
`Imports` returns the set of packages directly imported by this
one, and  may be useful for computing dependencies
(see [Initialization Order](#initialization-order)).



# Objects


The task of identifier resolution is to map every identifier in the
syntax tree, that is, every `ast.Ident`, to an object.
For our purposes, an _object_ is a named entity created by a
declaration, such as a `var`, `type`, or `func`
declaration.
(This is different from the everyday meaning of object in
object-oriented programming.)



Objects are represented by the `Object` interface:


	type Object interface {
	    Name() string   // package-local object name
	    Exported() bool // reports whether the name starts with a capital letter
	    Type() Type     // object type
	    Pos() token.Pos // position of object identifier in declaration
	
	    Parent() *Scope // scope in which this object is declared
	    Pkg() *Package  // nil for objects in the Universe scope and labels
	    Id() string     // object id (see Ids section below)
	}


The first four methods are straightforward; we'll explain the other
three later.
`Name` returns the object's name---an identifier.
`Exported` is a convenience method that reports whether the first
letter of `Name` is a capital, indicating that the object may be
visible from outside the package.
It's a shorthand for `ast.IsExported(obj.Name())`.
`Type` returns the object's type; we'll come back to that in
[Types](#types).



`Pos` returns the source position of the object's declaring identifier.
To make sense of a `token.Pos`, we need to call the
`(*token.FileSet).Position` method, which returns a struct with
individual fields for the file name, line number, column, and byte
offset, though usually we just call its `String` method:


	fmt.Println(fset.Position(obj.Pos())) // "hello.go:10:6"


Objects for predeclared functions and types such as `len` and `int`
do not have a valid (non-zero) position: `!obj.Pos().IsValid()`.


There are eight kinds of objects in the Go type checker.
Most familiar are the kinds that can be declared at package level:
constants, variables, functions, and types.
Less familiar are statement labels, imported package names
(such as `json` in a file containing an `import "encoding/json"`
declaration), built-in functions (such as `append` and
`len`), and the pre-declared `nil`.
The eight types shown below are the only concrete types that satisfy
the `Object` interface.
In other words, `Object` is a _discriminated union_ of 8
possible types, and we commonly use a type switch to distinguish them.


	Object = *Func         // function, concrete method, or abstract method
	       | *Var          // variable, parameter, result, or struct field
	       | *Const        // constant
	       | *TypeName     // type name
	       | *Label        // statement label
	       | *PkgName      // package name, e.g. json after import "encoding/json"
	       | *Builtin      // predeclared function such as append or len
	       | *Nil          // predeclared nil



Objects are canonical.
That is, two Objects `x` and `y` denote the same entity if and only if `x==y`.
(This is generally true but beware that parameterized types complicate matters; see
https://github.com/golang/exp/tree/master/typeparams/example for details.)

Object identity is significant, and objects are routinely compared by
the addresses of the underlying pointers.
A package-level object (func/var/const/type) can be uniquely
identified by its name and enclosing package.
The [`golang.org/x/tools/go/types/objectpath`](https://pkg.go.dev/golang.org/x/tools/go/types/objectpath)
package defines a naming scheme for objects that are
[exported](#imports) from their package or are unexported but form part of the
type of an exported object.
But for most objects, including all function-local objects,
there is no simple way to obtain a string that uniquely identifies it.



The `Parent` method returns the `Scope` (lexical block) in
which the object was declared; we'll come back to this in
[Scopes](#scopes).
Fields and methods are not found in the lexical environment, so
their objects have no `Parent`.
<!-- TODO check this -->



The `Pkg` method returns the `Package` to which this object
belongs, even for objects not declared at package level.
Only predeclared objects have no package.
The `Id` method will be explained in [Ids](#ids).



Not all methods make sense for each kind of object.  For instance,
the last four kinds above have no meaningful `Type` method.
And some kinds of objects have methods in addition to those required by the
`Object` interface:


	func (*Func) Scope() *Scope
	func (*Var) Anonymous() bool
	func (*Var) IsField() bool
	func (*Const) Val() constant.Value
	func (*TypeName) IsAlias() bool
	func (*PkgName) Imported() *Package


`(*Func).Scope` returns the [lexical block](#scopes)
containing the function's type parameters, parameters, results,
and other local declarations.
`(*Var).IsField` distinguishes struct fields from ordinary
variables, and `(*Var).Anonymous` discriminates named fields like
the one in `struct{T T}` from anonymous fields like the one in `struct{T}`.
`(*Const).Val` returns the value of a named [constant](#constants).


`(*TypeName).IsAlias` reports whether the type name declares an alias
for an existing type (as in `type I = int`), as opposed to defining a new
[`Named`](#named-types) type, as in `type Celsius float64`.
(Most `TypeName`s for which `IsAlias()` is true have a `Type()` of
type `*types.Alias`, but `IsAlias()` is also true for the predeclared
`byte` and `rune` types, which are aliases for `uint8` and `int32`.)


`(*PkgName).Imported` returns the package (for instance,
`encoding/json`) denoted by a given import name such as `json`.
Each time a package is imported, a new `PkgName` object is
created, usually with the same name as the `Package` it
denotes, but not always, as in the case of a renaming import.
`PkgName`s are objects, but `Package`s are not.
We'll look more closely at this in [Imports](#imports).



All relationships between the syntax trees (`ast.Node`s) and type
checker data structures such as `Object`s and `Type`s are
stored in mappings outside the syntax tree itself.
Be aware that the `go/ast` package also defines an older deprecated
type called `Object` that resembles---and predates---the type
checker's `Object`, and that `ast.Object`s are held directly by
identifiers in the AST.
They are created by the parser, which has a necessarily limited view
of the package, so the information they represent is at best partial and
in some cases wrong, as in the `T{k: 0}` example mentioned above.
If you are using the type checker, there is no reason to use the older
`ast.Object` mechanism.



# Identifier Resolution


Identifier resolution computes the relationship between
identifiers and objects.
Its results are recorded in the `Info` struct optionally passed
to `Check`.
The fields related to identifier resolution are shown below.


	type Info struct {
		Defs       map[*ast.Ident]Object
		Uses       map[*ast.Ident]Object
		Implicits  map[ast.Node]Object
		Selections map[*ast.SelectorExpr]*Selection
		Scopes     map[ast.Node]*Scope
		...
	}


Since not all facts computed by the type checker are needed by every
client, the API lets clients control which components of the result
should be recorded and which discarded: only fields that hold a
non-nil map will be populated during the call to `Check`.



The two fields of type `map[*ast.Ident]Object` are the most important:
`Defs` records _declaring_ identifiers and
`Uses` records _referring_ identifiers.
In the example below, the comments indicate which identifiers are of
which kind.


	var x int        // def of x, use of int
	fmt.Println(x)   // uses of fmt, Println, and x
	type T struct{U} // def of T, use of U (type), def of U (field)


The final line above illustrates why we don't combine `Defs` and
`Uses` into one map.
In the anonymous field declaration `struct{U}`, the
identifier `U` is both a use of the type `U` (a
`TypeName`) and a definition of the anonymous field (a
`Var`).



The function below prints the location of each referring and defining
identifier in the input program, and the object it refers to.


	// go get golang.org/x/example/gotypes/defsuses

```go
func PrintDefsUses(fset *token.FileSet, files ...*ast.File) error {
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}
	_, err := conf.Check("hello", fset, files, info)
	if err != nil {
		return err // type error
	}

	for id, obj := range info.Defs {
		fmt.Printf("%s: %q defines %v\n",
			fset.Position(id.Pos()), id.Name, obj)
	}
	for id, obj := range info.Uses {
		fmt.Printf("%s: %q uses %v\n",
			fset.Position(id.Pos()), id.Name, obj)
	}
	return nil
}
```


Let's use the _hello, world_ program again as the input:


	// go get golang.org/x/example/gotypes/hello

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```


This is what it prints:


```go
$ go build golang.org/x/example/gotypes/defsuses
$ ./defsuses
hello.go:1:9: "main" defines <nil>
hello.go:5:6: "main" defines func hello.main()
hello.go:6:9: "fmt" uses package fmt
hello.go:6:13: "Println" uses func fmt.Println(a ...interface{}) (n int, err error)
```


Notice that the `Defs` mapping may contain nil entries in a few
cases.
The first line of output reports that the package identifier
`main` is present in the `Defs` mapping, but has no
associated object.



The `Implicits` mapping handles two cases of the syntax in
which an `Object` is declared without an `ast.Ident`, namely type
switches and import declarations.
<!--
A third case: an anonymous function parameter or result variable.
These objects are returned by Signature.Param and Signature.Result.
-->
In the type switch below, which declares a local variable `y`,
the type of `y` is different in each case of the switch:


	switch y := x.(type) {
	case int:
		fmt.Printf("%d", y)
	case string:
		fmt.Printf("%q", y)
	default:
		fmt.Print(y)
	}


To represent this, for each single-type case, the type checker creates
a separate `Var` object for `y` with the appropriate type,
and `Implicits` maps each `ast.CaseClause` to the `Var`
for that case.
The `default` case, the `nil` case, and cases with more than one
type all use the regular `Var` object that is associated with the
identifier `y`, which is found in the `Defs` mapping.



The import declaration below defines the name `json` without an
`ast.Ident`:


	import "encoding/json"


`Implicits` maps this `ast.ImportSpec` to the `PkgName`
object named `json` that it implicitly declares.



The `Selections` mapping, of type
`map[*ast.SelectorExpr]*Selection`, records the meaning of each
expression of the form _`expr`_`.f`, where _`expr`_ is
an expression or type and `f` is the name of a field or method.
These expressions, called _selections_, are represented by
`ast.SelectorExpr` nodes in the AST.
We'll talk more about the `Selection` type in [Selections](#selections).



Not all `ast.SelectorExpr` nodes represent selections.
Expressions like `fmt.Println`, in which a package name precedes
the dot, are _qualified identifiers_.
They do not appear in the `Selections` mapping, but their
constituent identifiers (such as `fmt` and `Println`) both
appear in `Uses`.



Referring identifiers that are not part of an `ast.SelectorExpr`
are _lexical references_.
That is, they are resolved to an object by searching for the
innermost enclosing lexical declaration of that name.
We'll see how that search works in the next section.


# Scopes


The `Scope` type is a mapping from names to objects.


	type Scope struct{ ... }
	
	func (s *Scope) Names() []string
	func (s *Scope) Lookup(name string) Object


`Names` returns the set of names in the mapping, in sorted order.
(It is not a simple accessor though, so call it sparingly.)
The `Lookup ` method returns the object for a given name, so we
can print all the entries or _bindings_ in a scope like this:


	for _, name := range scope.Names() {
		fmt.Println(scope.Lookup(name))
	}


The _scope_ of a declaration of a name is the region of
program source in which a reference to the name resolves to that
declaration.  That is, scope is a property of a declaration.
However, in the `go/types` API, the `Scope` type represents
a _lexical block_, which is one component of the lexical
environment.
Consider the _hello, world_ program again:


	package main
	
	import "fmt"
	
	func main() {
		const message = "hello, world"
		fmt.Println(message)
	}


There are four lexical blocks in this program.
The outermost one is the _universe block_, which maps the
pre-declared names like `int`, `true`, and `append` to
their objects---a `TypeName`, a `Const`, and a
`Builtin`, respectively.
The universe block is represented by the global variable
`Universe`, of type `*Scope`, although it's logically a
constant so you shouldn't modify it.



Next is the _package block_, which maps `"main"` to the
`main` function.
Following that is the _file block_, which maps `"fmt"` to
the `PkgName` object for this import of the `fmt` package.
And finally, the innermost block is that of function `main`, a
local block, which contains the declaration of `message`, a `Const`.
The `main` function is trivial, but many functions contain
several blocks since each `if`, `for`, `switch`,
`case`, or `select` statement creates at least one
additional block.
Local blocks nest to arbitrary depths.



The structure of the lexical environment thus forms a tree, with the
universe block at the root, the package blocks beneath it, the file
blocks beneath them, and then any number of local blocks beneath the
files.
We can access and navigate this tree structure with the following
methods of `Scope`:


	func (s *Scope) Parent() *Scope
	func (s *Scope) NumChildren() int
	func (s *Scope) Child(i int) *Scope


`Parent` lets us walk up the tree, and `Child`
lets us walk down it.
Note that although the `Parent` of every package `Scope` is
`Universe`, `Universe` has no children.
This asymmetry is a consequence of using a global variable to hold
`Universe`.



To obtain the universe block, we use the `Universe` global variable.
To obtain the lexical block of a `Package`, we call its
`Scope` method.
To obtain the scope of a file (`*ast.File`), or any smaller piece
of syntax such as an `*ast.IfStmt`, we consult the `Scopes`
mapping in the `Info` struct, which maps each block-creating
syntax node to its block.
The lexical block of a named function or method can also be obtained
by calling its `(*Func).Scope` method.


<!--
TODO: explain explicit and implicit blocks
TODO: explain Dot imports.
TODO: explain that Func blocks are associated with FuncType (not FuncDecl or FuncLit)
-->


To look up a name in the lexical environment, we must search the tree
of lexical blocks, starting at a particular `Scope` and walking
up to the root until a declaration of the name is found.
For convenience, the `LookupParent` method does this, returning
not just the object, if found, but also the `Scope` in which it was
declared, which may be an ancestor of the initial one:


	func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)


The `pos` parameter determines the position in the source code at
which the name should be resolved.
The effective lexical environment is different at each point in the
block because it depends on which local declarations appear
before or after that point.
(We'll see an illustration in a moment.)



`Scope` has several other methods relating to source positions:


	func (s *Scope) Pos() token.Pos
	func (s *Scope) End() token.Pos
	func (s *Scope) Contains(pos token.Pos) bool
	func (s *Scope) Innermost(pos token.Pos) *Scope


`Pos` and `End` report the `Scope`'s start and end
position which, for explicit blocks, coincide with its curly
braces.
`Contains` is a convenience method that reports whether a
position lies in this interval.
`Innermost` returns the innermost scope containing the specified
position, which may be a child or other descendent of the initial
scope.



These features are useful for tools that wish to resolve names or
evaluate constant expressions as if they had appeared at a particular
point within the program.
The next example program finds all the comments in the input,
treating the contents of each one as a name.  It looks up each name in
the environment at the position of the comment, and prints what it
finds.
Observe that the `ParseComments` flag directs the parser to
preserve comments in the input.


	// go get golang.org/x/example/gotypes/lookup

```go
func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.ParseComments)
	if err != nil {
		log.Fatal(err) // parse error
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}

	// Each comment contains a name.
	// Look up that name in the innermost scope enclosing the comment.
	for _, comment := range f.Comments {
		pos := comment.Pos()
		name := strings.TrimSpace(comment.Text())
		fmt.Printf("At %s,\t%q = ", fset.Position(pos), name)
		inner := pkg.Scope().Innermost(pos)
		if _, obj := inner.LookupParent(name, pos); obj != nil {
			fmt.Println(obj)
		} else {
			fmt.Println("not found")
		}
	}
}
```


The expression `pkg.Scope().Innermost(pos)` finds the innermost
`Scope` that encloses the comment, and `LookupParent(name, pos)`
does a name lookup at a specific position in that lexical block.



A typical input is shown below.
The first comment causes a lookup of `"append"` in the file block.
The second comment looks up `"fmt"` in the `main` function's block,
and so on.


```go
const hello = `
package main

import "fmt"

// append
func main() {
        // fmt
        fmt.Println("Hello, world")
        // main
        main, x := 1, 2
        // main
        print(main, x)
        // x
}
// x
`
```


Here's the output:


```go
$ go build golang.org/x/example/gotypes/lookup
$ ./lookup
At hello.go:6:1,        "append" = builtin append
At hello.go:8:9,        "fmt" = package fmt
At hello.go:10:9,       "main" = func cmd/hello.main()
At hello.go:12:9,       "main" = var main int
At hello.go:14:9,       "x" = var x int
At hello.go:16:1,       "x" = not found
```


Notice how the two lookups of `main` return different results,
even though they occur in the same block, because one precedes the
declaration of the local variable named `main` and the other
follows it.
Also notice that there are two lookups of the name `x` but only
the first one, in the function block, succeeds.



Download the program and modify both the input program and
the set of comments to get a better feel for how name resolution works.



The table below summarizes which kinds of objects may be declared at
each level of the tree of lexical blocks.


                Universe File Package Local
    Builtin        ✔
    Nil            ✔
    Const          ✔            ✔      ✔
    TypeName       ✔            ✔      ✔
    Func                        ✔
    Var                         ✔      ✔
    PkgName               ✔
    Label                              ✔


# Initialization Order


In the course of identifier resolution, the type checker constructs a
graph of references among declarations of package-level variables and
functions.
The type checker reports an error if the initializer expression for a
variable refers to that variable, whether directly or indirectly.



The reference graph determines the initialization order of the
package-level variables, as required by the Go spec, using a
breadth-first algorithm.
First, variables in the graph with no successors are removed, sorted
into the order in which they appear in the source code, then added
to a list.  This creates more variables that have no successors.
The process repeats until they have all been removed.



The result is available in the `InitOrder` field of the
`Info` struct, whose type is `[]Initializer`.


	type Info struct {
		...
		InitOrder []Initializer
		...
	}
	
	type Initializer struct {
		Lhs []*Var // var Lhs = Rhs
		Rhs ast.Expr
	}


Each element of the list represents a single initializer expression
that must be executed, and the variables to which it is assigned.
The variables may number zero, one, or more, as in these examples:


	var _ io.Writer = new(bytes.Buffer)
	var rx = regexp.MustCompile("^b(an)*a$")
	var cwd, cwdErr = os.Getwd()


This process governs the initialization order of variables within a
package.
Across packages, dependencies must be initialized first, although the
order among them is not specified.
That is, any topological order of the import graph will do.
The `(*Package).Imports` method returns the set of direct
dependencies of a package.


# Types


The main job of the type checker is, of course, to deduce the type
of each expression and to report type errors.
Like `Object`, `Type` is an interface type used as a
discriminated union of several concrete types but, unlike
`Object`, `Type` has very few methods because types have
little in common with each other.
Here is the interface:


	type Type interface {
		Underlying() Type
	}


And here are the 14 concrete types that satisfy it:


	Type = *Basic
	     | *Pointer
	     | *Array
	     | *Slice
	     | *Map
	     | *Chan
	     | *Struct
	     | *Tuple
	     | *Signature
	     | *Alias
	     | *Named
	     | *Interface
	     | *Union
	     | *TypeParam


With the exception of `Named` types, instances of `Type` are
not canonical.
(Even for `Named` types, parameterized types complicate matters; see
https://github.com/golang/exp/tree/master/typeparams/example.)
That is, it is usually a mistake to compare types using `t1==t2`
since this equivalence is not the same as the
[type identity relation](https://golang.org/ref/spec#Type_identity)
defined by the Go spec.
Use this function instead:


	func Identical(t1, t2 Type) bool


For the same reason, you should not use a `Type` as a key in a map.
The [`golang.org/x/tools/go/types/typeutil` package](https://pkg.go.dev/golang.org/x/tools/go/types/typeutil)
provides a map keyed by types that uses the correct
equivalence relation.


The Go spec defines three relations over types.
[_Assignability_](https://golang.org/ref/spec#Assignability)
governs which pairs of types may appear on the
left- and right-hand side of an assignment, including implicit
assignments such as function calls, map and channel operations, and so
on.
[_Comparability_](https://golang.org/ref/spec#Comparison_operators)
determines which types may appear in a comparison `x==y` or a
switch case or may be used as a map key.
[_Convertibility_](https://golang.org/ref/spec#Conversions)
governs which pairs of types are allowed in a conversion operation
`T(v)`.
You can query these relations with the following predicate functions:


	func AssignableTo(V, T Type) bool
	func Comparable(T Type) bool
	func ConvertibleTo(V, T Type) bool


Let's take a look at each kind of type.


## Basic types


`Basic` represents all types that are not composed from simpler
types.
This is essentially the set of underlying types that a constant expression is
permitted to have--strings, booleans, and numbers---but it also
includes `unsafe.Pointer` and untyped nil.


	type Basic struct{...}
	func (*Basic) Kind() BasicKind
	func (*Basic) Name() string
	func (*Basic) Info() BasicInfo


The `Kind` method returns an "enum" value that indicates which
basic type this is.
The kinds `Bool`, `String`, `Int16`, and so on,
represent the corresponding predeclared boolean, string, or numeric
types.
There are two aliases: `Byte` is an alias for `Uint8`
and `Rune` is an alias for `Int32`.
The kind `UnsafePointer` represents `unsafe.Pointer`.
The kinds `UntypedBool`, `UntypedInt` and so on represent
the six kinds of "untyped" constant types: boolean, integer, rune,
float, complex, and string.
The kind `UntypedNil` represents the type of the predeclared
`nil` value.
And the kind `Invalid` indicates the invalid type, which is used
for expressions containing errors, or for objects without types, like
`Label`, `Builtin`, or `PkgName`.



The `Name` method returns the name of the type, such as
`"float64"`, and the `Info` method returns a bitfield that
encodes information about the type, such as whether it is signed or
unsigned, integer or floating point, or real or complex.



`Typ` is a table of canonical basic types, indexed by
kind, so `Typ[String]` returns the `*Basic` that represents
`string`, for instance.
Like `Universe`, `Typ` is logically a constant, so don't
modify it.



A few minor subtleties:

- According to the Go spec, pre-declared types such as `int` are
  named types for the purposes of assignability, even though the type
  checker does not represent them using `Named`.

- `unsafe.Pointer` is a pointer type for the purpose of
  determining whether the receiver type of a method is legal, even
  though the type checker does not represent it using `Pointer`.

- The "untyped" types are usually only ascribed to constant expressions,
  but there is one exception.
  A comparison `x==y` has type "untyped bool", so the result of
  this expression may be assigned to a variable of type `bool` or
  any other named boolean type.


## Simple Composite Types


The types `Pointer`, `Array`, `Slice`, `Map`,
and `Chan` are pretty self-explanatory.
All have an `Elem` method that returns the element type `T`
for a pointer `*T`, an array `[n]T`, a slice `[]T`, a
map `map[K]T`, or a channel `chan T`.
This should feel familiar if you've used the `reflect.Value` API.



In addition, the `*Map`, `*Chan`, and `*Array` types
have accessor methods that return their key type, direction, and
length, respectively:


	func (*Map) Key() Type
	func (*Chan) Dir() ChanDir      // = Send | Recv | SendRecv
	func (*Array) Len() int64



## Struct Types


A struct type has an ordered list of fields and a corresponding
ordered list of field tags.


	type Struct struct{ ... } 
	func (*Struct) NumFields() int
	func (*Struct) Field(i int) *Var
	func (*Struct) Tag(i int) string


Each field is a `Var` object whose `IsField` method returns true.
Field objects have no `Parent` scope, because they are
resolved through selections, not through the lexical environment.
<!-- TODO check this -->



Thanks to embedding, the expression `new(S).f` may be a shorthand
for a longer expression such as `new(S).d.e.f`, but in the
representation of `Struct` types, these field selection
operations are explicit.
That is, the set of fields of struct type `S` does not include `f`.
An anonymous field is represented like a regular field, but its
`Anonymous` method returns true.



One subtlety is relevant to tools that generate documentation.
When analyzing a declaration such as this,


	type T struct{x int}


it may be tempting to consider the `Var` object for field `x` as if it
had the name `"T.x"`, but beware: field objects do not have
canonical names and there is no way to obtain the name `"T"`
from the `Var` for `x`.
That's because several types may have the same underlying struct type,
as in this code:


	type T struct{x int}
	type U T


Here, the `Var` for field `x` belongs equally to `T`
and to `U`, and short of inspecting source positions or walking
the AST---neither of which is possible for objects loaded from compiler
export data---it is not possible to ascertain that `x` was declared as
part of `T`.
The type checker builds the exact same data structures given this input:


	type T U
	type U struct{x int}


A similar issue applies to the methods of named interface types.


## Tuple Types


Like a struct, a tuple type has an ordered list of fields, and fields
may be named.


	type Tuple struct{ ... }
	func (*Tuple) Len() int
	func (*Tuple) At(i int) *Var


Although tuples are not the type of any variable in Go, they are
the type of some expressions, such as the right-hand sides of these
assignments:


	v, ok = m[key]
	v, ok = <-ch
	v, ok = x.(T)
	f, err = os.Open(filename)


Tuples also represent the types of the parameter list and the result
list of a function, as we will see.



Since empty tuples are common, the nil `*Tuple` pointer is a valid empty tuple.



## Function and Method Types


The types of functions and methods are represented by a `Signature`,
which has a tuple of parameter types and a tuple of result types.


	type Signature struct{ ... }
	func (*Signature) Recv() *Var
	func (*Signature) Params() *Tuple
	func (*Signature) Results() *Tuple
	func (*Signature) Variadic() bool


Variadic functions such as `fmt.Println` have the `Variadic`
flag set.
The final parameter of such functions is always a slice, or in the
special case of certain calls to `append`, a string.



A `Signature` for a method, whether concrete or abstract, has a
non-nil receiver parameter, `Recv`.
The type of the receiver is usually a named type or a pointer to a named type, 
but it may be an unnamed struct or interface type in some cases.
Method types are rather second-class: they are only used for the
`Func` objects created by method declarations, and no Go
expression has a method type.
When printing a method type, the receiver does not appear, and the
`Identical` predicate ignores the receiver.



The types of `Builtin` objects like `append` cannot be
expressed as a `Signature` since those types require parametric
polymorphism.
`Builtin` objects are thus ascribed the `Invalid` basic type.
However, the type of each _call_ to a built-in function has a specific
and expressible Go type.
These types are recorded during type checking for later use
([TypeAndValue](#typeandvalue)).



## Alias Types

Type declarations come in two forms, aliases and defined types.

Aliases, though introduced only in Go 1.9 and not very common, are
simplest, so we'll present them first and explain defined types in
the next section ("Named Types").

An alias type declaration declares an alternative name for an existing
type. For example, this declaration lets you use the type `Dictionary`
as a synonym for `map[string]string`:

	type Dictionary = map[string]string

The declaration creates a `TypeName` object for `Dictionary`.
The object's `IsAlias` method returns true,
and its `Type` method returns an `Alias`:

        type Alias struct{ ... }
	func (a *Alias) Obj() *TypeName
	func (a *Alias) Origin() *Alias
	func (a *Alias) Rhs() Type
	func (a *Alias) SetTypeParams(tparams []*TypeParam)
	func (a *Alias) TypeArgs() *TypeList
	func (a *Alias) TypeParams() *TypeParamList

The type on the right-hand side of an alias declaration,
such as `map[string]string` in the example above,
can be accessed using the `Rhs()` method.
The `types.Unalias(t)` helper function recursively applies `Rhs`,
removing all `Alias` types from the operand t and returning the
outermost non-alias type.

The `Obj` method returns the declaring `TypeName` object, such as
`Dictionary`; it provides the name, position, and other properties of
the declaration. Conversely, the `TypeName` object's `Type` method
returns the `Alias` type.

Starting with Go 1.24, alias types may have type parameters.
For example, this declaration creates an Alias type with
a type parameter:

	type Set[T comparable] = map[T]bool

Each instantiation such as `Set[string]` is identical to the
corresponding instantiation of the alias' right-hand side type, such
as `map[string]bool`.

The remaining methods--Origin, SetTypeParams, TypeArgs,
TypeParams--are all concerned with type parameters. For now, see
https://github.com/golang/exp/tree/master/typeparams/example.

Prior to Go 1.22, aliases were not materialized as `Alias` types:
each reference to an alias type such as `Dictionary` would be
immediately replaced by its right-hand side type, leaving no
indication in the output of the type checker that an alias was
present.
By materializing alias types, optionally in Go 1.22 and by default
starting in Go 1.23, we can more faithfully record the structure of
the program, which improves the quality of diagnostic messages and
enables certain analyses and code transformations. And, crucially, it
enabled the addition of parameterized aliases in Go 1.24.)



## Named Types



The second form of type declaration, and the only kind prior to Go
1.9, does not use an equals sign:

	type Celsius float64

This declaration does more than just give a name to a type.
It first defines a new `Named` type
whose underlying type is `float64`; this `Named` type is different
from any other type, including `float64`.  The declaration binds the
`TypeName` object to the `Named` type.

Since Go 1.9, the Go language specification has used the term _defined
types_ instead of named types;
the essential property of a defined type is not that it has a name
(aliases and type parameters also have names)
but that it is a distinct type with its own method set.
However, the type checker API predates that
change and instead calls defined types `Named` types.

	type Named struct{ ... }
	func (*Named) NumMethods() int
	func (*Named) Method(i int) *Func
	func (*Named) Obj() *TypeName
	func (*Named) Underlying() Type

The `Named` type's `Obj` method returns the `TypeName` object, which
provides the name, position, and other properties of the declaration.
Conversely, the `TypeName` object's `Type` method returns the `Named` type.

A `Named` type may appear as the receiver type in a method declaration.
Methods are associated with the `Named` type, not the name (the
`TypeName` object); it's possible---though cryptic---to declare a
method on a `Named` type using one of its aliases.
The `NumMethods` and `Method` methods enumerate the declared
methods associated with this `Named` type (or a pointer to it),
in the order they were declared.
However, due to the subtleties of anonymous fields and the difference
between value and pointer receivers, a named type may have more or fewer
methods than this list.  We'll return to this in [Method Sets](#method-sets).



Every `Type` has an `Underlying` method, but for all of them
except `*Named` and `*Alias`, it is simply the identity function.
For a named or alias type, `Underlying` returns its underlying type, which
is always an unnamed type.
Thus `Underlying` returns `int` for both `T` and
`U` below.


	type T int
	type U T


Clients of the type checker often use type assertions or type switches
with a `Type` operand.
When doing so, it is often necessary to switch on the type that
underlies the type of interest, and failure to do so may be a bug.

This is a common pattern:


	// handle types of composite literal
	switch u := t.Underlying().(type) {  // remove any *Named and *Alias types
	case *Struct:        // ...
	case *Map:           // ...
	case *Array, *Slice: // ...
	default:
		panic("impossible")
	}

## Interface Types


Interface types are represented by `Interface`.


	type Interface struct{ ... }
	func (*Interface) Empty() bool
	func (*Interface) NumMethods() int
	func (*Interface) Method(i int) *Func
	func (*Interface) NumEmbeddeds() int
	func (*Interface) Embedded(i int) *Named
	func (*Interface) NumExplicitMethods() int
	func (*Interface) ExplicitMethod(i int) *Func


Syntactically, an interface type has a list of explicitly declared
methods (`ExplicitMethod`), and a list of embedded named
interface types (`Embedded`), but many clients care only about
the complete set of methods, which can be enumerated via
`Method`.
All three lists are ordered by name.
Since the empty interface is an important special case, the
`Empty` predicate provides a shorthand for `NumMethods() ==
0`.



As with the fields of structs (see above), the methods of interfaces
may belong equally to more than one interface type.
The `Func` object for method `f` in the code below is shared
by `I` and `J`:


	type I interface { f() }
	type J I


Because the difference between interface (abstract) and
non-interface (concrete) types is so important in Go, the
`IsInterface` predicate is provided for convenience.


	func IsInterface(Type) bool


The type checker provides three utility methods relating to interface
satisfaction:


	func Implements(V Type, T *Interface) bool
	func AssertableTo(V *Interface, T Type) bool
	func MissingMethod(V Type, T *Interface, static bool) (method *Func, wrongType bool)


The `Implements` predicate reports whether a type satisfies an
interface type.
`MissingMethod` is like `Implements`, but instead of
returning false, it explains why a type does not satisfy the
interface, for use in diagnostics.



`AssertableTo` reports whether a type assertion `v.(T)` is legal.
If `T` is a concrete type that doesn't have all the methods of
interface `v`, then the type assertion is not legal, as in this example:


	// error: io.Writer is not assertible to int
	func f(w io.Writer) int { return w.(int) }




## TypeParam types


A `TypeParam` is the type of a type parameter.
For example, the type of the variable `x` in the `identity` function
below is a `TypeParam`:

    func identity[T any](x T) T { return x }

As with `Alias` and `Named` types, each `TypeParam` has an associated
`TypeName` object that provides its name, position, and other
properties of the declaration.

See https://github.com/golang/exp/tree/master/typeparams/example
for a more thorough treatment of parameterized types.




## Union types

A `Union` is the type of type-parameter constraint of the form `func
f[T int | string]`.

See https://github.com/golang/exp/tree/master/typeparams/example
for a more thorough treatment of parameterized types.



## TypeAndValue


The type checker records the type of each expression in another field
of the `Info` struct, namely `Types`:


	type Info struct {
		...
		Types map[ast.Expr]TypeAndValue
	}


No entries are recorded for identifiers since the `Defs` and
`Uses` maps provide more information about them.
Also, no entries are recorded for pseudo-expressions like
`*ast.KeyValuePair` or `*ast.Ellipsis`.



The value of the `Types` map is a `TypeAndValue`, which
(unsurprisingly) holds the type and value of the expression, and in
addition, its _mode_.
The mode is opaque, but has predicates to answer questions such as:
Does this expression denote a value or a type?  Does this value have an
address?  Does this expression appear on the left-hand side of an
assignment?  Does this expression appear in a context that expects two
results?
The comments in the code below give examples of expressions that
satisfy each predicate.


	type TypeAndValue struct {
		Type  Type
		Value constant.Value // for constant expressions only
		...
	}
	
	func (TypeAndValue) IsVoid() bool      // e.g. "main()"
	func (TypeAndValue) IsType() bool      // e.g. "*os.File"
	func (TypeAndValue) IsBuiltin() bool   // e.g. "len(x)"
	func (TypeAndValue) IsValue() bool     // e.g. "*os.Stdout"
	func (TypeAndValue) IsNil() bool       // e.g. "nil"
	func (TypeAndValue) Addressable() bool // e.g. "a[i]" but not "f()", "m[key]"
	func (TypeAndValue) Assignable() bool  // e.g. "a[i]", "m[key]"
	func (TypeAndValue) HasOk() bool       // e.g. "<-ch", "m[key]"


The statement below inspects every expression within the AST of a single
type-checked file and prints its type, value, and mode:


	// go get golang.org/x/example/gotypes/typeandvalue

```go
// f is a parsed, type-checked *ast.File.
ast.Inspect(f, func(n ast.Node) bool {
	if expr, ok := n.(ast.Expr); ok {
		if tv, ok := info.Types[expr]; ok {
			fmt.Printf("%-24s\tmode:  %s\n", nodeString(expr), mode(tv))
			fmt.Printf("\t\t\t\ttype:  %v\n", tv.Type)
			if tv.Value != nil {
				fmt.Printf("\t\t\t\tvalue: %v\n", tv.Value)
			}
		}
	}
	return true
})
```


It makes use of these two helper functions, which are not shown:


	// nodeString formats a syntax tree in the style of gofmt.
	func nodeString(n ast.Node) string
	
	// mode returns a string describing the mode of an expression.
	func mode(tv types.TypeAndValue) string


Given this input:


```go
const input = `
package main

var m = make(map[string]int)

func main() {
	v, ok := m["hello, " + "world"]
	print(rune(v), ok)
}
`
```


the program prints:


```go
$ go build golang.org/x/example/gotypes/typeandvalue
$ ./typeandvalue
make(map[string]int)            mode:  value
                                type:  map[string]int
make                            mode:  builtin
                                type:  func(map[string]int) map[string]int
map[string]int                  mode:  type
                                type:  map[string]int
string                          mode:  type
                                type:  string
int                             mode:  type
                                type:  int
m["hello, "+"world"]            mode:  value,assignable,ok
                                type:  (int, bool)
m                               mode:  value,addressable,assignable
                                type:  map[string]int
"hello, " + "world"             mode:  value
                                type:  string
                                value: "hello, world"
"hello, "                       mode:  value
                                type:  untyped string
                                value: "hello, "
"world"                         mode:  value
                                type:  untyped string
                                value: "world"
print(rune(v), ok)              mode:  void
                                type:  ()
print                           mode:  builtin
                                type:  func(rune, bool)
rune(v)                         mode:  value
                                type:  rune
rune                            mode:  type
                                type:  rune
...more not shown...
```


Notice that the identifiers for the built-ins `make` and
`print` have types that are specific to the particular calls in
which they appear.
Also notice `m["hello"]` has a 2-tuple type `(int, bool)`
and that it is assignable, but unlike the variable `m`, it is not
addressable.



Download the example and vary the inputs and see what the program prints.



Here's another example, adapted from the `govet` static checking tool.
It checks for accidental uses of a method value `x.f` when a
call `x.f()` was intended;
comparing a method `x.f` against nil is a common mistake.


	// go get golang.org/x/example/gotypes/nilfunc

```go
// CheckNilFuncComparison reports unintended comparisons
// of functions against nil, e.g., "if x.Method == nil {".
func CheckNilFuncComparison(info *types.Info, n ast.Node) {
	e, ok := n.(*ast.BinaryExpr)
	if !ok {
		return // not a binary operation
	}
	if e.Op != token.EQL && e.Op != token.NEQ {
		return // not a comparison
	}

	// If this is a comparison against nil, find the other operand.
	var other ast.Expr
	if info.Types[e.X].IsNil() {
		other = e.Y
	} else if info.Types[e.Y].IsNil() {
		other = e.X
	} else {
		return // not a comparison against nil
	}

	// Find the object.
	var obj types.Object
	switch v := other.(type) {
	case *ast.Ident:
		obj = info.Uses[v]
	case *ast.SelectorExpr:
		obj = info.Uses[v.Sel]
	default:
		return // not an identifier or selection
	}

	if _, ok := obj.(*types.Func); !ok {
		return // not a function or method
	}

	fmt.Printf("%s: comparison of function %v %v nil is always %v\n",
		fset.Position(e.Pos()), obj.Name(), e.Op, e.Op == token.NEQ)
}
```


Given this input,


```go
const input = `package main

import "bytes"

func main() {
	var buf bytes.Buffer
	if buf.Bytes == nil && bytes.Repeat != nil && main == nil {
		// ...
	}
}
`
```


the program reports these errors:


```go
$ go build golang.org/x/example/gotypes/nilfunc
$ ./nilfunc
input.go:7:5: comparison of function Bytes == nil is always false
input.go:7:25: comparison of function Repeat != nil is always true
input.go:7:48: comparison of function main == nil is always false
```

# Selections


A _selection_ is an expression _`expr`_`.f` in which
`f` denotes either a struct field or a method.
A selection is resolved not by looking for a name in the lexical
environment, but by looking within a _type_.
The type checker ascertains the meaning of each selection in the
package---a surprisingly tricky business---and records it in the
`Selections` mapping of the `Info` struct, whose values are
of type `Selection`:


	type Selection struct{ ... }
	func (s *Selection) Kind() SelectionKind // = FieldVal | MethodVal | MethodExpr
	func (s *Selection) Recv() Type
	func (s *Selection) Obj() Object
	func (s *Selection) Type() Type
	func (s *Selection) Index() []int
	func (s *Selection) Indirect() bool


The `Kind` method discriminates between the three (legal) kinds
of selections, as indicated by the comments below.


	type T struct{Field int}
	func (T) Method() {}
	var v T
	
                         // Kind            Type
        var _ = v.Field  // FieldVal        int
        var _ = v.Method // MethodVal       func()
        var _ = T.Method // MethodExpr      func(T)


Because of embedding, a selection may denote more than one field or
method, in which case it is ambiguous, and no `Selection` is
recorded for it.



The `Obj` method returns the `Object` for the selected field
(`*Var`) or method (`*Func`).
Due to embedding, the object may belong to a different type than that
of the receiver expression _`expr`_.
The `Type` method returns the type of the selection.  For a field
selection, this is the type of the field, but for method selections,
the result is a function type that is not the same as the type of the
method.
For a `MethodVal`, the receiver parameter is dropped, and
for a `MethodExpr`, the receiver parameter becomes a regular
parameter, as shown in the example above.



The `Index` and `Indirect` methods report information about
implicit operations occurring during the selection that a compiler
would need to know about.
Because of embedding, a selection _`expr`_`.f` may be
shorthand for a sequence containing several implicit field selections,
_`expr`_`.d.e.f`, and `Index` reports the complete
sequence.
And because of automatic pointer dereferencing during struct field
accesses and method calls, a selection may imply one or more indirect
loads from memory; `Indirect` reports whether this occurs.



Clients of the type checker can call `LookupFieldOrMethod` to
look up a name within a type, as if by a selection.
This function has an intimidating signature, but conceptually it
accepts just a `Type` and a name, and returns a `Selection`:


	func LookupFieldOrMethod(T Type, addressable bool, pkg *Package, name string) \
	    (obj Object, index []int, indirect bool)


The result is not actually a `Selection`, but it contains the
three main components of one: `Obj`, `Index`,
and `Indirect`.



The `addressable` flag should be set if the receiver is a
_variable_ of type `T`, since in a method selection on a
variable, an implicit address-of operation (`&`) may occur.
The flag indicates whether the methods of type `*T` should be
considered during the lookup.
(You may wonder why this parameter is necessary.  Couldn't clients
instead call `LookupFieldOrMethod` on the pointer type `*T`
if the receiver is a `T` variable?  The answer is that if
`T` is an interface type, the type `*T` has no methods at
all.)



The final two parameters of `LookupFieldOrMethod` are `(pkg
*Package, name string)`.
Together they specify the name of the field or method to look up.
This brings us to `Id`s.



# Ids


`LookupFieldOrMethod`'s need for a `Package` parameter
is a subtle consequence of the
[_Uniqueness of identifiers_](https://golang.org/ref/spec#Uniqueness_of_identifiers)
section in the Go spec: "Two
identifiers are different if they are spelled differently, or if they
appear in different packages and are not exported."
In practical terms, this means that a type may have two methods
(or two fields, or one of each) both named `f` so long as those
methods are defined in different packages, as in this example:


	package a
	type A int
	func (A) f()
	
	package b
	type B int
	func (B) f()
	
	package c
	import ( "a"; "b" )
	type C struct{a.A; b.B} // C has two methods called f


The type `c.C` has two methods named `f`, but there is
no ambiguity because the two `f`s are distinct
identifiers---think of them as `fᵃ` and `fᵇ`.
For an exported method, this situation _would_ be ambiguous
because there is no distinction between `Fᵃ` and `Fᵇ`; there
is only `F`.



Despite having two methods called `f`, neither of them can be
called from within package `c` because `c` has no way to
identify them.
Within `c`,  `f` is the identifier `fᶜ`, and
type `C` has no method of that name.
But if we pass an instance of `C` to code in package `a`
and call its `f` method via an interface, `fᵃ` is called.



The practical consequence for tool builders is that any time you need
to look up a field or method by name, or construct a map of fields and/or
methods keyed by name, it is not sufficient to use the object's name
as a key.
Instead, you must call the `Object.Id` method, which returns
a string that incorporates the object name, and for unexported
objects, the package path too.
There is also a standalone function `Id` that combines a name and
the package path in the same way:


	func Id(pkg *Package, name string) string


This distinction applies to selections _`expr`_`.f`, but not
to lexical references `x` because for unexported identifiers,
declarations and references always appear in the same package.



Fun fact: the `reflect.StructField` type records both the
`Name` and the `PkgPath` strings for the same reason.
The `FieldByName` methods of `reflect.Value` and
`reflect.Type` match field names without regard to the package.
If there is more than one match, they return an invalid value.




# Method Sets


The _method set_ of a type is the set of methods that can be
called on any value of that type.
(A variable of type `T` has access to all the methods of type
`*T` as well, due to the implicit address-of operation during
method calls, but those extra methods are not part of the method set
of `T`.)



Clients can request the method set of a type `T` by calling
`NewMethodSet(T)`:


	type MethodSet struct{ ... }
	func NewMethodSet(T Type) *MethodSet
	func (s *MethodSet) Len() int
	func (s *MethodSet) At(i int) *Selection
	func (s *MethodSet) Lookup(pkg *Package, name string) *Selection


The `Len` and `At` methods access a list of
`Selections`, all of kind `MethodVal`, ordered by `Id`.
The `Lookup` function allows lookup of a single method by
name (and package path, as explained in the previous section).



`NewMethodSet` can be expensive, so for applications that compute
method sets repeatedly, `golang.org/x/tools/go/types/typeutil`
provides a `MethodSetCache` type that records previous results.
If you only need a single method, don't construct the
`MethodSet` at all; it's cheaper to use
`LookupFieldOrMethod`.



The next program generates a boilerplate
declaration of a new concrete type that satisfies an existing
interface.
Here's an example:


```go
$ ./skeleton io ReadWriteCloser buffer
// *buffer implements io.ReadWriteCloser.
type buffer struct{}
func (b *buffer) Close() error {
	panic("unimplemented")
}
func (b *buffer) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}
func (b *buffer) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}
```


The three arguments are the package and the name of the existing
interface type, and the name of the new concrete type.
The `main` function (not shown) loads the specified package and
calls `PrintSkeleton` with the remaining two arguments:


	// go get golang.org/x/example/gotypes/skeleton

```go
func PrintSkeleton(pkg *types.Package, ifacename, concname string) error {
	obj := pkg.Scope().Lookup(ifacename)
	if obj == nil {
		return fmt.Errorf("%s.%s not found", pkg.Path(), ifacename)
	}
	if _, ok := obj.(*types.TypeName); !ok {
		return fmt.Errorf("%v is not a named type", obj)
	}
	iface, ok := obj.Type().Underlying().(*types.Interface)
	if !ok {
		return fmt.Errorf("type %v is a %T, not an interface",
			obj, obj.Type().Underlying())
	}
	// Use first letter of type name as receiver parameter.
	if !isValidIdentifier(concname) {
		return fmt.Errorf("invalid concrete type name: %q", concname)
	}
	r, _ := utf8.DecodeRuneInString(concname)

	fmt.Printf("// *%s implements %s.%s.\n", concname, pkg.Path(), ifacename)
	fmt.Printf("type %s struct{}\n", concname)
	mset := types.NewMethodSet(iface)
	for i := 0; i < mset.Len(); i++ {
		meth := mset.At(i).Obj()
		sig := types.TypeString(meth.Type(), (*types.Package).Name)
		fmt.Printf("func (%c *%s) %s%s {\n\tpanic(\"unimplemented\")\n}\n",
			r, concname, meth.Name(),
			strings.TrimPrefix(sig, "func"))
	}
	return nil
}
```


First, `PrintSkeleton` locates the package-level named interface
type, handling various error cases.
Then it chooses the name for the receiver of the new methods: the
first letter of the concrete type.
Finally, it iterates over the method set of the interface, printing
the corresponding concrete method declarations.



There's a subtlety in the declaration of `sig`, which is the
string form of the method signature.
We could have obtained this string from `meth.Type().String()`,
but this would cause any named types within it to be formatted with
the complete package path, for instance
`net/http.ResponseWriter`, which is informative in diagnostics
but not legal Go syntax.
The `TypeString` function (explained in [Formatting Values](#formatting-values)) allows the
caller to control how packages are printed.
Passing `(*types.Package).Name` causes only the package name
`http` to be printed, not the complete path.
Here's another example that illustrates it:


```go
$ ./skeleton net/http Handler myHandler
// *myHandler implements net/http.Handler.
type myHandler struct{}
func (m *myHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}
```


The following program inspects all pairs of package-level named types
in `pkg`, and reports the types that satisfy each interface type.


	// go get golang.org/x/example/gotypes/implements

```go
// Find all named types at package level.
var allNamed []*types.Named
for _, name := range pkg.Scope().Names() {
	if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
		allNamed = append(allNamed, obj.Type().(*types.Named))
	}
}

// Test assignability of all distinct pairs of
// named types (T, U) where U is an interface.
for _, T := range allNamed {
	for _, U := range allNamed {
		if T == U || !types.IsInterface(U) {
			continue
		}
		if types.AssignableTo(T, U) {
			fmt.Printf("%s satisfies %s\n", T, U)
		} else if !types.IsInterface(T) &&
			types.AssignableTo(types.NewPointer(T), U) {
			fmt.Printf("%s satisfies %s\n", types.NewPointer(T), U)
		}
	}
}
```


Given this input,


	// go get golang.org/x/example/gotypes/implements

```go
const input = `package main

type A struct{}
func (*A) f()

type B int
func (B) f()
func (*B) g()

type I interface { f() }
type J interface { g() }
`
```


the program prints:


```go
$ go build golang.org/x/example/gotypes/implements
$ ./implements
*hello.A satisfies hello.I
hello.B satisfies hello.I
*hello.B satisfies hello.J
```


Notice that the method set of `B` does not include `g`, but
the method set of `*B` does.
That's why we needed the second assignability check, using the pointer
type `types.NewPointer(T)`.



# Constants


A constant expression is one whose value is guaranteed to be computed at
compile time.
Constant expressions may appear in types, specifically as the length
of an array type such as `[16]byte`, so one of the jobs of the
type checker is to compute the value of each constant expression.



As we saw in the `typeandvalue` example, the type checker records
the value of each constant expression like `"Hello, " + "world"`,
storing it in the `Value` field of the `TypeAndValue` struct.
Constants are represented using the `Value` interface from the
`go/constant` package.


	package constant // go/constant
	
	type Value interface {
		Kind() Kind
	}
	
	type Kind int // one of Unknown, Bool, String, Int, Float, Complex


The interface has only one method, for discriminating the various
kinds of constants, but the package provides many functions for
inspecting a value of a known kind,


	// Accessors
	func BoolVal(x Value) bool
	func Float32Val(x Value) (float32, bool)
	func Float64Val(x Value) (float64, bool)
	func Int64Val(x Value) (int64, bool)
	func StringVal(x Value) string
	func Uint64Val(x Value) (uint64, bool)
	func Bytes(x Value) []byte
	func BitLen(x Value) int
	func Sign(x Value) int


for performing arithmetic on values,


	// Operations
	func Compare(x Value, op token.Token, y Value) bool
	func UnaryOp(op token.Token, y Value, prec uint) Value
	func BinaryOp(x Value, op token.Token, y Value) Value
	func Shift(x Value, op token.Token, s uint) Value
	func Denom(x Value) Value
	func Num(x Value) Value
	func Real(x Value) Value
	func Imag(x Value) Value


and for constructing new values:


	// Constructors
	func MakeBool(b bool) Value
	func MakeFloat64(x float64) Value
	func MakeFromBytes(bytes []byte) Value
	func MakeFromLiteral(lit string, tok token.Token, prec uint) Value
	func MakeImag(x Value) Value
	func MakeInt64(x int64) Value
	func MakeString(s string) Value
	func MakeUint64(x uint64) Value
	func MakeUnknown() Value


All numeric `Value`s, whether integer or floating-point, signed or
unsigned, or real or complex, are represented more precisely than
ordinary Go types like `int64` and `float64`.
Internally, the `go/constant` package uses multi-precision data types
like `Int`, `Rat`, and `Float` from the `math/big` package so that
`Values` and their arithmetic operations are accurate to at least 256
bits, as required by the Go specification.


<!-- TODO example -->


# Size and Alignment


Because the calls `unsafe.Sizeof(v)`, `unsafe.Alignof(v)`,
and `unsafe.Offsetof(v.f)` are all constant expressions, the type
checker must be able to compute the memory layout of any value
`v`.



By default, the type checker uses the same layout algorithm as the Go
1.5 `gc` compiler targeting `amd64`.
Clients can configure the type checker to use a different algorithm by
providing an instance of the `types.Sizes` interface in the
`types.Config` struct:


	package types
	
	type Sizes interface {
		Alignof(T Type) int64
		Offsetsof(fields []*Var) []int64
		Sizeof(T Type) int64
	}



For common changes, like reducing the word size to 32 bits, clients
can use an instance of `StdSizes`:


	type StdSizes struct {
		WordSize int64
		MaxAlign int64
	}


This type has two basic size and alignment parameters from which it
derives all the other values using common assumptions.
For example, pointers, functions, maps, and channels fit in one word,
strings and interfaces require two words, and slices need three.
The default behaviour is equivalent to `StdSizes{8, 8}`.
For more esoteric layout changes, you'll need to write a new
implementation of the `Sizes` interface.



The `hugeparam` program below prints all function parameters and
results whose size exceeds a threshold.
By default, the threshold is 48 bytes, but you can set it via the
`-bytes` command-line flag.
Such a tool could help identify inefficient parameter passing in your
programs.


	// go get golang.org/x/example/gotypes/hugeparam

```go
var bytesFlag = flag.Int("bytes", 48, "maximum parameter size in bytes")

func PrintHugeParams(fset *token.FileSet, info *types.Info, sizes types.Sizes, files []*ast.File) {
	checkTuple := func(descr string, tuple *types.Tuple) {
		for i := 0; i < tuple.Len(); i++ {
			v := tuple.At(i)
			if sz := sizes.Sizeof(v.Type()); sz > int64(*bytesFlag) {
				fmt.Printf("%s: %q %s: %s = %d bytes\n",
					fset.Position(v.Pos()),
					v.Name(), descr, v.Type(), sz)
			}
		}
	}
	checkSig := func(sig *types.Signature) {
		checkTuple("parameter", sig.Params())
		checkTuple("result", sig.Results())
	}
	for _, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncDecl:
				checkSig(info.Defs[n.Name].Type().(*types.Signature))
			case *ast.FuncLit:
				checkSig(info.Types[n.Type].Type.(*types.Signature))
			}
			return true
		})
	}
}
```


As before, `Inspect` applies a function to every node in the AST.
The function cares about two kinds of nodes: declarations of named
functions and methods (`*ast.FuncDecl`) and function literals
(`*ast.FuncLit`).
Observe the two cases' different logic to obtain the type of each
function.



Here's a typical invocation on the standard `encoding/xml` package.
It reports a number of places where the 7-word
[`StartElement` type](https://pkg.go.dev/encoding/xml#StartElement)
is copied.


```go
% ./hugeparam encoding/xml
/go/src/encoding/xml/marshal.go:167:50: "start" parameter: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/marshal.go:734:97: "" result: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/marshal.go:761:51: "start" parameter: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/marshal.go:781:68: "start" parameter: encoding/xml.StartElement = 56 bytes
/go/src/encoding/xml/xml.go:72:30: "" result: encoding/xml.StartElement = 56 bytes
```

# Imports


The type checker's `Check` function processes a slice of parsed
files (`[]*ast.File`) that make up one package.
When the type checker encounters an import declaration, it needs the
type information for the objects in the imported package.
It gets it by calling the `Import` method of the `Importer`
interface shown below, an instance of which must be provided by the
`Config`.
This separation of concerns relieves the type checker from having to
know any of the details of Go workspace organization, `GOPATH`,
compiler file formats, and so on.


	type Importer interface {
		Import(path string) (*Package, error)
	}


Most of our examples used the trivial `Importer` implementation,
`importer.Default()`, provided by the `go/importer` package.
This importer looks for `.a` files written by the compiler
that was used to build the program.
In addition to object code, these files contain _export data_,
that is, a description of all the objects declared by the package, and
also of any objects from other packages that were referred to indirectly.
Because export data includes information about dependencies, the type
checker need load at most one file per import, instead of one per
transitive dependency.



Compiler export data is compact and efficient to locate, load, and
parse, but it has some shortcomings.
First, it does not contain complete syntax trees nor semantic
information about the bodies of all functions, so it is not
suitable for interprocedural analyses.
Second, compiler object data may be stale.  Nothing detects or ensures
that the object files are more recent than the source files from which
they were derived.
Generally, object data for standard packages is likely to be
up-to-date, but for user packages, it depends on how recently the user
ran a `go install` or `go build -i` command.



The [`golang.org/tools/x/go/packages`
package](https://pkg.go.dev/golang.org/x/tools/go/packages) provides
a comprehensive means of loading packages from source.
It runs `go list` to query the project metadata,
performs [`cgo`](https://golang.org/cmd/cgo/cgo) preprocessing if necessary,
reads and parses the source files,
and optionally type-checks each package.
It can load a whole program from source, or load just the initial
packages from source and load all their dependencies from export data.
It loads independent packages in parallel to hide I/O latency, and
detects and reports import cycles.
For each package, it provides the `types.Package` containing the
package's lexical environment, the list of `ast.File` syntax
trees for each file in the package, the `types.Info` containing
type information for each syntax node, a list of type errors
associated with that package, and other information too.
Since some of this information is more costly to compute,
the API allows you to select which parts you need,
but since this is a tutorial we'll generally request complete
information so that it is easier to explore.

The `doc` program below demonstrates a simple use of `go/packages`.
It is a rudimentary implementation of `go doc` that prints the type,
methods, and documentation of the package-level object specified on
the command line.
Here's an example:


```go
$ ./doc net/http File
type net/http.File interface{Readdir(count int) ([]os.FileInfo, error); Seek(offset int64, whence int) (int64, error); Stat() (os.FileInfo, error); io.Closer; io.Reader}
$GOROOT/src/io/io.go:92:2: method (net/http.File) Close() error
$GOROOT/src/io/io.go:71:2: method (net/http.File) Read(p []byte) (n int, err error)
/go/src/net/http/fs.go:65:2: method (net/http.File) Readdir(count int) ([]os.FileInfo, error)
$GOROOT/src/net/http/fs.go:66:2: method (net/http.File) Seek(offset int64, whence int) (int64, error)
/go/src/net/http/fs.go:67:2: method (net/http.File) Stat() (os.FileInfo, error)

 A File is returned by a FileSystem's Open method and can be
served by the FileServer implementation.

The methods should behave the same as those on an *os.File.
```


Observe that it prints the correct location of each method
declaration, even though, due to embedding, some of
`http.File`'s methods were declared in another package.
Here's the first part of the program, showing how to load
complete type information including typed syntax,
for a single package `pkgpath`,
plus exported type information for its dependencies.


	// go get golang.org/x/example/gotypes/doc

```go
pkgpath, name := os.Args[1], os.Args[2]

// Load complete type information for the specified packages,
// along with type-annotated syntax.
// Types for dependencies are loaded from export data.
conf := &packages.Config{Mode: packages.LoadSyntax}
pkgs, err := packages.Load(conf, pkgpath)
if err != nil {
	log.Fatal(err) // failed to load anything
}
if packages.PrintErrors(pkgs) > 0 {
	os.Exit(1) // some packages contained errors
}

// Find the package and package-level object.
pkg := pkgs[0]
obj := pkg.Types.Scope().Lookup(name)
if obj == nil {
	log.Fatalf("%s.%s not found", pkg.Types.Path(), name)
}
```


By default, `go/packages`, instructs the parser to retain comments during parsing.
The rest of the program prints the output:


	// go get golang.org/x/example/gotypes/doc

```go
// Print the object and its methods (incl. location of definition).
fmt.Println(obj)
for _, sel := range typeutil.IntuitiveMethodSet(obj.Type(), nil) {
	fmt.Printf("%s: %s\n", pkg.Fset.Position(sel.Obj().Pos()), sel)
}

// Find the path from the root of the AST to the object's position.
// Walk up to the enclosing ast.Decl for the doc comment.
for _, file := range pkg.Syntax {
	pos := obj.Pos()
	if !(file.FileStart <= pos && pos < file.FileEnd) {
		continue // not in this file
	}
	path, _ := astutil.PathEnclosingInterval(file, pos, pos)
	for _, n := range path {
		switch n := n.(type) {
		case *ast.GenDecl:
			fmt.Println("\n", n.Doc.Text())
			return
		case *ast.FuncDecl:
			fmt.Println("\n", n.Doc.Text())
			return
		}
	}
}
```


We used `IntuitiveMethodSet` to compute the method set, instead
of `NewMethodSet`.
The result of this convenience function, which is intended for use in
user interfaces, includes methods of `*T` as well as those of
`T`, since that matches most users' intuition about the method
set of a type.
(Our example, `http.File`, didn't illustrate the difference, but  try
running it on a type with both value and pointer methods.)



Also notice `PathEnclosingInterval`, which finds the set of AST
nodes that enclose a particular point, in this case, the object's
declaring identifier.
By walking up with path, we find the enclosing declaration, to which
the documentation is attached.



# Formatting support


All types that satisfy `Type` or `Object` define a
`String` method that formats the type or object in a readable
notation.   `Selection` also provides a `String` method.
All package-level objects within these data structures are
printed with the complete package path, as in these examples:


	[]encoding/json.Marshaler                                     // a *Slice type
	encoding/json.Marshal                                         // a *Func object
	(*encoding/json.Encoder).Encode                               // a *Func object (method)
	func (enc *encoding/json.Encoder) Encode(v interface{}) error // a method *Signature
	func NewEncoder(w io.Writer) *encoding/json.Encoder           // a function *Signature


This notation is unambiguous, but it is not legal Go syntax.
Also, package paths may be long, and the same package path may appear
many times in a single string, for instance, when formatting a
function of several parameters.
Because these strings often form part of a tool's user interface---as
with the diagnostic messages of `hugeparam` or the code generated
by `skeleton`---many clients want more control over the
formatting of package names.



The `go/types` package provides these alternatives to the
`String` methods:


	func ObjectString(obj Object, qf Qualifier) string
	func TypeString(typ Type, qf Qualifier) string
	func SelectionString(s *Selection, qf Qualifier) string
	
	type Qualifier func(*Package) string


The `TypeString`, `ObjectString`, and `SelectionString`
functions are like the `String` methods of the respective types,
but they accept an additional argument, a `Qualifier`.



A `Qualifier` is a client-provided function that determines how a
package name is rendered as a string.
If it is nil, the default behavior is to print the package's
path, just like the `String` methods do.
If a caller passes `(*Package).Name` as the qualifier, that is, a
function that accepts a package and returns its `Name`, then
objects are qualified only by the package name.
The above examples would look like this:


	[]json.Marshaler
	json.Marshal
	(*json.Encoder).Encode
	func (enc *json.Encoder) Encode(v interface{}) error
	func NewEncoder(w io.Writer) *json.Encoder


Often when a tool prints some output, it is implicitly in the
context of a particular package, perhaps one specified by the
command line or HTTP request.
In that case, it is more natural to omit the package qualification
altogether for objects belonging to that package, but to qualify all
other objects by their package's path.
That's what the `RelativeTo(pkg)` qualifier does:


	func RelativeTo(pkg *Package) Qualifier


The examples below show how `json.NewEncoder` would be printed
using three qualifiers, each relative to a different package:


	// RelativeTo "encoding/json":
	func NewEncoder(w io.Writer) *Encoder
	
	// RelativeTo "io":
	func NewEncoder(w Writer) *encoding/json.Encoder
	
	// RelativeTo any other package:
	func NewEncoder(w io.Writer) *encoding/json.Encoder


Another qualifier that may be relevant to refactoring tools (but is
not currently provided by the type checker) is one that renders each
package name using the locally appropriate name within a given source
file.
Its behavior would depend on the set of import declarations, including
renaming imports, within that source file.



# Getting from A to B


The type checker and its related packages represent many aspects of a
Go program in many different ways, and analysis tools must often map
between them.
For instance, a named entity may be identified by its `Object`;
by its declaring identifier (`ast.Ident`) or by any referring
identifier; by its declaring `ast.Node`; by the position
(`token.Pos`) of any those nodes; or by the filename and
line/column number (or byte offset) of those `token.Pos` values.



In this section, we'll list solutions to a number of common problems
of the form "I have an A; I need the corresponding B".



To map **from a `token.Pos` to an `ast.Node`**, call the
helper function
[`astutil.PathEnclosingInterval`](https://pkg.go.dev/golang.org/x/tools/go/ast/astutil#PathEnclosingInterval).
It returns the enclosing `ast.Node`, and all its ancestors up to
the root of the file.
If you don't know which file `*ast.File` the `token.Pos` belongs to,
you can iterate over the parsed files of the package and quickly test
whether its position falls within the file's range,
from `File.FileStart` to `File.FileEnd`.


To map **from an `Object` to its declaring syntax**, call
`Pos` to get its position, then use `PathEnclosingInterval` as before.
This approach is suitable for a one-off query.  For repeated use, it
may be more efficient to visit the syntax tree and construct the
mapping between declarations and objects.



To map **from an `ast.Ident` to the `Object`** it refers to (or
declares), consult the `Uses` or `Defs` map for the
package, as shown in [Identifier Resolution](#identifier-resolution).



To map **from an `Object` to its documentation**, find the
object's declaration, and look at the attached `Doc` field.
You must have set the parser's `ParseComments` flag.
See the `doc` example in [Imports](#imports).
```

### hello/reverse/example\_test.go

```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse_test

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func ExampleString() {
	fmt.Println(reverse.String("hello"))
	// Output: olleh
}
```

### hello/reverse/reverse.go

```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reverse can reverse things, particularly strings.
package reverse

// String returns its argument string reversed rune-wise left to right.
func String(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

### hello/reverse/reverse\_test.go

```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse

import "testing"

func TestString(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := String(c.in)
		if got != c.want {
			t.Errorf("String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

### hello/go.mod

```
module golang.org/x/example/hello

go 1.19

```

### hello/hello.go

```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a hello, world program, demonstrating
// how to write a simple command-line program.
//
// Usage:
//
//	hello [options] [name]
//
// The options are:
//
//	-g greeting
//		Greet with the given greeting, instead of "Hello".
//
//	-r
//		Greet in reverse.
//
// By default, hello greets the world.
// If a name is specified, hello greets that name instead.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/example/hello/reverse"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: hello [options] [name]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	greeting    = flag.String("g", "Hello", "Greet with `greeting`")
	reverseFlag = flag.Bool("r", false, "Greet in reverse")
)

func main() {
	// Configure logging for a command-line program.
	log.SetFlags(0)
	log.SetPrefix("hello: ")

	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments.
	name := "world"
	args := flag.Args()
	if len(args) >= 2 {
		usage()
	}
	if len(args) >= 1 {
		name = args[0]
	}
	if name == "" { // hello '' is an error
		log.Fatalf("invalid name %q", name)
	}

	// Run actual logic.
	if *reverseFlag {
		fmt.Printf("%s, %s!\n", reverse.String(*greeting), reverse.String(name))
		return
	}
	fmt.Printf("%s, %s!\n", *greeting, name)
}
```

### helloserver/go.mod

```
module golang.org/x/example/helloserver

go 1.19

```

### helloserver/server.go

```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
//
// See golang.org/x/example/outyet for a more sophisticated server.
package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	greeting = flag.String("g", "Hello", "Greet with `greeting`")
	addr     = flag.String("addr", "localhost:8080", "address to serve")
)

func main() {
	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments (none).
	args := flag.Args()
	if len(args) != 0 {
		usage()
	}

	// Register handlers.
	// All requests not otherwise mapped with go to greet.
	// /version is mapped specifically to version.
	http.HandleFunc("/", greet)
	http.HandleFunc("/version", version)

	log.Printf("serving http://%s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher"
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "%s, %s!\n", *greeting, html.EscapeString(name))
}
```

### internal/cmd/weave/weave.go

```go
// The weave command is a simple preprocessor for markdown files.
// It builds a table of contents and processes %include directives.
//
// Example usage:
//
//	$ go run internal/cmd/weave go-types.md > README.md
//
// The weave command copies lines of the input file to standard output, with two
// exceptions:
//
// If a line begins with "%toc", it is replaced with a table of contents
// consisting of links to the top two levels of headers below the %toc symbol.
//
// If a line begins with "%include FILENAME TAG", it is replaced with the lines
// of the file between lines containing "!+TAG" and  "!-TAG". TAG can be omitted,
// in which case the delimiters are simply "!+" and "!-".
//
// Before the included lines, a line of the form
//
//	// go get PACKAGE
//
// is output, where PACKAGE is constructed from the module path, the
// base name of the current directory, and the directory of FILENAME.
// This caption can be suppressed by putting "-" as the final word of the %include line.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var output = flag.String("o", "", "output file (empty means stdout)")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: weave [flags] <input.md>\n\nflags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	log.SetFlags(0)
	log.SetPrefix("weave: ")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	curDir := filepath.Base(wd)

	in, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out := os.Stdout
	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := out.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}

	printf := func(format string, args ...any) {
		if _, err := fmt.Fprintf(out, format, args...); err != nil {
			log.Fatalf("writing failed: %v", err)
		}
	}

	printf("<!-- Autogenerated by weave; DO NOT EDIT -->\n")

	// Pass 1: extract table of contents.
	type tocEntry struct {
		depth  int
		text   string
		anchor string
	}
	var (
		toc []tocEntry
		// We indent toc items according to their header depth, so that nested
		// headers result in nested lists. However, we want the lowest header depth
		// to correspond to the root of the list. Otherwise, the entire list is
		// indented, which turns it into a code block rather than an outline.
		minTocDepth int
	)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || (line[0] != '#' && line[0] != '%') {
			continue
		}
		line = strings.TrimSpace(line)
		if line == "%toc" {
			toc = nil
			minTocDepth = 0
		} else if strings.HasPrefix(line, "# ") ||
			strings.HasPrefix(line, "## ") ||
			strings.HasPrefix(line, "### ") ||
			strings.HasPrefix(line, "#### ") {

			words := strings.Fields(line)
			depth := len(words[0])
			if minTocDepth == 0 || depth < minTocDepth {
				minTocDepth = depth
			}
			words = words[1:]
			text := strings.Join(words, " ")
			anchor := strings.Join(words, "-")
			anchor = strings.ToLower(anchor)
			anchor = strings.ReplaceAll(anchor, "**", "")
			anchor = strings.ReplaceAll(anchor, "`", "")
			anchor = strings.ReplaceAll(anchor, "_", "")
			toc = append(toc, tocEntry{depth: depth, text: text, anchor: anchor})
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Pass 2.
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("can't rewind input: %v", err)
	}
	scanner = bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "%toc"): // ToC
			for _, h := range toc {
				// Only print two levels of headings.
				if h.depth-minTocDepth <= 1 {
					printf("%s1. [%s](#%s)\n", strings.Repeat("\t", h.depth-minTocDepth), h.text, h.anchor)
				}
			}
		case strings.HasPrefix(line, "%include"):
			words := strings.Fields(line)
			var section string
			caption := true
			switch len(words) {
			case 2: // %include filename
			// Nothing to do.
			case 3: // %include filename section OR %include filename -
				if words[2] == "-" {
					caption = false
				} else {
					section = words[2]
				}
			case 4: // %include filename section -
				section = words[2]
				if words[3] != "-" {
					log.Fatalf("last word is not '-': %s", line)
				}
				caption = false
			default:
				log.Fatalf("wrong # words (want 2-4): %s", line)
			}
			filename := words[1]

			if caption {
				printf("	// go get golang.org/x/example/%s/%s\n\n",
					curDir, filepath.Dir(filename))
			}

			s, err := include(filename, section)
			if err != nil {
				log.Fatal(err)
			}
			printf("```go\n")
			printf("%s\n", cleanListing(s)) // TODO(adonovan): escape /^```/ in s
			printf("```\n")
		default:
			printf("%s\n", line)
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

// include processes an included file, and returns the included text.
// Only lines between those matching !+tag and !-tag will be returned.
// This is true even if tag=="".
func include(file, tag string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	startre, err := regexp.Compile("!\\+" + tag + "$")
	if err != nil {
		return "", err
	}
	endre, err := regexp.Compile("!\\-" + tag + "$")
	if err != nil {
		return "", err
	}

	var text bytes.Buffer
	in := bufio.NewScanner(f)
	var on bool
	for in.Scan() {
		line := in.Text()
		switch {
		case startre.MatchString(line):
			on = true
		case endre.MatchString(line):
			on = false
		case on:
			text.WriteByte('\t')
			text.WriteString(line)
			text.WriteByte('\n')
		}
	}
	if in.Err() != nil {
		return "", in.Err()
	}
	if text.Len() == 0 {
		return "", fmt.Errorf("no lines of %s matched tag %q", file, tag)
	}
	return text.String(), nil
}

// cleanListing removes entirely blank leading and trailing lines from
// text, and removes n leading tabs.
func cleanListing(text string) string {
	lines := strings.Split(text, "\n")

	// remove minimum number of leading tabs from all non-blank lines
	tabs := 999
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			lines[i] = ""
		} else {
			if n := leadingTabs(line); n < tabs {
				tabs = n
			}
		}
	}
	for i, line := range lines {
		if line != "" {
			line := line[tabs:]
			lines[i] = line // remove leading tabs
		}
	}

	// remove leading blank lines
	for len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}
	// remove trailing blank lines
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return strings.Join(lines, "\n")
}

func leadingTabs(s string) int {
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] != '\t' {
			break
		}
	}
	return i
}
```

### ragserver/ragserver/go.mod

```
module golang.org/x/example/ragserver/ragserver

go 1.23.0

require (
	github.com/google/generative-ai-go v0.17.0
	github.com/weaviate/weaviate v1.26.1
	github.com/weaviate/weaviate-go-client/v4 v4.15.1
	google.golang.org/api v0.194.0
)

require (
	cloud.google.com/go v0.115.1 // indirect
	cloud.google.com/go/ai v0.8.0 // indirect
	cloud.google.com/go/auth v0.9.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.4 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	cloud.google.com/go/longrunning v0.5.7 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.21.2 // indirect
	github.com/go-openapi/errors v0.22.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/strfmt v0.23.0 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-openapi/validate v0.21.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.mongodb.org/mongo-driver v1.14.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240725223205-93522f1f2a9f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

### ragserver/ragserver/go.sum

```
cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
cloud.google.com/go v0.115.1 h1:Jo0SM9cQnSkYfp44+v+NQXHpcHqlnRJk2qxh6yvxxxQ=
cloud.google.com/go v0.115.1/go.mod h1:DuujITeaufu3gL68/lOFIirVNJwQeyf5UXyi+Wbgknc=
cloud.google.com/go/ai v0.8.0 h1:rXUEz8Wp2OlrM8r1bfmpF2+VKqc1VJpafE3HgzRnD/w=
cloud.google.com/go/ai v0.8.0/go.mod h1:t3Dfk4cM61sytiggo2UyGsDVW3RF1qGZaUKDrZFyqkE=
cloud.google.com/go/auth v0.9.1 h1:+pMtLEV2k0AXKvs/tGZojuj6QaioxfUjOpMsG5Gtx+w=
cloud.google.com/go/auth v0.9.1/go.mod h1:Sw8ocT5mhhXxFklyhT12Eiy0ed6tTrPMCJjSI8KhYLk=
cloud.google.com/go/auth/oauth2adapt v0.2.4 h1:0GWE/FUsXhf6C+jAkWgYm7X9tK8cuEIfy19DBn6B6bY=
cloud.google.com/go/auth/oauth2adapt v0.2.4/go.mod h1:jC/jOpwFP6JBxhB3P5Rr0a9HLMC/Pe3eaL4NmdvqPtc=
cloud.google.com/go/compute/metadata v0.5.0 h1:Zr0eK8JbFv6+Wi4ilXAR8FJ3wyNdpxHKJNPos6LTZOY=
cloud.google.com/go/compute/metadata v0.5.0/go.mod h1:aHnloV2TPI38yx4s9+wAZhHykWvVCfu7hQbF+9CWoiY=
cloud.google.com/go/longrunning v0.5.7 h1:WLbHekDbjK1fVFD3ibpFFVoyizlLRl73I7YKuAKilhU=
cloud.google.com/go/longrunning v0.5.7/go.mod h1:8GClkudohy1Fxm3owmBGid8W0pSgodEMwEAztp38Xng=
github.com/BurntSushi/toml v0.3.1/go.mod h1:xHWCNGjB5oqiDr8zfno3MHue2Ht5sIBksp03qcyfWMU=
github.com/PuerkitoBio/purell v1.1.1 h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=
github.com/PuerkitoBio/purell v1.1.1/go.mod h1:c11w/QuzBsJSee3cPx9rAFu61PvFxuPbtSwDGJws/X0=
github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=
github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578/go.mod h1:uGdkoq3SwY9Y+13GIhn11/XLaGBb4BfwItxLd5jeuXE=
github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef/go.mod h1:WaHUgvxTVq04UNunO+XhnAqY/wQc+bxr74GqbsZ/Jqw=
github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 h1:DklsrG3dyBCFEj5IhUbnKptjxatkF07cF2ak3yi77so=
github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2/go.mod h1:WaHUgvxTVq04UNunO+XhnAqY/wQc+bxr74GqbsZ/Jqw=
github.com/census-instrumentation/opencensus-proto v0.2.1/go.mod h1:f6KPmirojxKA12rnyqOA5BBL4O983OfeGPqjHWSTneU=
github.com/client9/misspell v0.3.4/go.mod h1:qj6jICC3Q7zFZvVWo7KLAzC3yx5G7kyvSDkc90ppPyw=
github.com/cncf/udpa/go v0.0.0-20191209042840-269d4d468f6f/go.mod h1:M8M6+tZqaGXZJjfX53e64911xZQV5JYwmTeXPW+k8Sc=
github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/envoyproxy/go-control-plane v0.9.0/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
github.com/envoyproxy/go-control-plane v0.9.1-0.20191026205805-5f8ba28d4473/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
github.com/envoyproxy/go-control-plane v0.9.4/go.mod h1:6rpuAdCZL397s3pYoYcLgu1mIlRU8Am5FuJP05cCM98=
github.com/envoyproxy/protoc-gen-validate v0.1.0/go.mod h1:iSmxcyjqTsJpI2R4NaDN7+kN2VEUnK/pcBlmesArF7c=
github.com/felixge/httpsnoop v1.0.4 h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=
github.com/felixge/httpsnoop v1.0.4/go.mod h1:m8KPJKqk1gH5J9DgRY2ASl2lWCfGKXixSwevea8zH2U=
github.com/go-logr/logr v1.2.2/go.mod h1:jdQByPbusPIv2/zmleS9BjJVeZ6kBagPoEUsqbVz/1A=
github.com/go-logr/logr v1.4.2 h1:6pFjapn8bFcIbiKo3XT4j/BhANplGihG6tvd+8rYgrY=
github.com/go-logr/logr v1.4.2/go.mod h1:9T104GzyrTigFIr8wt5mBrctHMim0Nb2HLGrmQ40KvY=
github.com/go-logr/stdr v1.2.2 h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=
github.com/go-logr/stdr v1.2.2/go.mod h1:mMo/vtBO5dYbehREoey6XUKy/eSumjCCveDpRre4VKE=
github.com/go-openapi/analysis v0.21.2 h1:hXFrOYFHUAMQdu6zwAiKKJHJQ8kqZs1ux/ru1P1wLJU=
github.com/go-openapi/analysis v0.21.2/go.mod h1:HZwRk4RRisyG8vx2Oe6aqeSQcoxRp47Xkp3+K6q+LdY=
github.com/go-openapi/errors v0.19.8/go.mod h1:cM//ZKUKyO06HSwqAelJ5NsEMMcpa6VpXe8DOa1Mi1M=
github.com/go-openapi/errors v0.19.9/go.mod h1:cM//ZKUKyO06HSwqAelJ5NsEMMcpa6VpXe8DOa1Mi1M=
github.com/go-openapi/errors v0.22.0 h1:c4xY/OLxUBSTiepAg3j/MHuAv5mJhnf53LLMWFB+u/w=
github.com/go-openapi/errors v0.22.0/go.mod h1:J3DmZScxCDufmIMsdOuDHxJbdOGC0xtUynjIx092vXE=
github.com/go-openapi/jsonpointer v0.19.3/go.mod h1:Pl9vOtqEWErmShwVjC8pYs9cog34VGT37dQOVbmoatg=
github.com/go-openapi/jsonpointer v0.19.5 h1:gZr+CIYByUqjcgeLXnQu2gHYQC9o73G2XUeOFYEICuY=
github.com/go-openapi/jsonpointer v0.19.5/go.mod h1:Pl9vOtqEWErmShwVjC8pYs9cog34VGT37dQOVbmoatg=
github.com/go-openapi/jsonreference v0.19.6 h1:UBIxjkht+AWIgYzCDSv2GN+E/togfwXUJFRTWhl2Jjs=
github.com/go-openapi/jsonreference v0.19.6/go.mod h1:diGHMEHg2IqXZGKxqyvWdfWU/aim5Dprw5bqpKkTvns=
github.com/go-openapi/loads v0.21.1 h1:Wb3nVZpdEzDTcly8S4HMkey6fjARRzb7iEaySimlDW0=
github.com/go-openapi/loads v0.21.1/go.mod h1:/DtAMXXneXFjbQMGEtbamCZb+4x7eGwkvZCvBmwUG+g=
github.com/go-openapi/spec v0.20.4 h1:O8hJrt0UMnhHcluhIdUgCLRWyM2x7QkBXRvOs7m+O1M=
github.com/go-openapi/spec v0.20.4/go.mod h1:faYFR1CvsJZ0mNsmsphTMSoRrNV3TEDoAM7FOEWeq8I=
github.com/go-openapi/strfmt v0.21.0/go.mod h1:ZRQ409bWMj+SOgXofQAGTIo2Ebu72Gs+WaRADcS5iNg=
github.com/go-openapi/strfmt v0.21.1/go.mod h1:I/XVKeLc5+MM5oPNN7P6urMOpuLXEcNrCX/rPGuWb0k=
github.com/go-openapi/strfmt v0.23.0 h1:nlUS6BCqcnAk0pyhi9Y+kdDVZdZMHfEKQiS4HaMgO/c=
github.com/go-openapi/strfmt v0.23.0/go.mod h1:NrtIpfKtWIygRkKVsxh7XQMDQW5HKQl6S5ik2elW+K4=
github.com/go-openapi/swag v0.19.5/go.mod h1:POnQmlKehdgb5mhVOsnJFsivZCEZ/vjK9gh66Z9tfKk=
github.com/go-openapi/swag v0.19.15/go.mod h1:QYRuS/SOXUCsnplDa677K7+DxSOj6IPNl/eQntq43wQ=
github.com/go-openapi/swag v0.21.1/go.mod h1:QYRuS/SOXUCsnplDa677K7+DxSOj6IPNl/eQntq43wQ=
github.com/go-openapi/swag v0.22.3 h1:yMBqmnQ0gyZvEb/+KzuWZOXgllrXT4SADYbvDaXHv/g=
github.com/go-openapi/swag v0.22.3/go.mod h1:UzaqsxGiab7freDnrUUra0MwWfN/q7tE4j+VcZ0yl14=
github.com/go-openapi/validate v0.21.0 h1:+Wqk39yKOhfpLqNLEC0/eViCkzM5FVXVqrvt526+wcI=
github.com/go-openapi/validate v0.21.0/go.mod h1:rjnrwK57VJ7A8xqfpAOEKRH8yQSGUriMu5/zuPSQ1hg=
github.com/go-stack/stack v1.8.0/go.mod h1:v0f6uXyyMGvRgIKkXu+yp6POWl0qKG85gN/melR3HDY=
github.com/gobuffalo/attrs v0.0.0-20190224210810-a9411de4debd/go.mod h1:4duuawTqi2wkkpB4ePgWMaai6/Kc6WEz83bhFwpHzj0=
github.com/gobuffalo/depgen v0.0.0-20190329151759-d478694a28d3/go.mod h1:3STtPUQYuzV0gBVOY3vy6CfMm/ljR4pABfrTeHNLHUY=
github.com/gobuffalo/depgen v0.1.0/go.mod h1:+ifsuy7fhi15RWncXQQKjWS9JPkdah5sZvtHc2RXGlg=
github.com/gobuffalo/envy v1.6.15/go.mod h1:n7DRkBerg/aorDM8kbduw5dN3oXGswK5liaSCx4T5NI=
github.com/gobuffalo/envy v1.7.0/go.mod h1:n7DRkBerg/aorDM8kbduw5dN3oXGswK5liaSCx4T5NI=
github.com/gobuffalo/flect v0.1.0/go.mod h1:d2ehjJqGOH/Kjqcoz+F7jHTBbmDb38yXA598Hb50EGs=
github.com/gobuffalo/flect v0.1.1/go.mod h1:8JCgGVbRjJhVgD6399mQr4fx5rRfGKVzFjbj6RE/9UI=
github.com/gobuffalo/flect v0.1.3/go.mod h1:8JCgGVbRjJhVgD6399mQr4fx5rRfGKVzFjbj6RE/9UI=
github.com/gobuffalo/genny v0.0.0-20190329151137-27723ad26ef9/go.mod h1:rWs4Z12d1Zbf19rlsn0nurr75KqhYp52EAGGxTbBhNk=
github.com/gobuffalo/genny v0.0.0-20190403191548-3ca520ef0d9e/go.mod h1:80lIj3kVJWwOrXWWMRzzdhW3DsrdjILVil/SFKBzF28=
github.com/gobuffalo/genny v0.1.0/go.mod h1:XidbUqzak3lHdS//TPu2OgiFB+51Ur5f7CSnXZ/JDvo=
github.com/gobuffalo/genny v0.1.1/go.mod h1:5TExbEyY48pfunL4QSXxlDOmdsD44RRq4mVZ0Ex28Xk=
github.com/gobuffalo/gitgen v0.0.0-20190315122116-cc086187d211/go.mod h1:vEHJk/E9DmhejeLeNt7UVvlSGv3ziL+djtTr3yyzcOw=
github.com/gobuffalo/gogen v0.0.0-20190315121717-8f38393713f5/go.mod h1:V9QVDIxsgKNZs6L2IYiGR8datgMhB577vzTDqypH360=
github.com/gobuffalo/gogen v0.1.0/go.mod h1:8NTelM5qd8RZ15VjQTFkAW6qOMx5wBbW4dSCS3BY8gg=
github.com/gobuffalo/gogen v0.1.1/go.mod h1:y8iBtmHmGc4qa3urIyo1shvOD8JftTtfcKi+71xfDNE=
github.com/gobuffalo/logger v0.0.0-20190315122211-86e12af44bc2/go.mod h1:QdxcLw541hSGtBnhUc4gaNIXRjiDppFGaDqzbrBd3v8=
github.com/gobuffalo/mapi v1.0.1/go.mod h1:4VAGh89y6rVOvm5A8fKFxYG+wIW6LO1FMTG9hnKStFc=
github.com/gobuffalo/mapi v1.0.2/go.mod h1:4VAGh89y6rVOvm5A8fKFxYG+wIW6LO1FMTG9hnKStFc=
github.com/gobuffalo/packd v0.0.0-20190315124812-a385830c7fc0/go.mod h1:M2Juc+hhDXf/PnmBANFCqx4DM3wRbgDvnVWeG2RIxq4=
github.com/gobuffalo/packd v0.1.0/go.mod h1:M2Juc+hhDXf/PnmBANFCqx4DM3wRbgDvnVWeG2RIxq4=
github.com/gobuffalo/packr/v2 v2.0.9/go.mod h1:emmyGweYTm6Kdper+iywB6YK5YzuKchGtJQZ0Odn4pQ=
github.com/gobuffalo/packr/v2 v2.2.0/go.mod h1:CaAwI0GPIAv+5wKLtv8Afwl+Cm78K/I/VCm/3ptBN+0=
github.com/gobuffalo/syncx v0.0.0-20190224160051-33c29581e754/go.mod h1:HhnNqWY95UYwwW3uSASeV7vtgYkT2t16hJgV3AEPUpw=
github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b/go.mod h1:SBH7ygxi8pfUlaOkMMuAQtPIUF8ecWP5IEl/CR7VP2Q=
github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=
github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
github.com/golang/mock v1.1.1/go.mod h1:oTYuIxOrZwtPieC+H1uAHpcLFnEyAGVDL/k47Jfbm0A=
github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.3.2/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.4.0-rc.1/go.mod h1:ceaxUfeHdC40wWswd/P6IGgMaK3YpKi5j83Wpe3EHw8=
github.com/golang/protobuf v1.4.0-rc.1.0.20200221234624-67d41d38c208/go.mod h1:xKAWHe0F5eneWXFV3EuXVDTCmh+JuBKY0li0aMyXATA=
github.com/golang/protobuf v1.4.0-rc.2/go.mod h1:LlEzMj4AhA7rCAGe4KMBDvJI+AwstrUpVNzEA03Pprs=
github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0/go.mod h1:WU3c8KckQ9AFe+yFwt9sWVRKCVIyN9cPHBJSNnbL67w=
github.com/golang/protobuf v1.4.0/go.mod h1:jodUvKwWbYaEsadDk5Fwe5c77LiNKVO9IDvqG2KuDX0=
github.com/golang/protobuf v1.4.1/go.mod h1:U8fpvMrcmy5pZrNK1lt4xCsGvpyWQ/VVv6QDs8UjoX8=
github.com/golang/protobuf v1.4.3/go.mod h1:oDoupMAO8OvCJWAcko0GGGIgR6R6ocIYbsSw735rRwI=
github.com/golang/protobuf v1.5.4 h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=
github.com/golang/protobuf v1.5.4/go.mod h1:lnTiLA8Wa4RWRcIUkrtSVa5nRhsEGBg48fD6rSs7xps=
github.com/golang/snappy v0.0.1/go.mod h1:/XxbfmMg8lxefKM7IXC3fBNl/7bRcc72aCRzEWrmP2Q=
github.com/google/generative-ai-go v0.17.0 h1:kUmCXUIwJouD7I7ev3OmxzzQVICyhIWAxaXk2yblCMY=
github.com/google/generative-ai-go v0.17.0/go.mod h1:JYolL13VG7j79kM5BtHz4qwONHkeJQzOCkKXnpqtS/E=
github.com/google/go-cmp v0.2.0/go.mod h1:oXzfMopK8JAjlY9xF4vHSVASa0yLyX7SntLO5aqRK0M=
github.com/google/go-cmp v0.3.0/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.3.1/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.4.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.2/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.3/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.6.0 h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=
github.com/google/go-cmp v0.6.0/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
github.com/google/s2a-go v0.1.8 h1:zZDs9gcbt9ZPLV0ndSyQk6Kacx2g/X+SKYovpnz3SMM=
github.com/google/s2a-go v0.1.8/go.mod h1:6iNWHTpQ+nfNRN5E00MSdfDwVesa8hhS32PhPO8deJA=
github.com/google/uuid v1.1.1/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/google/uuid v1.1.2/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/googleapis/enterprise-certificate-proxy v0.3.2 h1:Vie5ybvEvT75RniqhfFxPRy3Bf7vr3h0cechB90XaQs=
github.com/googleapis/enterprise-certificate-proxy v0.3.2/go.mod h1:VLSiSSBs/ksPL8kq3OBOQ6WRI2QnaFynd1DCjZ62+V0=
github.com/googleapis/gax-go/v2 v2.13.0 h1:yitjD5f7jQHhyDsnhKEBU52NdvvdSeGzlAnDPT0hH1s=
github.com/googleapis/gax-go/v2 v2.13.0/go.mod h1:Z/fvTZXF8/uw7Xu5GuslPw+bplx6SS338j1Is2S+B7A=
github.com/inconshreveable/mousetrap v1.0.0/go.mod h1:PxqpIevigyE2G7u3NXJIT2ANytuPF1OarO4DADm73n8=
github.com/joho/godotenv v1.3.0/go.mod h1:7hK45KPybAkOC6peb+G5yklZfMxEjkZhHbwpqxOKXbg=
github.com/josharian/intern v1.0.0 h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=
github.com/josharian/intern v1.0.0/go.mod h1:5DoeVV0s6jJacbCEi61lwdGj/aVlrQvzHFFd8Hwg//Y=
github.com/karrick/godirwalk v1.8.0/go.mod h1:H5KPZjojv4lE+QYImBI8xVtrBRgYrIVsaRPx4tDPEn4=
github.com/karrick/godirwalk v1.10.3/go.mod h1:RoGL9dQei4vP9ilrpETWE8CLOZ1kiN0LhBygSwrAsHA=
github.com/klauspost/compress v1.13.6/go.mod h1:/3/Vjq9QcHkK5uEr5lBEmyoZ1iFhe47etQ6QUkpK6sk=
github.com/konsorten/go-windows-terminal-sequences v1.0.1/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/konsorten/go-windows-terminal-sequences v1.0.2/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/kr/pretty v0.1.0 h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=
github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
github.com/mailru/easyjson v0.0.0-20190614124828-94de47d64c63/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
github.com/mailru/easyjson v0.7.6/go.mod h1:xzfreul335JAWq5oZzymOObrkdz5UnU4kGfJJLY9Nlc=
github.com/mailru/easyjson v0.7.7 h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=
github.com/mailru/easyjson v0.7.7/go.mod h1:xzfreul335JAWq5oZzymOObrkdz5UnU4kGfJJLY9Nlc=
github.com/markbates/oncer v0.0.0-20181203154359-bf2de49a0be2/go.mod h1:Ld9puTsIW75CHf65OeIOkyKbteujpZVXDpWK6YGZbxE=
github.com/markbates/safe v1.0.1/go.mod h1:nAqgmRi7cY2nqMc92/bSEeQA+R4OheNU2T1kNSCBdG0=
github.com/mitchellh/mapstructure v1.3.3/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/mapstructure v1.4.1/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/mapstructure v1.5.0 h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=
github.com/mitchellh/mapstructure v1.5.0/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe/go.mod h1:wL8QJuTMNUDYhXwkmfOly8iTdp5TEcJFWZD2D7SIkUc=
github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e/go.mod h1:zD1mROLANZcx1PVRCS0qkT7pwLkGfwJo4zjcN/Tysno=
github.com/oklog/ulid v1.3.1 h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=
github.com/oklog/ulid v1.3.1/go.mod h1:CirwcVhetQ6Lv90oh/F+FBtV6XMibvdAFo93nm5qn4U=
github.com/pelletier/go-toml v1.7.0/go.mod h1:vwGMzjaWMwyfHwgIBhI2YUM4fB6nL6lVAvS1LBMMhTE=
github.com/pkg/errors v0.8.0/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkg/errors v0.8.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4/go.mod h1:xMI15A0UPsDsEKsMN9yxemIoYk6Tm2C1GtYGdfGttqA=
github.com/rogpeppe/go-internal v1.1.0/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/rogpeppe/go-internal v1.2.2/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/rogpeppe/go-internal v1.3.0/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/sirupsen/logrus v1.4.0/go.mod h1:LxeOpSwHxABJmUn/MG1IvRgCAasNZTLOkJPxbbu5VWo=
github.com/sirupsen/logrus v1.4.1/go.mod h1:ni0Sbl8bgC9z8RoU9G6nDWqqs/fq4eDPysMBDgk/93Q=
github.com/sirupsen/logrus v1.4.2/go.mod h1:tLMulIdttU9McNUspp0xgXVQah82FyeX6MwdIuYE2rE=
github.com/spf13/cobra v0.0.3/go.mod h1:1l0Ry5zgKvJasoi3XT1TypsSe7PqH0Sj9dhYf7v3XqQ=
github.com/spf13/pflag v1.0.3/go.mod h1:DYY7MBk1bdzusC3SYhjObp+wFpr4gzcvqqNjLnInEg4=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/objx v0.1.1/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/objx v0.4.0/go.mod h1:YvHI0jy2hoMjB+UWwv71VJQ9isScKT/TqJzVSSt89Yw=
github.com/stretchr/objx v0.5.0/go.mod h1:Yh+to48EsGEfYuaHDzXPcE3xhTkx73EhmCGUpEOglKo=
github.com/stretchr/objx v0.5.2 h1:xuMeJ0Sdp5ZMRXx/aWO6RZxdr3beISkG5/G/aIRr3pY=
github.com/stretchr/objx v0.5.2/go.mod h1:FRsXN1f5AsAjCGJKqEizvkpNtU+EGNCLh3NxZ/8L+MA=
github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.6.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.8.0/go.mod h1:yNjHg4UonilssWZ8iaSj1OCr/vHnekPRkoO+kdMU+MU=
github.com/stretchr/testify v1.8.1/go.mod h1:w2LPCIKwWwSfY2zedu0+kehJoqGctiVI29o6fzry7u4=
github.com/stretchr/testify v1.9.0 h1:HtqpIVDClZ4nwg75+f6Lvsy/wHu+3BoSGCbBAcpTsTg=
github.com/stretchr/testify v1.9.0/go.mod h1:r2ic/lqez/lEtzL7wO/rwa5dbSLXVDPFyf8C91i36aY=
github.com/tidwall/pretty v1.0.0/go.mod h1:XNkn88O1ChpSDQmQeStsy+sBenx6DDtFZJxhVysOjyk=
github.com/weaviate/weaviate v1.26.1 h1:Vecl/mX5dt6z6E27k5bjAnkNS1hvWNiVB0+yMY29BsU=
github.com/weaviate/weaviate v1.26.1/go.mod h1:o6nFEB4UozA+B2fnAWyF0HZqB2ab44pRhJHYwvxVyyA=
github.com/weaviate/weaviate-go-client/v4 v4.15.1 h1:zw+aw8DsZZkSiFyZOkuKKK49y0jnNr+4nKXCUKsyR/Q=
github.com/weaviate/weaviate-go-client/v4 v4.15.1/go.mod h1:0aUsHXNfsdRfbO46t4Us5KUHqkhY+hlHevK6fBu72J4=
github.com/xdg-go/pbkdf2 v1.0.0/go.mod h1:jrpuAogTd400dnrH08LKmI/xc1MbPOebTwRqcT5RDeI=
github.com/xdg-go/scram v1.0.2/go.mod h1:1WAq6h33pAW+iRreB34OORO2Nf7qel3VV3fjBj+hCSs=
github.com/xdg-go/stringprep v1.0.2/go.mod h1:8F9zXuvzgwmyT5DUm4GUfZGDdT3W+LCvS6+da4O5kxM=
github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d/go.mod h1:rHwXgn7JulP+udvsHwJoVG1YGAP6VLg4y9I5dyZdqmA=
go.mongodb.org/mongo-driver v1.7.3/go.mod h1:NqaYOwnXWr5Pm7AOpO5QFxKJ503nbMse/R79oO62zWg=
go.mongodb.org/mongo-driver v1.7.5/go.mod h1:VXEWRZ6URJIkUq2SCAyapmhH0ZLRBP+FT4xhp5Zvxng=
go.mongodb.org/mongo-driver v1.14.0 h1:P98w8egYRjYe3XDjxhYJagTokP/H6HzlsnojRgZRd80=
go.mongodb.org/mongo-driver v1.14.0/go.mod h1:Vzb0Mk/pa7e6cWw85R4F/endUC3u0U9jGcNU603k65c=
go.opencensus.io v0.24.0 h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=
go.opencensus.io v0.24.0/go.mod h1:vNK8G9p7aAivkbmorf4v+7Hgx+Zs0yY+0fOtgBfjQKo=
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 h1:A3SayB3rNyt+1S6qpI9mHPkeHTZbD7XILEqWnYZb2l0=
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0/go.mod h1:27iA5uvhuRNmalO+iEUdVn5ZMj2qy10Mm+XRIpRmyuU=
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 h1:Xs2Ncz0gNihqu9iosIZ5SkBbWo5T8JhhLJFMQL1qmLI=
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0/go.mod h1:vy+2G/6NvVMpwGX/NyLqcC41fxepnuKHk16E6IZUcJc=
go.opentelemetry.io/otel v1.26.0 h1:LQwgL5s/1W7YiiRwxf03QGnWLb2HW4pLiAhaA5cZXBs=
go.opentelemetry.io/otel v1.26.0/go.mod h1:UmLkJHUAidDval2EICqBMbnAd0/m2vmpf/dAM+fvFs4=
go.opentelemetry.io/otel/metric v1.26.0 h1:7S39CLuY5Jgg9CrnA9HHiEjGMF/X2VHvoXGgSllRz30=
go.opentelemetry.io/otel/metric v1.26.0/go.mod h1:SY+rHOI4cEawI9a7N1A4nIg/nTQXe1ccCNWYOJUrpX4=
go.opentelemetry.io/otel/trace v1.26.0 h1:1ieeAUb4y0TE26jUFrCIXKpTuVK7uJGN9/Z/2LP5sQA=
go.opentelemetry.io/otel/trace v1.26.0/go.mod h1:4iDxvGDQuUkHve82hJJ8UqrwswHYsZuWCBllGV2U2y0=
golang.org/x/crypto v0.0.0-20180904163835-0709b304e793/go.mod h1:6SG95UA2DQfeDnfUPMdvaQW0Q7yPrPDi9nlGo2tz2b4=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20190422162423-af44ce270edf/go.mod h1:WFFai1msRO1wXaEeE5yQxYXgSfI8pQAWXbQop6sCtWE=
golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/crypto v0.26.0 h1:RrRspgV4mU+YwB4FYnuBoKsUapNIL5cohGAmSH3azsw=
golang.org/x/crypto v0.26.0/go.mod h1:GY7jblb9wI+FOo5y8/S2oY4zWP07AkOJ4+jxCqdqn54=
golang.org/x/exp v0.0.0-20190121172915-509febef88a4/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3/go.mod h1:UVdnD1Gm6xHRNCYTkRU2/jEulfH38KcIWyp/GAMgvoE=
golang.org/x/lint v0.0.0-20190227174305-5b3e6a55c961/go.mod h1:wehouNa3lNwaWXcvxsM5YxQ5yQlVC4a0KAMCusXpPoU=
golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
golang.org/x/net v0.0.0-20180724234803-3673e40ba225/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20180826012351-8a410e7b638d/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190213061140-3a22650c66bd/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190311183353-d8887717615a/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20201110031124-69a78807bb2b/go.mod h1:sp8m0HH+o8qH0wwXwYZr8TS3Oi6o0r6Gce1SSxlDquU=
golang.org/x/net v0.0.0-20210421230115-4e50805a0758/go.mod h1:72T/g9IO56b78aLF+1Kcs5dz7/ng1VjMUvfKvpfy+jM=
golang.org/x/net v0.28.0 h1:a9JDOJc5GMUJ0+UDqmLT86WiEy7iWyIhz8gz8E4e5hE=
golang.org/x/net v0.28.0/go.mod h1:yqtgsTWOOnlGLG9GFRrK3++bGOUEkNBoHZc8MEDWPNg=
golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be/go.mod h1:N/0e6XlmueqKjAGxoOufVs8QHGRruUQn6yWY3a++T0U=
golang.org/x/oauth2 v0.22.0 h1:BzDx2FehcG7jJwgWLELCdmLuxk2i+x9UDpSiss2u0ZA=
golang.org/x/oauth2 v0.22.0/go.mod h1:XYTD2NtWslqkgxebSiOHnXEap4TF09sJSc7H1sXbhtI=
golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20181108010431-42b317875d0f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190412183630-56d357773e84/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.8.0 h1:3NFvSEYkUoMifnESzZl15y791HH1qU2xm6eCJU5ZPXQ=
golang.org/x/sync v0.8.0/go.mod h1:Czt+wKu1gCyEFDUtn0jG5QVvpJ6rzVqr5aXyt9drQfk=
golang.org/x/sys v0.0.0-20180830151530-49385e6e1522/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190403152447-81d4e9dc473e/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190419153524-e8e3143a4f4a/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190422165155-953cdadca894/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190531175056-4c3a928424d2/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210420072515-93ed5bcd2bfe/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.24.0 h1:Twjiwq9dn6R1fQcyiK+wQyHWfaz/BJB+YIpzU/Cv3Xg=
golang.org/x/sys v0.24.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.5/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
golang.org/x/text v0.17.0 h1:XtiM5bkSOt+ewxlOE/aE/AKEHibwj/6gvWMl9Rsh0Qc=
golang.org/x/text v0.17.0/go.mod h1:BuEKDfySbSR4drPmRPG/7iBdf8hvFMuRexcpahXilzY=
golang.org/x/time v0.6.0 h1:eTDhh4ZXt5Qf0augr54TN6suAUudPcawVZeIAPU7D4U=
golang.org/x/time v0.6.0/go.mod h1:3BpzKBy/shNhVucY/MWOyx10tF3SFh9QdLuxbVysPQM=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190114222345-bf090417da8b/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190226205152-f727befe758c/go.mod h1:9Yl7xja0Znq3iFh3HoIrodX9oNMXvdceNzlUR8zjMvY=
golang.org/x/tools v0.0.0-20190311212946-11955173bddd/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190329151228-23e29df326fe/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190416151739-9c9e1878f421/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190420181800-aa740d480789/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135/go.mod h1:RgjU9mgBXZiqYHBnxXauZ1Gv1EHHAz9KjViQ78xBX0Q=
golang.org/x/tools v0.0.0-20190531172133-b3315ee88b7d/go.mod h1:/rFqwRUd4F7ZHNgwSSTFct+R/Kf4OFW1sUzUTQQTgfc=
golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
google.golang.org/api v0.194.0 h1:dztZKG9HgtIpbI35FhfuSNR/zmaMVdxNlntHj1sIS4s=
google.golang.org/api v0.194.0/go.mod h1:AgvUFdojGANh3vI+P7EVnxj3AISHllxGCJSFmggmnd0=
google.golang.org/appengine v1.1.0/go.mod h1:EbEs0AVv82hx2wNQdGPgUI5lhzA/G0D9YwlJXL52JkM=
google.golang.org/appengine v1.4.0/go.mod h1:xpcJRLb0r/rnEns0DIKYYv+WjYCduHsrkT7/EB5XEv4=
google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8/go.mod h1:JiN7NxoALGmiZfu7CAH4rXhgtRTLTxftemlI0sWmxmc=
google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55/go.mod h1:DMBHOl98Agz4BDEuKkezgsaosCRResVns1a3J2ZsMNc=
google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013/go.mod h1:NbSheEEYHJ7i3ixzK3sjbqSGDJWnxyFXZblF3eUsNvo=
google.golang.org/genproto/googleapis/api v0.0.0-20240725223205-93522f1f2a9f h1:b1Ln/PG8orm0SsBbHZWke8dDp2lrCD4jSmfglFpTZbk=
google.golang.org/genproto/googleapis/api v0.0.0-20240725223205-93522f1f2a9f/go.mod h1:AHT0dDg3SoMOgZGnZk29b5xTbPHMoEC8qthmBLJCpys=
google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 h1:e7S5W7MGGLaSu8j3YjdezkZ+m1/Nm0uRVRMEMGk26Xs=
google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142/go.mod h1:UqMtugtsSgubUsoxbuAoiCXvqvErP7Gf0so0mK9tHxU=
google.golang.org/grpc v1.19.0/go.mod h1:mqu4LbDTu4XGKhr4mRzUsmM4RtVoemTSY81AxZiDr8c=
google.golang.org/grpc v1.23.0/go.mod h1:Y5yQAOtifL1yxbo5wqy6BxZv8vAUGQwXBOALyacEbxg=
google.golang.org/grpc v1.25.1/go.mod h1:c3i+UQWmh7LiEpx4sFZnkU36qjEYZ0imhYfXVyQciAY=
google.golang.org/grpc v1.27.0/go.mod h1:qbnxyOmOxrQa7FizSgH+ReBfzJrCY1pSN7KXBS8abTk=
google.golang.org/grpc v1.33.2/go.mod h1:JMHMWHQWaTccqQQlmk3MJZS+GWXOdAesneDmEnv2fbc=
google.golang.org/grpc v1.65.0 h1:bs/cUb4lp1G5iImFFd3u5ixQzweKizoZJAwBNLR42lc=
google.golang.org/grpc v1.65.0/go.mod h1:WgYC2ypjlB0EiQi6wdKixMqukr6lBc0Vo+oOgjrM5ZQ=
google.golang.org/protobuf v0.0.0-20200109180630-ec00e32a8dfd/go.mod h1:DFci5gLYBciE7Vtevhsrf46CRTquxDuWsQurQQe4oz8=
google.golang.org/protobuf v0.0.0-20200221191635-4d8936d0db64/go.mod h1:kwYJMbMJ01Woi6D6+Kah6886xMZcty6N08ah7+eCXa0=
google.golang.org/protobuf v0.0.0-20200228230310-ab0ca4ff8a60/go.mod h1:cfTl7dwQJ+fmap5saPgwCLgHXTUD7jkjRqWcaiX5VyM=
google.golang.org/protobuf v1.20.1-0.20200309200217-e05f789c0967/go.mod h1:A+miEFZTKqfCUM6K7xSMQL9OKL/b6hQv+e19PK+JZNE=
google.golang.org/protobuf v1.21.0/go.mod h1:47Nbq4nVaFHyn7ilMalzfO3qCViNmqZ2kzikPIcrTAo=
google.golang.org/protobuf v1.22.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.1-0.20200526195155-81db48ad09cc/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.25.0/go.mod h1:9JNX74DMeImyA3h4bdi1ymwjUzf21/xIlbajtzgsN7c=
google.golang.org/protobuf v1.34.2 h1:6xV6lTsCfpGD21XK49h7MhtcApnLqkfYgPcdHftf6hg=
google.golang.org/protobuf v1.34.2/go.mod h1:qYOHts0dSfpeUzUFpOMr/WGzszTmLH+DiWniOlNbLDw=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
gopkg.in/errgo.v2 v2.1.0/go.mod h1:hNsd1EY+bozCKY1Ytp96fpM3vjJbqLJn88ws8XvfDNI=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.4.0 h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=
gopkg.in/yaml.v2 v2.4.0/go.mod h1:RDklbk79AGWmwhnvt/jBztapEOGDOx6ZbXqjP6csGnQ=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20200605160147-a5ece683394c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
honnef.co/go/tools v0.0.0-20190102054323-c2f93a96b099/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
```

### ragserver/ragserver/json.go

```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

// readRequestJSON expects req to have a JSON content type with a body that
// contains a JSON-encoded value complying with the underlying type of target.
// It populates target, or returns an error.
func readRequestJSON(req *http.Request, target any) error {
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	if mediaType != "application/json" {
		return fmt.Errorf("expect application/json Content-Type, got %s", mediaType)
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(target)
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
```

### ragserver/ragserver/main.go

```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command ragserver is an HTTP server that implements RAG (Retrieval
// Augmented Generation) using the Gemini model and Weaviate. See the
// accompanying README file for additional details.
package main

import (
	"cmp"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
	"google.golang.org/api/option"
)

const generativeModelName = "gemini-1.5-flash"
const embeddingModelName = "text-embedding-004"

// This is a standard Go HTTP server. Server state is in the ragServer struct.
// The `main` function connects to the required services (Weaviate and Google
// AI), initializes the server state and registers HTTP handlers.
func main() {
	ctx := context.Background()
	wvClient, err := initWeaviate(ctx)
	if err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	genaiClient, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer genaiClient.Close()

	server := &ragServer{
		ctx:      ctx,
		wvClient: wvClient,
		genModel: genaiClient.GenerativeModel(generativeModelName),
		embModel: genaiClient.EmbeddingModel(embeddingModelName),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /add/", server.addDocumentsHandler)
	mux.HandleFunc("POST /query/", server.queryHandler)

	port := cmp.Or(os.Getenv("SERVERPORT"), "9020")
	address := "localhost:" + port
	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

type ragServer struct {
	ctx      context.Context
	wvClient *weaviate.Client
	genModel *genai.GenerativeModel
	embModel *genai.EmbeddingModel
}

func (rs *ragServer) addDocumentsHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type document struct {
		Text string
	}
	type addRequest struct {
		Documents []document
	}
	ar := &addRequest{}

	err := readRequestJSON(req, ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the batch embedding API to embed all documents at once.
	batch := rs.embModel.NewBatch()
	for _, doc := range ar.Documents {
		batch.AddContent(genai.Text(doc.Text))
	}
	log.Printf("invoking embedding model with %v documents", len(ar.Documents))
	rsp, err := rs.embModel.BatchEmbedContents(rs.ctx, batch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(rsp.Embeddings) != len(ar.Documents) {
		http.Error(w, "embedded batch size mismatch", http.StatusInternalServerError)
		return
	}

	// Convert our documents - along with their embedding vectors - into types
	// used by the Weaviate client library.
	objects := make([]*models.Object, len(ar.Documents))
	for i, doc := range ar.Documents {
		objects[i] = &models.Object{
			Class: "Document",
			Properties: map[string]any{
				"text": doc.Text,
			},
			Vector: rsp.Embeddings[i].Values,
		}
	}

	// Store documents with embeddings in the Weaviate DB.
	log.Printf("storing %v objects in weaviate", len(objects))
	_, err = rs.wvClient.Batch().ObjectsBatcher().WithObjects(objects...).Do(rs.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs *ragServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type queryRequest struct {
		Content string
	}
	qr := &queryRequest{}
	err := readRequestJSON(req, qr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Embed the query contents.
	rsp, err := rs.embModel.EmbedContent(rs.ctx, genai.Text(qr.Content))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Search weaviate to find the most relevant (closest in vector space)
	// documents to the query.
	gql := rs.wvClient.GraphQL()
	result, err := gql.Get().
		WithNearVector(
			gql.NearVectorArgBuilder().WithVector(rsp.Embedding.Values)).
		WithClassName("Document").
		WithFields(graphql.Field{Name: "text"}).
		WithLimit(3).
		Do(rs.ctx)
	if werr := combinedWeaviateError(result, err); werr != nil {
		http.Error(w, werr.Error(), http.StatusInternalServerError)
		return
	}

	contents, err := decodeGetResults(result)
	if err != nil {
		http.Error(w, fmt.Errorf("reading weaviate response: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	// Create a RAG query for the LLM with the most relevant documents as
	// context.
	ragQuery := fmt.Sprintf(ragTemplateStr, qr.Content, strings.Join(contents, "\n"))
	resp, err := rs.genModel.GenerateContent(rs.ctx, genai.Text(ragQuery))
	if err != nil {
		log.Printf("calling generative model: %v", err.Error())
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	if len(resp.Candidates) != 1 {
		log.Printf("got %v candidates, expected 1", len(resp.Candidates))
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	var respTexts []string
	for _, part := range resp.Candidates[0].Content.Parts {
		if pt, ok := part.(genai.Text); ok {
			respTexts = append(respTexts, string(pt))
		} else {
			log.Printf("bad type of part: %v", pt)
			http.Error(w, "generative model error", http.StatusInternalServerError)
			return
		}
	}

	renderJSON(w, strings.Join(respTexts, "\n"))
}

const ragTemplateStr = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Question:
%s

Context:
%s
`

// decodeGetResults decodes the result returned by Weaviate's GraphQL Get
// query; these are returned as a nested map[string]any (just like JSON
// unmarshaled into a map[string]any). We have to extract all document contents
// as a list of strings.
func decodeGetResults(result *models.GraphQLResponse) ([]string, error) {
	data, ok := result.Data["Get"]
	if !ok {
		return nil, fmt.Errorf("Get key not found in result")
	}
	doc, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("Get key unexpected type")
	}
	slc, ok := doc["Document"].([]any)
	if !ok {
		return nil, fmt.Errorf("Document is not a list of results")
	}

	var out []string
	for _, s := range slc {
		smap, ok := s.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid element in list of documents")
		}
		s, ok := smap["text"].(string)
		if !ok {
			return nil, fmt.Errorf("expected string in list of documents")
		}
		out = append(out, s)
	}
	return out, nil
}
```

### ragserver/ragserver/weaviate.go

```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Utilities for working with Weaviate.

import (
	"cmp"
	"context"
	"fmt"
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

// initWeaviate initializes a weaviate client for our application.
func initWeaviate(ctx context.Context) (*weaviate.Client, error) {
	client, err := weaviate.NewClient(weaviate.Config{
		Host:   "localhost:" + cmp.Or(os.Getenv("WVPORT"), "9035"),
		Scheme: "http",
	})
	if err != nil {
		return nil, fmt.Errorf("initializing weaviate: %w", err)
	}

	// Create a new class (collection) in weaviate if it doesn't exist yet.
	cls := &models.Class{
		Class:      "Document",
		Vectorizer: "none",
	}
	exists, err := client.Schema().ClassExistenceChecker().WithClassName(cls.Class).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("weaviate error: %w", err)
	}
	if !exists {
		err = client.Schema().ClassCreator().WithClass(cls).Do(ctx)
		if err != nil {
			return nil, fmt.Errorf("weaviate error: %w", err)
		}
	}

	return client, nil
}

// combinedWeaviateError generates an error if err is non-nil or result has
// errors, and returns an error (or nil if there's no error). It's useful for
// the results of the Weaviate GraphQL API's "Do" calls.
func combinedWeaviateError(result *models.GraphQLResponse, err error) error {
	if err != nil {
		return err
	}
	if len(result.Errors) != 0 {
		var ss []string
		for _, e := range result.Errors {
			ss = append(ss, e.Message)
		}
		return fmt.Errorf("weaviate error: %v", ss)
	}
	return nil
}
```

### ragserver/ragserver-genkit/go.mod

```
module golang.org/x/example/ragserver/ragserver-genkit

go 1.23.0

require github.com/firebase/genkit/go v0.1.1

require (
	cloud.google.com/go v0.115.0 // indirect
	cloud.google.com/go/ai v0.8.1-0.20240711230438-265963bd5b91 // indirect
	cloud.google.com/go/auth v0.7.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/compute/metadata v0.4.0 // indirect
	cloud.google.com/go/longrunning v0.5.9 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.21.2 // indirect
	github.com/go-openapi/errors v0.22.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/strfmt v0.23.0 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-openapi/validate v0.21.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/generative-ai-go v0.16.1-0.20240711222609-09946422abc6 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.5 // indirect
	github.com/invopop/jsonschema v0.12.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/weaviate/weaviate v1.26.0-rc.1 // indirect
	github.com/weaviate/weaviate-go-client/v4 v4.15.0 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.14.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/exp v0.0.0-20240318143956-a85f2c67cd81 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/api v0.188.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240708141625-4ad9e859172b // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

### ragserver/ragserver-genkit/go.sum

```
cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
cloud.google.com/go v0.115.0 h1:CnFSK6Xo3lDYRoBKEcAtia6VSC837/ZkJuRduSFnr14=
cloud.google.com/go v0.115.0/go.mod h1:8jIM5vVgoAEoiVxQ/O4BFTfHqulPZgs/ufEzMcFMdWU=
cloud.google.com/go/ai v0.8.1-0.20240711230438-265963bd5b91 h1:VA80iXvWirtF1jQK5BQd7MPHvHOE+UZ2v4AJCcChHqk=
cloud.google.com/go/ai v0.8.1-0.20240711230438-265963bd5b91/go.mod h1:rVgd6oDdCDlN3mYqXqgE2nnzUblrwM/khbqLUXOJLeM=
cloud.google.com/go/auth v0.7.0 h1:kf/x9B3WTbBUHkC+1VS8wwwli9TzhSt0vSTVBmMR8Ts=
cloud.google.com/go/auth v0.7.0/go.mod h1:D+WqdrpcjmiCgWrXmLLxOVq1GACoE36chW6KXoEvuIw=
cloud.google.com/go/auth/oauth2adapt v0.2.2 h1:+TTV8aXpjeChS9M+aTtN/TjdQnzJvmzKFt//oWu7HX4=
cloud.google.com/go/auth/oauth2adapt v0.2.2/go.mod h1:wcYjgpZI9+Yu7LyYBg4pqSiaRkfEK3GQcpb7C/uyF1Q=
cloud.google.com/go/compute/metadata v0.4.0 h1:vHzJCWaM4g8XIcm8kopr3XmDA4Gy/lblD3EhhSux05c=
cloud.google.com/go/compute/metadata v0.4.0/go.mod h1:SIQh1Kkb4ZJ8zJ874fqVkslA29PRXuleyj6vOzlbK7M=
cloud.google.com/go/longrunning v0.5.9 h1:haH9pAuXdPAMqHvzX0zlWQigXT7B0+CL4/2nXXdBo5k=
cloud.google.com/go/longrunning v0.5.9/go.mod h1:HD+0l9/OOW0za6UWdKJtXoFAX/BGg/3Wj8p10NeWF7c=
github.com/BurntSushi/toml v0.3.1/go.mod h1:xHWCNGjB5oqiDr8zfno3MHue2Ht5sIBksp03qcyfWMU=
github.com/PuerkitoBio/purell v1.1.1 h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=
github.com/PuerkitoBio/purell v1.1.1/go.mod h1:c11w/QuzBsJSee3cPx9rAFu61PvFxuPbtSwDGJws/X0=
github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=
github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578/go.mod h1:uGdkoq3SwY9Y+13GIhn11/XLaGBb4BfwItxLd5jeuXE=
github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef/go.mod h1:WaHUgvxTVq04UNunO+XhnAqY/wQc+bxr74GqbsZ/Jqw=
github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 h1:DklsrG3dyBCFEj5IhUbnKptjxatkF07cF2ak3yi77so=
github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2/go.mod h1:WaHUgvxTVq04UNunO+XhnAqY/wQc+bxr74GqbsZ/Jqw=
github.com/bahlo/generic-list-go v0.2.0 h1:5sz/EEAK+ls5wF+NeqDpk5+iNdMDXrh3z3nPnH1Wvgk=
github.com/bahlo/generic-list-go v0.2.0/go.mod h1:2KvAjgMlE5NNynlg/5iLrrCCZ2+5xWbdbCW3pNTGyYg=
github.com/buger/jsonparser v1.1.1 h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=
github.com/buger/jsonparser v1.1.1/go.mod h1:6RYKKt7H4d4+iWqouImQ9R2FZql3VbhNgx27UK13J/0=
github.com/census-instrumentation/opencensus-proto v0.2.1/go.mod h1:f6KPmirojxKA12rnyqOA5BBL4O983OfeGPqjHWSTneU=
github.com/client9/misspell v0.3.4/go.mod h1:qj6jICC3Q7zFZvVWo7KLAzC3yx5G7kyvSDkc90ppPyw=
github.com/cncf/udpa/go v0.0.0-20191209042840-269d4d468f6f/go.mod h1:M8M6+tZqaGXZJjfX53e64911xZQV5JYwmTeXPW+k8Sc=
github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/envoyproxy/go-control-plane v0.9.0/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
github.com/envoyproxy/go-control-plane v0.9.1-0.20191026205805-5f8ba28d4473/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
github.com/envoyproxy/go-control-plane v0.9.4/go.mod h1:6rpuAdCZL397s3pYoYcLgu1mIlRU8Am5FuJP05cCM98=
github.com/envoyproxy/protoc-gen-validate v0.1.0/go.mod h1:iSmxcyjqTsJpI2R4NaDN7+kN2VEUnK/pcBlmesArF7c=
github.com/felixge/httpsnoop v1.0.4 h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=
github.com/felixge/httpsnoop v1.0.4/go.mod h1:m8KPJKqk1gH5J9DgRY2ASl2lWCfGKXixSwevea8zH2U=
github.com/firebase/genkit/go v0.1.1 h1:Cg1cjRvfJeZdCgs5JXYOV2JhXw64SZroJv/mOnTlVWw=
github.com/firebase/genkit/go v0.1.1/go.mod h1:XDbSDe348obNc/dqObxx7znYhANXoOQO9KdFamt1eS4=
github.com/go-logr/logr v1.2.2/go.mod h1:jdQByPbusPIv2/zmleS9BjJVeZ6kBagPoEUsqbVz/1A=
github.com/go-logr/logr v1.4.1 h1:pKouT5E8xu9zeFC39JXRDukb6JFQPXM5p5I91188VAQ=
github.com/go-logr/logr v1.4.1/go.mod h1:9T104GzyrTigFIr8wt5mBrctHMim0Nb2HLGrmQ40KvY=
github.com/go-logr/stdr v1.2.2 h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=
github.com/go-logr/stdr v1.2.2/go.mod h1:mMo/vtBO5dYbehREoey6XUKy/eSumjCCveDpRre4VKE=
github.com/go-openapi/analysis v0.21.2 h1:hXFrOYFHUAMQdu6zwAiKKJHJQ8kqZs1ux/ru1P1wLJU=
github.com/go-openapi/analysis v0.21.2/go.mod h1:HZwRk4RRisyG8vx2Oe6aqeSQcoxRp47Xkp3+K6q+LdY=
github.com/go-openapi/errors v0.19.8/go.mod h1:cM//ZKUKyO06HSwqAelJ5NsEMMcpa6VpXe8DOa1Mi1M=
github.com/go-openapi/errors v0.19.9/go.mod h1:cM//ZKUKyO06HSwqAelJ5NsEMMcpa6VpXe8DOa1Mi1M=
github.com/go-openapi/errors v0.22.0 h1:c4xY/OLxUBSTiepAg3j/MHuAv5mJhnf53LLMWFB+u/w=
github.com/go-openapi/errors v0.22.0/go.mod h1:J3DmZScxCDufmIMsdOuDHxJbdOGC0xtUynjIx092vXE=
github.com/go-openapi/jsonpointer v0.19.3/go.mod h1:Pl9vOtqEWErmShwVjC8pYs9cog34VGT37dQOVbmoatg=
github.com/go-openapi/jsonpointer v0.19.5 h1:gZr+CIYByUqjcgeLXnQu2gHYQC9o73G2XUeOFYEICuY=
github.com/go-openapi/jsonpointer v0.19.5/go.mod h1:Pl9vOtqEWErmShwVjC8pYs9cog34VGT37dQOVbmoatg=
github.com/go-openapi/jsonreference v0.19.6 h1:UBIxjkht+AWIgYzCDSv2GN+E/togfwXUJFRTWhl2Jjs=
github.com/go-openapi/jsonreference v0.19.6/go.mod h1:diGHMEHg2IqXZGKxqyvWdfWU/aim5Dprw5bqpKkTvns=
github.com/go-openapi/loads v0.21.1 h1:Wb3nVZpdEzDTcly8S4HMkey6fjARRzb7iEaySimlDW0=
github.com/go-openapi/loads v0.21.1/go.mod h1:/DtAMXXneXFjbQMGEtbamCZb+4x7eGwkvZCvBmwUG+g=
github.com/go-openapi/spec v0.20.4 h1:O8hJrt0UMnhHcluhIdUgCLRWyM2x7QkBXRvOs7m+O1M=
github.com/go-openapi/spec v0.20.4/go.mod h1:faYFR1CvsJZ0mNsmsphTMSoRrNV3TEDoAM7FOEWeq8I=
github.com/go-openapi/strfmt v0.21.0/go.mod h1:ZRQ409bWMj+SOgXofQAGTIo2Ebu72Gs+WaRADcS5iNg=
github.com/go-openapi/strfmt v0.21.1/go.mod h1:I/XVKeLc5+MM5oPNN7P6urMOpuLXEcNrCX/rPGuWb0k=
github.com/go-openapi/strfmt v0.23.0 h1:nlUS6BCqcnAk0pyhi9Y+kdDVZdZMHfEKQiS4HaMgO/c=
github.com/go-openapi/strfmt v0.23.0/go.mod h1:NrtIpfKtWIygRkKVsxh7XQMDQW5HKQl6S5ik2elW+K4=
github.com/go-openapi/swag v0.19.5/go.mod h1:POnQmlKehdgb5mhVOsnJFsivZCEZ/vjK9gh66Z9tfKk=
github.com/go-openapi/swag v0.19.15/go.mod h1:QYRuS/SOXUCsnplDa677K7+DxSOj6IPNl/eQntq43wQ=
github.com/go-openapi/swag v0.21.1/go.mod h1:QYRuS/SOXUCsnplDa677K7+DxSOj6IPNl/eQntq43wQ=
github.com/go-openapi/swag v0.22.3 h1:yMBqmnQ0gyZvEb/+KzuWZOXgllrXT4SADYbvDaXHv/g=
github.com/go-openapi/swag v0.22.3/go.mod h1:UzaqsxGiab7freDnrUUra0MwWfN/q7tE4j+VcZ0yl14=
github.com/go-openapi/validate v0.21.0 h1:+Wqk39yKOhfpLqNLEC0/eViCkzM5FVXVqrvt526+wcI=
github.com/go-openapi/validate v0.21.0/go.mod h1:rjnrwK57VJ7A8xqfpAOEKRH8yQSGUriMu5/zuPSQ1hg=
github.com/go-stack/stack v1.8.0/go.mod h1:v0f6uXyyMGvRgIKkXu+yp6POWl0qKG85gN/melR3HDY=
github.com/gobuffalo/attrs v0.0.0-20190224210810-a9411de4debd/go.mod h1:4duuawTqi2wkkpB4ePgWMaai6/Kc6WEz83bhFwpHzj0=
github.com/gobuffalo/depgen v0.0.0-20190329151759-d478694a28d3/go.mod h1:3STtPUQYuzV0gBVOY3vy6CfMm/ljR4pABfrTeHNLHUY=
github.com/gobuffalo/depgen v0.1.0/go.mod h1:+ifsuy7fhi15RWncXQQKjWS9JPkdah5sZvtHc2RXGlg=
github.com/gobuffalo/envy v1.6.15/go.mod h1:n7DRkBerg/aorDM8kbduw5dN3oXGswK5liaSCx4T5NI=
github.com/gobuffalo/envy v1.7.0/go.mod h1:n7DRkBerg/aorDM8kbduw5dN3oXGswK5liaSCx4T5NI=
github.com/gobuffalo/flect v0.1.0/go.mod h1:d2ehjJqGOH/Kjqcoz+F7jHTBbmDb38yXA598Hb50EGs=
github.com/gobuffalo/flect v0.1.1/go.mod h1:8JCgGVbRjJhVgD6399mQr4fx5rRfGKVzFjbj6RE/9UI=
github.com/gobuffalo/flect v0.1.3/go.mod h1:8JCgGVbRjJhVgD6399mQr4fx5rRfGKVzFjbj6RE/9UI=
github.com/gobuffalo/genny v0.0.0-20190329151137-27723ad26ef9/go.mod h1:rWs4Z12d1Zbf19rlsn0nurr75KqhYp52EAGGxTbBhNk=
github.com/gobuffalo/genny v0.0.0-20190403191548-3ca520ef0d9e/go.mod h1:80lIj3kVJWwOrXWWMRzzdhW3DsrdjILVil/SFKBzF28=
github.com/gobuffalo/genny v0.1.0/go.mod h1:XidbUqzak3lHdS//TPu2OgiFB+51Ur5f7CSnXZ/JDvo=
github.com/gobuffalo/genny v0.1.1/go.mod h1:5TExbEyY48pfunL4QSXxlDOmdsD44RRq4mVZ0Ex28Xk=
github.com/gobuffalo/gitgen v0.0.0-20190315122116-cc086187d211/go.mod h1:vEHJk/E9DmhejeLeNt7UVvlSGv3ziL+djtTr3yyzcOw=
github.com/gobuffalo/gogen v0.0.0-20190315121717-8f38393713f5/go.mod h1:V9QVDIxsgKNZs6L2IYiGR8datgMhB577vzTDqypH360=
github.com/gobuffalo/gogen v0.1.0/go.mod h1:8NTelM5qd8RZ15VjQTFkAW6qOMx5wBbW4dSCS3BY8gg=
github.com/gobuffalo/gogen v0.1.1/go.mod h1:y8iBtmHmGc4qa3urIyo1shvOD8JftTtfcKi+71xfDNE=
github.com/gobuffalo/logger v0.0.0-20190315122211-86e12af44bc2/go.mod h1:QdxcLw541hSGtBnhUc4gaNIXRjiDppFGaDqzbrBd3v8=
github.com/gobuffalo/mapi v1.0.1/go.mod h1:4VAGh89y6rVOvm5A8fKFxYG+wIW6LO1FMTG9hnKStFc=
github.com/gobuffalo/mapi v1.0.2/go.mod h1:4VAGh89y6rVOvm5A8fKFxYG+wIW6LO1FMTG9hnKStFc=
github.com/gobuffalo/packd v0.0.0-20190315124812-a385830c7fc0/go.mod h1:M2Juc+hhDXf/PnmBANFCqx4DM3wRbgDvnVWeG2RIxq4=
github.com/gobuffalo/packd v0.1.0/go.mod h1:M2Juc+hhDXf/PnmBANFCqx4DM3wRbgDvnVWeG2RIxq4=
github.com/gobuffalo/packr/v2 v2.0.9/go.mod h1:emmyGweYTm6Kdper+iywB6YK5YzuKchGtJQZ0Odn4pQ=
github.com/gobuffalo/packr/v2 v2.2.0/go.mod h1:CaAwI0GPIAv+5wKLtv8Afwl+Cm78K/I/VCm/3ptBN+0=
github.com/gobuffalo/syncx v0.0.0-20190224160051-33c29581e754/go.mod h1:HhnNqWY95UYwwW3uSASeV7vtgYkT2t16hJgV3AEPUpw=
github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b/go.mod h1:SBH7ygxi8pfUlaOkMMuAQtPIUF8ecWP5IEl/CR7VP2Q=
github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=
github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
github.com/golang/mock v1.1.1/go.mod h1:oTYuIxOrZwtPieC+H1uAHpcLFnEyAGVDL/k47Jfbm0A=
github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.3.2/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.4.0-rc.1/go.mod h1:ceaxUfeHdC40wWswd/P6IGgMaK3YpKi5j83Wpe3EHw8=
github.com/golang/protobuf v1.4.0-rc.1.0.20200221234624-67d41d38c208/go.mod h1:xKAWHe0F5eneWXFV3EuXVDTCmh+JuBKY0li0aMyXATA=
github.com/golang/protobuf v1.4.0-rc.2/go.mod h1:LlEzMj4AhA7rCAGe4KMBDvJI+AwstrUpVNzEA03Pprs=
github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0/go.mod h1:WU3c8KckQ9AFe+yFwt9sWVRKCVIyN9cPHBJSNnbL67w=
github.com/golang/protobuf v1.4.0/go.mod h1:jodUvKwWbYaEsadDk5Fwe5c77LiNKVO9IDvqG2KuDX0=
github.com/golang/protobuf v1.4.1/go.mod h1:U8fpvMrcmy5pZrNK1lt4xCsGvpyWQ/VVv6QDs8UjoX8=
github.com/golang/protobuf v1.4.3/go.mod h1:oDoupMAO8OvCJWAcko0GGGIgR6R6ocIYbsSw735rRwI=
github.com/golang/protobuf v1.5.4 h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=
github.com/golang/protobuf v1.5.4/go.mod h1:lnTiLA8Wa4RWRcIUkrtSVa5nRhsEGBg48fD6rSs7xps=
github.com/golang/snappy v0.0.1/go.mod h1:/XxbfmMg8lxefKM7IXC3fBNl/7bRcc72aCRzEWrmP2Q=
github.com/google/generative-ai-go v0.16.1-0.20240711222609-09946422abc6 h1:9nZYncjB1Wxlo7S+fM6YaXEWkK//zX9c0SExmzyctqA=
github.com/google/generative-ai-go v0.16.1-0.20240711222609-09946422abc6/go.mod h1:JYolL13VG7j79kM5BtHz4qwONHkeJQzOCkKXnpqtS/E=
github.com/google/go-cmp v0.2.0/go.mod h1:oXzfMopK8JAjlY9xF4vHSVASa0yLyX7SntLO5aqRK0M=
github.com/google/go-cmp v0.3.0/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.3.1/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.4.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.2/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.3/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.6.0 h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=
github.com/google/go-cmp v0.6.0/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
github.com/google/s2a-go v0.1.7 h1:60BLSyTrOV4/haCDW4zb1guZItoSq8foHCXrAnjBo/o=
github.com/google/s2a-go v0.1.7/go.mod h1:50CgR4k1jNlWBu4UfS4AcfhVe1r6pdZPygJ3R8F0Qdw=
github.com/google/uuid v1.1.1/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/google/uuid v1.1.2/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/googleapis/enterprise-certificate-proxy v0.3.2 h1:Vie5ybvEvT75RniqhfFxPRy3Bf7vr3h0cechB90XaQs=
github.com/googleapis/enterprise-certificate-proxy v0.3.2/go.mod h1:VLSiSSBs/ksPL8kq3OBOQ6WRI2QnaFynd1DCjZ62+V0=
github.com/googleapis/gax-go/v2 v2.12.5 h1:8gw9KZK8TiVKB6q3zHY3SBzLnrGp6HQjyfYBYGmXdxA=
github.com/googleapis/gax-go/v2 v2.12.5/go.mod h1:BUDKcWo+RaKq5SC9vVYL0wLADa3VcfswbOMMRmB9H3E=
github.com/inconshreveable/mousetrap v1.0.0/go.mod h1:PxqpIevigyE2G7u3NXJIT2ANytuPF1OarO4DADm73n8=
github.com/invopop/jsonschema v0.12.0 h1:6ovsNSuvn9wEQVOyc72aycBMVQFKz7cPdMJn10CvzRI=
github.com/invopop/jsonschema v0.12.0/go.mod h1:ffZ5Km5SWWRAIN6wbDXItl95euhFz2uON45H2qjYt+0=
github.com/joho/godotenv v1.3.0/go.mod h1:7hK45KPybAkOC6peb+G5yklZfMxEjkZhHbwpqxOKXbg=
github.com/josharian/intern v1.0.0 h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=
github.com/josharian/intern v1.0.0/go.mod h1:5DoeVV0s6jJacbCEi61lwdGj/aVlrQvzHFFd8Hwg//Y=
github.com/karrick/godirwalk v1.8.0/go.mod h1:H5KPZjojv4lE+QYImBI8xVtrBRgYrIVsaRPx4tDPEn4=
github.com/karrick/godirwalk v1.10.3/go.mod h1:RoGL9dQei4vP9ilrpETWE8CLOZ1kiN0LhBygSwrAsHA=
github.com/klauspost/compress v1.13.6/go.mod h1:/3/Vjq9QcHkK5uEr5lBEmyoZ1iFhe47etQ6QUkpK6sk=
github.com/konsorten/go-windows-terminal-sequences v1.0.1/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/konsorten/go-windows-terminal-sequences v1.0.2/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/kr/pretty v0.1.0 h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=
github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
github.com/mailru/easyjson v0.0.0-20190614124828-94de47d64c63/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
github.com/mailru/easyjson v0.7.6/go.mod h1:xzfreul335JAWq5oZzymOObrkdz5UnU4kGfJJLY9Nlc=
github.com/mailru/easyjson v0.7.7 h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=
github.com/mailru/easyjson v0.7.7/go.mod h1:xzfreul335JAWq5oZzymOObrkdz5UnU4kGfJJLY9Nlc=
github.com/markbates/oncer v0.0.0-20181203154359-bf2de49a0be2/go.mod h1:Ld9puTsIW75CHf65OeIOkyKbteujpZVXDpWK6YGZbxE=
github.com/markbates/safe v1.0.1/go.mod h1:nAqgmRi7cY2nqMc92/bSEeQA+R4OheNU2T1kNSCBdG0=
github.com/mitchellh/mapstructure v1.3.3/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/mapstructure v1.4.1/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/mapstructure v1.5.0 h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=
github.com/mitchellh/mapstructure v1.5.0/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe/go.mod h1:wL8QJuTMNUDYhXwkmfOly8iTdp5TEcJFWZD2D7SIkUc=
github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e/go.mod h1:zD1mROLANZcx1PVRCS0qkT7pwLkGfwJo4zjcN/Tysno=
github.com/oklog/ulid v1.3.1 h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=
github.com/oklog/ulid v1.3.1/go.mod h1:CirwcVhetQ6Lv90oh/F+FBtV6XMibvdAFo93nm5qn4U=
github.com/pelletier/go-toml v1.7.0/go.mod h1:vwGMzjaWMwyfHwgIBhI2YUM4fB6nL6lVAvS1LBMMhTE=
github.com/pkg/errors v0.8.0/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkg/errors v0.8.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4/go.mod h1:xMI15A0UPsDsEKsMN9yxemIoYk6Tm2C1GtYGdfGttqA=
github.com/rogpeppe/go-internal v1.1.0/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/rogpeppe/go-internal v1.2.2/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/rogpeppe/go-internal v1.3.0/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/sirupsen/logrus v1.4.0/go.mod h1:LxeOpSwHxABJmUn/MG1IvRgCAasNZTLOkJPxbbu5VWo=
github.com/sirupsen/logrus v1.4.1/go.mod h1:ni0Sbl8bgC9z8RoU9G6nDWqqs/fq4eDPysMBDgk/93Q=
github.com/sirupsen/logrus v1.4.2/go.mod h1:tLMulIdttU9McNUspp0xgXVQah82FyeX6MwdIuYE2rE=
github.com/spf13/cobra v0.0.3/go.mod h1:1l0Ry5zgKvJasoi3XT1TypsSe7PqH0Sj9dhYf7v3XqQ=
github.com/spf13/pflag v1.0.3/go.mod h1:DYY7MBk1bdzusC3SYhjObp+wFpr4gzcvqqNjLnInEg4=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/objx v0.1.1/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/objx v0.4.0/go.mod h1:YvHI0jy2hoMjB+UWwv71VJQ9isScKT/TqJzVSSt89Yw=
github.com/stretchr/objx v0.5.0/go.mod h1:Yh+to48EsGEfYuaHDzXPcE3xhTkx73EhmCGUpEOglKo=
github.com/stretchr/objx v0.5.2 h1:xuMeJ0Sdp5ZMRXx/aWO6RZxdr3beISkG5/G/aIRr3pY=
github.com/stretchr/objx v0.5.2/go.mod h1:FRsXN1f5AsAjCGJKqEizvkpNtU+EGNCLh3NxZ/8L+MA=
github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.6.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.8.0/go.mod h1:yNjHg4UonilssWZ8iaSj1OCr/vHnekPRkoO+kdMU+MU=
github.com/stretchr/testify v1.8.1/go.mod h1:w2LPCIKwWwSfY2zedu0+kehJoqGctiVI29o6fzry7u4=
github.com/stretchr/testify v1.9.0 h1:HtqpIVDClZ4nwg75+f6Lvsy/wHu+3BoSGCbBAcpTsTg=
github.com/stretchr/testify v1.9.0/go.mod h1:r2ic/lqez/lEtzL7wO/rwa5dbSLXVDPFyf8C91i36aY=
github.com/tidwall/pretty v1.0.0/go.mod h1:XNkn88O1ChpSDQmQeStsy+sBenx6DDtFZJxhVysOjyk=
github.com/weaviate/weaviate v1.26.0-rc.1 h1:p+8Cw4VfAbevtf90/sEN+43IHMBtdUc5ZH3r4NZKVR8=
github.com/weaviate/weaviate v1.26.0-rc.1/go.mod h1:o6nFEB4UozA+B2fnAWyF0HZqB2ab44pRhJHYwvxVyyA=
github.com/weaviate/weaviate-go-client/v4 v4.15.0 h1:+gSKFLpy6iXTDNtjgYFOuCj0RY7F+sICefKpZarnOuA=
github.com/weaviate/weaviate-go-client/v4 v4.15.0/go.mod h1:0aUsHXNfsdRfbO46t4Us5KUHqkhY+hlHevK6fBu72J4=
github.com/wk8/go-ordered-map/v2 v2.1.8 h1:5h/BUHu93oj4gIdvHHHGsScSTMijfx5PeYkE/fJgbpc=
github.com/wk8/go-ordered-map/v2 v2.1.8/go.mod h1:5nJHM5DyteebpVlHnWMV0rPz6Zp7+xBAnxjb1X5vnTw=
github.com/xdg-go/pbkdf2 v1.0.0/go.mod h1:jrpuAogTd400dnrH08LKmI/xc1MbPOebTwRqcT5RDeI=
github.com/xdg-go/scram v1.0.2/go.mod h1:1WAq6h33pAW+iRreB34OORO2Nf7qel3VV3fjBj+hCSs=
github.com/xdg-go/stringprep v1.0.2/go.mod h1:8F9zXuvzgwmyT5DUm4GUfZGDdT3W+LCvS6+da4O5kxM=
github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f/go.mod h1:N2zxlSyiKSe5eX1tZViRH5QA0qijqEDrYZiPEAiq3wU=
github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb h1:zGWFAtiMcyryUHoUjUJX0/lt1H2+i2Ka2n+D3DImSNo=
github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb/go.mod h1:N2zxlSyiKSe5eX1tZViRH5QA0qijqEDrYZiPEAiq3wU=
github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=
github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415/go.mod h1:GwrjFmJcFw6At/Gs6z4yjiIwzuJ1/+UwLxMQDVQXShQ=
github.com/xeipuuv/gojsonschema v1.2.0 h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=
github.com/xeipuuv/gojsonschema v1.2.0/go.mod h1:anYRn/JVcOK2ZgGU+IjEV4nwlhoK5sQluxsYJ78Id3Y=
github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d/go.mod h1:rHwXgn7JulP+udvsHwJoVG1YGAP6VLg4y9I5dyZdqmA=
go.mongodb.org/mongo-driver v1.7.3/go.mod h1:NqaYOwnXWr5Pm7AOpO5QFxKJ503nbMse/R79oO62zWg=
go.mongodb.org/mongo-driver v1.7.5/go.mod h1:VXEWRZ6URJIkUq2SCAyapmhH0ZLRBP+FT4xhp5Zvxng=
go.mongodb.org/mongo-driver v1.14.0 h1:P98w8egYRjYe3XDjxhYJagTokP/H6HzlsnojRgZRd80=
go.mongodb.org/mongo-driver v1.14.0/go.mod h1:Vzb0Mk/pa7e6cWw85R4F/endUC3u0U9jGcNU603k65c=
go.opencensus.io v0.24.0 h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=
go.opencensus.io v0.24.0/go.mod h1:vNK8G9p7aAivkbmorf4v+7Hgx+Zs0yY+0fOtgBfjQKo=
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 h1:A3SayB3rNyt+1S6qpI9mHPkeHTZbD7XILEqWnYZb2l0=
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0/go.mod h1:27iA5uvhuRNmalO+iEUdVn5ZMj2qy10Mm+XRIpRmyuU=
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 h1:Xs2Ncz0gNihqu9iosIZ5SkBbWo5T8JhhLJFMQL1qmLI=
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0/go.mod h1:vy+2G/6NvVMpwGX/NyLqcC41fxepnuKHk16E6IZUcJc=
go.opentelemetry.io/otel v1.26.0 h1:LQwgL5s/1W7YiiRwxf03QGnWLb2HW4pLiAhaA5cZXBs=
go.opentelemetry.io/otel v1.26.0/go.mod h1:UmLkJHUAidDval2EICqBMbnAd0/m2vmpf/dAM+fvFs4=
go.opentelemetry.io/otel/metric v1.26.0 h1:7S39CLuY5Jgg9CrnA9HHiEjGMF/X2VHvoXGgSllRz30=
go.opentelemetry.io/otel/metric v1.26.0/go.mod h1:SY+rHOI4cEawI9a7N1A4nIg/nTQXe1ccCNWYOJUrpX4=
go.opentelemetry.io/otel/sdk v1.26.0 h1:Y7bumHf5tAiDlRYFmGqetNcLaVUZmh4iYfmGxtmz7F8=
go.opentelemetry.io/otel/sdk v1.26.0/go.mod h1:0p8MXpqLeJ0pzcszQQN4F0S5FVjBLgypeGSngLsmirs=
go.opentelemetry.io/otel/trace v1.26.0 h1:1ieeAUb4y0TE26jUFrCIXKpTuVK7uJGN9/Z/2LP5sQA=
go.opentelemetry.io/otel/trace v1.26.0/go.mod h1:4iDxvGDQuUkHve82hJJ8UqrwswHYsZuWCBllGV2U2y0=
golang.org/x/crypto v0.0.0-20180904163835-0709b304e793/go.mod h1:6SG95UA2DQfeDnfUPMdvaQW0Q7yPrPDi9nlGo2tz2b4=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20190422162423-af44ce270edf/go.mod h1:WFFai1msRO1wXaEeE5yQxYXgSfI8pQAWXbQop6sCtWE=
golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/crypto v0.25.0 h1:ypSNr+bnYL2YhwoMt2zPxHFmbAN1KZs/njMG3hxUp30=
golang.org/x/crypto v0.25.0/go.mod h1:T+wALwcMOSE0kXgUAnPAHqTLW+XHgcELELW8VaDgm/M=
golang.org/x/exp v0.0.0-20190121172915-509febef88a4/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
golang.org/x/exp v0.0.0-20240318143956-a85f2c67cd81 h1:6R2FC06FonbXQ8pK11/PDFY6N6LWlf9KlzibaCapmqc=
golang.org/x/exp v0.0.0-20240318143956-a85f2c67cd81/go.mod h1:CQ1k9gNrJ50XIzaKCRR2hssIjF07kZFEiieALBM/ARQ=
golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3/go.mod h1:UVdnD1Gm6xHRNCYTkRU2/jEulfH38KcIWyp/GAMgvoE=
golang.org/x/lint v0.0.0-20190227174305-5b3e6a55c961/go.mod h1:wehouNa3lNwaWXcvxsM5YxQ5yQlVC4a0KAMCusXpPoU=
golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
golang.org/x/net v0.0.0-20180724234803-3673e40ba225/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20180826012351-8a410e7b638d/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190213061140-3a22650c66bd/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190311183353-d8887717615a/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20201110031124-69a78807bb2b/go.mod h1:sp8m0HH+o8qH0wwXwYZr8TS3Oi6o0r6Gce1SSxlDquU=
golang.org/x/net v0.0.0-20210421230115-4e50805a0758/go.mod h1:72T/g9IO56b78aLF+1Kcs5dz7/ng1VjMUvfKvpfy+jM=
golang.org/x/net v0.27.0 h1:5K3Njcw06/l2y9vpGCSdcxWOYHOUk3dVNGDXN+FvAys=
golang.org/x/net v0.27.0/go.mod h1:dDi0PyhWNoiUOrAS8uXv/vnScO4wnHQO4mj9fn/RytE=
golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be/go.mod h1:N/0e6XlmueqKjAGxoOufVs8QHGRruUQn6yWY3a++T0U=
golang.org/x/oauth2 v0.21.0 h1:tsimM75w1tF/uws5rbeHzIWxEqElMehnc+iW793zsZs=
golang.org/x/oauth2 v0.21.0/go.mod h1:XYTD2NtWslqkgxebSiOHnXEap4TF09sJSc7H1sXbhtI=
golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20181108010431-42b317875d0f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190412183630-56d357773e84/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.7.0 h1:YsImfSBoP9QPYL0xyKJPq0gcaJdG3rInoqxTWbfQu9M=
golang.org/x/sync v0.7.0/go.mod h1:Czt+wKu1gCyEFDUtn0jG5QVvpJ6rzVqr5aXyt9drQfk=
golang.org/x/sys v0.0.0-20180830151530-49385e6e1522/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190403152447-81d4e9dc473e/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190419153524-e8e3143a4f4a/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190422165155-953cdadca894/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190531175056-4c3a928424d2/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210420072515-93ed5bcd2bfe/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.22.0 h1:RI27ohtqKCnwULzJLqkv897zojh5/DwS/ENaMzUOaWI=
golang.org/x/sys v0.22.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.5/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
golang.org/x/text v0.16.0 h1:a94ExnEXNtEwYLGJSIUxnWoxoRz/ZcCsV63ROupILh4=
golang.org/x/text v0.16.0/go.mod h1:GhwF1Be+LQoKShO3cGOHzqOgRrGaYc9AvblQOmPVHnI=
golang.org/x/time v0.5.0 h1:o7cqy6amK/52YcAKIPlM3a+Fpj35zvRj2TP+e1xFSfk=
golang.org/x/time v0.5.0/go.mod h1:3BpzKBy/shNhVucY/MWOyx10tF3SFh9QdLuxbVysPQM=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190114222345-bf090417da8b/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190226205152-f727befe758c/go.mod h1:9Yl7xja0Znq3iFh3HoIrodX9oNMXvdceNzlUR8zjMvY=
golang.org/x/tools v0.0.0-20190311212946-11955173bddd/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190329151228-23e29df326fe/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190416151739-9c9e1878f421/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190420181800-aa740d480789/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135/go.mod h1:RgjU9mgBXZiqYHBnxXauZ1Gv1EHHAz9KjViQ78xBX0Q=
golang.org/x/tools v0.0.0-20190531172133-b3315ee88b7d/go.mod h1:/rFqwRUd4F7ZHNgwSSTFct+R/Kf4OFW1sUzUTQQTgfc=
golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
google.golang.org/api v0.188.0 h1:51y8fJ/b1AaaBRJr4yWm96fPcuxSo0JcegXE3DaHQHw=
google.golang.org/api v0.188.0/go.mod h1:VR0d+2SIiWOYG3r/jdm7adPW9hI2aRv9ETOSCQ9Beag=
google.golang.org/appengine v1.1.0/go.mod h1:EbEs0AVv82hx2wNQdGPgUI5lhzA/G0D9YwlJXL52JkM=
google.golang.org/appengine v1.4.0/go.mod h1:xpcJRLb0r/rnEns0DIKYYv+WjYCduHsrkT7/EB5XEv4=
google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8/go.mod h1:JiN7NxoALGmiZfu7CAH4rXhgtRTLTxftemlI0sWmxmc=
google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55/go.mod h1:DMBHOl98Agz4BDEuKkezgsaosCRResVns1a3J2ZsMNc=
google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013/go.mod h1:NbSheEEYHJ7i3ixzK3sjbqSGDJWnxyFXZblF3eUsNvo=
google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094 h1:0+ozOGcrp+Y8Aq8TLNN2Aliibms5LEzsq99ZZmAGYm0=
google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094/go.mod h1:fJ/e3If/Q67Mj99hin0hMhiNyCRmt6BQ2aWIJshUSJw=
google.golang.org/genproto/googleapis/rpc v0.0.0-20240708141625-4ad9e859172b h1:04+jVzTs2XBnOZcPsLnmrTGqltqJbZQ1Ey26hjYdQQ0=
google.golang.org/genproto/googleapis/rpc v0.0.0-20240708141625-4ad9e859172b/go.mod h1:Ue6ibwXGpU+dqIcODieyLOcgj7z8+IcskoNIgZxtrFY=
google.golang.org/grpc v1.19.0/go.mod h1:mqu4LbDTu4XGKhr4mRzUsmM4RtVoemTSY81AxZiDr8c=
google.golang.org/grpc v1.23.0/go.mod h1:Y5yQAOtifL1yxbo5wqy6BxZv8vAUGQwXBOALyacEbxg=
google.golang.org/grpc v1.25.1/go.mod h1:c3i+UQWmh7LiEpx4sFZnkU36qjEYZ0imhYfXVyQciAY=
google.golang.org/grpc v1.27.0/go.mod h1:qbnxyOmOxrQa7FizSgH+ReBfzJrCY1pSN7KXBS8abTk=
google.golang.org/grpc v1.33.2/go.mod h1:JMHMWHQWaTccqQQlmk3MJZS+GWXOdAesneDmEnv2fbc=
google.golang.org/grpc v1.65.0 h1:bs/cUb4lp1G5iImFFd3u5ixQzweKizoZJAwBNLR42lc=
google.golang.org/grpc v1.65.0/go.mod h1:WgYC2ypjlB0EiQi6wdKixMqukr6lBc0Vo+oOgjrM5ZQ=
google.golang.org/protobuf v0.0.0-20200109180630-ec00e32a8dfd/go.mod h1:DFci5gLYBciE7Vtevhsrf46CRTquxDuWsQurQQe4oz8=
google.golang.org/protobuf v0.0.0-20200221191635-4d8936d0db64/go.mod h1:kwYJMbMJ01Woi6D6+Kah6886xMZcty6N08ah7+eCXa0=
google.golang.org/protobuf v0.0.0-20200228230310-ab0ca4ff8a60/go.mod h1:cfTl7dwQJ+fmap5saPgwCLgHXTUD7jkjRqWcaiX5VyM=
google.golang.org/protobuf v1.20.1-0.20200309200217-e05f789c0967/go.mod h1:A+miEFZTKqfCUM6K7xSMQL9OKL/b6hQv+e19PK+JZNE=
google.golang.org/protobuf v1.21.0/go.mod h1:47Nbq4nVaFHyn7ilMalzfO3qCViNmqZ2kzikPIcrTAo=
google.golang.org/protobuf v1.22.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.1-0.20200526195155-81db48ad09cc/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.25.0/go.mod h1:9JNX74DMeImyA3h4bdi1ymwjUzf21/xIlbajtzgsN7c=
google.golang.org/protobuf v1.34.2 h1:6xV6lTsCfpGD21XK49h7MhtcApnLqkfYgPcdHftf6hg=
google.golang.org/protobuf v1.34.2/go.mod h1:qYOHts0dSfpeUzUFpOMr/WGzszTmLH+DiWniOlNbLDw=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
gopkg.in/errgo.v2 v2.1.0/go.mod h1:hNsd1EY+bozCKY1Ytp96fpM3vjJbqLJn88ws8XvfDNI=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.4.0 h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=
gopkg.in/yaml.v2 v2.4.0/go.mod h1:RDklbk79AGWmwhnvt/jBztapEOGDOx6ZbXqjP6csGnQ=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20200605160147-a5ece683394c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
honnef.co/go/tools v0.0.0-20190102054323-c2f93a96b099/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
```

### ragserver/ragserver-genkit/json.go

```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

// readRequestJSON expects req to have a JSON content type with a body that
// contains a JSON-encoded value complying with the underlying type of target.
// It populates target, or returns an error.
func readRequestJSON(req *http.Request, target any) error {
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	if mediaType != "application/json" {
		return fmt.Errorf("expect application/json Content-Type, got %s", mediaType)
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(target)
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
```

### ragserver/ragserver-genkit/main.go

```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command ragserver is an HTTP server that implements RAG (Retrieval
// Augmented Generation) using the Gemini model and Weaviate, which
// are accessed using the Genkit package. See the accompanying README file for
// additional details.
package main

import (
	"cmp"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/plugins/googleai"
	"github.com/firebase/genkit/go/plugins/weaviate"
)

const generativeModelName = "gemini-1.5-flash"
const embeddingModelName = "text-embedding-004"

// This is a standard Go HTTP server. Server state is in the ragServer struct.
// The `main` function connects to the required services (Weaviate and Google
// AI), initializes the server state and registers HTTP handlers.
func main() {
	ctx := context.Background()
	err := googleai.Init(ctx, &googleai.Config{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	wvConfig := &weaviate.ClientConfig{
		Scheme: "http",
		Addr:   "localhost:" + cmp.Or(os.Getenv("WVPORT"), "9035"),
	}
	_, err = weaviate.Init(ctx, wvConfig)
	if err != nil {
		log.Fatal(err)
	}

	classConfig := &weaviate.ClassConfig{
		Class:    "Document",
		Embedder: googleai.Embedder(embeddingModelName),
	}
	indexer, retriever, err := weaviate.DefineIndexerAndRetriever(ctx, *classConfig)
	if err != nil {
		log.Fatal(err)
	}

	model := googleai.Model(generativeModelName)
	if model == nil {
		log.Fatal("unable to set up gemini-1.5-flash model")
	}

	server := &ragServer{
		ctx:       ctx,
		indexer:   indexer,
		retriever: retriever,
		model:     model,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /add/", server.addDocumentsHandler)
	mux.HandleFunc("POST /query/", server.queryHandler)

	port := cmp.Or(os.Getenv("SERVERPORT"), "9020")
	address := "localhost:" + port
	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

type ragServer struct {
	ctx       context.Context
	indexer   ai.Indexer
	retriever ai.Retriever
	model     ai.Model
}

func (rs *ragServer) addDocumentsHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type document struct {
		Text string
	}
	type addRequest struct {
		Documents []document
	}
	ar := &addRequest{}
	err := readRequestJSON(req, ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert request documents into Weaviate documents used for embedding.
	var wvDocs []*ai.Document
	for _, doc := range ar.Documents {
		wvDocs = append(wvDocs, ai.DocumentFromText(doc.Text, nil))
	}

	// Index the requested documents.
	err = ai.Index(rs.ctx, rs.indexer, ai.WithIndexerDocs(wvDocs...))
	if err != nil {
		http.Error(w, fmt.Errorf("indexing: %w", err).Error(), http.StatusInternalServerError)
		return
	}
}

func (rs *ragServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type queryRequest struct {
		Content string
	}
	qr := &queryRequest{}
	err := readRequestJSON(req, qr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the most similar documents using the retriever.
	resp, err := ai.Retrieve(rs.ctx,
		rs.retriever,
		ai.WithRetrieverDoc(ai.DocumentFromText(qr.Content, nil)),
		ai.WithRetrieverOpts(&weaviate.RetrieverOptions{
			Count: 3,
		}))
	if err != nil {
		http.Error(w, fmt.Errorf("retrieval: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	var docsContents []string
	for _, d := range resp.Documents {
		docsContents = append(docsContents, d.Content[0].Text)
	}

	// Create a RAG query for the LLM with the most relevant documents as
	// context.
	ragQuery := fmt.Sprintf(ragTemplateStr, qr.Content, strings.Join(docsContents, "\n"))
	genResp, err := ai.Generate(rs.ctx, rs.model, ai.WithTextPrompt(ragQuery))
	if err != nil {
		log.Printf("calling generative model: %v", err.Error())
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	if len(genResp.Candidates) != 1 {
		log.Printf("got %v candidates, expected 1", len(genResp.Candidates))
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	renderJSON(w, genResp.Text())
}

const ragTemplateStr = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Question:
%s

Context:
%s
`
```

### ragserver/ragserver-langchaingo/go.mod

```
module golang.org/x/example/ragserver/ragserver-langchaingo

go 1.23.0

require github.com/tmc/langchaingo v0.1.12

require (
	cloud.google.com/go v0.113.0 // indirect
	cloud.google.com/go/ai v0.6.0 // indirect
	cloud.google.com/go/aiplatform v1.67.0 // indirect
	cloud.google.com/go/auth v0.4.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	cloud.google.com/go/iam v1.1.7 // indirect
	cloud.google.com/go/longrunning v0.5.7 // indirect
	cloud.google.com/go/vertexai v0.10.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.21.2 // indirect
	github.com/go-openapi/errors v0.22.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/strfmt v0.21.3 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-openapi/validate v0.21.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/generative-ai-go v0.14.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.4 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkoukk/tiktoken-go v0.1.6 // indirect
	github.com/weaviate/weaviate v1.24.1 // indirect
	github.com/weaviate/weaviate-go-client/v4 v4.13.1 // indirect
	go.mongodb.org/mongo-driver v1.14.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/api v0.180.0 // indirect
	google.golang.org/genproto v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.67.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

### ragserver/ragserver-langchaingo/go.sum

```
cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
cloud.google.com/go v0.113.0 h1:g3C70mn3lWfckKBiCVsAshabrDg01pQ0pnX1MNtnMkA=
cloud.google.com/go v0.113.0/go.mod h1:glEqlogERKYeePz6ZdkcLJ28Q2I6aERgDDErBg9GzO8=
cloud.google.com/go/ai v0.6.0 h1:QWjb2UoaM15e51IMeLuIUFyWxooKOKDb66Mk47zZ2/g=
cloud.google.com/go/ai v0.6.0/go.mod h1:6/mrRq6aJdK7MZH76ZvcMpESiAiha5aRvurmroiOrgI=
cloud.google.com/go/aiplatform v1.67.0 h1:YWeqD4BjYwrmY4fa+isGcw0P81lJ3dKVxbWxdBchoiU=
cloud.google.com/go/aiplatform v1.67.0/go.mod h1:s/sJ6btBEr6bKnrNWdK9ZgHCvwbZNdP90b3DDtxxw+Y=
cloud.google.com/go/auth v0.4.1 h1:Z7YNIhlWRtrnKlZke7z3GMqzvuYzdc2z98F9D1NV5Hg=
cloud.google.com/go/auth v0.4.1/go.mod h1:QVBuVEKpCn4Zp58hzRGvL0tjRGU0YqdRTdCHM1IHnro=
cloud.google.com/go/auth/oauth2adapt v0.2.2 h1:+TTV8aXpjeChS9M+aTtN/TjdQnzJvmzKFt//oWu7HX4=
cloud.google.com/go/auth/oauth2adapt v0.2.2/go.mod h1:wcYjgpZI9+Yu7LyYBg4pqSiaRkfEK3GQcpb7C/uyF1Q=
cloud.google.com/go/compute/metadata v0.5.0 h1:Zr0eK8JbFv6+Wi4ilXAR8FJ3wyNdpxHKJNPos6LTZOY=
cloud.google.com/go/compute/metadata v0.5.0/go.mod h1:aHnloV2TPI38yx4s9+wAZhHykWvVCfu7hQbF+9CWoiY=
cloud.google.com/go/iam v1.1.7 h1:z4VHOhwKLF/+UYXAJDFwGtNF0b6gjsW1Pk9Ml0U/IoM=
cloud.google.com/go/iam v1.1.7/go.mod h1:J4PMPg8TtyurAUvSmPj8FF3EDgY1SPRZxcUGrn7WXGA=
cloud.google.com/go/longrunning v0.5.7 h1:WLbHekDbjK1fVFD3ibpFFVoyizlLRl73I7YKuAKilhU=
cloud.google.com/go/longrunning v0.5.7/go.mod h1:8GClkudohy1Fxm3owmBGid8W0pSgodEMwEAztp38Xng=
cloud.google.com/go/vertexai v0.10.0 h1:k157bLrtyajGtAAZnqdEn8lwFlUTG3BgHc7kvWbP/3s=
cloud.google.com/go/vertexai v0.10.0/go.mod h1:w/Zb22QvOVvxx5CGM4fPzH3WA6gwUkId9juA7pigzFI=
dario.cat/mergo v1.0.0 h1:AGCNq9Evsj31mOgNPcLyXc+4PNABt905YmuqPYYpBWk=
dario.cat/mergo v1.0.0/go.mod h1:uNxQE+84aUszobStD9th8a29P2fMDhsBdgRYvZOxGmk=
github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 h1:L/gRVlceqvL25UVaW/CKtUDjefjrs0SPonmDGUVOYP0=
github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161/go.mod h1:xomTg63KZ2rFqZQzSB4Vz2SUXa1BpHTVz9L5PTmPC4E=
github.com/BurntSushi/toml v0.3.1/go.mod h1:xHWCNGjB5oqiDr8zfno3MHue2Ht5sIBksp03qcyfWMU=
github.com/Masterminds/goutils v1.1.1 h1:5nUrii3FMTL5diU80unEVvNevw1nH4+ZV4DSLVJLSYI=
github.com/Masterminds/goutils v1.1.1/go.mod h1:8cTjp+g8YejhMuvIA5y2vz3BpJxksy863GQaJW2MFNU=
github.com/Masterminds/semver v1.5.0 h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=
github.com/Masterminds/semver/v3 v3.2.0 h1:3MEsd0SM6jqZojhjLWWeBY+Kcjy9i6MQAeY7YgDP83g=
github.com/Masterminds/semver/v3 v3.2.0/go.mod h1:qvl/7zhW3nngYb5+80sSMF+FG2BjYrf8m9wsX0PNOMQ=
github.com/Masterminds/sprig/v3 v3.2.3 h1:eL2fZNezLomi0uOLqjQoN6BfsDD+fyLtgbJMAj9n6YA=
github.com/Masterminds/sprig/v3 v3.2.3/go.mod h1:rXcFaZ2zZbLRJv/xSysmlgIM1u11eBaRMhvYXJNkGuM=
github.com/Microsoft/go-winio v0.6.1 h1:9/kr64B9VUZrLm5YYwbGtUJnMgqWVOdUAXu6Migciow=
github.com/Microsoft/go-winio v0.6.1/go.mod h1:LRdKpFKfdobln8UmuiYcKPot9D2v6svN5+sAH+4kjUM=
github.com/Microsoft/hcsshim v0.11.4 h1:68vKo2VN8DE9AdN4tnkWnmdhqdbpUFM8OF3Airm7fz8=
github.com/Microsoft/hcsshim v0.11.4/go.mod h1:smjE4dvqPX9Zldna+t5FG3rnoHhaB7QYxPRqGcpAD9w=
github.com/PuerkitoBio/purell v1.1.1 h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=
github.com/PuerkitoBio/purell v1.1.1/go.mod h1:c11w/QuzBsJSee3cPx9rAFu61PvFxuPbtSwDGJws/X0=
github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=
github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578/go.mod h1:uGdkoq3SwY9Y+13GIhn11/XLaGBb4BfwItxLd5jeuXE=
github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef/go.mod h1:WaHUgvxTVq04UNunO+XhnAqY/wQc+bxr74GqbsZ/Jqw=
github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 h1:DklsrG3dyBCFEj5IhUbnKptjxatkF07cF2ak3yi77so=
github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2/go.mod h1:WaHUgvxTVq04UNunO+XhnAqY/wQc+bxr74GqbsZ/Jqw=
github.com/cenkalti/backoff v2.2.1+incompatible h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=
github.com/cenkalti/backoff/v4 v4.2.1 h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=
github.com/cenkalti/backoff/v4 v4.2.1/go.mod h1:Y3VNntkOUPxTVeUxJ/G5vcM//AlwfmyYozVcomhLiZE=
github.com/census-instrumentation/opencensus-proto v0.2.1/go.mod h1:f6KPmirojxKA12rnyqOA5BBL4O983OfeGPqjHWSTneU=
github.com/client9/misspell v0.3.4/go.mod h1:qj6jICC3Q7zFZvVWo7KLAzC3yx5G7kyvSDkc90ppPyw=
github.com/cncf/udpa/go v0.0.0-20191209042840-269d4d468f6f/go.mod h1:M8M6+tZqaGXZJjfX53e64911xZQV5JYwmTeXPW+k8Sc=
github.com/containerd/containerd v1.7.15 h1:afEHXdil9iAm03BmhjzKyXnnEBtjaLJefdU7DV0IFes=
github.com/containerd/containerd v1.7.15/go.mod h1:ISzRRTMF8EXNpJlTzyr2XMhN+j9K302C21/+cr3kUnY=
github.com/containerd/log v0.1.0 h1:TCJt7ioM2cr/tfR8GPbGf9/VRAX8D2B4PjzCpfX540I=
github.com/containerd/log v0.1.0/go.mod h1:VRRf09a7mHDIRezVKTRCrOq78v577GXq3bSa3EhrzVo=
github.com/cpuguy83/dockercfg v0.3.1 h1:/FpZ+JaygUR/lZP2NlFI2DVfrOEMAIKP5wWEJdoYe9E=
github.com/cpuguy83/dockercfg v0.3.1/go.mod h1:sugsbF4//dDlL/i+S+rtpIWp+5h0BHJHfjj5/jFyUJc=
github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/distribution/reference v0.5.0 h1:/FUIFXtfc/x2gpa5/VGfiGLuOIdYa1t65IKK2OFGvA0=
github.com/distribution/reference v0.5.0/go.mod h1:BbU0aIcezP1/5jX/8MP0YiH4SdvB5Y4f/wlDRiLyi3E=
github.com/dlclark/regexp2 v1.10.0 h1:+/GIL799phkJqYW+3YbOd8LCcbHzT0Pbo8zl70MHsq0=
github.com/dlclark/regexp2 v1.10.0/go.mod h1:DHkYz0B9wPfa6wondMfaivmHpzrQ3v9q8cnmRbL6yW8=
github.com/docker/docker v25.0.5+incompatible h1:UmQydMduGkrD5nQde1mecF/YnSbTOaPeFIeP5C4W+DE=
github.com/docker/docker v25.0.5+incompatible/go.mod h1:eEKB0N0r5NX/I1kEveEz05bcu8tLC/8azJZsviup8Sk=
github.com/docker/go-connections v0.5.0 h1:USnMq7hx7gwdVZq1L49hLXaFtUdTADjXGp+uj1Br63c=
github.com/docker/go-connections v0.5.0/go.mod h1:ov60Kzw0kKElRwhNs9UlUHAE/F9Fe6GLaXnqyDdmEXc=
github.com/docker/go-units v0.5.0 h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=
github.com/docker/go-units v0.5.0/go.mod h1:fgPhTUdO+D/Jk86RDLlptpiXQzgHJF7gydDDbaIK4Dk=
github.com/dustin/go-humanize v1.0.1 h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=
github.com/dustin/go-humanize v1.0.1/go.mod h1:Mu1zIs6XwVuF/gI1OepvI0qD18qycQx+mFykh5fBlto=
github.com/envoyproxy/go-control-plane v0.9.0/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
github.com/envoyproxy/go-control-plane v0.9.1-0.20191026205805-5f8ba28d4473/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
github.com/envoyproxy/go-control-plane v0.9.4/go.mod h1:6rpuAdCZL397s3pYoYcLgu1mIlRU8Am5FuJP05cCM98=
github.com/envoyproxy/protoc-gen-validate v0.1.0/go.mod h1:iSmxcyjqTsJpI2R4NaDN7+kN2VEUnK/pcBlmesArF7c=
github.com/felixge/httpsnoop v1.0.4 h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=
github.com/felixge/httpsnoop v1.0.4/go.mod h1:m8KPJKqk1gH5J9DgRY2ASl2lWCfGKXixSwevea8zH2U=
github.com/go-logr/logr v1.2.2/go.mod h1:jdQByPbusPIv2/zmleS9BjJVeZ6kBagPoEUsqbVz/1A=
github.com/go-logr/logr v1.4.1 h1:pKouT5E8xu9zeFC39JXRDukb6JFQPXM5p5I91188VAQ=
github.com/go-logr/logr v1.4.1/go.mod h1:9T104GzyrTigFIr8wt5mBrctHMim0Nb2HLGrmQ40KvY=
github.com/go-logr/stdr v1.2.2 h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=
github.com/go-logr/stdr v1.2.2/go.mod h1:mMo/vtBO5dYbehREoey6XUKy/eSumjCCveDpRre4VKE=
github.com/go-ole/go-ole v1.2.6 h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=
github.com/go-ole/go-ole v1.2.6/go.mod h1:pprOEPIfldk/42T2oK7lQ4v4JSDwmV0As9GaiUsvbm0=
github.com/go-openapi/analysis v0.21.2 h1:hXFrOYFHUAMQdu6zwAiKKJHJQ8kqZs1ux/ru1P1wLJU=
github.com/go-openapi/analysis v0.21.2/go.mod h1:HZwRk4RRisyG8vx2Oe6aqeSQcoxRp47Xkp3+K6q+LdY=
github.com/go-openapi/errors v0.19.8/go.mod h1:cM//ZKUKyO06HSwqAelJ5NsEMMcpa6VpXe8DOa1Mi1M=
github.com/go-openapi/errors v0.19.9/go.mod h1:cM//ZKUKyO06HSwqAelJ5NsEMMcpa6VpXe8DOa1Mi1M=
github.com/go-openapi/errors v0.20.2/go.mod h1:cM//ZKUKyO06HSwqAelJ5NsEMMcpa6VpXe8DOa1Mi1M=
github.com/go-openapi/errors v0.22.0 h1:c4xY/OLxUBSTiepAg3j/MHuAv5mJhnf53LLMWFB+u/w=
github.com/go-openapi/errors v0.22.0/go.mod h1:J3DmZScxCDufmIMsdOuDHxJbdOGC0xtUynjIx092vXE=
github.com/go-openapi/jsonpointer v0.19.3/go.mod h1:Pl9vOtqEWErmShwVjC8pYs9cog34VGT37dQOVbmoatg=
github.com/go-openapi/jsonpointer v0.19.5/go.mod h1:Pl9vOtqEWErmShwVjC8pYs9cog34VGT37dQOVbmoatg=
github.com/go-openapi/jsonpointer v0.19.6 h1:eCs3fxoIi3Wh6vtgmLTOjdhSpiqphQ+DaPn38N2ZdrE=
github.com/go-openapi/jsonpointer v0.19.6/go.mod h1:osyAmYz/mB/C3I+WsTTSgw1ONzaLJoLCyoi6/zppojs=
github.com/go-openapi/jsonreference v0.19.6 h1:UBIxjkht+AWIgYzCDSv2GN+E/togfwXUJFRTWhl2Jjs=
github.com/go-openapi/jsonreference v0.19.6/go.mod h1:diGHMEHg2IqXZGKxqyvWdfWU/aim5Dprw5bqpKkTvns=
github.com/go-openapi/loads v0.21.1 h1:Wb3nVZpdEzDTcly8S4HMkey6fjARRzb7iEaySimlDW0=
github.com/go-openapi/loads v0.21.1/go.mod h1:/DtAMXXneXFjbQMGEtbamCZb+4x7eGwkvZCvBmwUG+g=
github.com/go-openapi/spec v0.20.4 h1:O8hJrt0UMnhHcluhIdUgCLRWyM2x7QkBXRvOs7m+O1M=
github.com/go-openapi/spec v0.20.4/go.mod h1:faYFR1CvsJZ0mNsmsphTMSoRrNV3TEDoAM7FOEWeq8I=
github.com/go-openapi/strfmt v0.21.0/go.mod h1:ZRQ409bWMj+SOgXofQAGTIo2Ebu72Gs+WaRADcS5iNg=
github.com/go-openapi/strfmt v0.21.1/go.mod h1:I/XVKeLc5+MM5oPNN7P6urMOpuLXEcNrCX/rPGuWb0k=
github.com/go-openapi/strfmt v0.21.3 h1:xwhj5X6CjXEZZHMWy1zKJxvW9AfHC9pkyUjLvHtKG7o=
github.com/go-openapi/strfmt v0.21.3/go.mod h1:k+RzNO0Da+k3FrrynSNN8F7n/peCmQQqbbXjtDfvmGg=
github.com/go-openapi/swag v0.19.5/go.mod h1:POnQmlKehdgb5mhVOsnJFsivZCEZ/vjK9gh66Z9tfKk=
github.com/go-openapi/swag v0.19.15/go.mod h1:QYRuS/SOXUCsnplDa677K7+DxSOj6IPNl/eQntq43wQ=
github.com/go-openapi/swag v0.21.1/go.mod h1:QYRuS/SOXUCsnplDa677K7+DxSOj6IPNl/eQntq43wQ=
github.com/go-openapi/swag v0.22.3/go.mod h1:UzaqsxGiab7freDnrUUra0MwWfN/q7tE4j+VcZ0yl14=
github.com/go-openapi/swag v0.22.4 h1:QLMzNJnMGPRNDCbySlcj1x01tzU8/9LTTL9hZZZogBU=
github.com/go-openapi/swag v0.22.4/go.mod h1:UzaqsxGiab7freDnrUUra0MwWfN/q7tE4j+VcZ0yl14=
github.com/go-openapi/validate v0.21.0 h1:+Wqk39yKOhfpLqNLEC0/eViCkzM5FVXVqrvt526+wcI=
github.com/go-openapi/validate v0.21.0/go.mod h1:rjnrwK57VJ7A8xqfpAOEKRH8yQSGUriMu5/zuPSQ1hg=
github.com/go-stack/stack v1.8.0/go.mod h1:v0f6uXyyMGvRgIKkXu+yp6POWl0qKG85gN/melR3HDY=
github.com/gobuffalo/attrs v0.0.0-20190224210810-a9411de4debd/go.mod h1:4duuawTqi2wkkpB4ePgWMaai6/Kc6WEz83bhFwpHzj0=
github.com/gobuffalo/depgen v0.0.0-20190329151759-d478694a28d3/go.mod h1:3STtPUQYuzV0gBVOY3vy6CfMm/ljR4pABfrTeHNLHUY=
github.com/gobuffalo/depgen v0.1.0/go.mod h1:+ifsuy7fhi15RWncXQQKjWS9JPkdah5sZvtHc2RXGlg=
github.com/gobuffalo/envy v1.6.15/go.mod h1:n7DRkBerg/aorDM8kbduw5dN3oXGswK5liaSCx4T5NI=
github.com/gobuffalo/envy v1.7.0/go.mod h1:n7DRkBerg/aorDM8kbduw5dN3oXGswK5liaSCx4T5NI=
github.com/gobuffalo/flect v0.1.0/go.mod h1:d2ehjJqGOH/Kjqcoz+F7jHTBbmDb38yXA598Hb50EGs=
github.com/gobuffalo/flect v0.1.1/go.mod h1:8JCgGVbRjJhVgD6399mQr4fx5rRfGKVzFjbj6RE/9UI=
github.com/gobuffalo/flect v0.1.3/go.mod h1:8JCgGVbRjJhVgD6399mQr4fx5rRfGKVzFjbj6RE/9UI=
github.com/gobuffalo/genny v0.0.0-20190329151137-27723ad26ef9/go.mod h1:rWs4Z12d1Zbf19rlsn0nurr75KqhYp52EAGGxTbBhNk=
github.com/gobuffalo/genny v0.0.0-20190403191548-3ca520ef0d9e/go.mod h1:80lIj3kVJWwOrXWWMRzzdhW3DsrdjILVil/SFKBzF28=
github.com/gobuffalo/genny v0.1.0/go.mod h1:XidbUqzak3lHdS//TPu2OgiFB+51Ur5f7CSnXZ/JDvo=
github.com/gobuffalo/genny v0.1.1/go.mod h1:5TExbEyY48pfunL4QSXxlDOmdsD44RRq4mVZ0Ex28Xk=
github.com/gobuffalo/gitgen v0.0.0-20190315122116-cc086187d211/go.mod h1:vEHJk/E9DmhejeLeNt7UVvlSGv3ziL+djtTr3yyzcOw=
github.com/gobuffalo/gogen v0.0.0-20190315121717-8f38393713f5/go.mod h1:V9QVDIxsgKNZs6L2IYiGR8datgMhB577vzTDqypH360=
github.com/gobuffalo/gogen v0.1.0/go.mod h1:8NTelM5qd8RZ15VjQTFkAW6qOMx5wBbW4dSCS3BY8gg=
github.com/gobuffalo/gogen v0.1.1/go.mod h1:y8iBtmHmGc4qa3urIyo1shvOD8JftTtfcKi+71xfDNE=
github.com/gobuffalo/logger v0.0.0-20190315122211-86e12af44bc2/go.mod h1:QdxcLw541hSGtBnhUc4gaNIXRjiDppFGaDqzbrBd3v8=
github.com/gobuffalo/mapi v1.0.1/go.mod h1:4VAGh89y6rVOvm5A8fKFxYG+wIW6LO1FMTG9hnKStFc=
github.com/gobuffalo/mapi v1.0.2/go.mod h1:4VAGh89y6rVOvm5A8fKFxYG+wIW6LO1FMTG9hnKStFc=
github.com/gobuffalo/packd v0.0.0-20190315124812-a385830c7fc0/go.mod h1:M2Juc+hhDXf/PnmBANFCqx4DM3wRbgDvnVWeG2RIxq4=
github.com/gobuffalo/packd v0.1.0/go.mod h1:M2Juc+hhDXf/PnmBANFCqx4DM3wRbgDvnVWeG2RIxq4=
github.com/gobuffalo/packr/v2 v2.0.9/go.mod h1:emmyGweYTm6Kdper+iywB6YK5YzuKchGtJQZ0Odn4pQ=
github.com/gobuffalo/packr/v2 v2.2.0/go.mod h1:CaAwI0GPIAv+5wKLtv8Afwl+Cm78K/I/VCm/3ptBN+0=
github.com/gobuffalo/syncx v0.0.0-20190224160051-33c29581e754/go.mod h1:HhnNqWY95UYwwW3uSASeV7vtgYkT2t16hJgV3AEPUpw=
github.com/gogo/protobuf v1.3.2 h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=
github.com/gogo/protobuf v1.3.2/go.mod h1:P1XiOD3dCwIKUDQYPy72D8LYyHL2YPYrpS2s69NZV8Q=
github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b/go.mod h1:SBH7ygxi8pfUlaOkMMuAQtPIUF8ecWP5IEl/CR7VP2Q=
github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=
github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da/go.mod h1:cIg4eruTrX1D+g88fzRXU5OdNfaM+9IcxsU14FzY7Hc=
github.com/golang/mock v1.1.1/go.mod h1:oTYuIxOrZwtPieC+H1uAHpcLFnEyAGVDL/k47Jfbm0A=
github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.3.2/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.4.0-rc.1/go.mod h1:ceaxUfeHdC40wWswd/P6IGgMaK3YpKi5j83Wpe3EHw8=
github.com/golang/protobuf v1.4.0-rc.1.0.20200221234624-67d41d38c208/go.mod h1:xKAWHe0F5eneWXFV3EuXVDTCmh+JuBKY0li0aMyXATA=
github.com/golang/protobuf v1.4.0-rc.2/go.mod h1:LlEzMj4AhA7rCAGe4KMBDvJI+AwstrUpVNzEA03Pprs=
github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0/go.mod h1:WU3c8KckQ9AFe+yFwt9sWVRKCVIyN9cPHBJSNnbL67w=
github.com/golang/protobuf v1.4.0/go.mod h1:jodUvKwWbYaEsadDk5Fwe5c77LiNKVO9IDvqG2KuDX0=
github.com/golang/protobuf v1.4.1/go.mod h1:U8fpvMrcmy5pZrNK1lt4xCsGvpyWQ/VVv6QDs8UjoX8=
github.com/golang/protobuf v1.4.3/go.mod h1:oDoupMAO8OvCJWAcko0GGGIgR6R6ocIYbsSw735rRwI=
github.com/golang/protobuf v1.5.4 h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=
github.com/golang/protobuf v1.5.4/go.mod h1:lnTiLA8Wa4RWRcIUkrtSVa5nRhsEGBg48fD6rSs7xps=
github.com/golang/snappy v0.0.1/go.mod h1:/XxbfmMg8lxefKM7IXC3fBNl/7bRcc72aCRzEWrmP2Q=
github.com/google/generative-ai-go v0.14.0 h1:2GwFKXui9LmG+PukQwYk9KpJUIemmQ9NJ46BV9VIw38=
github.com/google/generative-ai-go v0.14.0/go.mod h1:hOzbW3cB5hRV2x05McOwJS4GsqSluYwejjk5tSfb6YY=
github.com/google/go-cmp v0.2.0/go.mod h1:oXzfMopK8JAjlY9xF4vHSVASa0yLyX7SntLO5aqRK0M=
github.com/google/go-cmp v0.3.0/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.3.1/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.4.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.2/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.3/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.6.0 h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=
github.com/google/go-cmp v0.6.0/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
github.com/google/s2a-go v0.1.7 h1:60BLSyTrOV4/haCDW4zb1guZItoSq8foHCXrAnjBo/o=
github.com/google/s2a-go v0.1.7/go.mod h1:50CgR4k1jNlWBu4UfS4AcfhVe1r6pdZPygJ3R8F0Qdw=
github.com/google/uuid v1.1.1/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/google/uuid v1.1.2/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/googleapis/enterprise-certificate-proxy v0.3.2 h1:Vie5ybvEvT75RniqhfFxPRy3Bf7vr3h0cechB90XaQs=
github.com/googleapis/enterprise-certificate-proxy v0.3.2/go.mod h1:VLSiSSBs/ksPL8kq3OBOQ6WRI2QnaFynd1DCjZ62+V0=
github.com/googleapis/gax-go/v2 v2.12.4 h1:9gWcmF85Wvq4ryPFvGFaOgPIs1AQX0d0bcbGw4Z96qg=
github.com/googleapis/gax-go/v2 v2.12.4/go.mod h1:KYEYLorsnIGDi/rPC8b5TdlB9kbKoFubselGIoBMCwI=
github.com/goph/emperror v0.17.2 h1:yLapQcmEsO0ipe9p5TaN22djm3OFV/TfM/fcYP0/J18=
github.com/goph/emperror v0.17.2/go.mod h1:+ZbQ+fUNO/6FNiUo0ujtMjhgad9Xa6fQL9KhH4LNHic=
github.com/huandu/xstrings v1.3.3 h1:/Gcsuc1x8JVbJ9/rlye4xZnVAbEkGauT8lbebqcQws4=
github.com/huandu/xstrings v1.3.3/go.mod h1:y5/lhBue+AyNmUVz9RLU9xbLR0o4KIIExikq4ovT0aE=
github.com/imdario/mergo v0.3.13 h1:lFzP57bqS/wsqKssCGmtLAb8A0wKjLGrve2q3PPVcBk=
github.com/imdario/mergo v0.3.13/go.mod h1:4lJ1jqUDcsbIECGy0RUJAXNIhg+6ocWgb1ALK2O4oXg=
github.com/inconshreveable/mousetrap v1.0.0/go.mod h1:PxqpIevigyE2G7u3NXJIT2ANytuPF1OarO4DADm73n8=
github.com/joho/godotenv v1.3.0/go.mod h1:7hK45KPybAkOC6peb+G5yklZfMxEjkZhHbwpqxOKXbg=
github.com/josharian/intern v1.0.0 h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=
github.com/josharian/intern v1.0.0/go.mod h1:5DoeVV0s6jJacbCEi61lwdGj/aVlrQvzHFFd8Hwg//Y=
github.com/json-iterator/go v1.1.12 h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=
github.com/json-iterator/go v1.1.12/go.mod h1:e30LSqwooZae/UwlEbR2852Gd8hjQvJoHmT4TnhNGBo=
github.com/karrick/godirwalk v1.8.0/go.mod h1:H5KPZjojv4lE+QYImBI8xVtrBRgYrIVsaRPx4tDPEn4=
github.com/karrick/godirwalk v1.10.3/go.mod h1:RoGL9dQei4vP9ilrpETWE8CLOZ1kiN0LhBygSwrAsHA=
github.com/klauspost/compress v1.13.6/go.mod h1:/3/Vjq9QcHkK5uEr5lBEmyoZ1iFhe47etQ6QUkpK6sk=
github.com/klauspost/compress v1.17.6 h1:60eq2E/jlfwQXtvZEeBUYADs+BwKBWURIY+Gj2eRGjI=
github.com/klauspost/compress v1.17.6/go.mod h1:/dCuZOvVtNoHsyb+cuJD3itjs3NbnF6KH9zAO4BDxPM=
github.com/konsorten/go-windows-terminal-sequences v1.0.1/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/konsorten/go-windows-terminal-sequences v1.0.2/go.mod h1:T0+1ngSBFLxvqU3pZ+m/2kptfBszLMUkC4ZK/EgS/cQ=
github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
github.com/kr/pretty v0.2.1/go.mod h1:ipq/a2n7PKx3OHsz4KJII5eveXtPO4qwEXGdVfWzfnI=
github.com/kr/pretty v0.3.1 h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=
github.com/kr/pretty v0.3.1/go.mod h1:hoEshYVHaxMs3cyo3Yncou5ZscifuDolrwPKZanG3xk=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 h1:6E+4a0GO5zZEnZ81pIr0yLvtUWk2if982qA3F3QD6H4=
github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0/go.mod h1:zJYVVT2jmtg6P3p1VtQj7WsuWi/y4VnjVBn7F8KPB3I=
github.com/magiconair/properties v1.8.7 h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=
github.com/magiconair/properties v1.8.7/go.mod h1:Dhd985XPs7jluiymwWYZ0G4Z61jb3vdS329zhj2hYo0=
github.com/mailru/easyjson v0.0.0-20190614124828-94de47d64c63/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e/go.mod h1:C1wdFJiN94OJF2b5HbByQZoLdCWB1Yqtg26g4irojpc=
github.com/mailru/easyjson v0.7.6/go.mod h1:xzfreul335JAWq5oZzymOObrkdz5UnU4kGfJJLY9Nlc=
github.com/mailru/easyjson v0.7.7 h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=
github.com/mailru/easyjson v0.7.7/go.mod h1:xzfreul335JAWq5oZzymOObrkdz5UnU4kGfJJLY9Nlc=
github.com/markbates/oncer v0.0.0-20181203154359-bf2de49a0be2/go.mod h1:Ld9puTsIW75CHf65OeIOkyKbteujpZVXDpWK6YGZbxE=
github.com/markbates/safe v1.0.1/go.mod h1:nAqgmRi7cY2nqMc92/bSEeQA+R4OheNU2T1kNSCBdG0=
github.com/mitchellh/copystructure v1.0.0 h1:Laisrj+bAB6b/yJwB5Bt3ITZhGJdqmxquMKeZ+mmkFQ=
github.com/mitchellh/copystructure v1.0.0/go.mod h1:SNtv71yrdKgLRyLFxmLdkAbkKEFWgYaq1OVrnRcwhnw=
github.com/mitchellh/mapstructure v1.3.3/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/mapstructure v1.4.1/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/mapstructure v1.5.0 h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=
github.com/mitchellh/mapstructure v1.5.0/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/reflectwalk v1.0.0 h1:9D+8oIskB4VJBN5SFlmc27fSlIBZaov1Wpk/IfikLNY=
github.com/mitchellh/reflectwalk v1.0.0/go.mod h1:mSTlrgnPZtwu0c4WaC2kGObEpuNDbx0jmZXqmk4esnw=
github.com/moby/patternmatcher v0.6.0 h1:GmP9lR19aU5GqSSFko+5pRqHi+Ohk1O69aFiKkVGiPk=
github.com/moby/patternmatcher v0.6.0/go.mod h1:hDPoyOpDY7OrrMDLaYoY3hf52gNCR/YOUYxkhApJIxc=
github.com/moby/sys/sequential v0.5.0 h1:OPvI35Lzn9K04PBbCLW0g4LcFAJgHsvXsRyewg5lXtc=
github.com/moby/sys/sequential v0.5.0/go.mod h1:tH2cOOs5V9MlPiXcQzRC+eEyab644PWKGRYaaV5ZZlo=
github.com/moby/sys/user v0.1.0 h1:WmZ93f5Ux6het5iituh9x2zAG7NFY9Aqi49jjE1PaQg=
github.com/moby/sys/user v0.1.0/go.mod h1:fKJhFOnsCN6xZ5gSfbM6zaHGgDJMrqt9/reuj4T7MmU=
github.com/moby/term v0.5.0 h1:xt8Q1nalod/v7BqbG21f8mQPqH+xAaC9C3N3wfWbVP0=
github.com/moby/term v0.5.0/go.mod h1:8FzsFHVUBGZdbDsJw/ot+X+d5HLUbvklYLJ9uGfcI3Y=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
github.com/modern-go/reflect2 v1.0.2 h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=
github.com/modern-go/reflect2 v1.0.2/go.mod h1:yWuevngMOJpCy52FWWMvUC8ws7m/LJsjYzDa0/r8luk=
github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe/go.mod h1:wL8QJuTMNUDYhXwkmfOly8iTdp5TEcJFWZD2D7SIkUc=
github.com/morikuni/aec v1.0.0 h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=
github.com/morikuni/aec v1.0.0/go.mod h1:BbKIizmSmc5MMPqRYbxO4ZU0S0+P200+tUnFx7PXmsc=
github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e/go.mod h1:zD1mROLANZcx1PVRCS0qkT7pwLkGfwJo4zjcN/Tysno=
github.com/nikolalohinski/gonja v1.5.3 h1:GsA+EEaZDZPGJ8JtpeGN78jidhOlxeJROpqMT9fTj9c=
github.com/nikolalohinski/gonja v1.5.3/go.mod h1:RmjwxNiXAEqcq1HeK5SSMmqFJvKOfTfXhkJv6YBtPa4=
github.com/oklog/ulid v1.3.1 h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=
github.com/oklog/ulid v1.3.1/go.mod h1:CirwcVhetQ6Lv90oh/F+FBtV6XMibvdAFo93nm5qn4U=
github.com/opencontainers/go-digest v1.0.0 h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=
github.com/opencontainers/go-digest v1.0.0/go.mod h1:0JzlMkj0TRzQZfJkVvzbP0HBR3IKzErnv2BNG4W4MAM=
github.com/opencontainers/image-spec v1.1.0 h1:8SG7/vwALn54lVB/0yZ/MMwhFrPYtpEHQb2IpWsCzug=
github.com/opencontainers/image-spec v1.1.0/go.mod h1:W4s4sFTMaBeK1BQLXbG4AdM2szdn85PY75RI83NrTrM=
github.com/pelletier/go-toml v1.7.0 h1:7utD74fnzVc/cpcyy8sjrlFr5vYpypUixARcHIMIGuI=
github.com/pelletier/go-toml v1.7.0/go.mod h1:vwGMzjaWMwyfHwgIBhI2YUM4fB6nL6lVAvS1LBMMhTE=
github.com/pelletier/go-toml/v2 v2.0.9 h1:uH2qQXheeefCCkuBBSLi7jCiSmj3VRh2+Goq2N7Xxu0=
github.com/pelletier/go-toml/v2 v2.0.9/go.mod h1:tJU2Z3ZkXwnxa4DPO899bsyIoywizdUvyaeZurnPPDc=
github.com/pkg/errors v0.8.0/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkg/errors v0.8.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pkoukk/tiktoken-go v0.1.6 h1:JF0TlJzhTbrI30wCvFuiw6FzP2+/bR+FIxUdgEAcUsw=
github.com/pkoukk/tiktoken-go v0.1.6/go.mod h1:9NiV+i9mJKGj1rYOT+njbv+ZwA/zJxYdewGl6qVatpg=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c h1:ncq/mPwQF4JjgDlrVEn3C11VoGHZN7m8qihwgMEtzYw=
github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c/go.mod h1:OmDBASR4679mdNQnz2pUhc2G8CO2JrUAVFDRBDP/hJE=
github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4/go.mod h1:xMI15A0UPsDsEKsMN9yxemIoYk6Tm2C1GtYGdfGttqA=
github.com/rogpeppe/go-internal v1.1.0/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/rogpeppe/go-internal v1.2.2/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/rogpeppe/go-internal v1.3.0/go.mod h1:M8bDsm7K2OlrFYOpmOWEs/qY81heoFRclV5y23lUDJ4=
github.com/rogpeppe/go-internal v1.11.0 h1:cWPaGQEPrBb5/AsnsZesgZZ9yb1OQ+GOISoDNXVBh4M=
github.com/rogpeppe/go-internal v1.11.0/go.mod h1:ddIwULY96R17DhadqLgMfk9H9tvdUzkipdSkR5nkCZA=
github.com/shirou/gopsutil/v3 v3.23.12 h1:z90NtUkp3bMtmICZKpC4+WaknU1eXtp5vtbQ11DgpE4=
github.com/shirou/gopsutil/v3 v3.23.12/go.mod h1:1FrWgea594Jp7qmjHUUPlJDTPgcsb9mGnXDxavtikzM=
github.com/shoenig/go-m1cpu v0.1.6 h1:nxdKQNcEB6vzgA2E2bvzKIYRuNj7XNJ4S/aRSwKzFtM=
github.com/shoenig/go-m1cpu v0.1.6/go.mod h1:1JJMcUBvfNwpq05QDQVAnx3gUHr9IYF7GNg9SUEw2VQ=
github.com/shopspring/decimal v1.2.0 h1:abSATXmQEYyShuxI4/vyW3tV1MrKAJzCZ/0zLUXYbsQ=
github.com/shopspring/decimal v1.2.0/go.mod h1:DKyhrW/HYNuLGql+MJL6WCR6knT2jwCFRcu2hWCYk4o=
github.com/sirupsen/logrus v1.4.0/go.mod h1:LxeOpSwHxABJmUn/MG1IvRgCAasNZTLOkJPxbbu5VWo=
github.com/sirupsen/logrus v1.4.1/go.mod h1:ni0Sbl8bgC9z8RoU9G6nDWqqs/fq4eDPysMBDgk/93Q=
github.com/sirupsen/logrus v1.4.2/go.mod h1:tLMulIdttU9McNUspp0xgXVQah82FyeX6MwdIuYE2rE=
github.com/sirupsen/logrus v1.9.3 h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=
github.com/sirupsen/logrus v1.9.3/go.mod h1:naHLuLoDiP4jHNo9R0sCBMtWGeIprob74mVsIT4qYEQ=
github.com/spf13/cast v1.3.1 h1:nFm6S0SMdyzrzcmThSipiEubIDy8WEXKNZ0UOgiRpng=
github.com/spf13/cast v1.3.1/go.mod h1:Qx5cxh0v+4UWYiBimWS+eyWzqEqokIECu5etghLkUJE=
github.com/spf13/cobra v0.0.3/go.mod h1:1l0Ry5zgKvJasoi3XT1TypsSe7PqH0Sj9dhYf7v3XqQ=
github.com/spf13/pflag v1.0.3/go.mod h1:DYY7MBk1bdzusC3SYhjObp+wFpr4gzcvqqNjLnInEg4=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/objx v0.1.1/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/objx v0.4.0/go.mod h1:YvHI0jy2hoMjB+UWwv71VJQ9isScKT/TqJzVSSt89Yw=
github.com/stretchr/objx v0.5.0/go.mod h1:Yh+to48EsGEfYuaHDzXPcE3xhTkx73EhmCGUpEOglKo=
github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.6.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.7.1/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
github.com/stretchr/testify v1.8.0/go.mod h1:yNjHg4UonilssWZ8iaSj1OCr/vHnekPRkoO+kdMU+MU=
github.com/stretchr/testify v1.8.1/go.mod h1:w2LPCIKwWwSfY2zedu0+kehJoqGctiVI29o6fzry7u4=
github.com/stretchr/testify v1.9.0 h1:HtqpIVDClZ4nwg75+f6Lvsy/wHu+3BoSGCbBAcpTsTg=
github.com/stretchr/testify v1.9.0/go.mod h1:r2ic/lqez/lEtzL7wO/rwa5dbSLXVDPFyf8C91i36aY=
github.com/testcontainers/testcontainers-go v0.31.0 h1:W0VwIhcEVhRflwL9as3dhY6jXjVCA27AkmbnZ+UTh3U=
github.com/testcontainers/testcontainers-go v0.31.0/go.mod h1:D2lAoA0zUFiSY+eAflqK5mcUx/A5hrrORaEQrd0SefI=
github.com/testcontainers/testcontainers-go/modules/weaviate v0.31.0 h1:iVJX9O12GHRhqPgIuz/eE8BsNEwyrUMJnWgduBt8quc=
github.com/testcontainers/testcontainers-go/modules/weaviate v0.31.0/go.mod h1:WNc2XhLphiLdNJdjJZvUtRj08ThLY8FL60y7FQSJTPQ=
github.com/tidwall/pretty v1.0.0/go.mod h1:XNkn88O1ChpSDQmQeStsy+sBenx6DDtFZJxhVysOjyk=
github.com/tklauser/go-sysconf v0.3.12 h1:0QaGUFOdQaIVdPgfITYzaTegZvdCjmYO52cSFAEVmqU=
github.com/tklauser/go-sysconf v0.3.12/go.mod h1:Ho14jnntGE1fpdOqQEEaiKRpvIavV0hSfmBq8nJbHYI=
github.com/tklauser/numcpus v0.6.1 h1:ng9scYS7az0Bk4OZLvrNXNSAO2Pxr1XXRAPyjhIx+Fk=
github.com/tklauser/numcpus v0.6.1/go.mod h1:1XfjsgE2zo8GVw7POkMbHENHzVg3GzmoZ9fESEdAacY=
github.com/tmc/langchaingo v0.1.12 h1:yXwSu54f3b1IKw0jJ5/DWu+qFVH1NBblwC0xddBzGJE=
github.com/tmc/langchaingo v0.1.12/go.mod h1:cd62xD6h+ouk8k/QQFhOsjRYBSA1JJ5UVKXSIgm7Ni4=
github.com/weaviate/weaviate v1.24.1 h1:Cl/NnqgFlNfyC7KcjFtETf1bwtTQPLF3oz5vavs+Jq0=
github.com/weaviate/weaviate v1.24.1/go.mod h1:wcg1vJgdIQL5MWBN+871DFJQa+nI2WzyXudmGjJ8cG4=
github.com/weaviate/weaviate-go-client/v4 v4.13.1 h1:7PuK/hpy6Q0b9XaVGiUg5OD1MI/eF2ew9CJge9XdBEE=
github.com/weaviate/weaviate-go-client/v4 v4.13.1/go.mod h1:B2m6g77xWDskrCq1GlU6CdilS0RG2+YXEgzwXRADad0=
github.com/xdg-go/pbkdf2 v1.0.0/go.mod h1:jrpuAogTd400dnrH08LKmI/xc1MbPOebTwRqcT5RDeI=
github.com/xdg-go/scram v1.0.2/go.mod h1:1WAq6h33pAW+iRreB34OORO2Nf7qel3VV3fjBj+hCSs=
github.com/xdg-go/scram v1.1.1/go.mod h1:RaEWvsqvNKKvBPvcKeFjrG2cJqOkHTiyTpzz23ni57g=
github.com/xdg-go/stringprep v1.0.2/go.mod h1:8F9zXuvzgwmyT5DUm4GUfZGDdT3W+LCvS6+da4O5kxM=
github.com/xdg-go/stringprep v1.0.3/go.mod h1:W3f5j4i+9rC0kuIEJL0ky1VpHXQU3ocBgklLGvcBnW8=
github.com/yargevad/filepathx v1.0.0 h1:SYcT+N3tYGi+NvazubCNlvgIPbzAk7i7y2dwg3I5FYc=
github.com/yargevad/filepathx v1.0.0/go.mod h1:BprfX/gpYNJHJfc35GjRRpVcwWXS89gGulUIU5tK3tA=
github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d/go.mod h1:rHwXgn7JulP+udvsHwJoVG1YGAP6VLg4y9I5dyZdqmA=
github.com/yusufpapurcu/wmi v1.2.3 h1:E1ctvB7uKFMOJw3fdOW32DwGE9I7t++CRUEMKvFoFiw=
github.com/yusufpapurcu/wmi v1.2.3/go.mod h1:SBZ9tNy3G9/m5Oi98Zks0QjeHVDvuK0qfxQmPyzfmi0=
go.mongodb.org/mongo-driver v1.7.3/go.mod h1:NqaYOwnXWr5Pm7AOpO5QFxKJ503nbMse/R79oO62zWg=
go.mongodb.org/mongo-driver v1.7.5/go.mod h1:VXEWRZ6URJIkUq2SCAyapmhH0ZLRBP+FT4xhp5Zvxng=
go.mongodb.org/mongo-driver v1.10.0/go.mod h1:wsihk0Kdgv8Kqu1Anit4sfK+22vSFbUrAVEYRhCXrA8=
go.mongodb.org/mongo-driver v1.14.0 h1:P98w8egYRjYe3XDjxhYJagTokP/H6HzlsnojRgZRd80=
go.mongodb.org/mongo-driver v1.14.0/go.mod h1:Vzb0Mk/pa7e6cWw85R4F/endUC3u0U9jGcNU603k65c=
go.opencensus.io v0.24.0 h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=
go.opencensus.io v0.24.0/go.mod h1:vNK8G9p7aAivkbmorf4v+7Hgx+Zs0yY+0fOtgBfjQKo=
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 h1:A3SayB3rNyt+1S6qpI9mHPkeHTZbD7XILEqWnYZb2l0=
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0/go.mod h1:27iA5uvhuRNmalO+iEUdVn5ZMj2qy10Mm+XRIpRmyuU=
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 h1:Xs2Ncz0gNihqu9iosIZ5SkBbWo5T8JhhLJFMQL1qmLI=
go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0/go.mod h1:vy+2G/6NvVMpwGX/NyLqcC41fxepnuKHk16E6IZUcJc=
go.opentelemetry.io/otel v1.26.0 h1:LQwgL5s/1W7YiiRwxf03QGnWLb2HW4pLiAhaA5cZXBs=
go.opentelemetry.io/otel v1.26.0/go.mod h1:UmLkJHUAidDval2EICqBMbnAd0/m2vmpf/dAM+fvFs4=
go.opentelemetry.io/otel/metric v1.26.0 h1:7S39CLuY5Jgg9CrnA9HHiEjGMF/X2VHvoXGgSllRz30=
go.opentelemetry.io/otel/metric v1.26.0/go.mod h1:SY+rHOI4cEawI9a7N1A4nIg/nTQXe1ccCNWYOJUrpX4=
go.opentelemetry.io/otel/trace v1.26.0 h1:1ieeAUb4y0TE26jUFrCIXKpTuVK7uJGN9/Z/2LP5sQA=
go.opentelemetry.io/otel/trace v1.26.0/go.mod h1:4iDxvGDQuUkHve82hJJ8UqrwswHYsZuWCBllGV2U2y0=
go.starlark.net v0.0.0-20230302034142-4b1e35fe2254 h1:Ss6D3hLXTM0KobyBYEAygXzFfGcjnmfEJOBgSbemCtg=
go.starlark.net v0.0.0-20230302034142-4b1e35fe2254/go.mod h1:jxU+3+j+71eXOW14274+SmmuW82qJzl6iZSeqEtTGds=
golang.org/x/crypto v0.0.0-20180904163835-0709b304e793/go.mod h1:6SG95UA2DQfeDnfUPMdvaQW0Q7yPrPDi9nlGo2tz2b4=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20190422162423-af44ce270edf/go.mod h1:WFFai1msRO1wXaEeE5yQxYXgSfI8pQAWXbQop6sCtWE=
golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d/go.mod h1:IxCIyHEi3zRg3s0A5j5BB6A9Jmi73HwBIUl50j+osU4=
golang.org/x/crypto v0.26.0 h1:RrRspgV4mU+YwB4FYnuBoKsUapNIL5cohGAmSH3azsw=
golang.org/x/crypto v0.26.0/go.mod h1:GY7jblb9wI+FOo5y8/S2oY4zWP07AkOJ4+jxCqdqn54=
golang.org/x/exp v0.0.0-20190121172915-509febef88a4/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 h1:MGwJjxBy0HJshjDNfLsYO8xppfqWlA5ZT9OhtUUhTNw=
golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1/go.mod h1:FXUEEKJgO7OQYeo8N01OfiKP8RXMtf6e8aTskBGqWdc=
golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3/go.mod h1:UVdnD1Gm6xHRNCYTkRU2/jEulfH38KcIWyp/GAMgvoE=
golang.org/x/lint v0.0.0-20190227174305-5b3e6a55c961/go.mod h1:wehouNa3lNwaWXcvxsM5YxQ5yQlVC4a0KAMCusXpPoU=
golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
golang.org/x/mod v0.17.0 h1:zY54UmvipHiNd+pm+m0x9KhZ9hl1/7QNMyxXbc6ICqA=
golang.org/x/mod v0.17.0/go.mod h1:hTbmBsO62+eylJbnUtE2MGJUyE7QWk4xUqPFrRgJ+7c=
golang.org/x/net v0.0.0-20180724234803-3673e40ba225/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20180826012351-8a410e7b638d/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190213061140-3a22650c66bd/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190311183353-d8887717615a/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20201110031124-69a78807bb2b/go.mod h1:sp8m0HH+o8qH0wwXwYZr8TS3Oi6o0r6Gce1SSxlDquU=
golang.org/x/net v0.0.0-20210421230115-4e50805a0758/go.mod h1:72T/g9IO56b78aLF+1Kcs5dz7/ng1VjMUvfKvpfy+jM=
golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
golang.org/x/net v0.28.0 h1:a9JDOJc5GMUJ0+UDqmLT86WiEy7iWyIhz8gz8E4e5hE=
golang.org/x/net v0.28.0/go.mod h1:yqtgsTWOOnlGLG9GFRrK3++bGOUEkNBoHZc8MEDWPNg=
golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be/go.mod h1:N/0e6XlmueqKjAGxoOufVs8QHGRruUQn6yWY3a++T0U=
golang.org/x/oauth2 v0.22.0 h1:BzDx2FehcG7jJwgWLELCdmLuxk2i+x9UDpSiss2u0ZA=
golang.org/x/oauth2 v0.22.0/go.mod h1:XYTD2NtWslqkgxebSiOHnXEap4TF09sJSc7H1sXbhtI=
golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20181108010431-42b317875d0f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190412183630-56d357773e84/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20210220032951-036812b2e83c/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.8.0 h1:3NFvSEYkUoMifnESzZl15y791HH1qU2xm6eCJU5ZPXQ=
golang.org/x/sync v0.8.0/go.mod h1:Czt+wKu1gCyEFDUtn0jG5QVvpJ6rzVqr5aXyt9drQfk=
golang.org/x/sys v0.0.0-20180830151530-49385e6e1522/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190403152447-81d4e9dc473e/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190419153524-e8e3143a4f4a/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190422165155-953cdadca894/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20190531175056-4c3a928424d2/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210420072515-93ed5bcd2bfe/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.24.0 h1:Twjiwq9dn6R1fQcyiK+wQyHWfaz/BJB+YIpzU/Cv3Xg=
golang.org/x/sys v0.24.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.5/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
golang.org/x/text v0.17.0 h1:XtiM5bkSOt+ewxlOE/aE/AKEHibwj/6gvWMl9Rsh0Qc=
golang.org/x/text v0.17.0/go.mod h1:BuEKDfySbSR4drPmRPG/7iBdf8hvFMuRexcpahXilzY=
golang.org/x/time v0.5.0 h1:o7cqy6amK/52YcAKIPlM3a+Fpj35zvRj2TP+e1xFSfk=
golang.org/x/time v0.5.0/go.mod h1:3BpzKBy/shNhVucY/MWOyx10tF3SFh9QdLuxbVysPQM=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190114222345-bf090417da8b/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190226205152-f727befe758c/go.mod h1:9Yl7xja0Znq3iFh3HoIrodX9oNMXvdceNzlUR8zjMvY=
golang.org/x/tools v0.0.0-20190311212946-11955173bddd/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190329151228-23e29df326fe/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190416151739-9c9e1878f421/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190420181800-aa740d480789/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135/go.mod h1:RgjU9mgBXZiqYHBnxXauZ1Gv1EHHAz9KjViQ78xBX0Q=
golang.org/x/tools v0.0.0-20190531172133-b3315ee88b7d/go.mod h1:/rFqwRUd4F7ZHNgwSSTFct+R/Kf4OFW1sUzUTQQTgfc=
golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d h1:vU5i/LfpvrRCpgM/VPfJLg5KjxD3E+hfT1SH+d9zLwg=
golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d/go.mod h1:aiJjzUbINMkxbQROHiO6hDPo2LHcIPhhQsa9DLh0yGk=
golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
google.golang.org/api v0.180.0 h1:M2D87Yo0rGBPWpo1orwfCLehUUL6E7/TYe5gvMQWDh4=
google.golang.org/api v0.180.0/go.mod h1:51AiyoEg1MJPSZ9zvklA8VnRILPXxn1iVen9v25XHAE=
google.golang.org/appengine v1.1.0/go.mod h1:EbEs0AVv82hx2wNQdGPgUI5lhzA/G0D9YwlJXL52JkM=
google.golang.org/appengine v1.4.0/go.mod h1:xpcJRLb0r/rnEns0DIKYYv+WjYCduHsrkT7/EB5XEv4=
google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8/go.mod h1:JiN7NxoALGmiZfu7CAH4rXhgtRTLTxftemlI0sWmxmc=
google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55/go.mod h1:DMBHOl98Agz4BDEuKkezgsaosCRResVns1a3J2ZsMNc=
google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013/go.mod h1:NbSheEEYHJ7i3ixzK3sjbqSGDJWnxyFXZblF3eUsNvo=
google.golang.org/genproto v0.0.0-20240401170217-c3f982113cda h1:wu/KJm9KJwpfHWhkkZGohVC6KRrc1oJNr4jwtQMOQXw=
google.golang.org/genproto v0.0.0-20240401170217-c3f982113cda/go.mod h1:g2LLCvCeCSir/JJSWosk19BR4NVxGqHUC6rxIRsd7Aw=
google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142 h1:wKguEg1hsxI2/L3hUYrpo1RVi48K+uTyzKqprwLXsb8=
google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142/go.mod h1:d6be+8HhtEtucleCbxpPW9PA9XwISACu8nvpPqF0BVo=
google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 h1:e7S5W7MGGLaSu8j3YjdezkZ+m1/Nm0uRVRMEMGk26Xs=
google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142/go.mod h1:UqMtugtsSgubUsoxbuAoiCXvqvErP7Gf0so0mK9tHxU=
google.golang.org/grpc v1.19.0/go.mod h1:mqu4LbDTu4XGKhr4mRzUsmM4RtVoemTSY81AxZiDr8c=
google.golang.org/grpc v1.23.0/go.mod h1:Y5yQAOtifL1yxbo5wqy6BxZv8vAUGQwXBOALyacEbxg=
google.golang.org/grpc v1.25.1/go.mod h1:c3i+UQWmh7LiEpx4sFZnkU36qjEYZ0imhYfXVyQciAY=
google.golang.org/grpc v1.27.0/go.mod h1:qbnxyOmOxrQa7FizSgH+ReBfzJrCY1pSN7KXBS8abTk=
google.golang.org/grpc v1.33.2/go.mod h1:JMHMWHQWaTccqQQlmk3MJZS+GWXOdAesneDmEnv2fbc=
google.golang.org/grpc v1.67.0 h1:IdH9y6PF5MPSdAntIcpjQ+tXO41pcQsfZV2RxtQgVcw=
google.golang.org/grpc v1.67.0/go.mod h1:1gLDyUQU7CTLJI90u3nXZ9ekeghjeM7pTDZlqFNg2AA=
google.golang.org/protobuf v0.0.0-20200109180630-ec00e32a8dfd/go.mod h1:DFci5gLYBciE7Vtevhsrf46CRTquxDuWsQurQQe4oz8=
google.golang.org/protobuf v0.0.0-20200221191635-4d8936d0db64/go.mod h1:kwYJMbMJ01Woi6D6+Kah6886xMZcty6N08ah7+eCXa0=
google.golang.org/protobuf v0.0.0-20200228230310-ab0ca4ff8a60/go.mod h1:cfTl7dwQJ+fmap5saPgwCLgHXTUD7jkjRqWcaiX5VyM=
google.golang.org/protobuf v1.20.1-0.20200309200217-e05f789c0967/go.mod h1:A+miEFZTKqfCUM6K7xSMQL9OKL/b6hQv+e19PK+JZNE=
google.golang.org/protobuf v1.21.0/go.mod h1:47Nbq4nVaFHyn7ilMalzfO3qCViNmqZ2kzikPIcrTAo=
google.golang.org/protobuf v1.22.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.1-0.20200526195155-81db48ad09cc/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.25.0/go.mod h1:9JNX74DMeImyA3h4bdi1ymwjUzf21/xIlbajtzgsN7c=
google.golang.org/protobuf v1.34.2 h1:6xV6lTsCfpGD21XK49h7MhtcApnLqkfYgPcdHftf6hg=
google.golang.org/protobuf v1.34.2/go.mod h1:qYOHts0dSfpeUzUFpOMr/WGzszTmLH+DiWniOlNbLDw=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
gopkg.in/errgo.v2 v2.1.0/go.mod h1:hNsd1EY+bozCKY1Ytp96fpM3vjJbqLJn88ws8XvfDNI=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.4.0 h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=
gopkg.in/yaml.v2 v2.4.0/go.mod h1:RDklbk79AGWmwhnvt/jBztapEOGDOx6ZbXqjP6csGnQ=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20200605160147-a5ece683394c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
honnef.co/go/tools v0.0.0-20190102054323-c2f93a96b099/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
sigs.k8s.io/yaml v1.3.0 h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=
sigs.k8s.io/yaml v1.3.0/go.mod h1:GeOyir5tyXNByN85N/dRIT9es5UQNerPYEKK56eTBm8=
```

### ragserver/ragserver-langchaingo/json.go

```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

// readRequestJSON expects req to have a JSON content type with a body that
// contains a JSON-encoded value complying with the underlying type of target.
// It populates target, or returns an error.
func readRequestJSON(req *http.Request, target any) error {
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	if mediaType != "application/json" {
		return fmt.Errorf("expect application/json Content-Type, got %s", mediaType)
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(target)
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
```

### ragserver/ragserver-langchaingo/main.go

```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command ragserver is an HTTP server that implements RAG (Retrieval
// Augmented Generation) using the Gemini model and Weaviate, which
// are accessed using LangChainGo. See the accompanying README file for
// additional details.
package main

import (
	"cmp"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

const generativeModelName = "gemini-1.5-flash"
const embeddingModelName = "text-embedding-004"

// This is a standard Go HTTP server. Server state is in the ragServer struct.
// The `main` function connects to the required services (Weaviate and Google
// AI), initializes the server state and registers HTTP handlers.
func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	geminiClient, err := googleai.New(ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultEmbeddingModel(embeddingModelName))
	if err != nil {
		log.Fatal(err)
	}

	emb, err := embeddings.NewEmbedder(geminiClient)
	if err != nil {
		log.Fatal(err)
	}

	wvStore, err := weaviate.New(
		weaviate.WithEmbedder(emb),
		weaviate.WithScheme("http"),
		weaviate.WithHost("localhost:"+cmp.Or(os.Getenv("WVPORT"), "9035")),
		weaviate.WithIndexName("Document"),
	)

	server := &ragServer{
		ctx:          ctx,
		wvStore:      wvStore,
		geminiClient: geminiClient,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /add/", server.addDocumentsHandler)
	mux.HandleFunc("POST /query/", server.queryHandler)

	port := cmp.Or(os.Getenv("SERVERPORT"), "9020")
	address := "localhost:" + port
	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

type ragServer struct {
	ctx          context.Context
	wvStore      weaviate.Store
	geminiClient *googleai.GoogleAI
}

func (rs *ragServer) addDocumentsHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type document struct {
		Text string
	}
	type addRequest struct {
		Documents []document
	}
	ar := &addRequest{}

	err := readRequestJSON(req, ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store documents and their embeddings in weaviate
	var wvDocs []schema.Document
	for _, doc := range ar.Documents {
		wvDocs = append(wvDocs, schema.Document{PageContent: doc.Text})
	}
	_, err = rs.wvStore.AddDocuments(rs.ctx, wvDocs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs *ragServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	// Parse HTTP request from JSON.
	type queryRequest struct {
		Content string
	}
	qr := &queryRequest{}
	err := readRequestJSON(req, qr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the most similar documents.
	docs, err := rs.wvStore.SimilaritySearch(rs.ctx, qr.Content, 3)
	if err != nil {
		http.Error(w, fmt.Errorf("similarity search: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	var docsContents []string
	for _, doc := range docs {
		docsContents = append(docsContents, doc.PageContent)
	}

	// Create a RAG query for the LLM with the most relevant documents as
	// context.
	ragQuery := fmt.Sprintf(ragTemplateStr, qr.Content, strings.Join(docsContents, "\n"))
	respText, err := llms.GenerateFromSinglePrompt(rs.ctx, rs.geminiClient, ragQuery, llms.WithModel(generativeModelName))
	if err != nil {
		log.Printf("calling generative model: %v", err.Error())
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	renderJSON(w, respText)
}

const ragTemplateStr = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Question:
%s

Context:
%s
`
```

### ragserver/tests/add-documents.sh

```bash
#!/bin/bash

# This script interacts with a running ragserver to add some example documents.

set -eux

echo '{
	"documents": [
	{"text": "TDXIRV is an environment variable for controlling throttle speed"},
	{"text": "some flags for setting acceleration are --accelxyzp and --acceljjrv"},
	{"text": "acceleration is also affected by the ACCUVI5 env var"},
	{"text": "/usr/local/fuel555 contains information about fuel capacity"},
	{"text": "we can control fuel savings with the --savemyfuelplease flag"},
	{"text": "fuel savings can be observed on local port 48332"}
]}' | tr -d "\n" | curl \
		-X POST \
    -H 'Content-Type: application/json' \
    -d @- \
    http://localhost:9020/add/
```

### ragserver/tests/docker-compose.yml

```yaml
---
version: '3.4'
services:
  weaviate:
    command:
    - --host
    - 0.0.0.0
    - --port
    - '9035'
    - --scheme
    - http
    image: cr.weaviate.io/semitechnologies/weaviate:1.25.2
    ports:
    - 9035:9035
    - 50051:50051
    restart: on-failure:0
    environment:
      QUERY_DEFAULTS_LIMIT: 25
      AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED: 'true'
      PERSISTENCE_DATA_PATH: '/var/lib/weaviate'
      CLUSTER_HOSTNAME: 'node1'
...
```

### ragserver/tests/query.sh

```bash
#!/bin/bash

set -eu

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 '<your query>'"
    exit 1
fi

set -x

# Capture the query from the command-line argument
QUERY_CONTENT=$1

# Build the JSON payload
PAYLOAD=$(echo "{\"content\": \"$QUERY_CONTENT\"}")

# Send the request
echo "$PAYLOAD" | tr -d "\n" | curl \
    -X POST \
    -H 'Content-Type: application/json' \
    -d @- \
    http://localhost:9020/query/ | sed 's/\\n/\n/g'
```

### ragserver/tests/weaviate-delete-objects.sh

```bash
#!/bin/bash

set -eux

curl \
  -X DELETE \
  http://localhost:9035/v1/schema/Document
```

### ragserver/tests/weaviate-show-all.sh

```bash
#!/bin/bash

set -eux

echo '{
  "query": "{
    Get {
      Document { 
        text
      }
    }
  }"
}' | tr -d "\n" | curl \
    -X POST \
    -H 'Content-Type: application/json' \
    -d @- \
    http://localhost:9035/v1/graphql | jq .
```

### ragserver/README.md

```markdown
# ragserver

*RAG stands for Retrieval Augmented Generation*

Demos of implementing a "RAG Server" in Go, using [Google
AI](https://ai.google.dev/) for embeddings and language models and
[Weaviate](https://weaviate.io/) as a vector database.


## How it works

The server we're developing is a standard Go HTTP server, listening on a local
port. See the next section for the request schema for this server. It supports
adding new documents to its context, and getting queries that would use this
context.

Weaviate has the be installed locally; the easiest way to do so is by using
`docker-compose` as described in the Usage section.

## Server request schema

```
/add/: POST {"documents": [{"text": "..."}, {"text": "..."}, ...]}
  response: OK (no body)

/query/: POST {"content": "..."}
  response: model response as a string
```

## Server variants

* `ragserver`: uses the Google AI Go SDK directly for LLM calls and embeddings,
  and the Weaviate Go client library directly for interacting with Weaviate.
* `ragserver-langchaingo`: uses [LangChain for Go](https://github.com/tmc/langchaingo)
  to interact with Weaviate and Google's LLM and embedding models.
* `ragserver-genkit`: uses [Genkit Go](https://firebase.google.com/docs/genkit-go/get-started-go)
  to interact with Weaviate and Google's LLM and embedding models.

## Usage

* In terminal window 1, `cd tests` and run `docker-compose up`;
  This will start the weaviate service in the foreground.
* In terminal window 2, run `GEMINI_API_KEY=... go run .` in the tested
  `ragserver` directory.
* In terminal window 3, we can now run scripts to clear/populate the
  weaviate DB and interact with `ragserver`. The following instructions are
  for terminal window 3.

Run `cd tests`; then we can clear out the weaviate DB with
`./weaviate-delete-objects.sh`. To add documents to the DB through `ragserver`,
run `./add-documents.sh`. For a sample query, run `./query.sh`
Adjust the contents of these scripts as needed.

## Environment variables

* `SERVERPORT`: the port this server is listening on (default 9020)
* `WVPORT`: the port Weaviate is listening on (default 9035)
* `GEMINI_API_KEY`: API key for the Gemini service at https://ai.google.dev
```

### slog-handler-guide/indenthandler1/indent\_handler.go

```go
//go:build go1.21

package indenthandler

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

// !+types
type IndentHandler struct {
	opts Options
	// TODO: state for WithGroup and WithAttrs
	mu  *sync.Mutex
	out io.Writer
}

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

//!-types

// !+enabled
func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

//!-enabled

func (h *IndentHandler) WithGroup(name string) slog.Handler {
	// TODO: implement.
	return h
}

func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// TODO: implement.
	return h
}

// !+handle
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	indentLevel := 0
	// TODO: output the Attrs and groups from WithAttrs and WithGroup.
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a, indentLevel)
		return true
	})
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

//!-handle

// !+appendAttr
func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else.
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs.
	if a.Equal(slog.Attr{}) {
		return buf
	}
	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339Nano))
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, indentLevel)
		}
	default:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
	}
	return buf
}

//!-appendAttr
```

### slog-handler-guide/indenthandler1/indent\_handler\_test.go

```go
//go:build go1.21

package indenthandler

import (
	"bytes"
	"regexp"
	"testing"

	"log/slog"
)

func Test(t *testing.T) {
	var buf bytes.Buffer
	l := slog.New(New(&buf, nil))
	l.Info("hello", "a", 1, "b", true, "c", 3.14, slog.Group("g", "h", 1, "i", 2), "d", "NO")
	got := buf.String()
	wantre := `time: [-0-9T:.]+Z?
level: INFO
source: ".*/indent_handler_test.go:\d+"
msg: "hello"
a: 1
b: true
c: 3.14
g:
    h: 1
    i: 2
d: "NO"
`
	re := regexp.MustCompile(wantre)
	if !re.MatchString(got) {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, wantre)
	}

	buf.Reset()
	l.Debug("test")
	if got := buf.Len(); got != 0 {
		t.Errorf("got buf.Len() = %d, want 0", got)
	}
}
```

### slog-handler-guide/indenthandler2/indent\_handler.go

```go
//go:build go1.21

package indenthandler

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

// !+IndentHandler
type IndentHandler struct {
	opts Options
	goas []groupOrAttrs
	mu   *sync.Mutex
	out  io.Writer
}

//!-IndentHandler

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

// !+gora
// groupOrAttrs holds either a group name or a list of slog.Attrs.
type groupOrAttrs struct {
	group string      // group name if non-empty
	attrs []slog.Attr // attrs if non-empty
}

//!-gora

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// !+withs
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{group: name})
}

func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{attrs: attrs})
}

//!-withs

// !+withgora
func (h *IndentHandler) withGroupOrAttrs(goa groupOrAttrs) *IndentHandler {
	h2 := *h
	h2.goas = make([]groupOrAttrs, len(h.goas)+1)
	copy(h2.goas, h.goas)
	h2.goas[len(h2.goas)-1] = goa
	return &h2
}

//!-withgora

// !+handle
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	indentLevel := 0
	// Handle state from WithGroup and WithAttrs.
	goas := h.goas
	if r.NumAttrs() == 0 {
		// If the record has no Attrs, remove groups at the end of the list; they are empty.
		for len(goas) > 0 && goas[len(goas)-1].group != "" {
			goas = goas[:len(goas)-1]
		}
	}
	for _, goa := range goas {
		if goa.group != "" {
			buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", goa.group)
			indentLevel++
		} else {
			for _, a := range goa.attrs {
				buf = h.appendAttr(buf, a, indentLevel)
			}
		}
	}
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a, indentLevel)
		return true
	})
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

//!-handle

func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else.
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs.
	if a.Equal(slog.Attr{}) {
		return buf
	}
	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339Nano))
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, indentLevel)
		}
	default:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
	}
	return buf
}
```

### slog-handler-guide/indenthandler2/indent\_handler\_test.go

```go
//go:build go1.21

package indenthandler

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"log/slog"
	"testing/slogtest"
)

func TestSlogtest(t *testing.T) {
	var buf bytes.Buffer
	err := slogtest.TestHandler(New(&buf, nil), func() []map[string]any {
		return parseLogEntries(buf.String())
	})
	if err != nil {
		t.Error(err)
	}
}

func Test(t *testing.T) {
	var buf bytes.Buffer
	l := slog.New(New(&buf, nil))
	l.Info("hello", "a", 1, "b", true, "c", 3.14, slog.Group("g", "h", 1, "i", 2), "d", "NO")
	got := buf.String()
	wantre := `time: [-0-9T:.]+Z?
level: INFO
source: ".*/indent_handler_test.go:\d+"
msg: "hello"
a: 1
b: true
c: 3.14
g:
    h: 1
    i: 2
d: "NO"
`
	re := regexp.MustCompile(wantre)
	if !re.MatchString(got) {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, wantre)
	}

	buf.Reset()
	l.Debug("test")
	if got := buf.Len(); got != 0 {
		t.Errorf("got buf.Len() = %d, want 0", got)
	}
}

func parseLogEntries(s string) []map[string]any {
	var ms []map[string]any
	scan := bufio.NewScanner(strings.NewReader(s))
	for scan.Scan() {
		m := parseGroup(scan)
		ms = append(ms, m)
	}
	if scan.Err() != nil {
		panic(scan.Err())
	}
	return ms
}

func parseGroup(scan *bufio.Scanner) map[string]any {
	m := map[string]any{}
	groupIndent := -1
	for {
		line := scan.Text()
		if line == "---" { // end of entry
			break
		}
		k, v, found := strings.Cut(line, ":")
		if !found {
			panic(fmt.Sprintf("no ':' in line %q", line))
		}
		indent := strings.IndexFunc(k, func(r rune) bool {
			return !unicode.IsSpace(r)
		})
		if indent < 0 {
			panic("blank line")
		}
		if groupIndent < 0 {
			// First line in group; remember the indent.
			groupIndent = indent
		} else if indent < groupIndent {
			// End of group
			break
		} else if indent > groupIndent {
			panic(fmt.Sprintf("indent increased on line %q", line))
		}

		key := strings.TrimSpace(k)
		if v == "" {
			// Just a key: start of a group.
			if !scan.Scan() {
				panic("empty group")
			}
			m[key] = parseGroup(scan)
		} else {
			v = strings.TrimSpace(v)
			if len(v) > 0 && v[0] == '"' {
				var err error
				v, err = strconv.Unquote(v)
				if err != nil {
					panic(err)
				}
			}
			m[key] = v
			if !scan.Scan() {
				break
			}
		}
	}
	return m
}

func TestParseLogEntries(t *testing.T) {
	in := `
a: 1
b: 2
c: 3
g:
    h: 4
    i: 5
d: 6
---
e: 7
---
`
	want := []map[string]any{
		{
			"a": "1",
			"b": "2",
			"c": "3",
			"g": map[string]any{
				"h": "4",
				"i": "5",
			},
			"d": "6",
		},
		{
			"e": "7",
		},
	}
	got := parseLogEntries(in[1:])
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
	}
}
```

### slog-handler-guide/indenthandler3/indent\_handler.go

```go
//go:build go1.21

package indenthandler

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"slices"
	"sync"
	"time"
)

// !+IndentHandler
type IndentHandler struct {
	opts           Options
	preformatted   []byte   // data from WithGroup and WithAttrs
	unopenedGroups []string // groups from WithGroup that haven't been opened
	indentLevel    int      // same as number of opened groups so far
	mu             *sync.Mutex
	out            io.Writer
}

//!-IndentHandler

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// !+WithGroup
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := *h
	// Add an unopened group to h2 without modifying h.
	h2.unopenedGroups = make([]string, len(h.unopenedGroups)+1)
	copy(h2.unopenedGroups, h.unopenedGroups)
	h2.unopenedGroups[len(h2.unopenedGroups)-1] = name
	return &h2
}

//!-WithGroup

// !+WithAttrs
func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := *h
	// Force an append to copy the underlying array.
	pre := slices.Clip(h.preformatted)
	// Add all groups from WithGroup that haven't already been added.
	h2.preformatted = h2.appendUnopenedGroups(pre, h2.indentLevel)
	// Each of those groups increased the indent level by 1.
	h2.indentLevel += len(h2.unopenedGroups)
	// Now all groups have been opened.
	h2.unopenedGroups = nil
	// Pre-format the attributes.
	for _, a := range attrs {
		h2.preformatted = h2.appendAttr(h2.preformatted, a, h2.indentLevel)
	}
	return &h2
}

func (h *IndentHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
	for _, g := range h.unopenedGroups {
		buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
		indentLevel++
	}
	return buf
}

//!-WithAttrs

// !+Handle
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	// Insert preformatted attributes just after built-in ones.
	buf = append(buf, h.preformatted...)
	if r.NumAttrs() > 0 {
		buf = h.appendUnopenedGroups(buf, h.indentLevel)
		r.Attrs(func(a slog.Attr) bool {
			buf = h.appendAttr(buf, a, h.indentLevel+len(h.unopenedGroups))
			return true
		})
	}
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

//!-Handle

func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else.
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs.
	if a.Equal(slog.Attr{}) {
		return buf
	}
	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339Nano))
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, indentLevel)
		}
	default:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
	}
	return buf
}
```

### slog-handler-guide/indenthandler3/indent\_handler\_test.go

```go
//go:build go1.21

package indenthandler

import (
	"bytes"
	"log/slog"
	"reflect"
	"regexp"
	"testing"
	"testing/slogtest"

	"gopkg.in/yaml.v3"
)

// !+TestSlogtest
func TestSlogtest(t *testing.T) {
	var buf bytes.Buffer
	err := slogtest.TestHandler(New(&buf, nil), func() []map[string]any {
		return parseLogEntries(t, buf.Bytes())
	})
	if err != nil {
		t.Error(err)
	}
}

// !-TestSlogtest

func Test(t *testing.T) {
	var buf bytes.Buffer
	l := slog.New(New(&buf, nil))
	l.Info("hello", "a", 1, "b", true, "c", 3.14, slog.Group("g", "h", 1, "i", 2), "d", "NO")
	got := buf.String()
	wantre := `time: [-0-9T:.]+Z?
level: INFO
source: ".*/indent_handler_test.go:\d+"
msg: "hello"
a: 1
b: true
c: 3.14
g:
    h: 1
    i: 2
d: "NO"
`
	re := regexp.MustCompile(wantre)
	if !re.MatchString(got) {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, wantre)
	}

	buf.Reset()
	l.Debug("test")
	if got := buf.Len(); got != 0 {
		t.Errorf("got buf.Len() = %d, want 0", got)
	}
}

// !+parseLogEntries
func parseLogEntries(t *testing.T, data []byte) []map[string]any {
	entries := bytes.Split(data, []byte("---\n"))
	entries = entries[:len(entries)-1] // last one is empty
	var ms []map[string]any
	for _, e := range entries {
		var m map[string]any
		if err := yaml.Unmarshal([]byte(e), &m); err != nil {
			t.Fatal(err)
		}
		ms = append(ms, m)
	}
	return ms
}

// !-parseLogEntries

func TestParseLogEntries(t *testing.T) {
	in := `
a: 1
b: 2
c: 3
g:
    h: 4
    i: five
d: 6
---
e: 7
---
`
	want := []map[string]any{
		{
			"a": 1,
			"b": 2,
			"c": 3,
			"g": map[string]any{
				"h": 4,
				"i": "five",
			},
			"d": 6,
		},
		{
			"e": 7,
		},
	}
	got := parseLogEntries(t, []byte(in[1:]))
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
	}
}
```

### slog-handler-guide/indenthandler4/indent\_handler.go

```go
//go:build go1.21

package indenthandler

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"slices"
	"strconv"
	"sync"
	"time"
)

// !+IndentHandler
type IndentHandler struct {
	opts           Options
	preformatted   []byte   // data from WithGroup and WithAttrs
	unopenedGroups []string // groups from WithGroup that haven't been opened
	indentLevel    int      // same as number of opened groups so far
	mu             *sync.Mutex
	out            io.Writer
}

//!-IndentHandler

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// !+WithGroup
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := *h
	// Add an unopened group to h2 without modifying h.
	h2.unopenedGroups = make([]string, len(h.unopenedGroups)+1)
	copy(h2.unopenedGroups, h.unopenedGroups)
	h2.unopenedGroups[len(h2.unopenedGroups)-1] = name
	return &h2
}

//!-WithGroup

// !+WithAttrs
func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := *h
	// Force an append to copy the underlying array.
	pre := slices.Clip(h.preformatted)
	// Add all groups from WithGroup that haven't already been added.
	h2.preformatted = h2.appendUnopenedGroups(pre, h2.indentLevel)
	// Each of those groups increased the indent level by 1.
	h2.indentLevel += len(h2.unopenedGroups)
	// Now all groups have been opened.
	h2.unopenedGroups = nil
	// Pre-format the attributes.
	for _, a := range attrs {
		h2.preformatted = h2.appendAttr(h2.preformatted, a, h2.indentLevel)
	}
	return &h2
}

func (h *IndentHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
	for _, g := range h.unopenedGroups {
		buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
		indentLevel++
	}
	return buf
}

//!-WithAttrs

// !+Handle
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	bufp := allocBuf()
	buf := *bufp
	defer func() {
		*bufp = buf
		freeBuf(bufp)
	}()
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		// Optimize to minimize allocation.
		srcbufp := allocBuf()
		defer freeBuf(srcbufp)
		*srcbufp = append(*srcbufp, f.File...)
		*srcbufp = append(*srcbufp, ':')
		*srcbufp = strconv.AppendInt(*srcbufp, int64(f.Line), 10)
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, string(*srcbufp)), 0)
	}

	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	// Insert preformatted attributes just after built-in ones.
	buf = append(buf, h.preformatted...)
	if r.NumAttrs() > 0 {
		buf = h.appendUnopenedGroups(buf, h.indentLevel)
		r.Attrs(func(a slog.Attr) bool {
			buf = h.appendAttr(buf, a, h.indentLevel+len(h.unopenedGroups))
			return true
		})
	}
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

//!-Handle

func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else.
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs.
	if a.Equal(slog.Attr{}) {
		return buf
	}
	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = strconv.AppendQuote(buf, a.Value.String())
		buf = append(buf, '\n')
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = a.Value.Time().AppendFormat(buf, time.RFC3339Nano)
		buf = append(buf, '\n')
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, indentLevel)
		}

	default:
		buf = append(buf, a.Key...)
		buf = append(buf, ": "...)
		buf = append(buf, a.Value.String()...)
		buf = append(buf, '\n')
	}
	return buf
}

// !+pool
var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 1024)
		return &b
	},
}

func allocBuf() *[]byte {
	return bufPool.Get().(*[]byte)
}

func freeBuf(b *[]byte) {
	// To reduce peak allocation, return only smaller buffers to the pool.
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}

//!-pool
```

### slog-handler-guide/indenthandler4/indent\_handler\_norace\_test.go

```go
//go:build go1.21 && !race

package indenthandler

import (
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"testing"
)

func TestAlloc(t *testing.T) {
	a := slog.String("key", "value")
	t.Run("Appendf", func(t *testing.T) {
		buf := make([]byte, 0, 100)
		g := testing.AllocsPerRun(2, func() {
			buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
		})
		if g, w := int(g), 2; g != w {
			t.Errorf("got %d, want %d", g, w)
		}
	})
	t.Run("appends", func(t *testing.T) {
		buf := make([]byte, 0, 100)
		g := testing.AllocsPerRun(2, func() {
			buf = append(buf, a.Key...)
			buf = append(buf, ": "...)
			buf = strconv.AppendQuote(buf, a.Value.String())
			buf = append(buf, '\n')
		})
		if g, w := int(g), 0; g != w {
			t.Errorf("got %d, want %d", g, w)
		}
	})

	t.Run("Handle", func(t *testing.T) {
		l := slog.New(New(io.Discard, nil))
		got := testing.AllocsPerRun(10, func() {
			l.LogAttrs(nil, slog.LevelInfo, "hello", slog.String("a", "1"))
		})
		if g, w := int(got), 6; g > w {
			t.Errorf("got %d, want at most %d", g, w)
		}
	})
}
```

### slog-handler-guide/indenthandler4/indent\_handler\_test.go

```go
//go:build go1.21

package indenthandler

import (
	"bytes"
	"log/slog"
	"reflect"
	"regexp"
	"testing"
	"testing/slogtest"

	"gopkg.in/yaml.v3"
)

// !+TestSlogtest
func TestSlogtest(t *testing.T) {
	var buf bytes.Buffer
	err := slogtest.TestHandler(New(&buf, nil), func() []map[string]any {
		return parseLogEntries(t, buf.Bytes())
	})
	if err != nil {
		t.Error(err)
	}
}

// !-TestSlogtest

func Test(t *testing.T) {
	var buf bytes.Buffer
	l := slog.New(New(&buf, nil))
	l.Info("hello", "a", 1, "b", true, "c", 3.14, slog.Group("g", "h", 1, "i", 2), "d", "NO")
	got := buf.String()
	wantre := `time: [-0-9T:.]+Z?
level: INFO
source: ".*/indent_handler_test.go:\d+"
msg: "hello"
a: 1
b: true
c: 3.14
g:
    h: 1
    i: 2
d: "NO"
`
	re := regexp.MustCompile(wantre)
	if !re.MatchString(got) {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, wantre)
	}

	buf.Reset()
	l.Debug("test")
	if got := buf.Len(); got != 0 {
		t.Errorf("got buf.Len() = %d, want 0", got)
	}
}

// !+parseLogEntries
func parseLogEntries(t *testing.T, data []byte) []map[string]any {
	entries := bytes.Split(data, []byte("---\n"))
	entries = entries[:len(entries)-1] // last one is empty
	var ms []map[string]any
	for _, e := range entries {
		var m map[string]any
		if err := yaml.Unmarshal([]byte(e), &m); err != nil {
			t.Fatal(err)
		}
		ms = append(ms, m)
	}
	return ms
}

// !-parseLogEntries

func TestParseLogEntries(t *testing.T) {
	in := `
a: 1
b: 2
c: 3
g:
    h: 4
    i: five
d: 6
---
e: 7
---
`
	want := []map[string]any{
		{
			"a": 1,
			"b": 2,
			"c": 3,
			"g": map[string]any{
				"h": 4,
				"i": "five",
			},
			"d": 6,
		},
		{
			"e": 7,
		},
	}
	got := parseLogEntries(t, []byte(in[1:]))
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
	}
}
```

### slog-handler-guide/guide.md

```markdown

# A Guide to Writing `slog` Handlers

This document is maintained by Jonathan Amsterdam `jba@google.com`.


# Contents

%toc


# Introduction

The standard library’s `log/slog` package has a two-part design.
A "frontend," implemented by the `Logger` type,
gathers structured log information like a message, level, and attributes,
and passes them to a "backend," an implementation of the `Handler` interface.
The package comes with two built-in handlers that should usually be adequate.
But you may need to write your own handler, and that is not always straightforward.
This guide is here to help.


# Loggers and their handlers

Writing a handler requires an understanding of how the `Logger` and `Handler`
types work together.

Each logger contains a handler. Certain `Logger` methods do some preliminary work,
such as gathering key-value pairs into `Attr`s, and then call one or more
`Handler` methods. These `Logger` methods are `With`, `WithGroup`,
and the output methods like `Info`, `Error` and so on.

An output method fulfills the main role of a logger: producing log output.
Here is a call to an output method:

    logger.Info("hello", "key", value)

There are two general output methods, `Log`, and `LogAttrs`. For convenience,
there is an output method for each of four common levels (`Debug`, `Info`,
`Warn` and `Error`), and corresponding methods that take a context (`DebugContext`,
`InfoContext`, `WarnContext` and `ErrorContext`).

Each `Logger` output method first calls its handler's `Enabled` method. If that call
returns true, the method constructs a `Record` from its arguments and calls
the handler's `Handle` method.

As a convenience and an optimization, attributes can be added to
`Logger` by calling the `With` method:

    logger = logger.With("k", v)

This call creates a new `Logger` value with the argument attributes; the
original remains unchanged.
All subsequent output from `logger` will include those attributes.
A logger's `With` method calls its handler's `WithAttrs` method.

The `WithGroup` method is used to avoid key collisions in large programs
by establishing separate namespaces.
This call creates a new `Logger` value with a group named "g":

    logger = logger.WithGroup("g")

All subsequent keys for `logger` will be qualified by the group name "g".
Exactly what "qualified" means depends on how the logger's handler formats the
output.
The built-in `TextHandler` treats the group as a prefix to the key, separated by
a dot: `g.k` for a key `k`, for example.
The built-in `JSONHandler` uses the group as a key for a nested JSON object:

    {"g": {"k": v}}

A logger's `WithGroup` method calls its handler's `WithGroup` method.


# Implementing `Handler` methods

We can now talk about the four `Handler` methods in detail.
Along the way, we will write a handler that formats logs using a format
reminiscent of YAML. It will display this log output call:

    logger.Info("hello", "key", 23)

something like this:

    time: 2023-05-15T16:29:00
    level: INFO
    message: "hello"
    key: 23
    ---

Although this particular output is valid YAML,
our implementation doesn't consider the subtleties of YAML syntax,
so it will sometimes produce invalid YAML.
For example, it doesn't quote keys that have colons in them.
We'll call it `IndentHandler` to forestall disappointment.

A brief aside before we start: it is tempting to embed `slog.Handler` in your
custom handler and implement only the methods that you need.
Loggers and handlers are too tightly coupled for that to work. You should
implement all four handler methods.

We begin with the `IndentHandler` type
and the `New` function that constructs it from an `io.Writer` and options:

%include indenthandler1/indent_handler.go types -

We'll support only one option, the ability to set a minimum level in order to
suppress detailed log output.
Handlers should always declare this option to be a `slog.Leveler`.
The `slog.Leveler` interface is implemented by both `Level` and `LevelVar`.
A `Level` value is easy for the user to provide,
but changing the level of multiple handlers requires tracking them all.
If the user instead passes a `LevelVar`, then a single change to that `LevelVar`
will change the behavior of all handlers that contain it.
Changes to `LevelVar`s are goroutine-safe.

You might also consider adding a `ReplaceAttr` option to your handler,
like the [one for the built-in
handlers](https://pkg.go.dev/log/slog#HandlerOptions.ReplaceAttr).
Although `ReplaceAttr` will complicate your implementation, it will also
make your handler more generally useful.

The mutex will be used to ensure that writes to the `io.Writer` happen atomically.
Unusually, `IndentHandler` holds a pointer to a `sync.Mutex` rather than holding a
`sync.Mutex` directly.
There is a good reason for that, which we'll explain [later](#getting-the-mutex-right).

Our handler will need additional state to track calls to `WithGroup` and `WithAttrs`.
We will describe that state when we get to those methods.

## The `Enabled` method

The `Enabled` method is an optimization that can avoid unnecessary work.
A `Logger` output method will call `Enabled` before it processes any of its arguments,
to see if it should proceed.

The signature is

    Enabled(context.Context, Level) bool

The context is available to allow decisions based on contextual information.
For example, a custom HTTP request header could specify a minimum level,
which the server adds to the context used for processing that request.
A handler's `Enabled` method could report whether the argument level
is greater than or equal to the context value, allowing the verbosity
of the work done by each request to be controlled independently.

Our `IndentHandler` doesn't use the context. It just compares the argument level
with its configured minimum level:

%include indenthandler1/indent_handler.go enabled -

## The `Handle` method

The `Handle` method is passed a `Record` containing all the information to be
logged for a single call to a `Logger` output method.
The `Handle` method should deal with it in some way.
One way is to output the `Record` in some format, as `TextHandler` and `JSONHandler` do.
But other options are to modify the `Record` and pass it on to another handler,
enqueue the `Record` for later processing, or ignore it.

The signature of `Handle` is

    Handle(context.Context, Record) error

The context is provided to support applications that provide logging information
along the call chain. In a break with usual Go practice, the `Handle` method
should not treat a canceled context as a signal to stop work.

If `Handle` processes its `Record`, it should follow the rules in the
[documentation](https://pkg.go.dev/log/slog#Handler.Handle).
For example, a zero `Time` field should be ignored, as should zero `Attr`s.

A `Handle` method that is going to produce output should carry out the following steps:

1. Allocate a buffer, typically a `[]byte`, to hold the output.
It's best to construct the output in memory first,
then write it with a single call to `io.Writer.Write`,
to minimize interleaving with other goroutines using the same writer.

2. Format the special fields: time, level, message, and source location (PC).
As a general rule, these fields should appear first and are not nested in
groups established by `WithGroup`.

3. Format the result of `WithGroup` and `WithAttrs` calls.

4. Format the attributes in the `Record`.

5. Output the buffer.

That is how `IndentHandler.Handle` is structured:

%include indenthandler1/indent_handler.go handle -

The first line allocates a `[]byte` that should be large enough for most log
output.
Allocating a buffer with some initial, fairly large capacity is a simple but
significant optimization: it avoids the repeated copying and allocation that
happen when the initial slice is empty or small.
We'll return to this line in the section on [speed](#speed)
and show how we can do even better.

The next part of our `Handle` method formats the special attributes,
observing the rules to ignore a zero time and a zero PC.

Next, the method processes the result of `WithAttrs` and `WithGroup` calls.
We'll skip that for now.

Then it's time to process the attributes in the argument record.
We use the `Record.Attrs` method to iterate over the attributes
in the order the user passed them to the `Logger` output method.
Handlers are free to reorder or de-duplicate the attributes,
but ours does not.

Lastly, after adding the line "---" to the output to separate log records,
our handler makes a single call to `h.out.Write` with the buffer we've accumulated.
We hold the lock for this write to make it atomic with respect to other
goroutines that may be calling `Handle` at the same time.

At the heart of the handler is the `appendAttr` method, responsible for
formatting a single attribute:

%include indenthandler1/indent_handler.go appendAttr -

It begins by resolving the attribute, to run the `LogValuer.LogValue` method of
the value if it has one. All handlers should resolve every attribute they
process.

Next, it follows the handler rule that says that empty attributes should be
ignored.

Then it switches on the attribute kind to determine what format to use. For most
kinds (the default case of the switch), it relies on `slog.Value`'s `String` method to
produce something reasonable. It handles strings and times specially:
strings by quoting them, and times by formatting them in a standard way.

When `appendAttr` sees a `Group`, it calls itself recursively on the group's
attributes, after applying two more handler rules.
First, a group with no attributes is ignored&mdash;not even its key is displayed.
Second, a group with an empty key is inlined: the group boundary isn't marked in
any way. In our case, that means the group's attributes aren't indented.

## The `WithAttrs` method

One of `slog`'s performance optimizations is support for pre-formatting
attributes. The `Logger.With` method converts key-value pairs into `Attr`s and
then calls `Handler.WithAttrs`.
The handler may store the attributes for later consumption by the `Handle` method,
or it may take the opportunity to format the attributes now, once,
rather than doing so repeatedly on each call to `Handle`.

The signature of the `WithAttrs` method is

    WithAttrs(attrs []Attr) Handler

The attributes are the processed key-value pairs passed to `Logger.With`.
The return value should be a new instance of your handler that contains
the attributes, possibly pre-formatted.

`WithAttrs` must return a new handler with the additional attributes, leaving
the original handler (its receiver) unchanged. For example, this call:

    logger2 := logger1.With("k", v)

creates a new logger, `logger2`, with an additional attribute, but has no
effect on `logger1`.

We will show example implementations of `WithAttrs` below, when we discuss `WithGroup`.

## The `WithGroup` method

`Logger.WithGroup` calls `Handler.WithGroup` directly, with the same
argument, the group name.
A handler should remember the name so it can use it to qualify all subsequent
attributes.

The signature of `WithGroup` is:

    WithGroup(name string) Handler

Like `WithAttrs`, the `WithGroup` method should return a new handler, not modify
the receiver.

The implementations of `WithGroup` and `WithAttrs` are intertwined.
Consider this statement:

    logger = logger.WithGroup("g1").With("k1", 1).WithGroup("g2").With("k2", 2)

Subsequent `logger` output should qualify key "k1" with group "g1",
and key "k2" with groups "g1" and "g2".
The order of the `Logger.WithGroup` and `Logger.With` calls must be respected by
the implementations of `Handler.WithGroup` and `Handler.WithAttrs`.

We will look at two implementations of `WithGroup` and `WithAttrs`, one that pre-formats and
one that doesn't.

### Without pre-formatting

Our first implementation will collect the information from `WithGroup` and
`WithAttrs` calls to build up a slice of group names and attribute lists,
and loop over that slice in `Handle`. We start with a struct that can hold
either a group name or some attributes:

%include indenthandler2/indent_handler.go gora -

Then we add a slice of `groupOrAttrs` to our handler:

%include indenthandler2/indent_handler.go IndentHandler -

As stated above, The `WithGroup` and `WithAttrs` methods should not modify their
receiver.
To that end, we define a method that will copy our handler struct
and append one `groupOrAttrs` to the copy:

%include indenthandler2/indent_handler.go withgora -

Most of the fields of `IndentHandler` can be copied shallowly, but the slice of
`groupOrAttrs` requires a deep copy, or the clone and the original will point to
the same underlying array. If we used `append` instead of making an explicit
copy, we would introduce that subtle aliasing bug.

The `With` methods are easy to write using `withGroupOrAttrs`:

%include indenthandler2/indent_handler.go withs -

The `Handle` method can now process the groupOrAttrs slice after
the built-in attributes and before the ones in the record:

%include indenthandler2/indent_handler.go handle -

You may have noticed that our algorithm for
recording `WithGroup` and `WithAttrs` information is quadratic in the
number of calls to those methods, because of the repeated copying.
That is unlikely to matter in practice, but if it bothers you,
you can use a linked list instead,
which `Handle` will have to reverse or visit recursively.
See the
[github.com/jba/slog/withsupport](https://github.com/jba/slog/tree/main/withsupport)
package for an implementation.

#### Getting the mutex right

Let us revisit the last few lines of `Handle`:

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
    return err

This code hasn't changed, but we can now appreciate why `h.mu` is a
pointer to a `sync.Mutex`. Both `WithGroup` and `WithAttrs` copy the handler.
Both copies point to the same mutex.
If the copy and the original used different mutexes and were used concurrently,
then their output could be interleaved, or some output could be lost.
Code like this:

    l2 := l1.With("a", 1)
    go l1.Info("hello")
    l2.Info("goodbye")

could produce output like this:

    hegoollo a=dbye1

See [this bug report](https://go.dev/issue/61321) for more detail.


### With pre-formatting

Our second version of the `WithGroup` and `WithAttrs` methods provides pre-formatting.
This implementation is more complicated than the previous one.
Is the extra complexity worth it?
That depends on your circumstances, but here is one circumstance where
it might be.
Say that you wanted your server to log a lot of information about an incoming
request with every log message that happens during that request. A typical
handler might look something like this:

    func (s *Server) handleWidgets(w http.ResponseWriter, r *http.Request) {
        logger := s.logger.With(
            "url", r.URL,
            "traceID": r.Header.Get("X-Cloud-Trace-Context"),
            // many other attributes
            )
        // ...
    }

A single `handleWidgets` request might generate hundreds of log lines.
For instance, it might contain code like this:

    for _, w := range widgets {
        logger.Info("processing widget", "name", w.Name)
        // ...
    }

For every such line, the `Handle` method we wrote above will format all
the attributes that were added using `With` above, in addition to the
ones on the log line itself.

Maybe all that extra work doesn't slow down your server significantly, because
it does so much other work that time spent logging is just noise.
But perhaps your server is fast enough that all that extra formatting appears high up
in your CPU profiles. That is when pre-formatting can make a big difference,
by formatting the attributes in a call to `With` just once.

To pre-format the arguments to `WithAttrs`, we need to keep track of some
additional state in the `IndentHandler` struct.

%include indenthandler3/indent_handler.go IndentHandler -

Mainly, we need a buffer to hold the pre-formatted data.
But we also need to keep track of which groups
we've seen but haven't output yet. We'll call those groups "unopened."
We also need to track how many groups we've opened, which we can do
with a simple counter, since an opened group's only effect is to change the
indentation level.

This `WithGroup` is a lot like the previous one: it just remembers the
new group, which is unopened initially.

%include indenthandler3/indent_handler.go WithGroup -

`WithAttrs` does all the pre-formatting:

%include indenthandler3/indent_handler.go WithAttrs -

It first opens any unopened groups. This handles calls like:

    logger.WithGroup("g").WithGroup("h").With("a", 1)

Here, `WithAttrs` must output "g" and "h" before "a". Since a group established
by `WithGroup` is in effect for the rest of the log line, `WithAttrs` increments
the indentation level for each group it opens.

Lastly, `WithAttrs` formats its argument attributes, using the same `appendAttr`
method we saw above.

It's the `Handle` method's job to insert the pre-formatted material in the right
place, which is after the built-in attributes and before the ones in the record:

%include indenthandler3/indent_handler.go Handle -

It must also open any groups that haven't yet been opened. The logic covers
log lines like this one:

    logger.WithGroup("g").Info("msg", "a", 1)

where "g" is unopened before `Handle` is called and must be written to produce
the correct output:

    level: INFO
    msg: "msg"
    g:
        a: 1

The check for `r.NumAttrs() > 0` handles this case:

    logger.WithGroup("g").Info("msg")

Here there are no record attributes, so no group to open.

## Testing

The [`Handler` contract](https://pkg.go.dev/log/slog#Handler) specifies several
constraints on handlers.
To verify that your handler follows these rules and generally produces proper
output, use the [testing/slogtest package](https://pkg.go.dev/testing/slogtest).

That package's `TestHandler` function takes an instance of your handler and
a function that returns its output formatted as a slice of maps. Here is the test function
for our example handler:

%include indenthandler3/indent_handler_test.go TestSlogtest -

Calling `TestHandler` is easy. The hard part is parsing your handler's output.
`TestHandler` calls your handler multiple times, resulting in a sequence of log
entries.
It is your job to parse each entry into a `map[string]any`.
A group in an entry should appear as a nested map.

If your handler outputs a standard format, you can use an existing parser.
For example, if your handler outputs one JSON object per line, then you
can split the output into lines and call `encoding/json.Unmarshal` on each.
Parsers for other formats that can unmarshal into a map can be used out
of the box.
Our example output is enough like YAML so that we can use the `gopkg.in/yaml.v3`
package to parse it:

%include indenthandler3/indent_handler_test.go parseLogEntries -

If you have to write your own parser, it can be far from perfect.
The `slogtest` package uses only a handful of simple attributes.
(It is testing handler conformance, not parsing.)
Your parser can ignore edge cases like whitespace and newlines in keys and
values. Before switching to a YAML parser, we wrote an adequate custom parser
in 65 lines.

# General considerations

## Copying records

Most handlers won't need to copy the `slog.Record` that is passed
to the `Handle` method.
Those that do must take special care in some cases.

A handler can make a single copy of a `Record` with an ordinary Go
assignment, channel send or function call if it doesn't retain the
original.
But if its actions result in more than one copy, it should call `Record.Clone`
to make the copies so that they don't share state.
This `Handle` method passes the record to a single handler, so it doesn't require `Clone`:

    type Handler1 struct {
        h slog.Handler
        // ...
    }

    func (h *Handler1) Handle(ctx context.Context, r slog.Record) error {
        return h.h.Handle(ctx, r)
    }

This `Handle` method might pass the record to more than one handler, so it
should use `Clone`:

    type Handler2 struct {
        hs []slog.Handler
        // ...
    }

    func (h *Handler2) Handle(ctx context.Context, r slog.Record) error {
        for _, hh := range h.hs {
            if err := hh.Handle(ctx, r.Clone()); err != nil {
                return err
            }
        }
        return nil
    }

## Concurrency safety

A handler must work properly when a single `Logger` is shared among several
goroutines.
That means that mutable state must be protected with a lock or some other mechanism.
In practice, this is not hard to achieve, because many handlers won't have any
mutable state.

- The `Enabled` method typically consults only its arguments and a configured
  level. The level is often either set once initially, or is held in a
  `LevelVar`, which is already concurrency-safe.

- The `WithAttrs` and `WithGroup` methods should not modify the receiver,
  for reasons discussed above.

- The `Handle` method typically works only with its arguments and stored fields.

Calls to output methods like `io.Writer.Write` should be synchronized unless
it can be verified that no locking is needed.
As we saw in our example, storing a pointer to a mutex enables a logger and
all of its clones to synchronize with each other.
Beware of facile claims like "Unix writes are atomic"; the situation is a lot more nuanced than that.

Some handlers have legitimate reasons for keeping state.
For example, a handler might support a `SetLevel` method to change its configured level
dynamically.
Or it might output the time between successive calls to `Handle`,
which requires a mutable field holding the last output time.
Synchronize all accesses to such fields, both reads and writes.

The built-in handlers have no directly mutable state.
They use a mutex only to sequence calls to their contained `io.Writer`.

## Robustness

Logging is often the debugging technique of last resort. When it is difficult or
impossible to inspect a system, as is typically the case with a production
server, logs provide the most detailed way to understand its behavior.
Therefore, your handler should be robust to bad input.

For example, the usual advice when a function discovers a problem,
like an invalid argument, is to panic or return an error.
The built-in handlers do not follow that advice.
Few things are more frustrating than being unable to debug a problem that
causes logging to fail;
it is better to produce some output, however imperfect, than to produce none at all.
That is why methods like `Logger.Info` convert programming bugs in their list of
key-value pairs, like missing values or malformed keys,
into `Attr`s that contain as much information as possible.

One place to avoid panics is in processing attribute values. A handler that wants
to format a value will probably switch on the value's kind:

    switch attr.Value.Kind() {
    case KindString: ...
    case KindTime: ...
    // all other Kinds
    default: ...
    }

What should happen in the default case, when the handler encounters a `Kind`
that it doesn't know about?
The built-in handlers try to muddle through by using the result of the value's
`String` method, as our example handler does.
They do not panic or return an error.
Your own handlers might in addition want to report the problem through your production monitoring
or error-tracking telemetry system.
The most likely explanation for the issue is that a newer version of the `slog` package added
a new `Kind`&mdash;a backwards-compatible change under the Go 1 Compatibility
Promise&mdash;and the handler wasn't updated.
That is certainly a problem, but it shouldn't deprive
readers from seeing the rest of the log output.

There is one circumstance where returning an error from `Handler.Handle` is appropriate.
If the output operation itself fails, the best course of action is to report
this failure by returning the error. For instance, the last two lines of the
built-in `Handle` methods are

    _, err := h.w.Write(*state.buf)
    return err

Although the output methods of `Logger` ignore the error, one could write a
handler that does something with it, perhaps falling back to writing to standard
error.

## Speed

Most programs don't need fast logging.
Before making your handler fast, gather data&mdash;preferably production data,
not benchmark comparisons&mdash;that demonstrates that it needs to be fast.
Avoid premature optimization.

If you need a fast handler, start with pre-formatting. It may provide dramatic
speed-ups in cases where a single call to `Logger.With` is followed by many
calls to the resulting logger.

If log output is the bottleneck, consider making your handler asynchronous.
Do the minimal amount of processing in the handler, then send the record and
other information over a channel. Another goroutine can collect the incoming log
entries and write them in bulk and in the background.
You might want to preserve the option to log synchronously
so you can see all the log output to debug a crash.

Allocation is often a major cause of a slow system.
The `slog` package already works hard at minimizing allocations.
If your handler does its own allocation, and profiling shows it to be
a problem, then see if you can minimize it.

One simple change you can make is to replace calls to `fmt.Sprintf` or `fmt.Appendf`
with direct appends to the buffer. For example, our IndentHandler appends string
attributes to the buffer like so:

	buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())

As of Go 1.21, that results in two allocations, one for each argument passed to
an `any` parameter. We can get that down to zero by using `append` directly:

	buf = append(buf, a.Key...)
	buf = append(buf, ": "...)
	buf = strconv.AppendQuote(buf, a.Value.String())
	buf = append(buf, '\n')

Another worthwhile change is to use a `sync.Pool` to manage the one chunk of
memory that most handlers need:
the `[]byte` buffer holding the formatted output.

Our example `Handle` method began with this line:

	buf := make([]byte, 0, 1024)

As we said above, providing a large initial capacity avoids repeated copying and
re-allocation of the slice as it grows, reducing the number of allocations to
one.
But we can get it down to zero in the steady state by keeping a global pool of buffers.
Initially, the pool will be empty and new buffers will be allocated.
But eventually, assuming the number of concurrent log calls reaches a steady
maximum, there will be enough buffers in the pool to share among all the
ongoing `Handler` calls. As long as no log entry grows past a buffer's capacity,
there will be no allocations from the garbage collector's point of view.

We will hide our pool behind a pair of functions, `allocBuf` and `freeBuf`.
The single line to get a buffer at the top of `Handle` becomes two lines:

	bufp := allocBuf()
	defer freeBuf(bufp)

One of the subtleties involved in making a `sync.Pool` of slices
is suggested by the variable name `bufp`: your pool must deal in
_pointers_ to slices, not the slices themselves.
Pooled values must always be pointers. If they aren't, then the `any` arguments
and return values of the `sync.Pool` methods will themselves cause allocations,
defeating the purpose of pooling.

There are two ways to proceed with our slice pointer: we can replace `buf`
with `*bufp` throughout our function, or we can dereference it and remember to
re-assign it before freeing:

	bufp := allocBuf()
	buf := *bufp
	defer func() {
		*bufp = buf
		freeBuf(bufp)
	}()


Here is our pool and its associated functions:

%include indenthandler4/indent_handler.go pool -

The pool's `New` function does the same thing as the original code:
create a byte slice with 0 length and plenty of capacity.
The `allocBuf` function just type-asserts the result of the pool's
`Get` method.

The `freeBuf` method truncates the buffer before putting it back
in the pool, so that `allocBuf` always returns zero-length slices.
It also implements an important optimization: it doesn't return
large buffers to the pool.
To see why this important, consider what would happen if there were a single,
unusually large log entry&mdash;say one that was a megabyte when formatted.
If that megabyte-sized buffer were put in the pool, it could remain
there indefinitely, constantly being reused, but with most of its capacity
wasted.
The extra memory might never be used again by the handler, and since it was in
the handler's pool, it might never be garbage-collected for reuse elsewhere.
We can avoid that situation by excluding large buffers from the pool.
```

### slog-handler-guide/Makefile

```

all: README.md
	go build ./...

README.md: guide.md ../internal/cmd/weave/*.go $(wildcard */*.go)
	go run ../internal/cmd/weave guide.md >README.md

# This is for previewing using github.
# $HOME/markdown must be a github client.
test: README.md
	cp README.md $$HOME/markdown/
	(cd $$HOME/markdown/ && git commit -m . README.md && git push)

```

### slog-handler-guide/README.md

```markdown
<!-- Autogenerated by weave; DO NOT EDIT -->

# A Guide to Writing `slog` Handlers

This document is maintained by Jonathan Amsterdam `jba@google.com`.


# Contents

1. [Introduction](#introduction)
1. [Loggers and their handlers](#loggers-and-their-handlers)
1. [Implementing `Handler` methods](#implementing-handler-methods)
	1. [The `Enabled` method](#the-enabled-method)
	1. [The `Handle` method](#the-handle-method)
	1. [The `WithAttrs` method](#the-withattrs-method)
	1. [The `WithGroup` method](#the-withgroup-method)
	1. [Testing](#testing)
1. [General considerations](#general-considerations)
	1. [Copying records](#copying-records)
	1. [Concurrency safety](#concurrency-safety)
	1. [Robustness](#robustness)
	1. [Speed](#speed)


# Introduction

The standard library’s `log/slog` package has a two-part design.
A "frontend," implemented by the `Logger` type,
gathers structured log information like a message, level, and attributes,
and passes them to a "backend," an implementation of the `Handler` interface.
The package comes with two built-in handlers that should usually be adequate.
But you may need to write your own handler, and that is not always straightforward.
This guide is here to help.


# Loggers and their handlers

Writing a handler requires an understanding of how the `Logger` and `Handler`
types work together.

Each logger contains a handler. Certain `Logger` methods do some preliminary work,
such as gathering key-value pairs into `Attr`s, and then call one or more
`Handler` methods. These `Logger` methods are `With`, `WithGroup`,
and the output methods like `Info`, `Error` and so on.

An output method fulfills the main role of a logger: producing log output.
Here is a call to an output method:

    logger.Info("hello", "key", value)

There are two general output methods, `Log`, and `LogAttrs`. For convenience,
there is an output method for each of four common levels (`Debug`, `Info`,
`Warn` and `Error`), and corresponding methods that take a context (`DebugContext`,
`InfoContext`, `WarnContext` and `ErrorContext`).

Each `Logger` output method first calls its handler's `Enabled` method. If that call
returns true, the method constructs a `Record` from its arguments and calls
the handler's `Handle` method.

As a convenience and an optimization, attributes can be added to
`Logger` by calling the `With` method:

    logger = logger.With("k", v)

This call creates a new `Logger` value with the argument attributes; the
original remains unchanged.
All subsequent output from `logger` will include those attributes.
A logger's `With` method calls its handler's `WithAttrs` method.

The `WithGroup` method is used to avoid key collisions in large programs
by establishing separate namespaces.
This call creates a new `Logger` value with a group named "g":

    logger = logger.WithGroup("g")

All subsequent keys for `logger` will be qualified by the group name "g".
Exactly what "qualified" means depends on how the logger's handler formats the
output.
The built-in `TextHandler` treats the group as a prefix to the key, separated by
a dot: `g.k` for a key `k`, for example.
The built-in `JSONHandler` uses the group as a key for a nested JSON object:

    {"g": {"k": v}}

A logger's `WithGroup` method calls its handler's `WithGroup` method.


# Implementing `Handler` methods

We can now talk about the four `Handler` methods in detail.
Along the way, we will write a handler that formats logs using a format
reminiscent of YAML. It will display this log output call:

    logger.Info("hello", "key", 23)

something like this:

    time: 2023-05-15T16:29:00
    level: INFO
    message: "hello"
    key: 23
    ---

Although this particular output is valid YAML,
our implementation doesn't consider the subtleties of YAML syntax,
so it will sometimes produce invalid YAML.
For example, it doesn't quote keys that have colons in them.
We'll call it `IndentHandler` to forestall disappointment.

A brief aside before we start: it is tempting to embed `slog.Handler` in your
custom handler and implement only the methods that you need.
Loggers and handlers are too tightly coupled for that to work. You should
implement all four handler methods.

We begin with the `IndentHandler` type
and the `New` function that constructs it from an `io.Writer` and options:

```go
type IndentHandler struct {
	opts Options
	// TODO: state for WithGroup and WithAttrs
	mu  *sync.Mutex
	out io.Writer
}

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}
```

We'll support only one option, the ability to set a minimum level in order to
suppress detailed log output.
Handlers should always declare this option to be a `slog.Leveler`.
The `slog.Leveler` interface is implemented by both `Level` and `LevelVar`.
A `Level` value is easy for the user to provide,
but changing the level of multiple handlers requires tracking them all.
If the user instead passes a `LevelVar`, then a single change to that `LevelVar`
will change the behavior of all handlers that contain it.
Changes to `LevelVar`s are goroutine-safe.

You might also consider adding a `ReplaceAttr` option to your handler,
like the [one for the built-in
handlers](https://pkg.go.dev/log/slog#HandlerOptions.ReplaceAttr).
Although `ReplaceAttr` will complicate your implementation, it will also
make your handler more generally useful.

The mutex will be used to ensure that writes to the `io.Writer` happen atomically.
Unusually, `IndentHandler` holds a pointer to a `sync.Mutex` rather than holding a
`sync.Mutex` directly.
There is a good reason for that, which we'll explain [later](#getting-the-mutex-right).

Our handler will need additional state to track calls to `WithGroup` and `WithAttrs`.
We will describe that state when we get to those methods.

## The `Enabled` method

The `Enabled` method is an optimization that can avoid unnecessary work.
A `Logger` output method will call `Enabled` before it processes any of its arguments,
to see if it should proceed.

The signature is

    Enabled(context.Context, Level) bool

The context is available to allow decisions based on contextual information.
For example, a custom HTTP request header could specify a minimum level,
which the server adds to the context used for processing that request.
A handler's `Enabled` method could report whether the argument level
is greater than or equal to the context value, allowing the verbosity
of the work done by each request to be controlled independently.

Our `IndentHandler` doesn't use the context. It just compares the argument level
with its configured minimum level:

```go
func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}
```

## The `Handle` method

The `Handle` method is passed a `Record` containing all the information to be
logged for a single call to a `Logger` output method.
The `Handle` method should deal with it in some way.
One way is to output the `Record` in some format, as `TextHandler` and `JSONHandler` do.
But other options are to modify the `Record` and pass it on to another handler,
enqueue the `Record` for later processing, or ignore it.

The signature of `Handle` is

    Handle(context.Context, Record) error

The context is provided to support applications that provide logging information
along the call chain. In a break with usual Go practice, the `Handle` method
should not treat a canceled context as a signal to stop work.

If `Handle` processes its `Record`, it should follow the rules in the
[documentation](https://pkg.go.dev/log/slog#Handler.Handle).
For example, a zero `Time` field should be ignored, as should zero `Attr`s.

A `Handle` method that is going to produce output should carry out the following steps:

1. Allocate a buffer, typically a `[]byte`, to hold the output.
It's best to construct the output in memory first,
then write it with a single call to `io.Writer.Write`,
to minimize interleaving with other goroutines using the same writer.

2. Format the special fields: time, level, message, and source location (PC).
As a general rule, these fields should appear first and are not nested in
groups established by `WithGroup`.

3. Format the result of `WithGroup` and `WithAttrs` calls.

4. Format the attributes in the `Record`.

5. Output the buffer.

That is how `IndentHandler.Handle` is structured:

```go
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	indentLevel := 0
	// TODO: output the Attrs and groups from WithAttrs and WithGroup.
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a, indentLevel)
		return true
	})
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}
```

The first line allocates a `[]byte` that should be large enough for most log
output.
Allocating a buffer with some initial, fairly large capacity is a simple but
significant optimization: it avoids the repeated copying and allocation that
happen when the initial slice is empty or small.
We'll return to this line in the section on [speed](#speed)
and show how we can do even better.

The next part of our `Handle` method formats the special attributes,
observing the rules to ignore a zero time and a zero PC.

Next, the method processes the result of `WithAttrs` and `WithGroup` calls.
We'll skip that for now.

Then it's time to process the attributes in the argument record.
We use the `Record.Attrs` method to iterate over the attributes
in the order the user passed them to the `Logger` output method.
Handlers are free to reorder or de-duplicate the attributes,
but ours does not.

Lastly, after adding the line "---" to the output to separate log records,
our handler makes a single call to `h.out.Write` with the buffer we've accumulated.
We hold the lock for this write to make it atomic with respect to other
goroutines that may be calling `Handle` at the same time.

At the heart of the handler is the `appendAttr` method, responsible for
formatting a single attribute:

```go
func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else.
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs.
	if a.Equal(slog.Attr{}) {
		return buf
	}
	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339Nano))
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, indentLevel)
		}
	default:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
	}
	return buf
}
```

It begins by resolving the attribute, to run the `LogValuer.LogValue` method of
the value if it has one. All handlers should resolve every attribute they
process.

Next, it follows the handler rule that says that empty attributes should be
ignored.

Then it switches on the attribute kind to determine what format to use. For most
kinds (the default case of the switch), it relies on `slog.Value`'s `String` method to
produce something reasonable. It handles strings and times specially:
strings by quoting them, and times by formatting them in a standard way.

When `appendAttr` sees a `Group`, it calls itself recursively on the group's
attributes, after applying two more handler rules.
First, a group with no attributes is ignored&mdash;not even its key is displayed.
Second, a group with an empty key is inlined: the group boundary isn't marked in
any way. In our case, that means the group's attributes aren't indented.

## The `WithAttrs` method

One of `slog`'s performance optimizations is support for pre-formatting
attributes. The `Logger.With` method converts key-value pairs into `Attr`s and
then calls `Handler.WithAttrs`.
The handler may store the attributes for later consumption by the `Handle` method,
or it may take the opportunity to format the attributes now, once,
rather than doing so repeatedly on each call to `Handle`.

The signature of the `WithAttrs` method is

    WithAttrs(attrs []Attr) Handler

The attributes are the processed key-value pairs passed to `Logger.With`.
The return value should be a new instance of your handler that contains
the attributes, possibly pre-formatted.

`WithAttrs` must return a new handler with the additional attributes, leaving
the original handler (its receiver) unchanged. For example, this call:

    logger2 := logger1.With("k", v)

creates a new logger, `logger2`, with an additional attribute, but has no
effect on `logger1`.

We will show example implementations of `WithAttrs` below, when we discuss `WithGroup`.

## The `WithGroup` method

`Logger.WithGroup` calls `Handler.WithGroup` directly, with the same
argument, the group name.
A handler should remember the name so it can use it to qualify all subsequent
attributes.

The signature of `WithGroup` is:

    WithGroup(name string) Handler

Like `WithAttrs`, the `WithGroup` method should return a new handler, not modify
the receiver.

The implementations of `WithGroup` and `WithAttrs` are intertwined.
Consider this statement:

    logger = logger.WithGroup("g1").With("k1", 1).WithGroup("g2").With("k2", 2)

Subsequent `logger` output should qualify key "k1" with group "g1",
and key "k2" with groups "g1" and "g2".
The order of the `Logger.WithGroup` and `Logger.With` calls must be respected by
the implementations of `Handler.WithGroup` and `Handler.WithAttrs`.

We will look at two implementations of `WithGroup` and `WithAttrs`, one that pre-formats and
one that doesn't.

### Without pre-formatting

Our first implementation will collect the information from `WithGroup` and
`WithAttrs` calls to build up a slice of group names and attribute lists,
and loop over that slice in `Handle`. We start with a struct that can hold
either a group name or some attributes:

```go
// groupOrAttrs holds either a group name or a list of slog.Attrs.
type groupOrAttrs struct {
	group string      // group name if non-empty
	attrs []slog.Attr // attrs if non-empty
}
```

Then we add a slice of `groupOrAttrs` to our handler:

```go
type IndentHandler struct {
	opts Options
	goas []groupOrAttrs
	mu   *sync.Mutex
	out  io.Writer
}
```

As stated above, The `WithGroup` and `WithAttrs` methods should not modify their
receiver.
To that end, we define a method that will copy our handler struct
and append one `groupOrAttrs` to the copy:

```go
func (h *IndentHandler) withGroupOrAttrs(goa groupOrAttrs) *IndentHandler {
	h2 := *h
	h2.goas = make([]groupOrAttrs, len(h.goas)+1)
	copy(h2.goas, h.goas)
	h2.goas[len(h2.goas)-1] = goa
	return &h2
}
```

Most of the fields of `IndentHandler` can be copied shallowly, but the slice of
`groupOrAttrs` requires a deep copy, or the clone and the original will point to
the same underlying array. If we used `append` instead of making an explicit
copy, we would introduce that subtle aliasing bug.

The `With` methods are easy to write using `withGroupOrAttrs`:

```go
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{group: name})
}

func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{attrs: attrs})
}
```

The `Handle` method can now process the groupOrAttrs slice after
the built-in attributes and before the ones in the record:

```go
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	indentLevel := 0
	// Handle state from WithGroup and WithAttrs.
	goas := h.goas
	if r.NumAttrs() == 0 {
		// If the record has no Attrs, remove groups at the end of the list; they are empty.
		for len(goas) > 0 && goas[len(goas)-1].group != "" {
			goas = goas[:len(goas)-1]
		}
	}
	for _, goa := range goas {
		if goa.group != "" {
			buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", goa.group)
			indentLevel++
		} else {
			for _, a := range goa.attrs {
				buf = h.appendAttr(buf, a, indentLevel)
			}
		}
	}
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a, indentLevel)
		return true
	})
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}
```

You may have noticed that our algorithm for
recording `WithGroup` and `WithAttrs` information is quadratic in the
number of calls to those methods, because of the repeated copying.
That is unlikely to matter in practice, but if it bothers you,
you can use a linked list instead,
which `Handle` will have to reverse or visit recursively.
See the
[github.com/jba/slog/withsupport](https://github.com/jba/slog/tree/main/withsupport)
package for an implementation.

#### Getting the mutex right

Let us revisit the last few lines of `Handle`:

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
    return err

This code hasn't changed, but we can now appreciate why `h.mu` is a
pointer to a `sync.Mutex`. Both `WithGroup` and `WithAttrs` copy the handler.
Both copies point to the same mutex.
If the copy and the original used different mutexes and were used concurrently,
then their output could be interleaved, or some output could be lost.
Code like this:

    l2 := l1.With("a", 1)
    go l1.Info("hello")
    l2.Info("goodbye")

could produce output like this:

    hegoollo a=dbye1

See [this bug report](https://go.dev/issue/61321) for more detail.


### With pre-formatting

Our second version of the `WithGroup` and `WithAttrs` methods provides pre-formatting.
This implementation is more complicated than the previous one.
Is the extra complexity worth it?
That depends on your circumstances, but here is one circumstance where
it might be.
Say that you wanted your server to log a lot of information about an incoming
request with every log message that happens during that request. A typical
handler might look something like this:

    func (s *Server) handleWidgets(w http.ResponseWriter, r *http.Request) {
        logger := s.logger.With(
            "url", r.URL,
            "traceID": r.Header.Get("X-Cloud-Trace-Context"),
            // many other attributes
            )
        // ...
    }

A single `handleWidgets` request might generate hundreds of log lines.
For instance, it might contain code like this:

    for _, w := range widgets {
        logger.Info("processing widget", "name", w.Name)
        // ...
    }

For every such line, the `Handle` method we wrote above will format all
the attributes that were added using `With` above, in addition to the
ones on the log line itself.

Maybe all that extra work doesn't slow down your server significantly, because
it does so much other work that time spent logging is just noise.
But perhaps your server is fast enough that all that extra formatting appears high up
in your CPU profiles. That is when pre-formatting can make a big difference,
by formatting the attributes in a call to `With` just once.

To pre-format the arguments to `WithAttrs`, we need to keep track of some
additional state in the `IndentHandler` struct.

```go
type IndentHandler struct {
	opts           Options
	preformatted   []byte   // data from WithGroup and WithAttrs
	unopenedGroups []string // groups from WithGroup that haven't been opened
	indentLevel    int      // same as number of opened groups so far
	mu             *sync.Mutex
	out            io.Writer
}
```

Mainly, we need a buffer to hold the pre-formatted data.
But we also need to keep track of which groups
we've seen but haven't output yet. We'll call those groups "unopened."
We also need to track how many groups we've opened, which we can do
with a simple counter, since an opened group's only effect is to change the
indentation level.

This `WithGroup` is a lot like the previous one: it just remembers the
new group, which is unopened initially.

```go
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := *h
	// Add an unopened group to h2 without modifying h.
	h2.unopenedGroups = make([]string, len(h.unopenedGroups)+1)
	copy(h2.unopenedGroups, h.unopenedGroups)
	h2.unopenedGroups[len(h2.unopenedGroups)-1] = name
	return &h2
}
```

`WithAttrs` does all the pre-formatting:

```go
func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := *h
	// Force an append to copy the underlying array.
	pre := slices.Clip(h.preformatted)
	// Add all groups from WithGroup that haven't already been added.
	h2.preformatted = h2.appendUnopenedGroups(pre, h2.indentLevel)
	// Each of those groups increased the indent level by 1.
	h2.indentLevel += len(h2.unopenedGroups)
	// Now all groups have been opened.
	h2.unopenedGroups = nil
	// Pre-format the attributes.
	for _, a := range attrs {
		h2.preformatted = h2.appendAttr(h2.preformatted, a, h2.indentLevel)
	}
	return &h2
}

func (h *IndentHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
	for _, g := range h.unopenedGroups {
		buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
		indentLevel++
	}
	return buf
}
```

It first opens any unopened groups. This handles calls like:

    logger.WithGroup("g").WithGroup("h").With("a", 1)

Here, `WithAttrs` must output "g" and "h" before "a". Since a group established
by `WithGroup` is in effect for the rest of the log line, `WithAttrs` increments
the indentation level for each group it opens.

Lastly, `WithAttrs` formats its argument attributes, using the same `appendAttr`
method we saw above.

It's the `Handle` method's job to insert the pre-formatted material in the right
place, which is after the built-in attributes and before the ones in the record:

```go
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	// Insert preformatted attributes just after built-in ones.
	buf = append(buf, h.preformatted...)
	if r.NumAttrs() > 0 {
		buf = h.appendUnopenedGroups(buf, h.indentLevel)
		r.Attrs(func(a slog.Attr) bool {
			buf = h.appendAttr(buf, a, h.indentLevel+len(h.unopenedGroups))
			return true
		})
	}
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}
```

It must also open any groups that haven't yet been opened. The logic covers
log lines like this one:

    logger.WithGroup("g").Info("msg", "a", 1)

where "g" is unopened before `Handle` is called and must be written to produce
the correct output:

    level: INFO
    msg: "msg"
    g:
        a: 1

The check for `r.NumAttrs() > 0` handles this case:

    logger.WithGroup("g").Info("msg")

Here there are no record attributes, so no group to open.

## Testing

The [`Handler` contract](https://pkg.go.dev/log/slog#Handler) specifies several
constraints on handlers.
To verify that your handler follows these rules and generally produces proper
output, use the [testing/slogtest package](https://pkg.go.dev/testing/slogtest).

That package's `TestHandler` function takes an instance of your handler and
a function that returns its output formatted as a slice of maps. Here is the test function
for our example handler:

```go
func TestSlogtest(t *testing.T) {
	var buf bytes.Buffer
	err := slogtest.TestHandler(New(&buf, nil), func() []map[string]any {
		return parseLogEntries(t, buf.Bytes())
	})
	if err != nil {
		t.Error(err)
	}
}
```

Calling `TestHandler` is easy. The hard part is parsing your handler's output.
`TestHandler` calls your handler multiple times, resulting in a sequence of log
entries.
It is your job to parse each entry into a `map[string]any`.
A group in an entry should appear as a nested map.

If your handler outputs a standard format, you can use an existing parser.
For example, if your handler outputs one JSON object per line, then you
can split the output into lines and call `encoding/json.Unmarshal` on each.
Parsers for other formats that can unmarshal into a map can be used out
of the box.
Our example output is enough like YAML so that we can use the `gopkg.in/yaml.v3`
package to parse it:

```go
func parseLogEntries(t *testing.T, data []byte) []map[string]any {
	entries := bytes.Split(data, []byte("---\n"))
	entries = entries[:len(entries)-1] // last one is empty
	var ms []map[string]any
	for _, e := range entries {
		var m map[string]any
		if err := yaml.Unmarshal([]byte(e), &m); err != nil {
			t.Fatal(err)
		}
		ms = append(ms, m)
	}
	return ms
}
```

If you have to write your own parser, it can be far from perfect.
The `slogtest` package uses only a handful of simple attributes.
(It is testing handler conformance, not parsing.)
Your parser can ignore edge cases like whitespace and newlines in keys and
values. Before switching to a YAML parser, we wrote an adequate custom parser
in 65 lines.

# General considerations

## Copying records

Most handlers won't need to copy the `slog.Record` that is passed
to the `Handle` method.
Those that do must take special care in some cases.

A handler can make a single copy of a `Record` with an ordinary Go
assignment, channel send or function call if it doesn't retain the
original.
But if its actions result in more than one copy, it should call `Record.Clone`
to make the copies so that they don't share state.
This `Handle` method passes the record to a single handler, so it doesn't require `Clone`:

    type Handler1 struct {
        h slog.Handler
        // ...
    }

    func (h *Handler1) Handle(ctx context.Context, r slog.Record) error {
        return h.h.Handle(ctx, r)
    }

This `Handle` method might pass the record to more than one handler, so it
should use `Clone`:

    type Handler2 struct {
        hs []slog.Handler
        // ...
    }

    func (h *Handler2) Handle(ctx context.Context, r slog.Record) error {
        for _, hh := range h.hs {
            if err := hh.Handle(ctx, r.Clone()); err != nil {
                return err
            }
        }
        return nil
    }

## Concurrency safety

A handler must work properly when a single `Logger` is shared among several
goroutines.
That means that mutable state must be protected with a lock or some other mechanism.
In practice, this is not hard to achieve, because many handlers won't have any
mutable state.

- The `Enabled` method typically consults only its arguments and a configured
  level. The level is often either set once initially, or is held in a
  `LevelVar`, which is already concurrency-safe.

- The `WithAttrs` and `WithGroup` methods should not modify the receiver,
  for reasons discussed above.

- The `Handle` method typically works only with its arguments and stored fields.

Calls to output methods like `io.Writer.Write` should be synchronized unless
it can be verified that no locking is needed.
As we saw in our example, storing a pointer to a mutex enables a logger and
all of its clones to synchronize with each other.
Beware of facile claims like "Unix writes are atomic"; the situation is a lot more nuanced than that.

Some handlers have legitimate reasons for keeping state.
For example, a handler might support a `SetLevel` method to change its configured level
dynamically.
Or it might output the time between successive calls to `Handle`,
which requires a mutable field holding the last output time.
Synchronize all accesses to such fields, both reads and writes.

The built-in handlers have no directly mutable state.
They use a mutex only to sequence calls to their contained `io.Writer`.

## Robustness

Logging is often the debugging technique of last resort. When it is difficult or
impossible to inspect a system, as is typically the case with a production
server, logs provide the most detailed way to understand its behavior.
Therefore, your handler should be robust to bad input.

For example, the usual advice when a function discovers a problem,
like an invalid argument, is to panic or return an error.
The built-in handlers do not follow that advice.
Few things are more frustrating than being unable to debug a problem that
causes logging to fail;
it is better to produce some output, however imperfect, than to produce none at all.
That is why methods like `Logger.Info` convert programming bugs in their list of
key-value pairs, like missing values or malformed keys,
into `Attr`s that contain as much information as possible.

One place to avoid panics is in processing attribute values. A handler that wants
to format a value will probably switch on the value's kind:

    switch attr.Value.Kind() {
    case KindString: ...
    case KindTime: ...
    // all other Kinds
    default: ...
    }

What should happen in the default case, when the handler encounters a `Kind`
that it doesn't know about?
The built-in handlers try to muddle through by using the result of the value's
`String` method, as our example handler does.
They do not panic or return an error.
Your own handlers might in addition want to report the problem through your production monitoring
or error-tracking telemetry system.
The most likely explanation for the issue is that a newer version of the `slog` package added
a new `Kind`&mdash;a backwards-compatible change under the Go 1 Compatibility
Promise&mdash;and the handler wasn't updated.
That is certainly a problem, but it shouldn't deprive
readers from seeing the rest of the log output.

There is one circumstance where returning an error from `Handler.Handle` is appropriate.
If the output operation itself fails, the best course of action is to report
this failure by returning the error. For instance, the last two lines of the
built-in `Handle` methods are

    _, err := h.w.Write(*state.buf)
    return err

Although the output methods of `Logger` ignore the error, one could write a
handler that does something with it, perhaps falling back to writing to standard
error.

## Speed

Most programs don't need fast logging.
Before making your handler fast, gather data&mdash;preferably production data,
not benchmark comparisons&mdash;that demonstrates that it needs to be fast.
Avoid premature optimization.

If you need a fast handler, start with pre-formatting. It may provide dramatic
speed-ups in cases where a single call to `Logger.With` is followed by many
calls to the resulting logger.

If log output is the bottleneck, consider making your handler asynchronous.
Do the minimal amount of processing in the handler, then send the record and
other information over a channel. Another goroutine can collect the incoming log
entries and write them in bulk and in the background.
You might want to preserve the option to log synchronously
so you can see all the log output to debug a crash.

Allocation is often a major cause of a slow system.
The `slog` package already works hard at minimizing allocations.
If your handler does its own allocation, and profiling shows it to be
a problem, then see if you can minimize it.

One simple change you can make is to replace calls to `fmt.Sprintf` or `fmt.Appendf`
with direct appends to the buffer. For example, our IndentHandler appends string
attributes to the buffer like so:

	buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())

As of Go 1.21, that results in two allocations, one for each argument passed to
an `any` parameter. We can get that down to zero by using `append` directly:

	buf = append(buf, a.Key...)
	buf = append(buf, ": "...)
	buf = strconv.AppendQuote(buf, a.Value.String())
	buf = append(buf, '\n')

Another worthwhile change is to use a `sync.Pool` to manage the one chunk of
memory that most handlers need:
the `[]byte` buffer holding the formatted output.

Our example `Handle` method began with this line:

	buf := make([]byte, 0, 1024)

As we said above, providing a large initial capacity avoids repeated copying and
re-allocation of the slice as it grows, reducing the number of allocations to
one.
But we can get it down to zero in the steady state by keeping a global pool of buffers.
Initially, the pool will be empty and new buffers will be allocated.
But eventually, assuming the number of concurrent log calls reaches a steady
maximum, there will be enough buffers in the pool to share among all the
ongoing `Handler` calls. As long as no log entry grows past a buffer's capacity,
there will be no allocations from the garbage collector's point of view.

We will hide our pool behind a pair of functions, `allocBuf` and `freeBuf`.
The single line to get a buffer at the top of `Handle` becomes two lines:

	bufp := allocBuf()
	defer freeBuf(bufp)

One of the subtleties involved in making a `sync.Pool` of slices
is suggested by the variable name `bufp`: your pool must deal in
_pointers_ to slices, not the slices themselves.
Pooled values must always be pointers. If they aren't, then the `any` arguments
and return values of the `sync.Pool` methods will themselves cause allocations,
defeating the purpose of pooling.

There are two ways to proceed with our slice pointer: we can replace `buf`
with `*bufp` throughout our function, or we can dereference it and remember to
re-assign it before freeing:

	bufp := allocBuf()
	buf := *bufp
	defer func() {
		*bufp = buf
		freeBuf(bufp)
	}()


Here is our pool and its associated functions:

```go
var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 1024)
		return &b
	},
}

func allocBuf() *[]byte {
	return bufPool.Get().(*[]byte)
}

func freeBuf(b *[]byte) {
	// To reduce peak allocation, return only smaller buffers to the pool.
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}
```

The pool's `New` function does the same thing as the original code:
create a byte slice with 0 length and plenty of capacity.
The `allocBuf` function just type-asserts the result of the pool's
`Get` method.

The `freeBuf` method truncates the buffer before putting it back
in the pool, so that `allocBuf` always returns zero-length slices.
It also implements an important optimization: it doesn't return
large buffers to the pool.
To see why this important, consider what would happen if there were a single,
unusually large log entry&mdash;say one that was a megabyte when formatted.
If that megabyte-sized buffer were put in the pool, it could remain
there indefinitely, constantly being reused, but with most of its capacity
wasted.
The extra memory might never be used again by the handler, and since it was in
the handler's pool, it might never be garbage-collected for reuse elsewhere.
We can avoid that situation by excluding large buffers from the pool.
```

### go.mod

```
module golang.org/x/example

go 1.23.0

require (
	golang.org/x/tools v0.33.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
)
```

### go.sum

```
github.com/google/go-cmp v0.6.0 h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=
github.com/google/go-cmp v0.6.0/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
golang.org/x/mod v0.25.0 h1:n7a+ZbQKQA/Ysbyb0/6IbB1H/X41mKgbhfv7AfG/44w=
golang.org/x/mod v0.25.0/go.mod h1:IXM97Txy2VM4PJ3gI61r1YEk/gAj6zAHN3AdZt6S9Ww=
golang.org/x/sync v0.15.0 h1:KWH3jNZsfyT6xfAfKiz6MRNmd46ByHDYaZ7KSkCtdW8=
golang.org/x/sync v0.15.0/go.mod h1:1dzgHSNfp02xaA81J2MS99Qcpr2w7fw1gpm99rleRqA=
golang.org/x/tools v0.33.0 h1:4qz2S3zmRxbGIhDIAgjxvFutSvH5EfnsYrRBj0UI0bc=
golang.org/x/tools v0.33.0/go.mod h1:CIJMaWEY88juyUfo7UbgPqbC8rU2OqfAV1h2Qp0oMYI=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
```

### LICENSE

```
Copyright 2009 The Go Authors.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google LLC nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

### PATENTS

```
Additional IP Rights Grant (Patents)

"This implementation" means the copyrightable works distributed by
Google as part of the Go project.

Google hereby grants to You a perpetual, worldwide, non-exclusive,
no-charge, royalty-free, irrevocable (except as stated in this section)
patent license to make, have made, use, offer to sell, sell, import,
transfer and otherwise run, modify and propagate the contents of this
implementation of Go, where such license applies only to those patent
claims, both currently owned or controlled by Google and acquired in
the future, licensable by Google that are necessarily infringed by this
implementation of Go.  This grant does not include claims that would be
infringed only as a consequence of further modification of this
implementation.  If you or your agent or exclusive licensee institute or
order or agree to the institution of patent litigation against any
entity (including a cross-claim or counterclaim in a lawsuit) alleging
that this implementation of Go or any code incorporated within this
implementation of Go constitutes direct or contributory patent
infringement, or inducement of patent infringement, then any patent
rights granted to you under this License for this implementation of Go
shall terminate as of the date such litigation is filed.
```

## Warnings

- appengine-hello/static/favicon.ico: binary file detected

---
Generated by **repo2page**
