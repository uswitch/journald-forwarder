# journald-forwarder
Forward systemd journals to Loggly

## Usage
`./journald-forwarder -token [LogglyToken] -logFile [where to place logs for this app]`

An example systemd script is available in `packaging/systemd`, which works on CoreOS.

## Building
The easiest way is to run `packaging/docker` which will statically compile journald-forwarder to a bin folder.

## License

journald-forwarder is released under the Apache 2.0 license. See the LICENSE file for details.

Specific components of journald-forwarder use code derivative from software distributed under other licenses; in those cases the appropriate licenses are stipulated alongside the code.
