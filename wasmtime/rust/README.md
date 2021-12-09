Install Rust: 

```shell
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Install wasmtime: 

```shell
curl https://wasmtime.dev/install.sh -sSf | bash
```

```shell
$ rustup target add wasm32-wasi
$ rustc main.rs --target wasm32-wasi
$ wasmtime main.wasm
```
