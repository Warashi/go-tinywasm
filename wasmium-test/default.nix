{
  pkgs,
  wasm-spec,
  wasmium-test,
  ...
}:
rec {
  testsuite-json = pkgs.runCommand "json-testsuite" { } ''
    mkdir -p $out
    for f in ${wasm-spec}/test/core/*.wast; do
      echo "Converting $f to json"
      ${pkgs.wabt}/bin/wast2json $f -o "$out/$(basename $f).json"
    done
  '';
  run = pkgs.writeShellApplication {
    name = "wasmium-test";
    text = ''
      echo "Running tests"
      for f in ${testsuite-json}/*; do
        echo "Running $f"
        ${wasmium-test}/bin/wasmium-test "$f"
      done
    '';
  };
}
