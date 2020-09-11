# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Hosts < Formula
  desc "hosts tool"
  homepage "https://github.com/noaway/hosts"
  url "https://github.com/noaway/hosts/releases/download/0.1.1/hosts"
  sha256 "1092a9b72b657c29e2f19551d428058a16fded155266fd15cd19fe41faf36a9c"
  version "0.1.1"
  license "Apache-2.0"
  
  def install
    bin.install "hosts"
  end

  test do
    system "false"
  end
end
