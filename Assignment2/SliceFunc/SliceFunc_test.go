package SliceFunc
import (
"fmt"
//"errors"
"testing"
)

func TestSLiceFunc(t *testing.T) {

  val := modifyvalue(-1)


  if val != true {

    t.Fatalf("Found error")

  }
fmt.Println("Success")
}
