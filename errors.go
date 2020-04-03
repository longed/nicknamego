package main

type Item struct {
	Code int
	Msg string
}

var (
	Success = &Item{Code: 0, Msg: "success"}

	UnsupportRestMethod = &Item{Code: 301001, Msg: "unsupport rest method"}
)