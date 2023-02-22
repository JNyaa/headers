package headers_test

import (
	"fmt"

	"github.com/JNyaa/headers"
)

func Example() {
	fmt.Println(headers.AcceptCharset)
	// -> "Accept-Charset"

	fmt.Println(headers.IfNoneMatch)
	// -> "If-None-Match"

	fmt.Println(headers.Normalize("content-type"))
	// -> "Content-Type"
}
