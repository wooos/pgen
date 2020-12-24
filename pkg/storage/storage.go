package storage

import (
    "bytes"
    "io/ioutil"
    "strings"
    "sort"
    "fmt"
)

type Storage struct {
    data    string
    path    string
}

func NewStorage(data, path string) Storage {
    return Storage{
        data: data,
        path: path,
    }
}

func (s Storage) Save() error {
    data, _ := ioutil.ReadFile(s.path)

    buffer := bytes.NewBuffer(data)
    buffer.WriteString(s.data + "\n")
    if err := ioutil.WriteFile(s.path, buffer.Bytes(), 0644); err != nil {
        return err
    }

   return nil
}

func (s Storage) List() {
    data, _ := ioutil.ReadFile(s.path)
    if len(data) <= 0 {
        fmt.Printf("%s\t%s\t%s\n", "Datetime", "Password", "Comment")
        return
    }
    
    lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

    var pl = []int{}
    var cl = []int{}
    for _, line := range lines {
        s := strings.Split(line, ",")
        pl = append(pl, len(s[1]))
        cl = append(cl, len(s[2]))
    }

    sort.Ints(pl)
    sort.Ints(cl)

    formatString := fmt.Sprintf("%%-%ds\t%%-%ds\t%%-%ds\n", 25, pl[len(pl)-1], cl[len(cl)-1])
    fmt.Printf(formatString, "Datetime", "Password", "Comment")
    for _, line := range lines {
        s := strings.Split(line, ",")
        fmt.Printf(formatString, s[0], s[1], s[2])
    }
}
