package classification

const (
	htmll = `<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <title>分类</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/xmdhs/searchqanda/style.css">
</head>

<body>
    <div class="container-lg px-3 my-5 markdown-body">
        <p>{{ .Text}}</p>
        <p><a href="{{ .Link}}">原帖</a></p>
        <hr>
        <p><a href="{{ .Yuanban}}">原版问答</a> <a href="{{ .Lianji}}">联机问答</a> <a href="{{ .Zoubian}}">周边问答</a> <a href="{{ .Mod}}">mod问答</a> 
        <a href="{{ .Be}}">基岩版问答</a> <a href="{{ .Water}}">灌水</a> <a href="{{ .Fuifu}}">回复</a> <a href="{{ .Tiaoguo}}">跳过</a></p>
    </div>
</body>
</html>`

	indexhtml = `<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <title>分类</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/xmdhs/searchqanda/style.css">
</head>

<body>
    <div class="container-lg px-3 my-5 markdown-body">
        <p><a href="/classification/add">分类</a></p>
        <p><a href="/classification/classification">评分</a></p>
    </div>
</body>
</html>`
	pinfen = `<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <title>分类</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/xmdhs/searchqanda/style.css">
</head>

<body>
    <div class="container-lg px-3 my-5 markdown-body">
        <h1>对文本分类</h1>
        <form action="/classification/classification" target="_blank"><input type="text" name="q"></form><br>
    </div>
</body>

</html>`
)
