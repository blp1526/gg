# gg

A CLI tool to search words by the default web browser

## Installation

```
$ make install
```

## Usage

### Format

```
$ gg [word word word...]
```

### Example

```
$ gg foo bar baz
```

This example runs below command.

#### Linux

```
$ xdg-open https://www.google.co.jp/search?q=foo+bar+baz
```

#### macOS

```
$ open https://www.google.co.jp/search?q=foo+bar+baz
```
