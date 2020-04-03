package main

type ErrItem struct {
	Code int
	Msg string
}

var (
	Success = &ErrItem{Code: 0, Msg: "success"}

	UnsupportRestMethod = &ErrItem{Code: 301001, Msg: "unsupport rest method"}
	ParseMultipartFormErr = &ErrItem{Code: 301002, Msg: "parse multi part form err. "}
	UnmarshalJsonErr = &ErrItem{Code: 301003, Msg: "unmarshal json err. "}
)