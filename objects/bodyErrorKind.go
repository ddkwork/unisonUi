package objects

//bodyErrorKind

type bodyErrorKind string

var NameBodyErrorKind bodyErrorKind

const errorBuf = "this buffer is not "

func (bodyErrorKind) HttpDump() string   { return errorBuf + NameBodyKind.HttpDump() }
func (bodyErrorKind) HexDump() string    { return errorBuf + NameBodyKind.HexDump() }
func (bodyErrorKind) Steam() string      { return errorBuf + NameBodyKind.Steam() }
func (bodyErrorKind) Head() string       { return errorBuf + NameBodyKind.Head() }
func (bodyErrorKind) Pb2() string        { return errorBuf + NameBodyKind.Pb2() }
func (bodyErrorKind) Pb3() string        { return errorBuf + NameBodyKind.Pb3() }
func (bodyErrorKind) Tdf() string        { return errorBuf + NameBodyKind.Tdf() }
func (bodyErrorKind) Taf() string        { return errorBuf + NameBodyKind.Taf() }
func (bodyErrorKind) Acc() string        { return errorBuf + NameBodyKind.Acc() }
func (bodyErrorKind) Text() string       { return errorBuf + NameBodyKind.Text() }
func (bodyErrorKind) Json() string       { return errorBuf + NameBodyKind.Json() }
func (bodyErrorKind) Html() string       { return errorBuf + NameBodyKind.Html() }
func (bodyErrorKind) Javascript() string { return errorBuf + NameBodyKind.Javascript() }
func (bodyErrorKind) Websocket() string  { return errorBuf + NameBodyKind.Websocket() }
func (bodyErrorKind) Msgpack() string    { return errorBuf + NameBodyKind.Msgpack() }
func (bodyErrorKind) Gzip() string       { return errorBuf + NameBodyKind.Gzip() }

//func (bodyErrorKind) Notes() string      { return errorBuf+NameBodyKind.Notes() }
//func (bodyErrorKind) GitProxy() string   { return errorBuf+NameBodyKind.GitProxy() }
