# repin

`repin` is a tool to replace strings between keyword pair.

## tl;dr

`repin` is a tool that makes it easy to write operations that can be written in GNU sed as follows.

~~~
$ cat README.md
# Hello

```console
```

$ cat README.md | sed -z 's/```console.*```/```console\n$ echo hello world!\n```/'
# Hello

```console
$ echo hello world!
```

$
~~~

~~~
$ repin README.md -k '```console' -k '```' -r '$ echo hello world!'
# Hello

```console
$ echo hello world!
```

$
~~~

## Install

**deb:**

Use [dpkg-i-from-url](https://github.com/k1LoW/dpkg-i-from-url)

``` console
$ export REPIN_VERSION=X.X.X
$ curl -L https://git.io/dpkg-i-from-url | bash -s -- https://github.com/k1LoW/repin/releases/download/v$REPIN_VERSION/repin_$REPIN_VERSION-1_amd64.deb
```

**RPM:**

``` console
$ export REPIN_VERSION=X.X.X
$ yum install https://github.com/k1LoW/repin/releases/download/v$REPIN_VERSION/repin_$REPIN_VERSION-1_amd64.rpm
```

**apk:**

Use [apk-add-from-url](https://github.com/k1LoW/apk-add-from-url)

``` console
$ export REPIN_VERSION=X.X.X
$ curl -L https://git.io/apk-add-from-url | sh -s -- https://github.com/k1LoW/repin/releases/download/v$REPIN_VERSION/repin_$REPIN_VERSION-1_amd64.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/repin
```

**go install:**

```console
$ go install github.com/k1LoW/repin/cmd/repin@vX.X.X
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/repin/releases)

**docker:**

```console
$ docker pull ghcr.io/k1low/repin:latest
```
