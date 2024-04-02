package packs

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// user directory (home)

var HomeDir, _ = os.UserHomeDir()

// mapping the map for all the inputs and stuff

var Nmap = make(map[string]string)

// note limit defined here, dont know

const NOTELIMIT = 1000 // apparently good ol' 9999 made it an infinite loop. hehe

// input function because im tired of having to repeat this in my code

func Inp(p string, au bool) string {
	s := bufio.NewScanner(os.Stdin)
	if au {
		fmt.Printf("%v ", p)
		s.Scan()
		return s.Text()
	} else {
		fmt.Print(p)
		s.Scan()
		return s.Text()
	}
	// i think this is only so the compiler is happy lol
	return "Failed to read"
}

// function to append a note to map

func Save(n int, v string) {
	Nmap[fmt.Sprintf("NOTE %v", n)] = v
}

// function to delete a note

func Delete(n int64) {
	Nmap[fmt.Sprintf("NOTE %v", n)] = "Deleted note"
	fmt.Printf("Note %v was deleted successfully!\n", n)
}

// searching func.

func Search(p string) (map[int64]string, int) {
	results := make(map[int64]string)
	i := 0
	for k, v := range Nmap {
		k = strings.TrimPrefix(k, "NOTE ")
		nk, _ := strconv.ParseInt(k, 10,  64)
		if strings.Contains(v, p) {
			results[nk] = v
			i++
		}
	}
	return results, i
}

// checking for invalid blah blah error thing

func Check_err(e error) {
	if e != nil {
		panic(e)
	}
}

// exportin' stuff to a directory ($HOME/Documents/qwriter/)
func Export(fname string, n string) {
	file, e := os.Create(fmt.Sprintf("%v/Documents/qwriter/%v", HomeDir, fname))
	Check_err(e)
    fmt.Printf("File %v has been created!\n", fname)
    _, err := file.WriteString(Nmap[fmt.Sprintf("NOTE %v", n)])
    if err != nil {
    	fmt.Println(err)
    	file.Close()
    	return
    }
    e = file.Close()
    Check_err(e)
    fmt.Println("Now written contents to the file. Find it at $HOME/Documents/qwriter/")
}

// print help message. would take up too much space on the other file lol

func PrintHelp() {
	fmt.Println(`

   ____                _ _               ___               __   _    _      _       
  / __ \              (_) |             / _ \         /\  /_ | | |  | |    | |      
 | |  | |_      ___ __ _| |_ ___ _ __  | | | |______ /  \  | | | |__| | ___| |_ __  
 | |  | \ \ /\ / / '__| | __/ _ \ '__| | | | |______/ /\ \ | | |  __  |/ _ \ | '_ \ 
 | |__| |\ V  V /| |  | | ||  __/ |    | |_| |     / ____ \| | | |  | |  __/ | |_) |
  \___\_\ \_/\_/ |_|  |_|\__\___|_|     \___/     /_/    \_\_| |_|  |_|\___|_| .__/ 
                                                                             | |    
                                                                             |_|  
 This is the help page for Qwriter 0-A1. It displays all the commands you can enter in this
 version and their syntax, if any.

 help ------------------------------------------------------- Displays this menu.

 note view [note] ------------------------------------------- Displays the value of that note.
 															  If you wanted the value of note 6,
 															  you would type note view 6.

 note view S ------------------------------------------------ Displays the first note, hence 'S',
                                                              for 'Start'.

 note view L ------------------------------------------------ Displays your latest note, hence 'L',
                                                              for 'Latest'.
 
 note del [note] -------------------------------------------- Deletes a note, making it void
                                                              from both the program and user.

 note search [keyword] -------------------------------------- Searches all notes to see if they
                                                              contain/match your keyword.

 note export [note] ----------------------------------------- Exports a note to your qwriter directory
                                                              at $HOME/Documents/qwriter/.

 note cache ------------------------------------------------- (Usually only for debugging)
 															  Prints the memory of notes.

 version ---------------------------------------------------- Prints the current version of
                                                              the program.

 Revision 1, by IDKnix
`)
}