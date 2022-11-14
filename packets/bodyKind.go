package packets

type bodyKind string

var NameBodyKind bodyKind

func (bodyKind) HttpDump() string  { return "HttpDump" }
func (bodyKind) HexDump() string   { return "HexDump" }
func (bodyKind) Steam() string     { return "Steam" }
func (bodyKind) Websocket() string { return "Websocket" }
func (bodyKind) ProtoBuf() string  { return "ProtoBuf" }
func (bodyKind) Tdf() string       { return "Tdf" }
func (bodyKind) Taf() string       { return "Taf" }
func (bodyKind) Acc() string       { return "Acc" }
func (bodyKind) Msgpack() string   { return "Msgpack" }
func (bodyKind) Notes() string     { return "Notes" }
func (bodyKind) UnitTest() string  { return "UnitTest" }
func (bodyKind) GitProxy() string  { return "GitProxy" }
