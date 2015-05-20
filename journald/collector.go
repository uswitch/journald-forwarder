package journald

import "os/exec"
import "encoding/json"
import "bufio"
import "strings"
import "log"

const DefaultSocket = "/var/run/journald-test.sock"

func CollectJournal(c chan JournalEntry) {
	cmd := exec.Command("journalctl", "--output", "json-sse", "--follow")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Could not run journalctl: %v", err)
	}
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		msg := scanner.Text()
		var entry JournalEntry
		err := json.Unmarshal([]byte(strings.Replace(msg, "data:", "", 1)), &entry)
		if err != nil {
			// Ignore blank lines
		} else {
			c <- (entry)
		}
	}
}
