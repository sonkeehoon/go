my study link : https://github.com/sonkeehoon/go
run : go run {filename}.go
build : go build {filename}.go


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




