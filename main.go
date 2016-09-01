package main

import (
	"fmt"
	"github.com/rhysd/dotfiles/src"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	cli = kingpin.New("dotfiles", "A dotfiles manager")

	clone       = cli.Command("clone", "Clone remote repository")
	clone_repo  = clone.Arg("repository", "Repository.  Format: 'user', 'user/repo-name', 'git@somewhere.com:repo.git, 'https://somewhere.com/repo.git'").Required().String()
	clone_path  = clone.Arg("path", "Path where repository cloned").String()
	clone_https = clone.Flag("https", "Use https:// instead of git@ protocol for `git clone`.").Short('h').Bool()

	link           = cli.Command("link", "Put symlinks to setup your configurations")
	link_dryrun    = link.Flag("dry", "Show what happens only").Bool()
	link_repo      = link.Arg("repo", "Path to your dotfiles repository.  If omitted, the current directory is assumed to be dotfiles repository.").String()
	link_specified = link.Arg("files", "Files to link. If you specify no file, all will be linked.").Strings()
	// TODO link_no_default = link.Flag("no-default", "Link files specified by mappings.json and mappings_*.json")

	list = cli.Command("list", "Show a list of symbolic link put by this command")

	clean      = cli.Command("clean", "Remove all symbolic links put by this command")
	clean_repo = clean.Arg("repo", "Path to your dotfiles repository.  If omitted, the current directory is assumed to be dotfiles repository.").String()

	update = cli.Command("update", "Update your dotfiles repository")

	version = cli.Command("version", "Show version")
)

func unimplemented(cmd string) {
	fmt.Fprintf(os.Stderr, "Command '%s' is not implemented yet!\n", cmd)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func main() {
	switch kingpin.MustParse(cli.Parse(os.Args[1:])) {
	case clone.FullCommand():
		handleError(dotfiles.Clone(*clone_repo, *clone_path, *clone_https))
	case link.FullCommand():
		handleError(dotfiles.Link(*link_repo, *link_specified, *link_dryrun))
	case list.FullCommand():
		unimplemented("list")
	case clean.FullCommand():
		handleError(dotfiles.Clean(*clean_repo))
	case update.FullCommand():
		unimplemented("update")
	case version.FullCommand():
		fmt.Println(dotfiles.Version())
	default:
		panic("Internal error: Unreachable! Please report this to https://github.com/rhysd/dotfiles/issues")
	}
}
