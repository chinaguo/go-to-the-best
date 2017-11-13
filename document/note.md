## golang 需要注意的问题
### [英文网址](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#mline_lit_comma)

应用于golang 1.5以下，经过测试1.8以下也要注意。从1.6开始用了那么长时间的go，感觉应该进行下总结，正好有个时间点，为自己做个笔记。
人生起起伏伏如山峦一样，又如驼峰的编码，有时高有时低...
### 汇总
+ 左边大括号不能另起一行

```js
  func main()
  {
  
  }
```
编译错误，应该为：

```js
  func main() {
  
  }
```

+ 变量必须要使用

```js
  //not an error, is global variable
  var gOK int 
   
  //is ok, have no error if the argument unused 
  func check(va string) {
      
  }
   
  func main() {
   	 var a int    //error, unused
  }
```

+ 变量作用域

```js
  func main() {
    x := 1					//prints 1
    {
       x := 2
       fmt.Println(x)   //prints 2
    }
    fmt.Println(x)      //prints 1
  }
```

+ slices and map 不能为nil

```js
  func main() {
	 var s []int
	 // panic error, s is nil
	 s = append(s, 1)
  }
```

+ map刚建立的时候，不能使用cap函数

```js
  func main() {
	 m := make(map[string]int, 99)
	 cap(m)    // is error
  }
```
+ 空字符串不是nil, string而是""，不能直接赋值nil

```js
  func main() {
	 var x string = nil //error, x is ""
	 if x == nil {
		x = "default"
	 }
  }
```
+ map中keys不存在的情况

```js
  func main() {
	 x := map[string]string{"one":"a", "two":"", "three":"c"}
	    
	 //如果key 不存在，需要利用ok判断，一定要去ok去判断
	 if _,ok := x["two"]; !ok {
	   fmt.Println("no entry")
	 }
  }
```

+ 更改字符串需要注意的地方

```js
  func main() {
	 //更改字符要是byte数组
	 x := "text"
	 xByte := []byte(x)
	 xByte[0] = 'T'
		
	 fmt.Println(string(xByte))
  }
```
+ interface看起来像指针，但是它不是指针，和指针是有差别的

```js
  package main
	
  import "fmt"
  func main() {
	 var data *byte
	 var in interface{}   //is interface{}
	 fmt.Println(data, data == nil) //is true
	 fmt.Println(in, in == nil)     //is true
		
	 in = data
	 fmt.Println(in, in == nil) //nil, false, data is nil, but 'in' is not nil ,大部分人对这里很奇怪，但是事实就是这样，打印出nil是因为fmt.Println的实现检测到指针指向是nil，但是in本身并不是nil
  }
```

+ 更新map的元素的时候需要注意的地方
map元素如果是个结构体，不能更新结构体的私有的成员变量。
slice的里面的元素是可以的，不一样哈

```js
  package main
	
  type data struct {
	 name string
  }
	
  func main() {
	 mt := map[string]data{"x":{"very good"}}
	 mt["x"].name = "no"    //这里是错误的，此变量无法访问
  }
```

+ 在进行http请求的时候，http.Get()或者Post()返回值，要判断error，如果不为nil不能使用defer resp.Body.Close()

```js
  package main

  import(
    "fmt"
    "net/http"
    "io/ioutil"
  )
  
  func main() {
    //简单的一个Get
    resp, err := http.Get("https://baidu.com")
  	 if err != nil {
  	    fmt.Println(err)
  	    return
  	 }
  	 
  	 defer resp.Body.Close()    //先判断error，此处才可以关闭
  	 body, err := ioutil.ReadAll(resp.Body)   //读取body内容
  	 if err != nil {
  	 	fmt.Println(err)
  	 	return
  	 }
  	 
  	 fmt.Println(string(body))
  }
```
+ json.Unmarshal的时候，如果result为map[string]interface{},返回的数字型应该用float64进行强转

```json
  data := []byte(`{"hello": 10}`)
  
  var result map[string]interface{}
  if err := json.Unmarshal(data, &result); err != nil {
    fmt.Println("error:", err)
    return
  }
  
  var hello = result["hello"].(float64)
  fmt.Println("status value:", status)
```
+ channel的使用

```json
  channel包括有缓冲和无缓冲，都是要使用make建立实例。
  channel使用完成后必须要进行关闭，否则将会导致内存的泄露。要使用close关闭。
```
+ channel的关闭(gracefully)

```js
  有几种规则其实应该注意一下：
  1.其实没有特别普遍的方法来关闭channel，一般是由发送者来关闭channel。
  2.不能关闭已经关闭的channel，会出现panic。
  3.发送数据给已经关闭的channel，会出现panic。
```
+ 关于sync.WaitGroup使用

```js
  package main
	
  import(
	 "time"
	 "fmt"
	 "sync" 
  )
  
  func main() {
    wg := sync.WaitGroup{}
    wg.Add(1)
    
    go func() {
       defer wg.Done()
       fmt.Println("ok")
    }()
    
    wg.Wait()
  }
```
+ golang中关于第三方包的使用，推荐go vendor。1.5之后就可以使用，1.6正式引入。

```js
  为什么要使用vendor呢？
  1.在远程服务器进行编译的时候不用再go get 外部依赖包了，如果在国外的服务器上下载是非常慢的，有时候可能会出现问题。
  2.在编译服务器上go get 的包有可能是最新的包，和你机器上的不一致就有可能有问题。
  3.有了vendor编译的时候会优先vendor目录下寻找依赖的包。
  4.看看docker吧，你会感觉选择没有错。
```
+ golang第三方包版本控制，golang没有集中的包管理的地方，和node的设计差别比较大。

```js
  使用gopkg.in的应用，这个是个开源的软件。
  比如：gopkg.in/mgo.v2就是使用mgo中v2的版本
```