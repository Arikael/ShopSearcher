import child_process from 'child_process'
import * as fs from 'fs';
console.log(process.argv)
const popupDir = './browser-extension/popup/';
const svelteDir = './dist/svelte';

fs.watch(svelteDir, {}, (eventType, filename) => {
    console.log(eventType)
})

fs.rmSync(popupDir, {recursive: true, force: true});
fs.mkdirSync(popupDir)
fs.cpSync(svelteDir, popupDir, {recursive: true})
console.info(`copied files from ${svelteDir} to ${popupDir}`);
