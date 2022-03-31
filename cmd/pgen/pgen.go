package main

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"
	"path"
	"time"

	"github.com/wooos/pgen/internal/version"
	"github.com/wooos/pgen/pkg/storage"
	"github.com/wooos/pgen/pkg/utils"
)

type options struct {
	size    int
	upper   bool
	lower   bool
	digit   bool
	comment string
	list    bool
	version bool
}

func main() {
	o := &options{}
	cmd := cobra.Command{
		Use:   "pgen",
		Short: "password generator",
		Run:   o.runCommand,
	}

	flags := cmd.Flags()
	flags.IntVarP(&o.size, "size", "s", 16, "password size")
	flags.StringVarP(&o.comment, "comment", "c", "", "password comment")
	flags.BoolVar(&o.upper, "upper", true, "include upper letters")
	flags.BoolVar(&o.lower, "lower", true, "include lower letters")
	flags.BoolVar(&o.digit, "digit", true, "include digit")
	flags.BoolVarP(&o.list, "list", "l", false, "list password")
	flags.BoolVarP(&o.version, "version", "v", false, "print version information")

	_ = cmd.Execute()
}

func (o *options) runCommand(cmd *cobra.Command, args []string) {
	storagePath := path.Join(os.Getenv("HOME"), ".pgen")

	if o.list {
		st := storage.NewStorage("", storagePath)
		st.List()
		return
	}

	if o.version {
		fmt.Printf("Pgen Version: %v\n", formatVersion())
		return
	}

	if o.comment == "" {
		fmt.Println(`Error: required flag(s) "comment" not set`)
		return
	}

	pass := utils.RandPassword(o.size, true, true, true)

	currentTime := time.Now().Format(time.RFC3339)
	storageData := fmt.Sprintf("%s,%s,%s", currentTime, pass, o.comment)

	st := storage.NewStorage(storageData, storagePath)
	if err := st.Save(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
}

func formatVersion() string {
	// return version.GetVersion()
	v := version.Get()

	return fmt.Sprintf("%#v", v)
}
