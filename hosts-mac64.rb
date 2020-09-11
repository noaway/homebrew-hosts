# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/0.1.2/hosts-mac64-0.1.2.tar.xz"
  sha256 "dfded6affc512c509c818b4fd4e36d450a3126a6a0507009b399090b15203ba8"
  version "0.1.2"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end
end