package proth_theorem

import (
  "github.com/obsessedyouth/Research-Assistant-Crypto/solutions/proth_theorem"
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestProth3(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(3))
}

func TestProth5(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(5))
}

func TestProth9(t *testing.T){
  assert.Equal(t, false,proth_theorem. Proth(9))
}

func TestProth13(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(13))
}

func TestProth17(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(17))
}

func TestProth25(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(25))
}

func TestProth33(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(33))
}

func TestProth41(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(41))
}

func TestProth49(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(49))
}

func TestProth57(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(57))
}

func TestProth65(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(65))
}

func TestProth81(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(81))
}

func TestProth97(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(97))
}

func TestProth113(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(113))
}

func TestProth129(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(129))
}

func TestProth145(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(145))
}

func TestProth161(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(161))
}

func TestProth177(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(177))
}

func TestProth193(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(193))
}

func TestProth209(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(209))
}

func TestProth225(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(225))
}

func TestProth241(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(241))
}

func TestProth257(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(257))
}

func TestProth289(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(289))
}

func TestProth321(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(321))
}

func TestProth353(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(353))
}

func TestProth385(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(385))
}

func TestProth417(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(417))
}

func TestProth449(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(449))
}

func TestProth481(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(481))
}

func TestProth513(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(513))
}

func TestProth545(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(545))
}

func TestProth577(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(577))
}

func TestProth609(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(609))
}

func TestProth641(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(641))
}

func TestProth673(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(673))
}

func TestProth705(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(745))
}

func TestProth737(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(737))
}

func TestProth769(t *testing.T){
  assert.Equal(t, true, proth_theorem.Proth(769))
}

func TestProth801(t *testing.T){
  assert.Equal(t, false, proth_theorem.Proth(801))
}
