# pinba-debugger

Server for debugging Pinba (https://github.com/tony2001/pinba_engine)

# Why?

Sometimes, when you send lots of timers with lots of tags, you want to see,
what you app is sending.

And this simple server is exactly for that.
It just shows data you app sending to Pinba.

# How?

Just `go get` it:
```
go get github.com/olegfedoseev/pinba-debugger
```

And run it:
```
pinba-debugger -in=0.0.0.0:30002
```

# TODO
- [ ]  More verbose output
- [ ]  Filters for `server_name`, `script_name` etc.
