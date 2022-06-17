package composite

import (
	"testing"
)

func TestFolders(t *testing.T) {
	game1 := &file{name: "diablo.exe"}
	game2 := &file{name: "CK3.exe"}
	program1 := &file{name: "vscode.exe"}
	program2 := &file{name: "opera.exe"}

	home := &folder{name: "home"}
	games := &folder{name: "games"}
	programs := &folder{name: "programs"}

	games.add(game1).add(game2)
	programs.add(program1).add(program2)
	home.add(games).add(programs)

	if home.search("pokemon") {
		t.Fatal("Not expected to have a file named pokemon")
	}
	if !home.search("diablo.exe") {
		t.Fatal("Expected to have a file named diablo.exe")
	}
}
