# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/v0.2.4/hosts-mac64-v0.2.4.tar.xz"
  sha256 "a812ddd183fa8936f69c3827dab90c84366c337b4db3b900fbb24223e04426d0"
  version "0.2.4"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end
end