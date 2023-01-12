package file

import "testing"

// cd file
// go test -v
func TestJsonFile(t *testing.T) {
	WriteJSON()
	ReadJSON()
}

func TestXmlFile(t *testing.T) {
	WriteXML()
	ReadXML()
}
func TestGobFile(t *testing.T) {
	WriteGob()
	ReadGob()
}
func TestTextFile(t *testing.T) {
	WriteText()
	ReadText()
}

func TestBinFile(t *testing.T) {
	WriteBin()
	ReadBin()
}

func TestZipFile(t *testing.T) {
	WriteZip()
	ReadZip()
}

func TestTarFile(t *testing.T) {
	WriteTar()
	ReadTar()
}

func TestBuffer(t *testing.T) {
	BufferWrite()
	BufferRead()
}
