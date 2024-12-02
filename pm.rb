class Pm < Formula
    desc "A simple command line tool to navigate and execute commands through your project"
    homepage "https://github.com/abroudoux/pm"
    url "https://github.com/abroudoux/pm/archive/1.1.0.tar.gz"
    sha256 "644bb29989e7554b2413d419549c46485b2678db"
    version "1.1.0"

    def install
        bin.install "pm"
    end

    def install
        system "#{bin}/install.sh"
    end

    test do
        system "pm --help"
    end
end