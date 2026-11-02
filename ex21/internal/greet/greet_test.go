package greet

import(
	"testing"
)

func TestHello(t *testing.T ){
	res := Hello("myname")
	need := "Halo, myname"
	if res != need{
		t.Errorf("Hello() = %q , need %q", res, need)
	}
}
func TestSello(t *testing.T ){
	res := shout("myname")
	need := "MYNAME"
	if res != need{
		t.Errorf("Hello() = %q , need %q", res, need)
	}
}