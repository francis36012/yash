# YASH - Yet Another Shell

---
This is a toy shell written for the purpose of learning.
YASH is written in [Golang](https://www.golang.org).

Programs can be executed as background processes by adding an ampersand (&) to
the command.<br/>
Eg: `shouldbebackground &`

A history of executed programs is kept and can be viewed by executing `history`.
The stored history can be cleared by executing `history -c`.

There is no redirection or pipes so the following:<br/>
`echo "something" > something.txt`<br/>
`somecommand | grep "important"`<br/>
simply won't work.
