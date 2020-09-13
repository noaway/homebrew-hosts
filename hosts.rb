# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/v0.1.6/hosts-mac64-v0.1.6.tar.xz"
  sha256 "b32814a9ae2c6c2a979fcb6388e90be4da738d5f6fc7c235b4e4931943845d96"
  version "0.1.6"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end
end