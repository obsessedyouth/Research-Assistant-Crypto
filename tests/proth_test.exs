ExUnit.start()

defmodule ProthTheoremTest do
  use ExUnit.Case

  test "3 is a prime number" do
    assert ProthTheorem.proth(3) == True
  end
  test "5 is a prime number" do
    assert ProthTheorem.proth(5) == True
  end
  test "9 is a prime number" do
    assert ProthTheorem.proth(9) == False
  end
  test "13 is a prime number" do
    assert ProthTheorem.proth(13) == True
  end
  test "17 is a prime number" do
    assert ProthTheorem.proth(17) == True
  end
  test "25 is a prime number" do
    assert ProthTheorem.proth(25) == False
  end
  test "33 is a prime number" do
    assert ProthTheorem.proth(33) == False
  end
  test "41 is a prime number" do
    assert ProthTheorem.proth(41) == True
  end
  test "49 is a prime number" do
    assert ProthTheorem.proth(49) == False
  end
  test "57 is a prime number" do
    assert ProthTheorem.proth(57) == False
  end
  test "65 is a prime number" do
    assert ProthTheorem.proth(65) == False
  end
  test "81 is a prime number" do
    assert ProthTheorem.proth(81) == False
  end
  test "97 is a prime number" do
    assert ProthTheorem.proth(97) == True
  end
  test "113 is a prime number" do
    assert ProthTheorem.proth(113) == True
  end
  test "129 is a prime number" do
    assert ProthTheorem.proth(129) == False
  end
  test "145 is a prime number" do
    assert ProthTheorem.proth(145) == False
  end
  test "161 is a prime number" do
    assert ProthTheorem.proth(161) == False
  end
  test "177 is a prime number" do
    assert ProthTheorem.proth(177) == False
  end
  test "193 is a prime number" do
    assert ProthTheorem.proth(193) == True
  end
  test "209 is a prime number" do
    assert ProthTheorem.proth(209) == False
  end
  test "225 is a prime number" do
    assert ProthTheorem.proth(225) == False
  end
  test "241 is a prime number" do
    assert ProthTheorem.proth(241) == True
  end
  test "257 is a prime number" do
    assert ProthTheorem.proth(257) == True
  end
  test "289 is a prime number" do
    assert ProthTheorem.proth(289) == False
  end
  test "321 is a prime number" do
    assert ProthTheorem.proth(321) == False
  end
  test "353 is a prime number" do
    assert ProthTheorem.proth(353) == True
  end
  test "385 is a prime number" do
    assert ProthTheorem.proth(385) == False
  end
  test "417 is a prime number" do
    assert ProthTheorem.proth(417) == False
  end
  test "449 is a prime number" do
    assert ProthTheorem.proth(449) == True
  end
  test "481 is a prime number" do
    assert ProthTheorem.proth(481) == False
  end
  test "513 is a prime number" do
    assert ProthTheorem.proth(513) == False
  end
  test "545 is a prime number" do
    assert ProthTheorem.proth(545) == False
  end
  test "577 is a prime number" do
    assert ProthTheorem.proth(577) == True
  end
  test "609 is a prime number" do
    assert ProthTheorem.proth(609) == False
  end
  test "641 is a prime number" do
    assert ProthTheorem.proth(641) == True
  end
  test "673 is a prime number" do
    assert ProthTheorem.proth(673) == True
  end
  test "705 is a prime number" do
    assert ProthTheorem.proth(705) == False
  end
  test "737 is a prime number" do
    assert ProthTheorem.proth(737) == False
  end
  test "769 is a prime number" do
    assert ProthTheorem.proth(769) == True
  end
  test "801 is a prime number" do
    assert ProthTheorem.proth(801) == False
  end
end
