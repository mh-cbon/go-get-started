# go-get-started
get started with go

# Install go >= 1.6.2

I recommend to use `gvm`

```sh
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.6.2
gvm use go1.6.2 --default
which go
  $HOME/.gvm/gos/go1.6.2/bin/go
go version
  go version go1.6.2 linux/amd64
```

See https://github.com/moovweb/gvm

# Configure the go workspace

#### Create a folder to host the workspace

```sh
mkdir $HOME/gow
```

see https://golang.org/doc/code.html#Workspaces

#### Setup the required environment variables

```sh
cat <<EOT >> ~/.bashrc
# Go variables
export GOPATH=$HOME/gow
export PATH=$PATH:$GOPATH/bin
# go 1.6.x and below only
export GOBIN=$GOPATH/bin
EOT
source ~/.bashrc
```

see https://golang.org/doc/code.html#GOPATH

#### Setup directories

```sh
mkdir -p $GOPATH/pkg
mkdir -p $GOPATH/src
mkdir -p $GOPATH/bin
```

# Get a package manager

I recommend `glide` at that time of writing

```sh
go get github.com/Masterminds/glide
cd $GOPATH/src/github.com/Masterminds/glide
make build
go install -ldflags "-X main.version=0.10.2-86-g5865b8e" glide.go
```

FYI `which glide -> $GOPATH/bin/glide`

see https://github.com/Masterminds/glide

# Create your first package

Let s create a first package `a`, hosted on `github.com`, as user `mh-cbon`.

At any time you can refer to [this folder](https://github.com/mh-cbon/go-get-started/blob/master/gow/) to see the expected result.

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon/a
cd $GOPATH/src/github.com/mh-cbon/a
glide create

mkdir sub
touch sub/lib.go
# File: $GOPATH/src/github.com/mh-cbon/a/sub/lib.go
cat <<EOT > sub/lib.go
package lib

import "fmt"

func Hello () {
  fmt.Println("Hello world! It's a/lib!")
}
EOT

touch main.go
# File: $GOPATH/src/github.com/mh-cbon/a/main.go
cat <<EOT > main.go
package main

import "github.com/mh-cbon/a/sub"

func main () {
  lib.Hello()
}
EOT

go run main.go
  Hello world! It's a/lib!
```

This step has created a package named `a` with a binary available at `main.go`.
`a` package has a lib `sub` loaded with `import "github.com/mh-cbon/a/sub"`.

Notice the source are created under `$GOPATH/src/` as `github.com/mh-cbon/a`

# Add a second package

Let s create a second package `b` to illustrate the setup and import of a library

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon/b
cd $GOPATH/src/github.com/mh-cbon/b
glide create
touch index.go

# File: $GOPATH/src/github.com/mh-cbon/b/index.go
cat <<EOT > index.go
package b

import "fmt"

func Hello () {
  fmt.Println("Hello world! It's b!")
}
EOT
```

This step created a `b` package. It is a library only. It declares a file `index.go` and is loadable via `import "github.com/mh-cbon/b"`.

Notice the source are created under `$GOPATH/src/` as `github.com/mh-cbon/b`

It is located at `$GOPATH/src`, which makes it loadable from your other go programs/libraries as `github.com/mh-cbon/b`.

# Import b package from a package

```sh
cd $GOPATH/src/github.com/mh-cbon/a

# File: $GOPATH/src/github.com/mh-cbon/a/main.go
cat <<EOT > main.go
package main

import (
  "github.com/mh-cbon/a/sub"
  "github.com/mh-cbon/b"
)

func main () {
  lib.Hello()
  b.Hello()
}
EOT

go run main.go
  Hello world! It's a/lib!
  Hello world! It's b!
```

In this step program `a` depends on the local library `github.com/mh-cbon/b` declared and available at `$GOPATH/src`

# Depend on a remote package

Let s finish this readme by importing ans consuming a remote dependency.

`a` will depend and consume `github.com/Masterminds/semver`

```sh
cd $GOPATH/src/github.com/mh-cbon/a

glide get github.com/Masterminds/semver

# File: $GOPATH/src/github.com/mh-cbon/a/main.go
cat <<EOT > main.go
package main

import (
  "fmt"

  "github.com/mh-cbon/a/sub"
  "github.com/mh-cbon/b"
  "github.com/Masterminds/semver"
)

func main () {
  lib.Hello()
  b.Hello()
  c, _ := semver.NewConstraint("<= 1.2.3, >= 1.4")
  fmt.Println(c)
}
EOT

go run main.go
  Hello world! It's a/lib!
  Hello world! It's b!
  &{[[0xc8200129c0 0xc820012a00]]}
```

In this step program `a` depends on a remote library installed via `glide`.
The library `semver` loaded via `import "github.com/Masterminds/semver"` is installed and available at `./vendor/`.
It is a local dependency.

Notice the new `glide.yaml`

```sh
cat $GOPATH/src/github.com/mh-cbon/a/glide.yaml
package: github.com/mh-cbon/a
import:
- package: github.com/Masterminds/semver
```

Notice the new `vendor` directory

```sh
ls -al .
 sub/
 glide.yaml
 main.go
 vendor/
```

# Other tools to use

#### go vet

`
Vet examines Go source code and reports suspicious constructs
`

```sh
$ go vet
can't load package: /home/mh-cbon/gow/src/github.com/mh-cbon/a/main.go:7:2: \
local import "./a" in non-local package
```

see [this](https://golang.org/cmd/vet/)


#### go fmt

`
Gofmt formats Go programs.
`

```sh
$ go fmt
main.go
```

see [this](https://golang.org/cmd/gofmt/)

#### go doc

`
Show the documentation for the package in the current directory.
`

Standard practices to document your code are available [here](https://blog.golang.org/godoc-documenting-go-code)

Go team also provide automatic package documentation generation at [godoc.org](https://godoc.org/)

#### There is more of it

Go team provides many tools that you can check [here](https://golang.org/cmd/)

# Other notes

Notice the environment variables
```sh
env | grep GO
GOROOT=$HOME/.gvm/gos/go1.6
GOPATH=$HOME/gow
GOBIN=$GOPATH/bin
```

Which means,
- when you run `go install....` it installs binary into `$GOPATH/bin`
- new packages are created under `$HOME/gow/src`
- do not use relative import paths (eg `import "./mypackage/"`) because they are not compatible with `go install`.
See [this](https://golang.org/cmd/go/#hdr-Relative_import_paths) and [that](http://stackoverflow.com/questions/30885098/go-local-import-in-non-local-package)

If you really want to understand go implementation of import, please check [Context.Import](https://github.com/golang/go/blob/master/src/go/build/build.go#L493) method.

## Continue the readings

- https://github.com/golang/example
- https://golanglibs.com

That s it !

~~ Happy coding
