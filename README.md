# godep2portmk

This is a really, *really*, tacky script to parse the Gopkg.lock file of Golang projects that use the new 
[golang/dep](https://github.com/golang/dep) dependancy manager thingy, and output the correct format for a 
FreeBSD Ports Makefile.

I've specifically created it for [gohugoio/hugo](https://github.com/gohugoio/hugo), 
and so there are a few things that may break.

## Setup
```
mkdir godep2portmk             
cd godep2portmk/
export GOPATH=${PWD}
mkdir src
cd src/
git clone https://github.com/forquare/godep2portmk
 Cloning into 'godep2portmk'...
 remote: Counting objects: 17, done.
 remote: Compressing objects: 100% (10/10), done.
 remote: Total 17 (delta 5), reused 17 (delta 5), pack-reused 0
 Unpacking objects: 100% (17/17), done.
 
cd godep2portmk/
glide update && glide install
 [INFO]	Waiting on Glide global cache access
 [WARN]	The name listed in the config file () does not match the current location (godep2portmk)
 [INFO]	Downloading dependencies. Please wait...
 [INFO]	No references set.
 [INFO]	Resolving imports
 [INFO]	Downloading dependencies. Please wait...
 [INFO]	--> Fetching updates for github.com/pelletier/go-toml
 [INFO]	Setting references for remaining imports
 [INFO]	Exporting resolved dependencies...
 [INFO]	--> Exporting github.com/pelletier/go-toml
 [INFO]	Replacing existing vendor dependencies
 [INFO]	Project relies on 1 dependencies.
 [WARN]	The name listed in the config file () does not match the current location (godep2portmk)
 [INFO]	Downloading dependencies. Please wait...
 [INFO]	--> Found desired version locally github.com/pelletier/go-toml 4e9e0ee19b60b13eb79915933f44d8ed5f268bdd!
 [INFO]	Setting references.
 [INFO]	--> Setting version for github.com/pelletier/go-toml to 4e9e0ee19b60b13eb79915933f44d8ed5f268bdd.
 [INFO]	Exporting resolved dependencies...
 [INFO]	--> Exporting github.com/pelletier/go-toml
 [INFO]	Replacing existing vendor dependencies
 
go install
cd ../../bin
```

## Usage
Either use `godep2portmk` on a file:
```
./godep2portmk /path/to/my/Gopkg.lock
```

Or use `godep2portmk` from stdin:
```
curl -s https://example.com/somerepo/Gopkg.lock | ./godep2portmk 
```
