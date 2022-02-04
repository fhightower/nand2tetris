# nand2tetris

## Intro

This repo contains scripts and content to interactively read the [nand2tetris book (a.k.a. "The Elements of Computing Systems")](https://www.nand2tetris.org/) while running the scripts and programs in a docker container.

## Setup

This setup is designed for users on MacOS; users on other OS-es will have to do analogous steps with different commands.

```bash
brew install socat
brew install xquartz

# log out and back in
open -a Xquartz
# update preferences as defined here: https://cntnr.io/running-guis-with-docker-on-mac-os-x-a14df6a76efc#:~:text=Now%20open%20up%20the%20preferences%20from%20the%20top%20menu%20and%20go%20to%20the%20last%20tab%20%E2%80%98security%E2%80%99.
```

## Usage

Run `socat` in a terminal window "to create a tunnel from an open X11 port (6000) through to the local UNIX socket where XQuartz is listening for connections" ([source](https://blog.alexellis.io/linux-desktop-on-mac/#:~:text=to%20create%20a%20tunnel%20from%20an%20open%20X11%20port%20(6000)%20through%20to%20the%20local%20UNIX%20socket%20where%20XQuartz%20is%20listening%20for%20connections)). Socat will block, so you'll need to open another window after running this command:

```bash
socat TCP-LISTEN:6000,reuseaddr,fork UNIX-CLIENT:\"$DISPLAY\"
```

Now, in another window, run this:

```bash
ifconfig en0 | grep "inet " | cut -d' ' -f2
```

Finally, we are ready to run our docker container (replacing `<YOUR-LOCAL-IP-HERE>` with the `$LOCAL_IP` from the previous code block):

```bash
docker-compose run --rm dev
export DISPLAY=<YOUR-LOCAL-IP-HERE>:0
bash nand2tetris/tools/HardwareSimulator.sh
```

## Credits

- [goropikari/nand2tetris](https://github.com/goropikari/nand2tetris) - This project contributed the docker image and the idea of using x11 to display a GUI in a docker container in the host
- [Running GUI’s with Docker on Mac OS X](https://cntnr.io/running-guis-with-docker-on-mac-os-x-a14df6a76efc) - Very helpful guide showing how to use `socat` and x11/xquartz
- [Bring Linux apps to the Mac Desktop with Docker](https://blog.alexellis.io/linux-desktop-on-mac/) - Ditto ^
- [nand2tetris website](https://www.nand2tetris.org/software) contains the base source code used in this project as well as guides for how to setup a computer to run these files
