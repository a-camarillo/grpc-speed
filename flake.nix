{
  description = "A microservice for calculating average speed";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-dawin" "x86_64-dawin" ];

      perSystem = { config, self', inputs', pkgs, system, ... }:
        {
          devShells = {
            default = pkgs.mkShell {
              packages = [
                pkgs.go
                pkgs.protoc-gen-go
                pkgs.protoc-gen-go-grpc
                pkgs.protobuf
              ];
            };
          };
        };
    };
}
