# emojilogs

## what is this?

A tiny app that emits some useless logs that happen to contain emoji characters.

Intended to help validate non-ASCII character handling in logging subsystems.

## how?

```
$ ./emojilogs
{"level":"info","time":"2018-10-18T14:01:44+11:00","message":"I am starving and require 🍕 "}
{"level":"info","time":"2018-10-18T14:01:49+11:00","message":"I am thirsty and require 🍺 "}
{"level":"info","time":"2018-10-18T14:01:54+11:00","message":"I am thirsty and require 🍺 "}
{"level":"info","time":"2018-10-18T14:01:59+11:00","message":"I am thirsty and require 🍺 "}
```

## more?

Nope, that's it.

## Docker?

`gcr.io/jsleeio-containers/emojilogs`
