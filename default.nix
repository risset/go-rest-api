with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "go-rest-api";
  buildInputs = with pkgs; [
    go
    gopkgs
    delve
    gopls
    go-tools
  ];
}
