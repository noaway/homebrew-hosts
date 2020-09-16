# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/v0.2.1/hosts-mac64-v0.2.1.tar.xz"
  sha256 "1c9e57ff984e11f3a98c2721de6944c92869e77ddb6d2eeaa0156ca004e50f9b"
  version "0.2.1"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end
end