package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func exercise001() {
	person1 := MakePerson("James", "Bond", 30)
	person2 := MakePersonPointer("James", "Bond", 30)
	fmt.Println(person1, *person2)
}

/*
# iv. について

##  出力結果

```
./exercise001.go:11:6: can inline MakePerson
./exercise001.go:19:6: can inline MakePersonPointer
./main.go:3:6: can inline main
./exercise001.go:28:23: inlining call to MakePerson
./exercise001.go:29:30: inlining call to MakePersonPointer
./exercise001.go:30:13: inlining call to fmt.Println
./exercise001.go:11:17: leaking param: firstName to result ~r0 level=0
./exercise001.go:11:28: leaking param: lastName to result ~r0 level=0
./exercise001.go:19:24: leaking param: firstName
./exercise001.go:19:35: leaking param: lastName
./exercise001.go:20:9: &Person{...} escapes to heap
./exercise001.go:29:30: &Person{...} does not escape
./exercise001.go:30:13: ... argument does not escape
./exercise001.go:30:14: person1 escapes to heap
./exercise001.go:30:23: *person2 escapes to heap
```


## 考察

1. `./exercise001.go:20:9: &Person{...} escapes to heap` について

無名のPerson型ローカル変数が作られ，その値が戻り値として戻ることになるので，このローカル変数はヒープに割り当てられる。

2. `./exercise001.go:30:14: person1 escapes to heap`, `./exercise001.go:30:23: *person2 escapes to heap` について

`fmt.Println` に渡される引数は ...any 型であり，Goではインターフェイス型の引数はヒープに割り当てられることになっている。
*/
