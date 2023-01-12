package file

//单元测试demo
import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func init() {
	log.Println("==============file package init")
}

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func WriteJSON() {
	info := []Website{
		{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}},
		{"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}
	// 创建文件
	filePtr, err := os.Create("info.json")
	if err != nil {
		fmt.Println("WriteJSON 文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("WriteJSON 编码错误", err.Error())
	} else {
		fmt.Println("WriteJSON 编码成功")
	}
}

func ReadJSON() {
	filePtr, err := os.Open("./info.json")
	if err != nil {
		fmt.Printf("ReadJSON 文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var info []Website
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("ReadJSON 解码失败", err.Error())
	} else {
		fmt.Println("ReadJSON 解码成功")
		fmt.Println("ReadJSON:", info)
	}
}

func WriteXML() {
	//实例化对象
	info := Website{"C语言中文网", "http://c.biancheng.net/golang/", []string{"Go语言入门教程", "Golang入门教程"}}
	f, err := os.Create("./info.xml")
	if err != nil {
		fmt.Println("WriteXML 文件创建失败", err.Error())
		return
	}
	defer f.Close()
	//序列化到文件中
	encoder := xml.NewEncoder(f)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("WriteXML 编码错误：", err.Error())
		return
	} else {
		fmt.Println("WriteXML 编码成功")
	}
}
func ReadXML() {
	//打开xml文件
	file, err := os.Open("./info.xml")
	if err != nil {
		fmt.Printf("文件打开失败：%v", err)
		return
	}
	defer file.Close()
	info := Website{}
	//创建 xml 解码器
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Printf("解码失败：%v", err)
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

// 为了让某个数据结构能够在网络上传输或能够保存至文件，它必须被编码然后再解码。当然已经有许多可用的编码方式了，比如 JSON、XML、Google 的 protocol buffers 等等。
// 而现在又多了一种，由Go语言 encoding/gob 包提供的方式。

// Gob 是Go语言自己以二进制形式序列化和反序列化程序数据的格式，可以在 encoding 包中找到。这种格式的数据简称为 Gob（即 Go binary 的缩写）。
// 类似于 Python 的“pickle”和 Java 的“Serialization”。

// Gob 和 JSON 的 pack 之类的方法一样，由发送端使用 Encoder 对数据结构进行编码。在接收端收到消息之后，接收端使用 Decoder 将序列化的数据变化成本地变量。

// Go语言可以通过 JSON 或 Gob 来序列化 struct 对象，虽然 JSON 的序列化更为通用，但利用 Gob 编码可以实现 JSON 所不能支持的 struct 的方法序列化，
// 利用 Gob 包序列化 struct 保存到本地也十分简单。

// Gob 不是可外部定义、语言无关的编码方式，它的首选的是二进制格式，而不是像 JSON 或 XML 那样的文本格式。Gob 并不是一种不同于 Go 的语言，而是在编码和解码过程中用到了 Go 的反射。

// Gob 通常用于远程方法调用参数和结果的传输，以及应用程序和机器之间的数据传输。它和 JSON 或 XML 有什么不同呢？Gob 特定的用于纯 Go 的环境中，
// 例如两个用Go语言写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。

// Gob 文件或流是完全自描述的，它里面包含的所有类型都有一个对应的描述，并且都是可以用Go语言解码，而不需要了解文件的内容。

// 只有可导出的字段会被编码，零值会被忽略。在解码结构体的时候，只有同时匹配名称和可兼容类型的字段才会被解码。
// 当源数据类型增加新字段后，Gob 解码客户端仍然可以以这种方式正常工作。解码客户端会继续识别以前存在的字段，
// 并且还提供了很大的灵活性，比如在发送者看来，整数被编码成没有固定长度的可变长度，而忽略具体的 Go 类型。

func WriteGob() {
	info := map[string]string{
		"name":    "C语言中文网",
		"website": "http://c.biancheng.net/golang/",
	}
	name := "info.gob"
	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer File.Close()
	enc := gob.NewEncoder(File)
	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
	fmt.Println("WriteGob 编码成功")
}
func ReadGob() {
	var M map[string]string
	File, _ := os.Open("info.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println("ReadGob 解码成功:", M)
}

func WriteText() {
	//创建一个新文件，写入内容
	filePath := "./info.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	//写入内容
	str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt文件要看到换行效果要用 \r\n
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
	fmt.Println("WriteText done")
}
func ReadText() {
	//打开文件
	file, err := os.Open("./info.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}

type WebsiteX struct {
	Url int32
}

func WriteBin() {
	file, err := os.Create("info.bin")
	for i := 1; i <= 10; i++ {
		info := WebsiteX{
			int32(i),
		}
		if err != nil {
			fmt.Println("WriteBin 文件创建失败 ", err.Error())
			return
		}
		defer file.Close()
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.LittleEndian, info)
		b := bin_buf.Bytes()
		_, err = file.Write(b)
		if err != nil {
			fmt.Println("WriteBin 编码失败", err.Error())
			return
		}
	}
	fmt.Println("WriteBin 编码成功")
}

func ReadBin() {
	file, err := os.Open("info.bin")
	defer file.Close()
	if err != nil {
		fmt.Println("ReadBin文件打开失败", err.Error())
		return
	}
	m := WebsiteX{}
	for i := 1; i <= 10; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("ReadBin二进制文件读取失败", err)
			return
		}
		fmt.Println("ReadBin第", i, "个值为：", m)
	}
}
func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("ReadBin解码失败", err)
	}
	return bytes
}
func WriteZip() {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)
	// 创建一个压缩文档
	w := zip.NewWriter(buf)
	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "http://c.biancheng.net/golang/"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println("WriteZip Create:", err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println("WriteZip Write:", err)
		}
	}
	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println("WriteZip Close:", err)
	}
	// 将压缩文档内容写入文件
	f, err := os.OpenFile("info.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("WriteZip Close:", err)
	}
	buf.WriteTo(f)
	fmt.Println("WriteZip 成功")
}

func ReadZip() {
	// 打开一个zip格式文件
	r, err := zip.OpenReader("info.zip")
	if err != nil {
		fmt.Println("ReadZip OpenReader error：", err.Error())
	}
	defer r.Close()
	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range r.File {
		fmt.Printf("ReadZip 文件名: %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println("ReadZip Open error：", err.Error())
		}
		//输出到控制台
		_, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
		if err != nil {
			fmt.Println("ReadZip CopyN error：", err.Error())
		}
		rc.Close()
	}
}

// 创建 tar 归档文件
// tar 是一种打包格式，但不对文件进行压缩，所以打包后的文档一般远远大于 zip 和 tar.gz，因为不需要压缩的原因，所以打包的速度是非常快的，打包时 CPU 占用率也很低。

// tar 的目的在于方便文件的管理，比如在我们的生活中，有很多小物品分散在房间的各个角落，为了方便整洁可以将这些零散的物品整理进一个箱子中，而 tar 的功能就类似这样。

// 创建 tar 归档文件与创建 .zip 归档文件非常类似，主要不同点在于我们将所有数据都写入相同的 writer 中，并且在写入文件的数据之前必须写入完整的头部，而非仅仅是一个文件名。

// tar 打包实现原理如下：
// 创建一个文件 x.tar，然后向 x.tar 写入 tar 头部信息；
// 打开要被 tar 的文件，向 x.tar 写入头部信息，然后向 x.tar 写入文件信息；
// 当有多个文件需要被 tar 时，重复第二步直到所有文件都被写入到 x.tar 中；
// 关闭 x.tar，完成打包。
func WriteTar() {
	f, err := os.Create("./info.tar") //创建一个 tar 文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	tw := tar.NewWriter(f)
	defer tw.Close()
	fileinfo, err := os.Stat("./info.txt") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err)
	}
	err = tw.WriteHeader(hdr) //写入头文件信息
	if err != nil {
		fmt.Println(err)
	}
	f1, err := os.Open("./info.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) //将main.exe文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

func ReadTar() {
	f, err := os.Open("info.tar")
	if err != nil {
		fmt.Println("ReadTar 文件打开失败", err)
		return
	}
	defer f.Close()
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println("ReadTar fileName", err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println("ReadTar fileName", fileinfo.Name())
		f, err := os.Create("123" + fileinfo.Name())
		if err != nil {
			fmt.Println("ReadTar Create error:", err)
		}
		defer f.Close()
		//文件拷贝
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println("ReadTar Copy error:", err)
		}
	}
}

// buffer 是缓冲器的意思，Go语言要实现缓冲读取需要使用到 bufio 包。bufio 包本身包装了 io.Reader 和 io.Writer 对象，
// 同时创建了另外的 Reader 和 Writer 对象，因此对于文本 I/O 来说，bufio 包提供了一定的便利性。

// buffer 缓冲器的实现原理就是，将文件读取进缓冲（内存）之中，再次读取的时候就可以避免文件系统的 I/O 从而提高速度。同理在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统。

func BufferWrite() {
	name := "buffer.txt"
	content := "http://c.biancheng.net/golang/"
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("BufferWrite 文件打开失败", err)
	}
	defer fileObj.Close()
	writeObj := bufio.NewWriterSize(fileObj, 4096)
	//使用 Write 方法,需要使用 Writer 对象的 Flush 方法将 buffer 中的数据刷到磁盘
	buf := []byte(content)
	if _, err := writeObj.Write(buf); err == nil {
		if err := writeObj.Flush(); err != nil {
			panic(err)
		}
		fmt.Println("BufferWrite 数据写入成功")
	}
}

func BufferRead() {
	fileObj, err := os.Open("buffer.txt")
	if err != nil {
		fmt.Println("BufferRead 文件打开失败：", err)
		return
	}
	defer fileObj.Close()
	//一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
	reader := bufio.NewReader(fileObj)
	buf := make([]byte, 1024)
	//读取 Reader 对象中的内容到 []byte 类型的 buf 中
	info, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("BufferRead 读取的字节数:" + strconv.Itoa(info))
	//这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
	fmt.Println("BufferRead 读取的文件内容:", string(buf))
}
