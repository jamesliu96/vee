type Bytes = Uint8Array | Uint8ClampedArray;

declare namespace Vee {
  const faux: (data: Bytes, key: Bytes) => Promise<Bytes>;
}
