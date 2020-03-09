package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

const filename string = "diagram.bpmn"

func Getcontent(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

type BpmnDiagram struct {
	XMLName   xml.Name `xml:"BPMNDiagram"`
	Id        string   `xml:"id,attr"`
	BpmnPlane []BpmnPlane
}

type BpmnPlane struct {
	XMLName     xml.Name `xml:"BPMNPlane"`
	Id          string   `xml:"id,attr"`
	BpmnElement string   `xml:"bpmnElement,attr"`
	BpmnShape   []BpmnShape
}

type BpmnShape struct {
	XMLName     xml.Name `xml:"BPMNShape"`
	Id          string   `xml:"id,attr"`
	BpmnElement string   `xml:"bpmnElement"`
	Bounds      []Bounds
}

type Bounds struct {
	XMLName xml.Name `xml:"Bounds"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
}

type StartEvent struct {
	XMLName xml.Name `xml:"startEvent"`
	Id      string   `xml:"id,attr"`
}

type Process struct {
	XMLName    xml.Name `xml:"process"`
	Id         string   `xml:"id,attr"`
	IsExe      string   `xml:"isExecutable,attr"`
	StartEvent StartEvent
}

type Definitions struct {
	XMLName     xml.Name `xml:"definitions"`
	Process     Process
	BpmnDiagram BpmnDiagram
}

func Parsecontent(content []byte) {
	d := Definitions{}
	err := xml.Unmarshal(content, &d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Process id = %v\n", d.Process.Id)
	fmt.Printf("IsExecutable = %v\n", d.Process.IsExe)
	fmt.Printf("StartEvent = %v\n", d.Process.StartEvent.Id)
	fmt.Printf("BPMNPlane id = %v\n", d.BpmnDiagram.Id)
	fmt.Printf("BPMNPlane element = %v\n", d.BpmnDiagram.Id)
}

func main() {
	content := Getcontent(filename)
	Parsecontent(content)
}
