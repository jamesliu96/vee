import './wasm_exec.js';

addEventListener('load', async () => {
  const go = new Go();
  const { instance } = await WebAssembly.instantiateStreaming(
    fetch('vee.wasm'),
    go.importObject
  );
  go.run(instance);

  document.body.style.margin = '0';
  const root = document.createElement('div');
  root.style.display = 'flex';
  const s0 = document.createElement('div');
  const s1 = document.createElement('div');
  const s2 = document.createElement('div');
  const s3 = document.createElement('div');
  s0.style.width = s1.style.width = s2.style.width = s3.style.width = '25vw';
  s2.style.display = 'flex';
  s2.style.flexDirection = 'column';
  document.body.append(root);
  root.append(s0, s1, s2, s3);
  const input = document.createElement('input');
  input.type = 'file';
  input.accept =
    'image/png,image/jpeg,image/gif,image/png,image/bmp,image/tiff,image/webp';
  input.addEventListener('change', () => {
    if (input.files?.length) {
      inputImg.src = URL.createObjectURL(input.files.item(0)!);
    }
  });
  s0.append(input);
  const inputImg = document.createElement('img');
  inputImg.style.width = '100%';
  s1.append(inputImg);
  const key = document.createElement('input');
  key.type = 'password';
  key.placeholder = 'passcode';
  s2.append(key);
  const button = document.createElement('button');
  button.textContent = 'RUN';
  button.addEventListener('click', async () => {
    const data = await input.files?.item(0)?.arrayBuffer();
    if (data) {
      const output = await Vee.faux(
        new Uint8Array(data),
        new TextEncoder().encode(key.value)
      );
      outputImg.src = URL.createObjectURL(new Blob([output]));
    }
  });
  s2.append(button);
  const outputImg = document.createElement('img');
  outputImg.style.width = '100%';
  s3.append(outputImg);
});
