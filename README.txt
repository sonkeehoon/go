my study link : https://github.com/sonkeehoon/go
reference : https://blog.naver.com/hojun0313/222500289501
run : go run {filename}.go
build : go build {filename}.go

- go 환경설정
go 를 설치한 디렉터리에 bin, pkg, src 3개의 폴더가 존재해야한다
코딩은 src 폴더에서 하면 된다

- print와 fmt 패키지
main 패키지에서 "fmt"패지키를 불러올때
import "fmt"
c언어의 <stdio.h>랑 비슷한 개념인듯 하다
프린트 할때는 fmt.Print(" "), fmt.Println(" "), fmt.Printf(" ") 이런식으로 표현한다 (fmt 일때는 대문자 P임에 주의)
그런데 go에서는 fmt 없이 그냥 출력해도 된다
c에서 사용하던 print(출력), println(출력하고 한줄 개행), printf("포매팅이 필요할때,%d%s 이런거")
세미콜론은 필요없다(파이썬과 유사)
변수 선언은 조금 특이하다
var num int = 1
num int : 변수 이름 다음에 타입이 온다는것이 타 언어들과는 다르다

- 상수(const)
상수는 선언과 동시에 초기화를 해야함
상수는 선언만 하고 사용하지 않아도 에러는 발생하지 않음
상수는 const키워드를 사용하므로 := 는 사용할수 없음

상수에 괄호를 이용해 여러개의 값을 초기화 할수 있고, 다른 자료형도 초기화가 가능하다
예를들면
const (
    val1 = 1
    val2 = "hello"
)
이런 표현도 가능

괄호 사용할때 주의사항
괄호 속 첫번째 상수는 무조건 선언되어야함
선언되지 않은 값은 직전 상수의 값과 같다
상수 값을 iota로 초기화하면 그 다음 상수들은 index값으로 자동 저장된다

- Go 연산자 우선순위

1순위   (), [], ->, ++, --

2순위   +, -, ~, (type), *, &, sizeof

3순위   *, /, %

4순위   +, -

5순위   <<, >>

6순위~10순위    <, <=, >, =>, ==, !=, &, ^, |

11순위~15순위   &&, ||, ?:, =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |=, ,

- 자료형
자료형은 코드보다 이론적인 내용이 많아서 아래 내용을 참조
https://blog.naver.com/hojun0313/222516858858

- 문자열
Raw String Literal은 ``으로 문자열을 표현한다
\n등을 넣어도 문자 그대로 출력시킨다
Interpreted String Literal은 ""으로 문자열을 표현한다
Raw와 다르게 \n같은 문자를 탈출문자로 인식한다




