// This CLI tool syncs content between two folders.
//
// The synchronization can be two-way or one-way. This behavior is
// controlled by the "bi" argument. If bi-directional synchronization
// is not enabled, it will make the destination folder exactly the same
// as the source folder. Otherwise, the latest change is populated to
// both of the folders.
//
// The synchronization is absolute. It will make sure the two folders
// are exactly the same
package main

import (
	"flag"
	"fmt"
	"time"
)

// DEFAULT serves as a placeholder value for default folder name.

var f1 string
var f2 string
var bi bool

func main() {
	flag.StringVar(&f1, "f1", "",
		"Folder 1 or source if bi is set to false.")
	flag.StringVar(&f2, "f2", "",
		"Folder 2 or destination if bi is set to false.")
	flag.BoolVar(&bi, "bi", true,
		"Whether the synchronization is bi-directional.")
	flag.Parse()

	if len(f1) == 0 || len(f2) == 0 {
		fmt.Printf("f1 or f2 must be set and cannot be empty.\n")
		return
	}

	// TODO: check whether two folders are actually the same folder.

	ticker := time.NewTicker(1 * time.Second)
	quiter := make(chan bool)
	for {
		select {
		case <-ticker.C:
			CheckAndSync(f1, f2, bi)
		case <-quiter:
			ticker.Stop()
			return
		}
	}
}

// CheckAndSync checks and synchronizes two folders f1 and f2.
// The synchronization is by default bi-directional,
// but you can use the "bi" flag to control this behavior.
// When set to true, f1 becomes the source folder and f2
// becomes the destination folder.
func CheckAndSync(f1 string, f2 string, bi bool) {
	fmt.Printf("sync %v and %v, bidirectional: %v\n", f1, f2, bi)
}
