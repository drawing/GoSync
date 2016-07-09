package sync

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Node struct {
	Name       string
	UpdateTime time.Time
	Md5        []byte
	IsDir      bool
	Children   map[string]*Node
}

type Outline struct {
	Name    string
	Root    string
	Driver  string
	Storage string
	Tree    *Node
}

func (o *Outline) LoadTree() error {
	treeData, err := ioutil.ReadFile(o.Storage)
	if err != nil {
		log.Println("read tree failed:", err)
		return err
	}

	err = json.Unmarshal(treeData, &o.Tree)
	if err != nil {
		log.Println("Unmarshal tree config failed:", err)
		return err
	}
	return nil
}

func (o *Outline) SaveTree() error {
	treeData, err := json.Marshal(o.Tree)
	if err != nil {
		log.Println("Marshal tree failed:", err)
		return err
	}

	err = ioutil.WriteFile(o.Storage, treeData, os.ModePerm)
	if err != nil {
		log.Println("write tree failed:", err)
		return err
	}
	return nil
}
