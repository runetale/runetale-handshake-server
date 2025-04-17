{
  description = "Runetale handshake server";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    let
      # Generate a user-friendly version number.
      version = builtins.substring 0 8 self.lastModifiedDate;
      # System types to support.
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      eachSystem = flake-utils.lib.eachSystem supportedSystems;
    in
    eachSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        defaultPackage = self.packages.${system}.runetale-handshake-server;
        packages = {
          containerImage = pkgs.dockerTools.buildLayeredImage
            {
              name = "runetale-handshake-server";
              contents = [
                self.packages.${system}.runetale-handshake-server
              ];
              config = {
                Entrypoint = [ "${pkgs.tini}/bin/tini" "--" ];
                Cmd = [ "${self.packages.${system}.runetale-handshake-server}" ];
              };
            };
          runetale-handshake-server =
            pkgs.buildGoModule {
              pname = "runetale-handshake-server";
              inherit version;
              src = self;
              vendorHash = "sha256-VGWmAAaakCvFOha1OGZu4A0z7WRlFJjHwZygCXX/OI1=";
              doCheck = false;
            };
        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs;
            [
              protoc-gen-go
              go_1_17
              goimports
              gopls
              protobuf
              protoc-gen-go-grpc
              docker-compose
              docker
              kustomize
            ];

          shellHook = ''
            export GOPATH=$GOPATH
            PATH=$PATH:$GOPATH/bin
            export GO111MODULE=on
          '';
        };
      }
    );
}
