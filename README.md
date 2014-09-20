Housekeeper
===========

Housekeeper is a small, dumb and Zero-Config™ Configuration Manager(sorry for the tautology, but it is).

####Installation

```
    git clone https://github.com/Like-all/housekeeper.git
    cd housekeeper
    go build
```

Prebuilt binaries for various operating systems will be available soon.

####Usage

+ Write down a JSON recipe like that:

```
{
    "preinstall": [
        "sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 8C2C65DE",
        "sudo echo \"deb http://wasteland.it-the-drote.tk/apps/debian/ jessie main contrib non-free\" > /etc/apt/sources.list.d/wasteland.it-the-drote.list",
        "sudo apt-get update"
],
    "packages": ["git", "vim", "screen", "htop", "zsh", "apt-file"],
    "dotfiles": [".vimrc", ".screenrc", ".gitconfig"],
    "postinstall": ["apt-file update"]
}
```

+ Put it in the place accessible via http from your computer(e.g. Github)
+ Invoke housekeeper: `./housekeeper recipe-name` (housekeeper omits '.json' extension)
+ Housekeeper will ask you about your http recipe repository URL
+ Configuration is done!
