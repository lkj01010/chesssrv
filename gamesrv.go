package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return `
		<!DOCTYPE HTML>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>文本输入框、密码输入框</title>
</head>
<body>
<form  method="post" action="save.php">
    账户:
	<input type="text" name="myName">
	<br>
	密码:

</form>
</body>
</html>
		`
	})
	m.Run()
}