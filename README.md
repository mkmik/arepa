# arepa

> a repetition for acme

Small helper for the [acme editor](http://acme.cat-v.org/) that re-runs a command
every time another (trigger) command exits.

Works well in combination with [waigo](https://github.com/mkmik/waigo) and other simple tools like it.

# Install

```
$ go get -u github.com/mkmik/arepa
```

# Usage

For example you can combine it with [waigo](https://github.com/mkmik/waigo) so it can show test and build errors for the package under the current directory:

```
arepa -t waigo go test
```

(which I wrap in a smaller shell script called `ago` which does `exec arepa -t waigo go "$@"`)
