---
layout: post
title:  "Initial macOS (Apple Silicon) Setup"
date:   2024-06-21 10:18:53 -0500
---

Sometimes I like to wipe my Macbook and start over but I forget some little things that make
starting over frustrating.

I am not going to give full detail instructions on how I install it. I usually just follow the steps
from the project site. I will only include nuances that I may do different or in addition to the normal steps.

<!--more-->

### App List:
* Homebrew - [https://brew.sh]()
* oh-my-zsh - [https://ohmyz.sh]()
* Docker - [https://docs.docker.com/desktop/install/mac-install/]()
* My dotfiles - [https://github.com/haydenk/dotfiles]()
* asdf - [https://asdf-vm.com]()
* 1Password - [https://1password.com]()

First thing first, `xcode-select --install`, we want to install the Command Line Tools to get some basic packages installed. More specifically,
we want git installed so we can clone Homebrew, oh-my-zsh, and the dotfiles.

While we're at it we will go ahead and install rosetta, `softwareupdate --install-rosetta`. I don't really like installing this if I don't have to but
there's some versions in asdf that do not have an arm version for Apple Silicon.

Homebrew, oh-my-zsh, and dotfiles are the first to be installed. There is no particular order. Homebrew gets installed in the user's home directory under
`~/.brew`, I don't have to worry about system conflicts or permission issues in the home directory because *I* have full access here.

I basically use the information from topbug.net with some revisions.

* [https://www.topbug.net/blog/2013/04/14/install-and-use-gnu-command-line-tools-in-mac-os-x/]()

### Homebrew Packages
* ansible
* asdf
* awscli
* binutils
* bzip2
* coreutils
* curl
* diffutils
* findutils
* gawk
* gnu-getopt
* gnu-sed
* gnu-tar
* gnu-which
* gzip
* grep
* gcc
* gh
* git
* git-flow-avh
* jq
* libpq
* make
* mysql-client
* mysql-client@5.7
* openssl
* pwgen
* session-manager-plugin
* sqlite
* tree
* vim
* wget
* zsh
* zsh-autosuggestions
* zsh-syntax-highlighting

### macOS UI Applications from Homebrew
* firefox
* github
* google-chrome
* slack
* visual-studio-code
* zoom