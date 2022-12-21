package talk

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartTalkServer() {
	log.Println("====start TalkServer===========")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster() //广播消息
	for {
		conn, err := listener.Accept() //监听
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) //处理客户端连接
	}
}

type client chan<- string // 对外发送消息的通道
var (
	entering = make(chan client) //连接客户端通道
	leaving  = make(chan client) //断开的客户端通道
	messages = make(chan string) // 所有连接的客户端
)

// 广播接收到的消息到所有客户端
// 使用一个 select 开启一个多路复用：
// 1. 每当有广播消息从 messages 发送进来，都会循环 cliens 对里面的每个 channel 发消息。
// 2. 每当有消息从 entering 里面发送过来，就生成一个新的 key - value，相当于给 clients 里面增加一个新的 client。
// 3. 每当有消息从 leaving 里面发送过来，就删掉这个 key - value 对，并关闭对应的 channel。
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 把所有接收到的消息广播给所有客户端
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering: //客户端建立连接
			clients[cli] = true
		case cli := <-leaving: //客户端断开
			delete(clients, cli)
			close(cli)
		}
	}
}

// 处理连接
// 1. 获取连接过来的 ip 地址和端口号；
// 2. 把欢迎信息写进 channel 返回给客户端；
// 3. 生成一条广播消息写进 messages 里；
// 4. 把这个 channel 加入到客户端集合，也就是 entering <- ch；
// 5. 监听客户端往 conn 里写的数据，每扫描到一条就将这条消息发送到广播 channel 中；
// 6. 如果关闭了客户端，那么把队列离开写入 leaving 交给广播函数去删除这个客户端并关闭这个客户端；
// 7. 广播通知其他客户端该客户端已关闭；
// 8.  最后关闭这个客户端的连接 Conn.Close()。
func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户消息的通道
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "欢迎 " + who       //发送下次到客户端通道
	messages <- who + " 上线" //发送广播上线消息到消息通道
	entering <- ch          //上线客户端
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// 注意：忽略 input.Err() 中可能的错误
	leaving <- ch           //下线客户端
	messages <- who + " 下线" //发送广播下线消息到消息通道
	conn.Close()            //关闭连接
}

// 发送通道消息到客户端
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意：忽略网络层面的错误
	}
}
