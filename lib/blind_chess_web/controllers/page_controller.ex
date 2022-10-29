defmodule BlindChessWeb.PageController do
  use BlindChessWeb, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
