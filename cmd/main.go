package main

import (
	"github.com/cch123/elasticsql"
	"github.com/dollarkillerx/erguotou"

	"log"
	"strings"
)

type Payload struct {
	SQL string `json:"sql"`
	Dsl string `json:"dsl"`
}

func main() {
	app := erguotou.New()

	app.Get("/", func(ctx *erguotou.Context) {
		ctx.Ctx.SetContentType("Content-Type: text/html; charset=utf-8")
		ctx.Write(200, []byte(html))
	})

	app.Post("/api", func(ctx *erguotou.Context) {
		p := Payload{}
		err := ctx.BindJson(&p)
		if err != nil {
			ctx.Json(400, erguotou.H{"msg": err.Error()})
			return
		}

		convert, _, err := elasticsql.Convert(strings.TrimSpace(p.SQL))
		if err != nil {
			ctx.Json(400, erguotou.H{"msg": err.Error()})
			return
		}

		ctx.Json(200, erguotou.H{"dsl": convert})
	})

	err := app.Run(erguotou.SetHost("0.0.0.0:8086"))
	if err != nil {
		log.Fatalln(err)
	}
}

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>EsQuery</title>
    <!-- CSS only -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
    <!-- JavaScript Bundle with Popper -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM"
            crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.min.js"
            integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>

</head>

<body>
<div class="container">
    <!-- Content here -->
    <div>
        <div class="page-header">
            <h1>This tool converts sql to elasticsearch dsl  <small>github.com/dollarkillerx/es_query</small></h1>
        </div>
        <hr>
        <br>
        <div class="mb-3">
            <label for="sql" class="form-label">SQL: </label>
            <input type="email" class="form-control" id="sql" placeholder="select * from ...">
        </div>
        <div class="mb-3">
            <label for="dsl" class="form-label">Elasticsearch DSL: </label>
            <textarea class="form-control" id="dsl" rows="3"></textarea>
        </div>
        <button type="button" class="btn btn-primary" onclick="Translate()">Translate</button>
    </div>
</div>

<script>
    function Translate() {
        let sql = $('#sql').val();
        if (sql == "") {
            alert("sql is null")
        }

        $.ajax({
            type: "POST",
            url: "/api",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify({"sql": sql}),
            dataType: "json",
            success: function (message) {
                $('#dsl').val(message.dsl);
            },
            error: function (message) {
                $('#dsl').val("error: " + message.responseJSON.msg);
            }
        });
    }
</script>

</body>
</html>
`
