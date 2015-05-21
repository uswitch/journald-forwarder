# journald-forwarder
Forward systemd journals to Loggly

## Usage
`./journald-forwarder -token [LogglyToken] -logFile [where to place logs for this app] -tag [Loggly tag] -logFile [path to logfile]`

### Defaults
`-logFile = /var/log/journald-forwarder.log`
`-tag = ""`

An example systemd script is available in [packaging/systemd](https://github.com/uswitch/journald-forwarder/blob/master/packaging/systemd/journald-forwarder.service), which works on CoreOS.

## Building
The easiest way is to run `packaging/docker` which will statically compile journald-forwarder to a bin folder.

## License

journald-forwarder is released under the Apache 2.0 license. See the LICENSE file for details.
