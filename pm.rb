class Pm < Formula
    desc "A simple command line tool to navigate and execute commands through your project"
    homepage "https://github.com/abroudoux/pm"
    url "https://github.com/abroudoux/pm/archive/1.0.0.tar.gz"
    sha256 "8cbd152034b765d203007b4bfe30ccc2353a7902"
    version "1.0.0"

    def install
    bin.install "pm.sh" => "pm"
    end

    test do
        system "pm --file"
    end
end