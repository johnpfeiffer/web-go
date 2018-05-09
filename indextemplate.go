package main

import "html/template"

// NoData is an empty struct as I do not pass anything into the template
type NoData struct{}

// GetIndexTemplate returns the index.html template https://golang.org/pkg/html/template/#Template
func GetIndexTemplate() *template.Template {
	var indexTemplate = template.Must(template.New("index").Parse(`<html>
<head>
<title>Example Go Web App</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<style type="text/css">
body{    
  font-family: "Georgia";
  font-size: 1.9em;
}
</style>
</head>
<body>
hi , try <a href="/hexcolors">hexcolors</a>
</body>
</html>
`))
	return indexTemplate
}
