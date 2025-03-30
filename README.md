# sponge

> Read in all available data from STDIN, then output it to a file named on the command line, overwriting the file if it exists.

## Usage

```sh
cat README.md | grep -v 'TODO' | sponge README.md
```

## Installation

```sh
go install github.com/mlbright/sponge@latest
```

## Why?

I only really use `sponge` from the [moreutils package][moreutils], and the package still conflicts with `parallel` on MacOS.
So ~~I wrote~~ chatGPT wrote this to implement the `sponge` functionality because I like Golang.
Sue me.

## Limitations

- [ ] Only uses main memory, not temporary files, to accumulate STDIN. Therefore it cannot be used on gargantuan accumulations of files.
- [ ] Does not implement the `-a` flag from the original `sponge` command, which appends to the file specified instead of overwriting it.
- [x] ~~Does not preserve the file permissions of the original file specified.~~

[moreutils]: https://joeyh.name/code/moreutils/
