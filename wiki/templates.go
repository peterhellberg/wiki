package wiki

const showTpl = `<!DOCTYPE html>
<html>
	<head>
	</head>
	<body style="padding: 0.5em;">
		<form action="{{.Path}}" method="GET">
    	<button type="submit">Edit</button>
		</form>
		<div class="container" style="padding: 0em 1em; border: 1px solid #ccc;">{{.Text}}</div>
	</body>
</html>`

const editTpl = `<!DOCTYPE html>
<html>
	<head>
	</head>
	<body style="padding: 0.5em;">
		<div class="container">
			<form action="{{.Path}}" method="POST">
				<input type="submit" value="Update" /><br />
				<textarea name="text" style="width: 100%; height: 40em;">{{.Text}}</textarea>
			</form>
		</div>
	</body>
</html>`

const emptyPageString = `# Empty page
So this is an empty page
`
