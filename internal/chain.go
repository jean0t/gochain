package internal

import (
    "encoding/json"
    "crypto/sha256"
    "time"
    "fmt"
    "os"
)

//====================================================== Data declaration
type Data struct {
    Value float64
    From string
    To string
}

type Node struct {
    Parent string
    Timestamp int64
    Data Data
    Signature string
}

type Chain struct {
    Nodes []Node
    Length int
}

func (c *Chain) GetLastElementSignature() string {
    return c.Nodes[c.Length-1].Signature
}

func (c *Chain) PrintChain() {
    jsonData, err := json.MarshalIndent(c.Nodes, "", " ")
    if err != nil {
        return
    }

    fmt.Println(string(jsonData))
}

func (c *Chain) AddNode(info Data) bool {
    var timeNow int64 = time.Now().UnixNano()
    var newNode Node = Node{
    Parent: c.GetLastElementSignature(),
    Timestamp: timeNow,
    Data: info,
    Signature: fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("from %s - to %s - value %.4f - time %d", info.From, info.To, info.Value, timeNow))))}


    c.Nodes = append(c.Nodes, newNode)
    c.Length++
    return true
}

func (c *Chain) Save() error {
    file, err := os.Create("GoChain.json")
    if err != nil {
        return err
    }
    defer file.Close()

    var encoder *json.Encoder = json.NewEncoder(file)
    encoder.SetIndent("", " ")
    return encoder.Encode(c)
}

func (c *Chain) Load() (error) {
    file, err := os.Open("GoChain.json")
    if err != nil {
        return err
    }
    defer file.Close()

    var decoder *json.Decoder = json.NewDecoder(file)
    _ = decoder.Decode(c)
    return nil
}

//======================================================| Creation of the chain
var genesisBlock = Node{
    Parent: "",
    Timestamp: time.Now().UnixNano(),
    Data: Data{Value: 30000, From: "creator", To: "jean0t"},
    Signature: fmt.Sprintf("%x", sha256.Sum256([]byte("this is the first operation of the chain :)"))),
}

func CreateData(fromAccount, toAccount string, valueExchanged float64) Data {
    var data Data = Data{From: fromAccount, To: toAccount, Value: valueExchanged}
    return data
}

func CreateChain() {
    var tmp Chain = Chain{Nodes: []Node{genesisBlock}, Length: 1}
    tmp.Save()
}
