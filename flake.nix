{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    testsuite = {
      url = "github:WebAssembly/spec?dir=test/core";
      flake = false;
    };
  };

  outputs =
    {
      flake-utils,
      nixpkgs,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        overlays = [ ];
        pkgs = import nixpkgs {
          inherit system overlays;
        };
      in
      {
        devShells.default =
          with pkgs;
          mkShell {
            GOTOOLCHAIN = "local";
            buildInputs = [
              go
              wabt
            ];
          };
      }
    );
}
