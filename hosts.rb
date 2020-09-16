# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/v0.2.2/hosts-mac64-v0.2.2.tar.xz"
  sha256 "6252b3e6e23564f48b667fe5e7b391e40629780e8adc364f32ce5b844b7f8d05"
  version "0.2.2"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end
end