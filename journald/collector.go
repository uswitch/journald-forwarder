package journald

import "os/exec"
import "encoding/json"
import "bufio"
import "log"

const DefaultSocket = "/var/run/journald-test.sock"

func CollectJournal(c chan JournalEntry) {
	cmd := exec.Command("journalctl", "--output", "json-sse", "--follow")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Could not run journalctl: %v", err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatalf("Could not run journalctl(2): %v", err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		msg := scanner.Text()
		var entry JournalEntry
		if len(msg) > 5 {
			err := json.Unmarshal([]byte(msg[5:]), &entry)
			if err != nil {
				log.Fatalln("unmarshal error", err)
			}
			c <- (entry)
		}
	}
}
