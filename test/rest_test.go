package test

import (
	"fmt"
	"testing"
)

// import (
// 	"testing"

// 	"github.com/raafly/realtime-app/db"
// 	restserver "github.com/raafly/realtime-app/rest-server"
// )

// var conn = db.NewDB()

// func TestRegister(t *testing.T) {
// 	restserver.
// }

func TestOTP(t *testing.T) {
	otp1, otp2 := 306136, 3061361
	if otp1 != otp2 {
		fmt.Println("tidak sama")
	}

	fmt.Println("sama")
}