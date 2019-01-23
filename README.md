# MdConv: a golang based markdown convertor

For the purpose of parsing and use a markdown document easier, this project provided a simple way, converting the markdown string to a AST, so that it could be used more flexiable.

Currently this project is in a very very early stage. Please ***DON'T*** use it yet.

Any bug report, suggestion, pull request is highly appreciated.


## Markdown Spec

This project is based on [GitHub Flavored Markdown Spec](https://github.github.com/gfm/) (aka gfm).

However, for the purpose of use the code widely, it could also support:

+ latex


## LaTeX

To make it easier, this project ONLY support the follow

```md
$inline-latex$
$$multiple-line-latex$$
```


## References

To prepare a project, a huge a mount of resources are 

+ [gfm](https://github.github.com/gfm/) -- GitHub Flavored Markdown Spec
+ [CommonMark](https://commonmark.org/) -- A strongly defined, highly compatible specification of Markdown
+ [blackfriday](https://github.com/russross/blackfriday) -- a markdown processor for Go
+ [sundown](https://github.com/vmg/sundown) -- Standards compliant, fast, secure markdown processing library in C
+ [cmark](https://github.com/commonmark/cmark) -- CommonMark parsing and rendering library and program in C
