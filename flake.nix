{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    wasm-spec = {
      url = "github:WebAssembly/spec";
      flake = false;
    };
  };

  outputs =
    {
      flake-utils,
      nixpkgs,
      wasm-spec,
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
              gotools
              wabt
            ];
          };
        packages = rec {
          wasmium-test = pkgs.callPackage ./wasmium-test.nix { };
          test = import ./wasmium-test {
            inherit pkgs wasm-spec wasmium-test;
          };
        };
      }
    );
}
