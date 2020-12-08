package Utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"../../Common/Message"
)

//这里将Utils工具包的方法关联到结构体中
//Transfer结构
type Transfer struct {
	 Conn net.Conn
	 Buff [8096]byte
}


//读取包的公共函数
func (this *Transfer)ReadPkg()(message Message.Message, err error){
	//先读取前4个因为发送的时候，前4字节是消息长度字段
	//buff:=make([]byte,8096)
	n,err :=this.Conn.Read(this.Buff[0:4])
	if(err!=nil){
		fmt.Printf("conn.Write(bytes)1111 n=%v err=%v\n",n,err)
		err = errors.New("Util：conn.Read(buff[:4])")
		return Message.Message{}, err
	}

	//读取的buff[:4]转化为unit32类型，用来比较传送的pkg是否由丢包 （发送长度与实际接收长度比较）
	var pkgLen  = binary.BigEndian.Uint32(this.Buff[0:4])
	fmt.Println("pkgLen:",pkgLen)
	//根据pkgLen读取消息内容 (从conn中读取到buff中)
	n,errr :=this.Conn.Read(this.Buff[:pkgLen])
	if n!=int(pkgLen) ||errr!=nil {
		errr = errors.New("Util：conn.Read(buff[:pkgLen])")
		return Message.Message{}, errr
	}

	//将buff中内容反序列化   buff--》message
	err =json.Unmarshal(this.Buff[:pkgLen],&message)
	if(err!=nil){
		err = errors.New("Util：json.Unmarshal(buff[:pkgLen],&message)")
		return Message.Message{}, err
	}

	//读包成功，返回反序列化的结果：message
	return message,nil
}

//写包的公共函数
func (this *Transfer)WtitePkg(data []byte)(err error){
	//1.1发送一个消息的长度给对方
	//conn.Write的参数是切片，我们发的长度是int，需要int转换切片
	var dataMessageLen =uint32(len(data))
	//var buff [4]byte
	binary.BigEndian.PutUint32(this.Buff[0:4],dataMessageLen)//int转换切片

	//2.1 发送dataMessage的长度发送给服务器
	n,err:=this.Conn.Write(this.Buff[:4])
	if(n!=4||err!=nil){
		fmt.Printf("Util：conn.Write(buff[:4]) n=%v err=%v\n",n,err)
		return err;
	}

	fmt.Println("发送消息信息的长度完成,长度为",dataMessageLen)

	//3.发送消息本身
	_,errr:=this.Conn.Write(data)
	if(errr!=nil){
		fmt.Println("Util：conn.Write(data) err")
		return errr;
	}
	fmt.Println("发送dataMessage信息完成")
	return nil
}