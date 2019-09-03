package debugger

var front = []byte(`
<html>
	<head>
		<title>Gose debugger</title>
	</head>
	<body>
		<button onClick="fetch('/pause')">toggle pause</button>
		<button onClick="fetch('/step')">step</button>
	</body>
</html>
`)
