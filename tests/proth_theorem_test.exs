ExUnit.start()

defmodule ProthTheoremTest do
  use ExUnit.Case

  test "greets the world" do
    assert ProthTheorem.proth() == "hello odd numbers"
  end
end
