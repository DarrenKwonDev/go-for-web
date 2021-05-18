# Go lectures  

## community
[golang bridge](https://forum.golangbridge.org/)  

## lecture source

https://github.com/GoesToEleven/GolangTraining  
https://github.com/GoesToEleven/golang-web-dev  
https://github.com/GoesToEleven/go-programming  


## resources

### official docs & spec  

[golang official doc](https://golang.org/doc/) => [effect go](https://gosudaweb.gitbooks.io/effective-go-in-korean/content/)는 한국어 번역이 있음  
[golang language spec](https://golang.org/ref/spec)  
[golang official doc (일부 한글)](https://github.com/golang-kr/golang-doc/wiki)  

### golang.org vs godoc.org  

golang.org : language, stdlib  
godoc.org : stdlib, third-party packages  
쉽게 말해 패키지 찾으려면 godoc, 일반 문서를 보려면 golang  

## book  

가장 빨리 만나는 Go  http://pyrasis.com/private/2015/06/01/publish-go-for-the-really-impatient-book  
Go 언어를 활용한 마이크로서비스  
Go 풀스택 웹 개발 006.76 G649g 2020    
Go 마스터하기  005.133 T786m 절판됨. 도서관에서 빌려볼 것  
Go in action   
[DI in go(한국어 책)](http://acornpub.co.kr/book/dependency-injection-go)

## std lib and framewok

[std lib](https://golang.org/pkg/)  
[gin](https://github.com/gin-gonic/gin)   
[gin-gonic](https://gin-gonic.com/)   
[echo](https://github.com/labstack/echo)   

## useful blog post  

[유연하고 테스트 가능한 Go 코드 작성하기](https://medium.com/daangn/how-to-write-a-testable-golang-code-4c0e67612bb8)  
[errgroup으로 goroutine 10배 잘 활용하기](https://devjin-blog.com/golang-errgroup-goroutine/)  
[rest-api-with-golang-using-gin-and-gorm](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/)  
[Go와 함께하는 전화망 서비스 구축 1편](https://d2.naver.com/helloworld/5827706)   
[Go와 함께하는 전화망 서비스 구축 2편](https://d2.naver.com/helloworld/0814313)   
[Go 전반적인 이야기와 웹 프레임워크 선택](https://kimtaekju-study.tistory.com/203)  
[Error는 검사만 하지말고, 우아하게 처리하세요.](http://cloudrain21.com/golang-graceful-error-handling#errors-just-values)  

[Go로 살피는 디자인 패턴](https://pi-sin0.medium.com/head-first-design-patterns-using-go-introduction-d9f05a4c243a) 
[medium의 rungo 모든 포스트들](https://medium.com/rungo)  

[Go로 구현하는 블록체인 1~7](https://mingrammer.com/categories/blockchain/)  
[Go 관련한 여러 글들](https://mingrammer.com/categories/go/)    


## good website  

[golangKorea](https://github.com/golangkorea)  
[effective go(이펙티브 고의 한국어 번역)](https://gosudaweb.gitbooks.io/effective-go-in-korean/content/)  
[gobyexample](https://gobyexample.com/) => [번역](https://joinc.co.kr/w/GoLang/example/)  
[awesome-go](https://awesome-go.com/#web-frameworks)  
[go in vscode](https://code.visualstudio.com/docs/languages/go)  

## go lectures  

[한 눈에 끝내는 고랭 기초](https://edu.goorm.io/lecture/2010/%25ED%2595%259C-%25EB%2588%2588%25EC%2597%2590-%25EB%2581%259D%25EB%2582%25B4%25EB%258A%2594-%25EA%25B3%25A0%25EB%259E%25AD-%25EA%25B8%25B0%25EC%25B4%2588)  
[An Introducing to programming in go 한국어 번역](http://codingnuri.com/golang-book/)  
[Go 101](https://go101.org/article/101.html)  
[예제로 공부하는 Go 언어(유명한 gobyexample의 번역)](https://joinc.co.kr/w/GoLang/example/)  
[pyrasis 님의 강의](http://pyrasis.com/private/2015/06/01/publish-go-for-the-really-impatient-book)  
[medium post](medium.com/qvault/learn-go-fast-best-courses-and-resources-3a42e70476c3)  
[udemy lecture](www.udemy.com/course/go-programming-language/)  
[ardan labs](https://www.ardanlabs.com/ultimate-go/)  
[마소에서 제공하는 go 강의](https://docs.microsoft.com/ko-kr/learn/paths/go-first-steps/)  

## usefull source  

https://github.com/gookit/goutil  


## go blockchain  

https://ethereum.org/en/developers/docs/programming-languages/golang/  
https://pkg.go.dev/github.com/ethereum/go-ethereum  
https://github.com/ethereum/go-ethereum   


## info  

- Everything in Go is **pass by value** 따라서 직접 값을 복사해주고 싶지 않다면 포인터를 넘겨야 한다.
아래 코드를 보면 File을 반환하는 것이 아니라 File의 메모리 주소값을 반환하고 있다. File을 직접 넘긴다면 File을 또 생성해야 하기 때문이다. 메모리 절약을 위해 포인터를 쓰는 것이 좋다
```go
func NewFile(fd uintptr, name string) *File
``` 
- Go는 OOP 기반 언어지만 class와 contructor가 없어 구조체(struct)와 struct factory를 이용해야 한다.

- array와 stuct는 value다. 따라서 함수의 인자로 사용할 때는 값복사가 되지 않게 포인터를 이용하는 것이 좋다. 특히 struct factory는 포인터를 반환하여 객체 복사가 이루어지지 않는 것이 좋다  

- 반면 slice와 map은 reference type이다. 복사하면 참조 복사를 하니까 그냥 사용해도 된다.  
  
- effective go에서는 array 대신 slice를 이용을 권장하고 있다.  

- map, slice는 reference type이다. 
초기화하지 않으면 기본값으로 자동 초기화되지 않고 nil이 된다.  

- A string is a sequence of bytes. 마셜링할 때 더욱 느끼게 된다. string을 []byte로 변환하는 일이 많다.

- 함수 리터럴과 함수 타입 변수(값으로서의 함수)는 다른겁니다.
```go
func main() {
	i := 0

	// 함수 타입 변수. 함수를 변수에 넣었을 뿐임. 
	// 값복사를 하고, 함수 스코프 벗어나면 가비지 컬렉팅됨
	innerAddTenFunc := func(a int) {
		a += 10
	}

	// 함수 리터럴. 참조로 변수 가져와 원본을 변화시킴.
	innerAddTenLiteral := func() {
		i += 10
	}

	innerAddTenFunc(i)
	fmt.Println(i) // 0. 원본 유지.
	
	innerAddTenLiteral()
	fmt.Println(i) // 10. 변화됨
}
```

- 슬라이스의 cap을 초과하게 되면 새로운 배열을 만들기 때문에 의도한대로 동작하지 않을 수 있다. 슬라이스는 항상 cap

```go
func main() {
	slice1 := []int{1, 2, 3} // len 3, cap 3
	slice2 := append(slice1, 4, 5) // 4, 5를 추가할 cap이 부족. 기존 cap의 2배인 새로운 배열을 만듦.

	fmt.Println(slice1, len(slice1), cap(slice1)) // [1 2 3] 3 3
	fmt.Println(slice2, len(slice2), cap(slice2)) // [1 2 3 4 5] 5 6
	
	slice1[1] = 100;
	fmt.Println(slice1) // [1 100 3] 예상대로 바뀜
	fmt.Println(slice2) // [1 2 3 4 5] 안 바뀜. cap이 부족해서 새 배열을 만들었기 때문.
}
```
