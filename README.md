goAnagrams
==========

This tool is inspired by this [tweet](https://twitter.com/fermatslibrary/status/875340896379817984).
It checks, if two given arguments are anagrams. As it is written in [golang](https://golang.org/)
it interprets each argument as a [Rune](https://golang.org/pkg/go/types/#Rune) slice.

Building
--------

        $ git clone git@github.com:florianl/goAnagrams.git
          Cloning into 'goAnagrams'...
          [...]
        $ cd goAnagrams/
        $ export GOPATH=$HOME/go
        $ go build *.go
          [...]
        $ ./goAnagrams
          [...]

Or you can get it directly via [golang](https://golang.org/):

        $ export GOPATH=$HOME/go
        $ go get github.com/florianl/goAnagrams
          [...]
        $ $GOPATH/bin/goAnagrams
          [...]

Examples
--------
        $ ./goAnagrams Hallo hallo
          "h" is missing in mapping
        $ ./goAnagrams "Hello, 世界" "世界, Hello"
          Hello, 世界 and 世界, Hello are probably anagrams
