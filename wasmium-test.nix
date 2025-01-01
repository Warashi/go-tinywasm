{
  lib,
  buildGoModule,
  ...
}:
buildGoModule {
  pname = "wasmium-test";
  version = "dev";
  vendorHash = null;

  src = ./.;

  subPackages = [ "wasmium-test" ];

  meta = with lib; {
    description = "wasmium test runner";
    homepage = "https://github.com/Warashi/wasmium";
    license = licenses.mit;
  };
}
