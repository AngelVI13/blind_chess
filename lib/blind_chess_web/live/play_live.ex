defmodule BlindChessWeb.PlayLive do
  use BlindChessWeb, :live_view

  def mount(_params, _session, socket) do
    {:ok, assign(socket, :name, "Angel")}
  end

  def render(assigns) do
    ~H"""
    <h1>Hello, <%= @name %></h1>
    """
  end
end
