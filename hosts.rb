# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/v0.2.3/hosts-mac64-v0.2.3.tar.xz"
  sha256 "0390d1ce1dd44a5abfb11479064209887242ba9d07baa29ef0fdf73edba33373"
  version "0.2.3"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end
end