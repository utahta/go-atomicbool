# Atomic Bool

[![Build Status](https://travis-ci.org/utahta/go-atomicbool.svg?branch=master)](https://travis-ci.org/utahta/go-atomicbool)

## Installing

```
$ go get -u github.com/utahta/go-atomicbool
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/utahta/go-atomicbool"
)

func main() {
	b := atomicbool.New(false)

	if b.Disabled() {
		fmt.Println("disabled")
	}

	b.Set(true)

	if b.Enabled() {
		fmt.Println("enabled")
	}
}
```

## Contributing

1. Fork it ( https://github.com/utahta/go-atomicbool/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

