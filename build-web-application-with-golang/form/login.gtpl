<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    ユーザ名:<input type="text" name="username">
    パスワード:<input type="password" name="password">
    <input type="submit" value="ログイン">
    <input type="hidden" name="token" value="{{.}}">
</form>
</body>
</html>
