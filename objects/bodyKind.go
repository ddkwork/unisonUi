package objects

type bodyKind string

var NameBodyKind bodyKind

func (bodyKind) HttpDump() string   { return "HttpDump" }
func (bodyKind) HexDump() string    { return "HexDumpEntry" }
func (bodyKind) Steam() string      { return "Steam" }
func (bodyKind) Head() string       { return "Head" }
func (bodyKind) Pb2() string        { return "Pb2" }
func (bodyKind) Pb3() string        { return "Pb3" }
func (bodyKind) Tdf() string        { return "Tdf" }
func (bodyKind) Taf() string        { return "Taf" }
func (bodyKind) Acc() string        { return "Acc" }
func (bodyKind) Text() string       { return "Text" }
func (bodyKind) Json() string       { return "Json" }
func (bodyKind) Html() string       { return "Html" }
func (bodyKind) Javascript() string { return "Javascript" }
func (bodyKind) Websocket() string  { return "Websocket" }
func (bodyKind) Msgpack() string    { return "Msgpack" }
func (bodyKind) Gzip() string       { return "Gzip" }
func (bodyKind) Notes() string      { return "Notes" }
func (bodyKind) GitProxy() string   { return "GitProxy" }
