# go-get-started
get started with go

# Install go >= 1.6.2

I recommend to use `gvm`

```sh
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.6.2
gvm use go1.6.2 --default
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
cat <<EOT >> ~/.bash_profile
# Go variables
export GOPATH=$HOME/gow
export PATH=PATH:$GOPATH/bin
EOT
source ~/.bash_profile
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
mdkir -p $GOPATH/src/github.com/Masterminds/glide
cd $GOPATH/src/github.com/Masterminds/glide
git clone https://github.com/Masterminds/glide.git .
make build
go install -ldflags "-X main.version=0.10.2-86-g5865b8e" glide.go
```

FYI `which glide -> $GOPATH/bin/glide`

see https://github.com/Masterminds/glide

# Create your first package

Let s create a first package `a`, hosted on `github.com`, as user `mh-cbon`

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon/a
cd $GOPATH/src/github.com/mh-cbon/a
glide create

mkdir a
touch a/a.go
cat <<EOT > a/a.go
package a

import "fmt"

func Hello () {
  fmt.Println("Hello world! It's a!")
}
EOT

touch main.go
cat <<EOT > main.go
package main

import "./a"

func main () {
  a.Hello()
}
EOT

go run main.go
  Hello world! It's a!
```

This step has created a package named `a` with a binary available at `main.go`.
`a` package has a lib `a` loaded with `import "./a"`.

Notice the source are created under `$GOPATH/src/` as `github.com/mh-cbon/a`

# Add a second package

Let s create a second package `b` to illustrate the setup and import of a library

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon/b
cd $GOPATH/src/github.com/mh-cbon/b
glide create
touch index.go

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

It is located at `$GOPATH/src`, which makes it loadable from your other go programs/libraries as `github.com/mh-cbon/b``.

# Import b package into a package

```sh
cd $GOPATH/github.com/mh-cbon/a

cat <<EOT > main.go
package main

import (
  "./a"
  "github.com/mh-cbon/b"
)

func main () {
  a.Hello()
  b.Hello()
}
EOT

go run main.go
  Hello world! It's a!
  Hello world! It's b!
```

In this step program `a` depends on the local library `github.com/mh-cbon/b` declared and available at `$GOPATH/src`

# Depend on a remote package

Let s finish this readme by importing ans consuming a remote dependency. 

`a` will depend and consume `github.com/Masterminds/semver`

```sh
cd $GOPATH/github.com/mh-cbon/a

glide get github.com/Masterminds/semver

cat <<EOT > main.go
package main

import (
  "./a"
  "github.com/mh-cbon/b"
  "github.com/Masterminds/semver"
)

func main () {
  a.Hello()
  b.Hello()
  c, _ := semver.NewConstraint("<= 1.2.3, >= 1.4")
  fmt.Println(c)
}
EOT

go run main.go
  Hello world! It's a!
  Hello world! It's b!
  &{[[0xc8200129c0 0xc820012a00]]}
```

In this step program `a` depends on a remote library installed via `glide`.
The library `semver` loaded via `import "github.com/Masterminds/semver"` is insalled and available at `./vendor/`.
It is a local dependency.

Notice the new `glide.yaml`

```sh
cat glide.yaml 
package: github.com/mh-cbon/a
import:
- package: github.com/mh-cbon/b
```

Notice the new `vendor` directory

```sh
ls -al .
 a/
 glide.yaml
 main.go
 vendor/
```

# Other notes

Notice the environment variables
```sh
env | grep GO
GOROOT=$HOME/.gvm/gos/go1.6
GOPATH=$HOME/gow
```

Which means,
- when you run `go install....` it installs binary into `$HOME/.gvm/gos/go1.6/bin`
- new pakages are created under `$HOME/gow/src`


That s it ! 

~~ Happy coding
