{
  description = "A Nix-flake-based go development environment";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { system = system; config.allowUnfree = true; };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [ go ];
        };
      }
    );
}
