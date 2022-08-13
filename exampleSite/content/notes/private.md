---
title: "Private Notes"
date: 2022-08-13T06:57:34-05:00
tags: ["private"]
categories:
  - reference
---

Sometimes you don't want to share

<!--more-->

## Premise

In these days of Zoom meetings, you may be sharing your screen.  When
navigating note to note, it might be nice to hide some entries by default.

## What it does

Private entries are only shown on the index if the little "private" toggle is
turned on at the bottom (off by default).

This is ***NOT*** a security measure.  It just changes the display on the
index. To provide *a* measure to keep from embarrassing yourself.

## How to

To have a hidden category, you could have the following in your site's
`config.yaml` file:

```yaml
params:
  privateCategories:
    - secrets
```

That will hide the category and any notes contained in it from the index.

Or if you just want to make an individual note private you can set it in the
front-matter like

```yaml
private: true
```

## Caveats

This does not change the listing in `/categories` and `/tags`, *only* the index
itself.

There are some weird side-effects.  If the category is not private, but all of
the notes in it are, then the category will still be listed, but with no
entries.

Tags are not dealt with at all (at least yet).  You can't hide a tag, and any
tag you use in a private note will still show on the index.

## Example

[This is an example private note](/notes/secrets)


