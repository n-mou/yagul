# YAGUL

<a href="https://pkg.go.dev/github.com/n-mou/yagul"><img src="https://pkg.go.dev/badge/github.com/n-mou/yagul.svg" alt="Go Reference"></a>

YAGUL (Yet Another Go-Utils Library) is my personal recopilation of helper and quality of life improvement modules that I regularly use in my go projects and I made it public because other devs may also find them useful. 

# Why?

I'm relatively new to Go development and I've seen **many devs create a Go-utils repo** containing helper libraries they find useful. Some repos are very extensive and others are a small compilation of simple and selective helpers. However, **none of these Go-utils package is popular enough to became a de-facto standard** (like JQuery was in early Javascript). 

I could make my own compilation but in most of these libraries there are some common functions that are rewritten over and over (like the functional `map()`, `filter()` and `reduce()`). So **instead of reinventing the wheel, I'll just post modules that I missed in those Go-utils repos** and that I think other Go devs might benefit from them.

Thus, this repo is my humble contribution of helper functions and dev patterns that I find useful and not very popular in the Go community instead of a full fledged Go-utils library that aims to be a JQuery for Go. If you like it and use it in your projects, please give a star.

# Current subpackages

- `g`: types and functions that could live as global variables instead of being in a module (the g is for "Global"). Currently it only has functions that serve as syntax sugar for calling functions and panicking when an  error is found in a single line.
- `fs`: useful functions for file and directories management that I personally miss in Go standard library. 
- `itertools`: some helper functions to work with iterators and an adapter function that takes a pull based iterator (defined with a next() and a stop() function) and returns a regular `iter.Seq` or `iter.Seq2` regular iterator.
- `list`: reimplementation of Go's standard library `container/list` using generics (the original used values of type any and every time a value is retrieved it must be type casted, reimplementing it with a generic type instead is more convenient).

# Roadmap (lack of)

This repo is a compilation of utilities I needed as a Go developer and wasn't able to find in third party projects. There are no projections on when or how will it be expanded or updated. However, **I'll try not to change the type signature of the existing functions** to avoid breaking somebody else's code that might rely on some submodule.

All commited code is documented and tested, and if eventually someone ends using this repo to the point that is willing to contribute, **PRs are appreciated and welcome**.

# License

In case of someone wanting to use some of this modules, this repo is released under the MIT license. So you can copy, alter and reditribute the code at your will. But please, give some credit and leave a link to this software repo somewhere in your own license.