package packets

type bodyErrorKind string

var NameBodyErrorKind bodyErrorKind

const errorBuf = "this buffer is not "

func (bodyErrorKind) HttpDump() string  { return errorBuf + NameBodyKind.HttpDump() }
func (bodyErrorKind) HexDump() string   { return errorBuf + NameBodyKind.HexDump() }
func (bodyErrorKind) Steam() string     { return errorBuf + NameBodyKind.Steam() }
func (bodyErrorKind) UnitTest() string  { return errorBuf + NameBodyKind.UnitTest() }
func (bodyErrorKind) ProtoBuf() string  { return errorBuf + NameBodyKind.ProtoBuf() }
func (bodyErrorKind) Tdf() string       { return errorBuf + NameBodyKind.Tdf() }
func (bodyErrorKind) Taf() string       { return errorBuf + NameBodyKind.Taf() }
func (bodyErrorKind) Acc() string       { return errorBuf + NameBodyKind.Acc() }
func (bodyErrorKind) Websocket() string { return errorBuf + NameBodyKind.Websocket() }
func (bodyErrorKind) Msgpack() string   { return errorBuf + NameBodyKind.Msgpack() }

//func (bodyErrorKind) Notes() string      { return errorBuf+NameBodyKind.Notes() }
//func (bodyErrorKind) GitProxy() string   { return errorBuf+NameBodyKind.GitProxy() }
