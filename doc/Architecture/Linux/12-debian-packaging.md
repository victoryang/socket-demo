# Debian Packaging

[Debian Pakcaging](https://wiki.debian.org/Packaging/Intro?action=show&redirect=IntroDebianPackaging)

## Introduction

### Requirements

- build-essential
- devscript
- debhelper

### Core Conception

- upstream tarball
- source package
    - upstream tarball
    - a debian directory, with changes made to upstream source, plus all the files created for the Debian package. This has a .debian.tar.gz ending
    - a descirption file(with .dsc ending), which list the other two files
- binary package

### The packaging work flow

#### Rename the upstream tarball