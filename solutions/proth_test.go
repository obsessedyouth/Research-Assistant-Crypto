package proth_theorem

import (
  "github.com/stretchr/testify/assert"
  "testing"
  )
/* import "github.com/obsessedyouth/Research-Assistant-Crypto/solutions/proth_theorem"
TODO: replace proth_theorem.Proth
*/

func TestProth3(t *testing.T){
  assert.Equal(t, true, Proth(3))
}

func TestProth5(t *testing.T){
  assert.Equal(t, true, Proth(5))
}

func TestProth9(t *testing.T){
  assert.Equal(t, false, Proth(9))
}

func TestProth13(t *testing.T){
  assert.Equal(t, true, Proth(13))
}

func TestProth17(t *testing.T){
  assert.Equal(t, true, Proth(17))
}

// this is where I pause

func TestProth25(t *testing.T){
  assert.Equal(t, false, Proth(25))
}

func TestProth33(t *testing.T){
  assert.Equal(t, false, Proth(33))
}

func TestProth41(t *testing.T){
  assert.Equal(t, true, Proth(41))
}

func TestProth49(t *testing.T){
  assert.Equal(t, false, Proth(49))
}

func TestProth57(t *testing.T){
  assert.Equal(t, false, Proth(57))
}

func TestProth65(t *testing.T){
  assert.Equal(t, false, Proth(65))
}

func TestProth81(t *testing.T){
  assert.Equal(t, false, Proth(81))
}

func TestProth97(t *testing.T){
  assert.Equal(t, true, Proth(97))
}

func TestProth113(t *testing.T){
  assert.Equal(t, true, Proth(113))
}

func TestProth129(t *testing.T){
  assert.Equal(t, false, Proth(129))
}

func TestProth145(t *testing.T){
  assert.Equal(t, false, Proth(145))
}

func TestProth161(t *testing.T){
  assert.Equal(t, false, Proth(161))
}

func TestProth177(t *testing.T){
  assert.Equal(t, false, Proth(177))
}

func TestProth193(t *testing.T){
  assert.Equal(t, true, Proth(193))
}

func TestProth209(t *testing.T){
  assert.Equal(t, false, Proth(209))
}

func TestProth225(t *testing.T){
  assert.Equal(t, false, Proth(225))
}

func TestProth241(t *testing.T){
  assert.Equal(t, true, Proth(241))
}

func TestProth257(t *testing.T){
  assert.Equal(t, true, Proth(257))
}

func TestProth289(t *testing.T){
  assert.Equal(t, false, Proth(289))
}

func TestProth321(t *testing.T){
  assert.Equal(t, false, Proth(321))
}

func TestProth353(t *testing.T){
  assert.Equal(t, true, Proth(353))
}

func TestProth385(t *testing.T){
  assert.Equal(t, false, Proth(385))
}

func TestProth417(t *testing.T){
  assert.Equal(t, false, Proth(417))
}

func TestProth449(t *testing.T){
  assert.Equal(t, true, Proth(449))
}

func TestProth481(t *testing.T){
  assert.Equal(t, false, Proth(481))
}

func TestProth513(t *testing.T){
  assert.Equal(t, false, Proth(513))
}

func TestProth545(t *testing.T){
  assert.Equal(t, false, Proth(545))
}

func TestProth577(t *testing.T){
  assert.Equal(t, true, Proth(577))
}

func TestProth609(t *testing.T){
  assert.Equal(t, false, Proth(609))
}

func TestProth641(t *testing.T){
  assert.Equal(t, true, Proth(641))
}

func TestProth673(t *testing.T){
  assert.Equal(t, true, Proth(673))
}

func TestProth705(t *testing.T){
  assert.Equal(t, false, Proth(745))
}

func TestProth737(t *testing.T){
  assert.Equal(t, false, Proth(737))
}

func TestProth769(t *testing.T){
  assert.Equal(t, true, Proth(769))
}

func TestProth801(t *testing.T){
  assert.Equal(t, false, Proth(801))
}

/* To benchmark or not to benchmark?
That is the question
*/