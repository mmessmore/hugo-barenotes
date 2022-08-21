# Messynotes: a Hugo theme for personal notes/TODO

## Example

You can [view the sample site](https://mmessmore.github.io/hugo-messynotes).
It contains notes with samples of features and some basic documentation of how
to use them.

## Warnings

### This needs a modern-ish browser

This will probably not render reasonably nicely and interactive things probably
won't work in IE or older versions of browsers.  I'm using this to learn
modern CSS/JavaScript, I don't want to clutter the code with hacks and
backfills, and I don't have any of those to test against and wouldn't spend
the time to do so anyway.

## Premise

I love Markdown based note management + revision control.  But all apps I had
used frustrated me.  I made this theme in combination with [a
CLI](https://github.com/mmessmore/messynotes) to make a system using Hugo to do
a lot of the heavy lifting.

It is easiest to consume via the [messynotes
command](https://github.com/mmessmore/messynotes), but this could be used
separately if you're into that.

### This is not yet stable

This theme is still in pretty heavy flux.  Current half-baked changes are
in the [develop](https://github.com/mmessmore/hugo-messynotes/tree/develop)
branch.  They'll get promoted to main as they're fully baked, or at least I
think they are.

Reasons for changes:

- Changing the way I like to view things as I expand my own content
- Adding functionality (bookmarking in the works)
- Learning modern CSS like 'grid' layout
- Learning modern JavaScript APIs

## Principles

All of these are about keeping as minimal as possible.

- **No frameworks**  Everything should be "normal" HTML, CSS, and JavaScript
- **Use hugo first** I'm using Hugo's templating over JavaScript whenever I can
- **Offline support** Everything must degrade gracefully when disconnected to
  the internet.  That means no CDN dependencies for JavaScript libraries.
  Currently the only external dependency is Google Fonts, which are "nice to
  have".  That does mean having to keep static copies of libraries that may get
  out of date.  I'll try to reasonably keep up.  PRs are welcome if I'm not and
  you need it.
- **Limit dependencies** "extra" JavaScript like Mermaid and ABC support should
  only be pulled in when needed.
- **Semantic HTML** avoid HTML hacks for layout, when possible.  Leverage HTML5
  tags.  I am not doing the best job of this and it needs to be improved.


## Features

### Index

The index provides a streamlined view of the content.  What that means to me is
changing as I flesh out both my usage and as I add functionality.  Currently
that means :

- keep TODO notes (I use 1, but it supports multiple) at the top.
- Notes are listed by category below
- Tags are listed as links to the tab pages

Coming soonish:

- A sidebar on the left with categorized bookmarks.

I'm experimenting with using Hugo's "data" support for this.  See the
'develop' branch for that. The example site doesn't have it yet, and I
need to implement CLI support for it as well.

### Categorization

*Tags* serve as a lookup for related notes. *Categories* function more as
a per-topic stream.

I have many tags per note, but one category.  In theory it will work with many
of each.  I think.

### "Private" Junk

Sometimes you share your screen and you don't want stuff showing up by default.
So you can tag a note as "private" and/or add a list of private categories in
your config.

You can toggle them visible at the bottom of the page.

### Light/Dark

The color scheme will change based on the system/browser settings for light and
dark mode.  This is done with a combination of CSS media queries and variables.

### Extended Markup Support

I implement these based on the mermaid example on the Hugo page.  It's pretty
nice to me.

I may add more as I go based on what I want/need.

#### Mermaid.js

Mermaid diagrams are rendered in "mermaid" code fences.  That's pretty normal.

#### ABC

For music/scores, [ABC Notation](https://abcnotation.com/) is rendered from
"abc" code fences.  This is handy for my to jot down song ideas as they come to
me. (I'm not actually good at this)

## Contribution

Contributions to clean it up or generally make things prettier are
appreciated.  I'm not a "frontend" guy, and learning modern stuff is part of
this exercise to me.  Corrections are welcome.

I'll try to stick things I want to fix or would appreciate help with in
[TODO.md](./TODO.md).

PRs are appreciated.  I'll try to deal with bugs in Issues.  Feature requests
without a PR will probably take some time and only be done if they seem useful
to me personally.

Pure bug fix PRs can go against 'main'.  All else should be based on and PRed
to 'develop'.

I have a pretty demanding "real job" as well as a family I enjoy being around.

## License

This is licensed under the [MIT License](./LICENSE).  You are welcome to do
anything within the terms of it: fork, make a product, whatever.  I will not be
offended.  That license was chosen with that in mind.
