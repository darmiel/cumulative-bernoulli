package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// res := findUpperBoundLe(12000, 0.2, 0.025)
	// if res != nil {
	// 	log.Println("Result:", *res)
	// } else {
	// 	log.Println("Nothing found :(")
	// }

	router := mux.NewRouter()
	router.HandleFunc("/calc/{mode}/{n}/{p}/{P}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		mode := vars["mode"]
		nStr := vars["n"]
		pStr := vars["p"]
		PStr := vars["P"]

		if mode != "le" && mode != "ge" {
			_, _ = fmt.Fprintln(writer, "Mode must be le (lower-equals) or ge (greater-equals)")
			return
		}

		n1, err := strconv.Atoi(nStr)
		if err != nil {
			_, _ = fmt.Fprintln(writer, err.Error())
			return
		}
		n := int64(n1)

		p, err := strconv.ParseFloat(pStr, 64)
		if err != nil {
			_, _ = fmt.Fprintln(writer, err.Error())
			return
		}

		P, err := strconv.ParseFloat(PStr, 64)
		if err != nil {
			_, _ = fmt.Fprintln(writer, err.Error())
			return
		}

		if n > 12000 {
			_, _ = fmt.Fprintln(writer, "n is too big")
			return
		} else if n <= 1 {
			_, _ = fmt.Fprintln(writer, "n is too small")
			return
		} else if p > 1 {
			_, _ = fmt.Fprintln(writer, "p is too big")
			return
		} else if p <= 0 {
			_, _ = fmt.Fprintln(writer, "p is too small")
			return
		} else if P > 1 {
			_, _ = fmt.Fprintln(writer, "P is too big")
			return
		} else if P <= 0 {
			_, _ = fmt.Fprintln(writer, "P is too small")
			return
		}

		var res *int64
		if mode == "ge" {
			res = findUpperBoundGe(writer, n, p, P)
		} else if mode == "le" {
			res = findUpperBoundLe(writer, n, p, P)
		}

		_, _ = fmt.Fprintln(writer)

		if res == nil {
			_, _ = fmt.Fprintln(writer, "😡 Result could not be determined")
		} else {
			_, _ = fmt.Fprintln(writer, "😊 Result:", *res)
		}
	})

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, `<html>

<head>
	<title>Ye boi calculator</title>
	<script src="https://code.jquery.com/jquery-3.5.1.min.js"></script>
</head>

<body>
	<form id="form">
		<input type="text" id="n" name="n" value="1000" placeholder="n... (int64)">
		<input type="text" id="p" name="p" value="0.2" placeholder="p... (float64)">
		<input type="text" id="P" name="P" value="0.025" placeholder="P... (float64)">
		<select id="mode" name="mode">
			<option>le (<= | lower-equals)</option>
			<option>ge (>= | greater-equals)</option>
		</select>
		<input type="submit" value="Calculate! (Can take a while)">
	</form>

	<script>
		const form = $("#form");

		const n = $("#n");
		const p = $("#p");
		const P = $("#P");

		const mode = $("#mode");

		form.on("submit", (event) => {
			event.preventDefault();
			const m = mode.val().substring(0, 2);

			$(location).attr("href", "/calc/" + m + "/" + n.val() + "/" + p.val() + "/" + P.val());
		});
	</script>
</body>

</html>`)
	})

	if err := http.ListenAndServe(":1339", router); err != nil {
		log.Fatalln("Error serving:", err)
	}
}
