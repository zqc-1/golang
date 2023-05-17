package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/**
  os.O_RDONLY: 只读模式(read-only)
  os.O_WRONLY: 只写模式(write-only)
  os.O_RDWR : 读写模式(read-write)
  os.O_APPEND: 追加模式(append)
  os.O_CREATE: ⽂件不存在就创建(create a new file if none exists.)
  os.O_TRUNC: 打开并清空⽂件（必须有写权限）
  os.O_EXCL: 如与 O_CREATE ⼀起⽤，构成⼀个新建⽂件的功能，它要求⽂件必须不存在(used with O_CREATE, file must not exist)
  os.O_SYNC：同步⽅式打开，即不使⽤缓存，直接写⼊硬盘

  0777：-rwxrwxrwx，创建了一个普通文件，所有人拥有所有的读、写、执行权限
  0666：-rw-rw-rw-，创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
  0644：-rw-r–r–，创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，没有执行权限
*/

func main() {
	file, err := os.OpenFile("满江红2", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	fmt.Println("file:", file)
	//file.WriteString("怒发冲冠，凭栏处、潇潇雨歇。\n")

	// 按字节或字符串写操作
	//file.Write([]byte("怒发冲冠，凭栏处、潇潇雨歇。\n"))
	//file.WriteString("抬望眼、仰天长啸，壮怀激烈。\n")

	// 缓冲写
	//writer := bufio.NewWriter(file)
	//writer.WriteString("三十功名尘与土，八千里路云和月。\n")
	//writer.Flush()

	// 写整个文件
	str := "怒发冲冠，凭栏处、潇潇雨歇。\n抬望眼、仰天长啸，壮怀激烈。\n三十功名尘与土，八千里路云和月。\n莫等闲、白了少年头，空悲切。"
	ioutil.WriteFile("满江红3", []byte(str), 0666)
}
