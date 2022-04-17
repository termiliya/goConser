package goConser

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const (
	FileCreate = 1
	FileAppend = 2
)

type Const2File struct {
	savePath    string
	packageName string
	flag        int
	blocks      []ConstBlock
}

func NewConst2File() *Const2File {
	return &Const2File{flag: FileCreate, blocks: make([]ConstBlock, 0)}
}

func (cf *Const2File) SetFlagAppend() *Const2File {
	cf.flag = FileAppend
	return cf
}

func (cf *Const2File) SetSavePath(path string) *Const2File {
	cf.savePath = path
	return cf
}

func (cf *Const2File) SetPackageName(pkgName string) *Const2File {
	cf.packageName = pkgName
	return cf
}

func (cf *Const2File) SetFlag(flag int) *Const2File {
	cf.flag = flag
	return cf
}

func (cf *Const2File) AddConstBlock(block ConstBlock) *Const2File {
	cf.blocks = append(cf.blocks, block)
	return cf
}

func (cf *Const2File) Run() error {
	if cf.savePath == "" {
		return errors.New("lack of save file path")
	}
	if cf.packageName == "" {
		return errors.New("lack of package name")
	}

	packageName := fmt.Sprintf("package %s\n\n", cf.packageName)
	body := cf.buildBody()
	var (
		f   *os.File
		err error
	)
	if cf.flag == FileCreate {
		if f, err = os.Create(cf.savePath); err != nil {
			return err
		}
	} else {
		if f, err = os.OpenFile(cf.savePath, os.O_WRONLY|os.O_APPEND, 0777); err != nil {
			return err
		}
	}
	defer f.Close()
	_, err = f.WriteString(packageName + body)
	if err != nil {
		return err
	}

	cmd := exec.Command("gofmt", "-w", cf.savePath)
	return cmd.Run()
}

func (cf *Const2File) buildBody() string {
	var body string
	for _, block := range cf.blocks {
		body += block.WriteBlock()
	}
	return body
}
