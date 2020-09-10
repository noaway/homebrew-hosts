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

  # depends_on "cmake" => :build

  def install
    # ENV.deparallelize  # if your formula fails when building in parallel
    # Remove unrecognized options if warned by configure
   # system "./configure", "--disable-debug",
   #                       "--disable-dependency-tracking",
   #                       "--disable-silent-rules",
   #                       "--prefix=#{prefix}"
    # system "cmake", ".", *std_cmake_args
    bin.install "hosts"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test celery`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
