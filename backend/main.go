package main

// Copyright 2019 Serge Bazanski
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY
// SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION
// OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN
// CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

var (
	flagBind            string
	flagUseForwardedFor bool
)

func cors(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
}

func noCache(w http.ResponseWriter) {
	w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, s-maxage=0")
	w.Header().Add("Cache-Control", " post-check=0, pre-check=0")
	w.Header().Add("Pragma", "no-cache")
}

func remoteIP(r *http.Request) net.IP {
	remote := strings.TrimSpace(r.RemoteAddr)
	if flagUseForwardedFor {
		remote = strings.TrimSpace(r.Header.Get("X-Forwarded-For"))
	}

	if remote == "" {
		return nil
	}

	if strings.HasPrefix(remote, "[") {
		// Looks like a [::1] v6 address
		inner := strings.Split(strings.Split(remote, "[")[1], "]")[0]
		return net.ParseIP(inner)
	}

	if strings.Count(remote, ":") == 1 {
		// Looks like a :4321 port suffix
		inner := strings.Split(remote, ":")[0]
		return net.ParseIP(inner)
	}

	return net.ParseIP(remote)
}

func main() {
	flag.StringVar(&flagBind, "bind", "0.0.0.0:8080", "Address at which to serve HTTP requests")
	flag.BoolVar(&flagUseForwardedFor, "use_forwarded_for", false, "Honor X-Forwarded-For headers to detect user IP")
	flag.Parse()
	http.HandleFunc("/empty", func(w http.ResponseWriter, r *http.Request) {
		cors(w)
		noCache(w)
		w.Header().Add("Connection", "keep-alive")
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/garbage", func(w http.ResponseWriter, r *http.Request) {
		cors(w)
		w.Header().Add("Content-Description", "File Tranfer")
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Disposition", "attachment; filename=random.dat")
		w.Header().Add("Content-Transfer-Encoding", "binary")
		noCache(w)
		w.WriteHeader(http.StatusOK)

		chunks := 4
		if val, err := strconv.Atoi(r.Header.Get("ckSize")); err == nil {
			if val > 0 && val <= 1024 {
				chunks = val
			}
		}

		chunk := make([]byte, 1024*1024)
		rand.Read(chunk)

		for i := 0; i < chunks; i += 1 {
			w.Write(chunk)
		}
	})
	http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		cors(w)
		noCache(w)

		res := struct {
			ProcessedString string `json:processedString`
			RawISPInfo      string `json:rawIspInfo`
		}{
			ProcessedString: "",
			RawISPInfo:      "",
		}

		ip := remoteIP(r)
		if ip != nil {
			res.ProcessedString = ip.String()
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(res)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/speedtest.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "speedtest.js")
	})
	http.HandleFunc("/speedtest_worker.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "speedtest_worker.js")
	})

	glog.Infof("Starting up...")
	err := http.ListenAndServe(flagBind, nil)
	if err != nil {
		glog.Exit(err)
	}
}
