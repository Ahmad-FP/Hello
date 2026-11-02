package greet

import(
	"testing"
)

func TestHello(t *testing.T ){
	res := Hello("Budi")
	need := "Halo, Budi"
	if res != need{
		t.Errorf("Hello() = %q , need %q", res, need)
	}
}
func TestSello(t *testing.T ){
	res := shout("Budi")
	need := "BUDI"
	if res != need{
		t.Errorf("Hello() = %q , need %q", res, need)
	}
}