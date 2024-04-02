package main

import (
	"fmt"
	"qwriter/packs"
	"os"
	"strings"
	"strconv"
)

func main() {
	note := 1 // probably the first difference from the original (it was 0)
	for note <= packs.NOTELIMIT {
		nval := packs.Inp(fmt.Sprintf("Note %v: ", note), false) // no comment for a while for now
		if nval == "q" {
			os.Exit(13) // see errrorcodes.txt for more info
		} else if nval == "note view L" {
			fmt.Printf("\n---NOTE VIEW L---\n\nNote %v = '%v'\n\n[If your output is empty, it was a command]\n", note-1, packs.Nmap[fmt.Sprintf("NOTE %v", note - 1)])
		} else if nval == "note view S" {
			fmt.Printf("\n---NOTE VIEW S---\n\nNote 1 = '%v'\n\n[If your output is empty, it was a command]\n", packs.Nmap["NOTE 1"])			
		} else if strings.HasPrefix(nval, "note view ") && note > 1 {
			toView := strings.TrimPrefix(nval, "note view ")
			fmt.Printf("\n---NOTE VIEW %v---\n\nNote %v = '%v'\n\n[If your output is empty, it was a command]\n", toView, toView, packs.Nmap[fmt.Sprintf("NOTE %v", toView)])
		} else if strings.HasPrefix(nval, "note del ") {
			toDel := strings.TrimPrefix(nval, "note del ")
			toDeli, _ := strconv.ParseInt(toDel, 10, 64)
			packs.Delete(toDeli)
		} else if nval == "note cache" {
			fmt.Println(packs.Nmap)
		} else if strings.HasPrefix(nval, "note search ") {
			toSrch := strings.TrimPrefix(nval, "note search ")
			lookup, matches := packs.Search(toSrch)
			fmt.Printf("\n---NOTE SEARCH '%v'---\n\n%v Results: %v\n\n", toSrch, matches, lookup)
		} else if strings.HasPrefix(nval, "note export ") {
			toExp := strings.TrimPrefix(nval, "note export ")
			packs.Export(fmt.Sprintf("EXPORT_%v.txt", toExp), toExp)
		} else if nval == "version" {
			fmt.Println("Version Linux 0-A1")
		} else if nval == "help" {
			packs.PrintHelp()
		} else {
			packs.Save(note, nval)
		}
		note++		
	}
}