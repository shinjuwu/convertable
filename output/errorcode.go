//This file is creating automatic, don't edit!
package constant
const (
    //成功
    Success int = 0
    //帳號或密碼錯誤
    LoginAccOrPassNotMatch int = 1
    //無效的帳號
    LoginAccIDInvalid int = 2
    //重複登入
    LoginRepeat int = 3
    //發生內部錯誤
    LoginInnerError int = 4
    //帳號或密碼為空
    LoginAccOrPassIsNull int = 5
)
//ErrorCodeName is a map to get error string by error code
var ErrorCodeName = map[int]string{
    0: "Success",
    1: "LoginAccOrPassNotMatch",
    2: "LoginAccIDInvalid",
    3: "LoginRepeat",
    4: "LoginInnerError",
    5: "LoginAccOrPassIsNull",
}
//ErrorCodeValue is a map to get error code by error string
var ErrorCodeValue = map[string]int{
    "Success":0,
    "LoginAccOrPassNotMatch":1,
    "LoginAccIDInvalid":2,
    "LoginRepeat":3,
    "LoginInnerError":4,
    "LoginAccOrPassIsNull":5,
}
