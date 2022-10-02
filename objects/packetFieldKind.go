package objects

type packetFieldKind string

var NamePacketField packetFieldKind

func (packetFieldKind) Index() string         { return "Index" }
func (packetFieldKind) Method() string        { return "Method" }
func (packetFieldKind) Scheme() string        { return "Scheme" }
func (packetFieldKind) Url() string           { return "Url" }
func (packetFieldKind) ContentType() string   { return "ContentType" } //todo TransferEncoding
func (packetFieldKind) ContentLength() string { return "ContentLength" }
func (packetFieldKind) Status() string        { return "Status" }
func (packetFieldKind) Notes() string         { return "Notes" }
func (packetFieldKind) StartTime() string     { return "StartTime" }
func (packetFieldKind) PadTime() string       { return "PadTime" }
func (packetFieldKind) Request() string       { return "Request" }
func (packetFieldKind) Response() string      { return "Response" }
