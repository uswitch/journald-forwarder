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

## Dev notes

The forwarder first marshals the json into a journald.JournalEntry struct. Then copies the data into loggly.JournalEntry struct and finally marshals into json. This is done to translate the field names from journal names (e.g. ___PID) into nicer loggly names (e.g. Pid) and additionally to convert the timestamp into miliseconds.

## License

journald-forwarder is released under the Apache 2.0 license. See the LICENSE file for details.
