# arepa

> a repetition for acme

Small helper for the [acme editor](http://acme.cat-v.org/) that re-runs a command
every time another (trigger) command exits.

Works well in combination with [waigo](https://github.com/mkmik/waigo) and other simple tools like it.

# Install

This project uses the experimental Go modules available in go 1.11.x.
While you can use go get to just install it with some random dependencies,
ideally you should use something like this to install it:

```
(T=$(mktemp); rm -rf "$T"; git clone https://github.com/mkmik/arepa "$T" && cd "$T"; go install; rm -rf "$T")
```

(Please let me know if there is a better oneliner to do it)

# Usage

For example you can combine it with [waigo](https://github.com/mkmik/waigo) so it can show 

```
arepa -t waigo go test
```

