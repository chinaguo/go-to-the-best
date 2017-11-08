## golang 需要注意的问题
### [英文网址](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#mline_lit_comma)

应用于golang 1.5以下，经过测试1.8以下也要注意。

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
