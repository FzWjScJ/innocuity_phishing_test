package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	mu          sync.Mutex
	hostInfoMap = make(map[string]int)
)

func main() {
	http.HandleFunc("/", displayStats)
	http.HandleFunc("/save", saveToFileAndReset)
	go http.ListenAndServe(":10004", nil)

	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	hostInfo := conn.RemoteAddr().String()
	hostName := make([]byte, 256)
	n, _ := conn.Read(hostName)
	hostInfo = string(hostName[:n])

	mu.Lock()
	hostInfoMap[hostInfo]++
	mu.Unlock()
}

func displayStats(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><head><title>Stats</title></head><body>"))
	w.Write([]byte("<table border='1'><tr><th>Hostname</th><th>IP</th><th>Count</th></tr>"))

	for info, count := range hostInfoMap {
		parts := strings.Split(info, "\t")
		if len(parts) >= 2 {
			w.Write([]byte(fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td></tr>", parts[0], parts[1], count)))
		} else {
			log.Printf("Unexpected format: %s\n", info)
		}
	}

	w.Write([]byte("</table>"))
	w.Write([]byte("<form action='/save' method='post'><input type='submit' value='Save to File and Reset'></form>"))
	w.Write([]byte("</body></html>"))
}

func saveToFileAndReset(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Create 'log' directory if it doesn't exist
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
	}

	// Format the current time to create a unique file name
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	filePath := filepath.Join(logDir, fmt.Sprintf("data_%s.txt", currentTime))

	file, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	for info, count := range hostInfoMap {
		file.WriteString(fmt.Sprintf("%s\t%d\n", info, count))
	}

	hostInfoMap = make(map[string]int)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
