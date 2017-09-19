package server

import "html/template"

var emptyPageText = []byte(`# Empty page
So this is an empty page`)

var show = template.Must(template.New("show").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>{{.Title}}</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<style>
		` + stylesheet + `
		</style>
	</head>
	<body>
		<form action="{{.Path}}" method="GET">
    	<button type="submit">Edit</button>
		</form>
		<article>{{.Text}}</article>
	</body>
</html>`))

var edit = template.Must(template.New("edit").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>{{.Title}}</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<style>
		` + stylesheet + `
		</style>
		<script>
			window.onload = function () {
				var contenteditableDiv = document.getElementById("textdiv");
				
				contenteditableDiv.addEventListener("input", function(){ 
      	 document.getElementById("text").value = contenteditableDiv.innerText;
      	});

				contenteditableDiv.focus();
			};
		</script>
	</head>
	<body>
		<form action="{{.Path}}" method="POST">
			<button type="submit">Save</button><br />
			<div id="textdiv" contentEditable autofocus>{{.Text}}</div>
			<textarea id="text" name="text">{{.Text}}</textarea>
		</form>
	</body>
</html>`))

const stylesheet = `
	a {
		color: #0366d6;
		text-decoration: none;
	}

	a:hover {
		text-decoration: underline;
	}

	body {
		font-family: -apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif;
		font-weight: 300;
		font-size: 16px;
		line-height: 1.5;
		margin: 0;
	}

	@media only screen and (min-device-width: 960px) {
		body {
			padding: 3rem 5rem;
		}

		button {
			position: fixed;
			bottom: 0px;
			right: 0px;
		}
	}

	button {
		margin: 1rem;
		background: #fafafa;
	  background-image: linear-gradient(to bottom, #fafafa, #dedede);
	  border-radius: 6px;
	  font-family: Helvetica;
	  color: #595959;
	  font-size: 14px;
	  padding: 6px 12px 6px 12px;
	  border: solid #b8b8b8 1px;
		cursor: pointer;
	}
	
	button:active {
		background: #f7f7f7;
	  background-image: linear-gradient(to bottom, #f7f7f7, #c4c2c4);
	  text-decoration: none;
	}

	button:focus, textarea:focus {
		outline:0;
	}

	h1 { font-size: 3.5em; } 
	h2 { font-size: 3em; }   
	h3 { font-size: 2.5em; } 
	h4 { font-size: 2em; }   
	h5 { font-size: 1.5em; } 
	h6 { font-size: 1em; }   

	h1, h2, h3, h4, h5, h6 {
    line-height: 1em; 
	}
 
	h1, h2, h3, h4, h5, h6, p {
    margin: 0 0 0.5rem 0; 
	}

	p:last-child, 
	ul:last-child, 
	ol:last-child, 
	dl:last-child, 
	blockquote:last-child, 
	pre:last-child, 
	table:last-child {
	    margin-bottom: 0;
	}

	hr {
 		border: 0; 
		height: 0; 
		border-top: 1px solid rgba(0, 0, 0, 0.1); 
		border-bottom: 1px solid rgba(255, 255, 255, 0.3);
	}

	p {
		margin: 0 0 9px;
	}

	strong {
		font-weight: 700;
	}

	blockquote {
		border-left: 0.4rem solid #dddfe1;
		padding-left: 1rem;
		margin-left: 0;
	}

	pre {
		font-size: 0.9em;
		border: 1px solid #dddfe1;
		border-radius: 3px;
		background: #f6f8fa;
		padding: 1rem;
		overflow: auto;
	}
	
	code {
		font-family: monaco,Consolas,Lucida Console,monospace;
	}

	textarea {
		display: none;
	}

	article {
		padding: 0 1rem;
	}

	#textdiv {
		background: #F6F8FA;
		border: 1px solid #dddfe1;
		border-radius: 6px;
		padding: 0.5rem;
		min-height: 10em;
		overflow: auto;
		white-space: pre;
		font-size: 0.8rem;
		font-family: monaco,Consolas,Lucida Console,monospace;
	}

	#textdiv:focus {
		outline: none;
	}
`
