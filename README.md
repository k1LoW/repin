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
