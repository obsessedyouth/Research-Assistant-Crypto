package proth_theorem

import "testing"
/* import "github.com/obsessedyouth/Research-Assistant-Crypto/solutions/proth_theorem"
*/


func TestProth(t *testing.T){
  expected := "hello odd numbers"
  actual := Proth()
  if actual != expected {
    t.Error("That's weird :( ")
  }
}
