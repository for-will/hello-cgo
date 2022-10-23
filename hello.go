package main

// /Users/xuyz/qt/5.15/lib/QtGui.framework/QtGui /Users/xuyz/qt/5.15/lib/QtCore.framework/QtCore

/*
#cgo LDFLAGS: -L${SRCDIR} -lmyapp -Wl,-rpath,/Users/xuyz/qt/5.15/lib -F/Users/xuyz/qt/5.15/lib/ -framework QtWidgets  -framework QtCore -framework QtGui
#include <stdio.h>

void SayHello(int argc, char* argv[]){
	for(int i = 0; i < argc; i++){
		printf("%s\n", argv[i]);
	}
}

int startApp(int argc, char* argv[]);

*/
import "C"
import "unsafe"

func main() {
	//var hi = "Hello"
	var argv = []*C.char{C.CString("hello"), C.CString("World"), C.CString("hello")}
	//C.startApp(1, argv)
	C.SayHello(C.int(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))
	C.startApp(C.int(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))
	//C.startApp()
}
