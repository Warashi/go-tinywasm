{
  pkgs,
  wasm-spec,
  ...
}:
let
  testsuite-json = pkgs.runCommand "json-testsuite" { } ''
    mkdir -p $out
    for f in ${wasm-spec}/test/core/*.wast; do
      echo "Converting $f to json"
      ${pkgs.wabt}/bin/wast2json $f -o "$out/$(basename $f).json"
    done
  '';
  wasmium-test = pkgs.callPackage ./wasmium-test.nix { };
in
pkgs.writeShellApplication {
  name = "wasmium-test";
  text = ''
    export WASMIUM_TEST_DIR="${testsuite-json}"
    for f in ${testsuite-json}/*.json; do
      echo "Running $f"
      ${wasmium-test}/bin/wasmium-test "$f"
    done
  '';
}
