<<<<<<< conflict 1 of 1
%%%%%%% diff from: knpxortk a32ea3a1 "run x/tools modernize and tidy prefix string usage (#7)" (parents of rebased revision)
\\\\\\\        to: urqqzzvl 40981e9e (rebase destination)
 package daemon
 
 import (
 	"os/exec"
 	"strings"
 )
 
 // isMounted checks whether the given path is an active mount point.
 // On macOS, mount(8) output format is: <device> on <path> (<options>)
 // We match the " on <path> (" pattern to avoid substring false positives.
 // macOS reports /private/tmp even for /tmp paths.
 func isMounted(path string) bool {
 	out, err := exec.Command("mount").Output()
 	if err != nil {
 		return false
 	}
-	for _, line := range strings.Split(string(out), "\n") {
+	for line := range strings.SplitSeq(string(out), "\n") {
 		if strings.Contains(line, " on "+path+" (") || strings.Contains(line, " on /private"+path+" (") {
 			return true
 		}
 	}
 	return false
 }
+++++++ rotlylmn f141c32a "Rename module to rektide/artifact-fs-extropic, move internal/ to pkg/" (rebased revision)
>>>>>>> conflict 1 of 1 ends
