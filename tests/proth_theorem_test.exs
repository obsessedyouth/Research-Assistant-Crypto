ExUnit.start()

defmodule ProthTheoremTest do
  use ExUnit.Case

  test "3 is a prime number" do
    assert ProthTheorem.proth(3) == true
  end
  test "5 is a prime number" do
    assert ProthTheorem.proth(5) == true
  end
  test "9 is a prime number" do
    assert ProthTheorem.proth(9) == false
  end
  test "13 is a prime number" do
    assert ProthTheorem.proth(13) == true
  end
  test "17 is a prime number" do
    assert ProthTheorem.proth(17) == true
  end
  test "25 is a prime number" do
    assert ProthTheorem.proth(25) == false
  end
  test "33 is a prime number" do
    assert ProthTheorem.proth(33) == false
  end
  test "41 is a prime number" do
    assert ProthTheorem.proth(41) == true
  end
  test "49 is a prime number" do
    assert ProthTheorem.proth(49) == false
  end
  test "57 is a prime number" do
    assert ProthTheorem.proth(57) == false
  end
  test "65 is a prime number" do
    assert ProthTheorem.proth(65) == false
  end
  test "81 is a prime number" do
    assert ProthTheorem.proth(81) == false
  end
  test "97 is a prime number" do
    assert ProthTheorem.proth(97) == true
  end
  test "113 is a prime number" do
    assert ProthTheorem.proth(113) == true
  end
  test "129 is a prime number" do
    assert ProthTheorem.proth(129) == false
  end
  test "145 is a prime number" do
    assert ProthTheorem.proth(145) == false
  end
  test "161 is a prime number" do
    assert ProthTheorem.proth(161) == false
  end
  test "177 is a prime number" do
    assert ProthTheorem.proth(177) == false
  end
  test "193 is a prime number" do
    assert ProthTheorem.proth(193) == true
  end
  test "209 is a prime number" do
    assert ProthTheorem.proth(209) == false
  end
  test "225 is a prime number" do
    assert ProthTheorem.proth(225) == false
  end
  test "241 is a prime number" do
    assert ProthTheorem.proth(241) == true
  end
  test "257 is a prime number" do
    assert ProthTheorem.proth(257) == true
  end
  test "289 is a prime number" do
    assert ProthTheorem.proth(289) == false
  end
  test "321 is a prime number" do
    assert ProthTheorem.proth(321) == false
  end
  test "353 is a prime number" do
    assert ProthTheorem.proth(353) == true
  end
  test "385 is a prime number" do
    assert ProthTheorem.proth(385) == false
  end
  test "417 is a prime number" do
    assert ProthTheorem.proth(417) == false
  end
  test "449 is a prime number" do
    assert ProthTheorem.proth(449) == true
  end
  test "481 is a prime number" do
    assert ProthTheorem.proth(481) == false
  end
  test "513 is a prime number" do
    assert ProthTheorem.proth(513) == false
  end
  test "545 is a prime number" do
    assert ProthTheorem.proth(545) == false
  end
  test "577 is a prime number" do
    assert ProthTheorem.proth(577) == true
  end
  test "609 is a prime number" do
    assert ProthTheorem.proth(609) == false
  end
  test "641 is a prime number" do
    assert ProthTheorem.proth(641) == true
  end
  test "673 is a prime number" do
    assert ProthTheorem.proth(673) == true
  end
  test "705 is a prime number" do
    assert ProthTheorem.proth(705) == false
  end
  test "737 is a prime number" do
    assert ProthTheorem.proth(737) == false
  end
  test "769 is a prime number" do
    assert ProthTheorem.proth(769) == true
  end
  test "801 is a prime number" do
    assert ProthTheorem.proth(801) == false
  end
end
