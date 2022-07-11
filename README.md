# Barenotes: a Hugo theme and stuff

I was frustrated with personal note taking systems and TODO list managers.
So I made a thing.

## Premises

These were the things I wanted

1. Use markdown
   - I want plain text, that can render into pages
   - I want to use git with it
2. Use vim
   - I'm tired of hitting ESC in some electron app
3. I want it to render on save, not as I type
4. I want to be able to super easily access it whenever I need to

## Story

I started thinking about writing a tool that would render the pages and do
the updates on file changes and, and, and...  I was about to try to write
[Hugo](https://gohugo.io/).  I love Hugo.  I'll just use that.

## CLI

The only issue with Hugo was that I wanted to be working on whatever and
then just open my TODO list, pull it up, create a note, etc.

So I made a [wrapper CLI](./cli) for that.  When I create a config
telling it where my notes repo is, I can type a simple command from
whatever directory I'm in and do that.  It's a glorified shell script.
It actually was a shell script and it's about 5% smarter now.  But it
allows me to be able to add more functionality, sane config, etc.

## The Theme

It's pretty spartan and that's intentional.  Just very simple HTML and
a sprinkle of CSS.  No JavaScript at this point.  I'll try to avoid it
if possible.  If I add any, it will be vanilla JS.  No frameworks.

## Install the Theme/Set up Hugo To set it up it's probably easiest to
do this:

```bash
mkdir ~/notes # or whatever directory
cd ~/notes
git init
mkdir themes
git submodule add https://github.com/mmessmore/hugo-barenotes.git themes/barenotes
cp -r themes/barenotes/exampleSite/* .
```

If you don't want to submodule and just want to download the theme you could do
something like
```bash
wget https://github.com/mmessmore/hugo-barenotes/archive/refs/heads/main.zip
hugo init notes
cd notes/themes
unzip ../../main.zip
rm ../../main.zip
```

Whatever you want, just have this in your `themes` directory and see the
`exampleSite` subdirectory for the expected layout, which is basically
just the config file and a `contents/notes/TODO.md` file with `weight: 1`
in the frontmatter.


## Install the cli tool
If you have `go` installed you can run
```bash
cd themes/barenotes/cli
go install # installs as "barenotes" wherever go sticks binaries

# or
make install # installs as "notes" in ~/bin or set BINDIR to an alternate path
```

The `barenotes` or `notes` cli is relatively feature complete, but I'll
probably test-drive it a bit more before I put up a release.


## CLI Usage
```text
The barenotes hugo theme is designed to be a minimalistic
system for maintaining personal notes and todo items.

This is a wrapper cli for providing simple access to maintain and use it
in this fashion.

This will try its best to choose the text editor of your choice.  The order
of precidence (first set wins): command line argument, config file,
$VISUAL, $EDITOR, editor command (for update-alternatives), and vi.  If
none are found, commands like 'new' will fail and you will need to
specify in the config or on the command line.

For the browser it will walk command line argument, config file, the 'open'
command, and then the 'xdg-open' command.  If none work, it will fail.

Usage:
  barenotes [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  new         Create a new note
  open        Open a web browser to the hugo url (http://localhost:1313)
  restart     Restart hugo server
  showConfig  display configuration specified or implied
  start       Run the hugo server and open the browser to it.
  stop        Stop hugo server
  todo        Edit TODO file

Flags:
  -b, --browser string   Web browser to use
      --config string    config file (default $HOME/.barenotes.yaml)
  -e, --editor string    Text editor to use
  -h, --help             help for barenotes
  -H, --hugo string      Hugo binary (default "hugo")
  -r, --root string      Root of hugo repository (default ".")

Use "barenotes [command] --help" for more information about a command.
```

## External tools and config

It will do its best to find a web browser and editor for you.

For the browser it will try to use the config, then the "system default"
a few ways and then wander into an arbitrary list of browsers.

1. command line option
2. config file
3. `open` for macos
4. `xdg-open` for Linux/BSD with XDG stuff
5. `x-www-browser` for /etc/alternatives (ie `update-alternatives`)
6. `firefox`
7. `google-chrome-stable`
8. `chromium-browser`

For the editor it will walk though:

1. command line option
2. config file
3. `$VISUAL` environment variable
4. `$EDITOR` environment variable
5. `editor` for /etc/alternatives
6. `vi`
7. `nano`

Without specifying a "root" in the config or on the command line, it
will assume that `.` is the root.  Setting that will let you just run the
command from anywhere.  You should do that.

You can configure those, the location of the Hugo binary and the root
directory of the repo in `~/.barenotes.yaml`.

To take the working config as a starting point you can run `barenotes
showConfig -y > ~/.barenotes.yaml` and edit it.

# Contributing

I'm open to making this prettier, but not adding external requirements.
Code fixes and improvements are great. So feel free to put in a PR.

You are also welcome to fork and dress it up past what I want to.

This is a spare-time hobby project. I do have a real job that is pretty
consuming and a family and responsibilities and stuff.  So I may not be
prompt in addressing things.

PRs will always be prioritized.  I'll try to address issues as well, when I
can.  Bugs will, well, bug me.  So I'll try to get to those (but a PR is
better).  Feature requests may take a while and my progress on this will
be bursty.

# Code of Conduct

Be nice to me.  Be nice to people submitting issues and PRs.  There are plenty
of other places to be angry or mean on the internet.  Aggressive or rude
behavior and language towards other people won't be tolerated even if you are
"right."  Foul words in comments is fine as long as they are not hateful to
other humans.

# License

This is licensed under the [MIT License](./LICENSE).

You are welcome to do anything under those provisions (fork, make a commercial
product, whatever).  If you want to include SweetFramework.js, cool icon sets,
and that sort of thing but have this structure as a starting point, I will be
flattered, not irritated, if you fork.  I'm not a front-end guy, and that's not
what I want for myself.  Everyone has their own tastes and mine are not
"right."
