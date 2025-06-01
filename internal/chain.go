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
    value float64
    from string
    to string
}

type Node struct {
    parent [32]byte
    timestamp int64
    data Data
    signature [32]byte
}

type Chain struct {
    nodes []Node
    length int
}

func (c *Chain) GetLastElementSignature() [32]byte {
    return c.nodes[c.length-1].signature
}

func (c *Chain) PrintChain() {
    jsonData, err := json.MarshalIndent(c.nodes, "", " ")
    if err != nil {
        return
    }

    fmt.Println(string(jsonData))
}

func (c *Chain) AddNode(info Data) bool {
    var timeNow int64 = time.Now().UnixNano()
    var newNode Node = Node{
    parent: c.GetLastElementSignature(),
    timestamp: timeNow,
    data: info,
    signature: sha256.Sum256([]byte(fmt.Sprintf("from %s - to %s - value %.4f - time %d", info.from, info.to, info.value, timeNow)))}


    c.nodes = append(c.nodes, newNode)
    c.length++
    return true
}

func (c *Chain) Save(data Chain) error {
    file, err := os.Create("GoChain.json")
    if err != nil {
        return err
    }
    defer file.Close()

    var encoder *json.Encoder = json.NewEncoder(file)
    encoder.SetIndent("", " ")
    return encoder.Encode(data)
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
    parent: [32]byte{},
    timestamp: time.Now().UnixNano(),
    data: Data{value: 30000, from: "jean0t", to: "jean0t"},
    signature: sha256.Sum256([]byte("this is the first operation of the chain :)")),
}

func CreateData(fromAccount, toAccount string, valueExchanged float64) Data {
    var data Data = Data{from: fromAccount, to: toAccount, value: valueExchanged}
    return data
}

func CreateChain() *Chain {
    return &Chain{nodes: []Node{genesisBlock}, length: 1}
}
