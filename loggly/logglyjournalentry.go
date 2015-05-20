package loggly

import "time"

type JournalEntry struct {
	Pid                     int       `json:"pid,omitempty"`
	Uid                     int       `json:"uid,omitempty"`
	Gid                     int       `json:"gid,omitempty"`
	Comm                    string    `json:"appName,omitempty"`
	Exe                     string    `json:"exe,omitempty"`
	Cmdline                 string    `json:"cmdline,omitempty"`
	CapEffective            string    `json:"capEffective,omitempty"`
	AuditSession            int       `json:"auditSession,omitempty"`
	AuditLoginId            string    `json:"auditLoginId,omitempty"`
	SystemdGroup            string    `json:"systemdCgroup,omitempty"`
	SystemdSession          string    `json:"systemdSession,omitempty"`
	SystemdUnit             string    `json:"systemdUnit,omitempty"`
	SystemdUserInit         string    `json:"systemdUserInit,omitempty"`
	SystemdOwnerUid         string    `json:"systemdOwnerUid,omitempty"`
	SystemdSlice            string    `json:"systemdSlice,omitempty"`
	SelinuxContext          string    `json:"selinuxContext,omitempty"`
	SourceRealtimeTimestamp int64     `json:"sourceRealtimeTimestamp,omitempty"`
	BootId                  string    `json:"bootId,omitempty"`
	MachineId               string    `json:"machineId,omitempty"`
	Hostname                string    `json:"hostname,omitempty"`
	Transport               string    `json:"transport,omitempty"`
	Timestamp               time.Time `json:"timestamp,omitempty"`
	MonotonicTimestamp      int64     `json:"monotonicTimestamp,omitempty"`
	CoredumpUnit            string    `json:"coredumpUnit,omitempty"`
	CoredumpUserInit        string    `json:"coredumpUserInit,omitempty"`
	ObjectPid               int       `json:"objectPid,omitempty"`
	ObjectUid               int       `json:"objectUid,omitempty"`
	ObjectGid               int       `json:"objectGid,omitempty"`
	ObjectComm              string    `json:"objectComm,omitempty"`
	ObjectExe               string    `json:"objectExe,omitempty"`
	ObjectCmdline           string    `json:"objectCmdline,omitempty"`
	ObjectAuditSession      string    `json:"objectAuditSession,omitempty"`
	ObjectAuditLoginId      string    `json:"objectAuditLoginId,omitempty"`
	ObjectSystemdCgroup     string    `json:"objectSystemdCgroup,omitempty"`
	ObjectSystemdSession    string    `json:"objectSystemdSession,omitempty"`
	ObjectSystemdUnit       string    `json:"objectSystemdUnit,omitempty"`
	ObjectSystemdUserInit   string    `json:"objectSystemdUserInit,omitempty"`
	ObjectSystemdOwnerUid   int       `json:"objectSystemdOwnerUid,omitempty"`
	Message                 string    `json:"message,omitempty"`
	MessageId               int       `json:"messageId,omitempty"`
	Priority                int       `json:"priority,omitempty"`
	CodeFile                string    `json:"codeFile,omitempty"`
	CodeLine                string    `json:"codeLine,omitempty"`
	CodeFunc                string    `json:"codeFunc,omitempty"`
	ErrNo                   int       `json:"errNo,omitempty"`
	SyslogFacility          string    `json:"facility,omitempty"`
	SyslogIdentifier        string    `json:"syslogIdentifier,omitempty"`
}
