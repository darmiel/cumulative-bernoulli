package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/calc/{mode}/{n}/{p}/{P}", func(writer http.ResponseWriter, request *http.Request) {

		// do not crash if stack overflows
		defer func() {
			if err := recover(); err != nil {
				_, _ = fmt.Fprint(writer, "Stack Overflow error.")
			}
		}()

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

    <link href="//cdn.jsdelivr.net/npm/@sweetalert2/theme-dark@4/dark.css" rel="stylesheet">
    <script src="//cdn.jsdelivr.net/npm/sweetalert2@10/dist/sweetalert2.min.js"></script>
</head>

<body>
<form id="form">
    <ul>
        <li><strong>n (int64): </strong><input type="text" id="n" name="n" value="1000" placeholder="n... (int64)"></li>
        <li><strong>p (float64): </strong><input type="text" id="p1" name="p" value="0.2" placeholder="p... (float64)"></li>
        <li><strong>P (float64): </strong><input type="text" id="p2" name="P" value="0.025" placeholder="P... (float64)"></li>
        
		<li><strong>Mode: </strong><select id="mode" name="mode">
            <option>le (<= | lower-equals)</option>
            <option>ge (>= | greater-equals)</option>
        </select></li>

        <li><input type="submit" value="Calculate! (Can take a while)"></li>
    </ul>
</form>

<script>
    const form = $("#form");

    const n = $("#n");
    const p = $("#p1");
    const P = $("#p2");

    const mode = $("#mode");

    form.on("submit", (event) => {
        event.preventDefault();

        const m = mode.val().substring(0, 2);
        const url = "/calc/" + m + "/" + n.val() + "/" + p.val() + "/" + P.val();

        // calculate
        Swal.queue([{
            title: 'Calculate',
            confirmButtonText: 'Calculate! Jetzt!',
            html: 'This can take a while.<br><span></span>',
            showLoaderOnConfirm: true,

            preConfirm: () => {
                // loading
                let a = 0;
                timerInterval = setInterval(() => {
                    a += 50;

                    const content = Swal.getContent()
                    if (content) {
                        const b = content.querySelector('span')
                        if (b) {
                            b.innerHTML = "<strong>" + a + "</strong>" + " ms";
                        }
                    }
                }, 50);

                return $.get(url, (data, status) => {
                	clearInterval(timerInterval);
					console.log("clear");
                    Swal.fire({
                        title: 'Result:',
                        html: '<pre style="text-align: left">' + data + '</pre>',
                        icon: 'success',
                        width: '50%'
                    });
                });
            }
        }]);
    });
</script>
</body>
</html>`)
	})

	router.HandleFunc("/calc/fac/{n}", func(writer http.ResponseWriter, request *http.Request) {

		// do not crash if stack overflows
		defer func() {
			if err := recover(); err != nil {
				_, _ = fmt.Fprint(writer, "Stack Overflow error.")
			}
		}()

		vars := mux.Vars(request)
		nStr := vars["n"]

		n, err := strconv.ParseFloat(nStr, 64)
		if err != nil {
			_, _ = fmt.Fprint(writer, err)
			return
		}

		if n > 500_000 {
			_, _ = fmt.Fprint(writer, "🧢 500.000")
			return
		}

		if n <= 0 {
			_, _ = fmt.Fprint(writer, "Holy fuck - what do you have in mind?!")
			return
		}

		start := time.Now()
		res := fac(big.NewFloat(n))

		milliseconds := time.Since(start).Milliseconds()
		text := []byte(res.Text('g', math.MaxInt32))

		_, _ = fmt.Fprintf(writer, "✅ %f! [%dms] [%d bytes] \n\n", n, milliseconds, len(text))

		start = time.Now()
		for i := 0; i < len(text); i++ {
			if i%3 == 0 {
				_, _ = fmt.Fprint(writer, " ")
			}
			_, _ = fmt.Fprint(writer, string(text[i]))
		}

		_, _ = fmt.Fprintf(writer, "\n\n💌 Text-Output took %dms", time.Since(start).Milliseconds())
	})

	if err := http.ListenAndServe(":1339", router); err != nil {
		log.Fatalln("Error serving:", err)
	}
}
