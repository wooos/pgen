package main

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"
	"path"
	"pgen/pkg/storage"
	"pgen/pkg/utils"
	"time"
)

type options struct {
	length  int
	upper   bool
	lower   bool
	digit   bool
	comment string
	list    bool
}

func main() {
	o := &options{}
	cmd := cobra.Command{
		Use:   "pgen",
		Short: "password generator",
		Run:   o.runCommand,
	}

	flags := cmd.Flags()
	flags.IntVarP(&o.length, "length", "l", 16, "password length")
	flags.StringVarP(&o.comment, "comment", "c", "", "password comment")
	flags.BoolVar(&o.upper, "upper", true, "include upper letters")
	flags.BoolVar(&o.lower, "lower", true, "include lower letters")
	flags.BoolVar(&o.digit, "digit", true, "include digit")
	flags.BoolVar(&o.list, "list", false, "list password")

	_ = cmd.Execute()
}

func (o *options) runCommand(cmd *cobra.Command, args []string) {
	storagePath := path.Join(os.Getenv("HOME"), ".pgen")

    if o.list {
        st := storage.NewStorage("", storagePath)
        st.List()
        return
    }

    if o.comment == "" {
        fmt.Println(`Error: required flag(s) "comment" not set`)
        return
    }

	pass := utils.RandPassword(o.length, true, true, true)

	currentTime := time.Now().Format(time.RFC3339)
	storageData := fmt.Sprintf("%s,%s,%s", currentTime, pass, o.comment)

	st := storage.NewStorage(storageData, storagePath)
	if err := st.Save(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
}

