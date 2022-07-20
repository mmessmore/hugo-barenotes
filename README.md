# Messynotes: a Hugo theme and stuff

This is a simple theme for personal note taking and TODO management.

I love Markdown based note management + revision control.  But all apps I had
used frustrated me.  I made this theme in combination with [a
CLI](https://github.com/mmessmore/messynotes) to make a system using Hugo to do
a lot of the heavy lifting.

It is easiest to consume via the [messynotes
command](https://github.com/mmessmore/messynotes), but this could be consumed
separately.

## Features

### Index

The index tracks recent notes on the left with the TODO always at the top.

One the right, there is a bare list of tags, then a list of categories with the
stream of titles for each.

### Categorization

Tags serve as a lookup for related notes.
Categories function more as a per-topic stream.

The idea is that you have multiple tags per note, and one category, but you can
do what you want.

### "Private" Junk

Sometimes you share your screen and you don't want stuff showing up by default.
So you can tag a note as "private" and/or add a list of private categories in
your config.

You can toggle them visible at the bottom of the page.

### Lightweight

This uses as little JavaScript as possible.  0 CSS/JavaScript frameworks.  Only
external dependency is fonts, which should degrade gracefully offline.  There
are no browser-specific features or workarounds, just standard stuff.

### Light/Dark

The color scheme will change based on the system/browser settings for light and
dark mode.  This is done with a combination of CSS media queries and variables.

## Contribution

My usage of CSS is probably awkward and falls into the "it works for me"
category.  Contributions to clean it up or generally make things prettier are
appreciated.  I'm not a "frontend" guy, and I'm learning this new fangled CSS
stuff that seems cool.

I'll try to stick things I want to fix or would appreciate help with in
[TODO.md](./TODO.md).

PRs are appreciated.  I'll try to deal with bugs in Issues.  Feature requests
without a PR will probably take some time and only be done if they seem useful
to me personally.

I have a pretty demanding "real job" as well as a family I enjoy being around.

## License

This is licensed under the [MIT License](./LICENSE).  You are welcome to do
anything within the terms of it: fork, make a product, whatever.  I will not be
offended.  I'm not sure why people do get offended when someone does something
within their license, but I am not one of them.
