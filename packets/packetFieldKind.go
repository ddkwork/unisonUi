package packets

type packetFieldKind string

var NamePacketField packetFieldKind

func (packetFieldKind) Method() string        { return "Method" }
func (packetFieldKind) Scheme() string        { return "Scheme" }
func (packetFieldKind) Host() string          { return "Host" }
func (packetFieldKind) Path() string          { return "Path" }
func (packetFieldKind) ContentType() string   { return "ContentType" } //todo TransferEncoding
func (packetFieldKind) ContentLength() string { return "ContentLength" }
func (packetFieldKind) Status() string        { return "Status" }
func (packetFieldKind) Notes() string         { return "Notes" }
func (packetFieldKind) Process() string       { return "Process" }
func (packetFieldKind) PadTime() string       { return "PadTime" }
