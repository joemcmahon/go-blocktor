# Blocktor
Simple application to block requests from TOR network, based from https://check.torproject.org/.

## Compile

```bash

$ git clone https://github.com/paulopinda/blocktor.git
$ cd blocktor
$ go build *.go

```

## How to use

List of all Tor exit nodes from the past 16 hours that can contact xxx.xxx.xxx.xxx on port 80 and 443.

```
$ sudo ./blocktor -address xxx.xxx.xxx.xxx
```


List of all Tor exit nodes from the past 16 hours that can contact xxx.xxx.xxx.xxx on port 80 and 443 and load script with initials rules.

``` bash
$ sudo ./blocktor -address xxx.xxx.xxx.xxx -initial-rules ./init.sh # default is ./start.sh

```
