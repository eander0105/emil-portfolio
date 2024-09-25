{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # General depedencies
    docker-compose
    docker
    git

    # Frontend depedencies
    nodejs
  ];
}
