package manipulador

import "html/template"

var Modelos = template.Must(template.ParseFiles("html/ola.html"))
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))