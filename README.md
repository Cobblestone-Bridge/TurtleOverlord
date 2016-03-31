# TurtleOverlord
This is a "small" Turtle Management System for the Computer Craft Minecraft Mod.
The web server is built with the [Goji](http://github.com/zenazn/goji) micro framework and the [Gorilla Toolkit](http://github.com/gorilla).

# Installing
```
go get github.com/Cobblestone-Bridge/TurtleOverlord
```

# Glide
To fetch these dependencies automatically we use [glide](http://github.com/Masterminds/glide).

## Install
```
go get github.com/Masterminds/glide
```

## Run
If you run inside the TurtleOverlord source folder:
```
glide install
```
The dependencies will automatically be installed.

# Dependencies
 ```
 go get github.com/golang/glog
 go get github.com/gorilla/context
 go get github.com/gorilla/sessions
 go get github.com/zenazn/goji
 go get gopkg.in/mgo.v2
 go get golang.org/x/crypto
 ```

# Continuous Development

For Continuous Development I recommend using `Fresh` - https://github.com/pilu/fresh

You can install `Fresh` by issuing:

```
go get github.com/pilu/fresh
```

If you run in the TurtleOverlord source folder:

```
fresh -c runner.conf
```

Project should automatically rebuild itself when a change occurs.

# Credit
The layout of this project was "borrowed" from https://github.com/elcct/defaultproject.
