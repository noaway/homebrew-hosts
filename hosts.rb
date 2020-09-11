# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/0.1.3/hosts-mac64-0.1.3.tar.xz"
  sha256 "4e503d8685c0dc9a5a89f651a800e7c296f9dae01ba899f18667ae39d6b853df"
  version "0.1.3"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end
end