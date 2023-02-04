import * as fs from 'fs';
console.log(process.argv)
const destinationDir = './browser-extension/popup/';
const sourceDir = './dist/svelte';

// fs.watch(sourceDir, {}, (eventType, filename) => {
//     console.log(eventType)
// })

fs.rmSync(destinationDir, {recursive: true, force: true});
fs.mkdirSync(destinationDir)
fs.cpSync(sourceDir, destinationDir, {recursive: true})
console.info(`copied files from ${sourceDir} to ${destinationDir}`);
