# ShopSearcher

This is a simple app and firefox extensio built with vite and svelte


## Build
### Prerequisites
- node v16
- npm v8

it has been developed on Windows 11 but should also work with Linux.

run `npm run build` to run the svelte build

the firefox extension resides in `./browser-extension`.  
Note that the `./popup` directory and its content are auto-generated since it's basically
the same files as the web app.  
you can either run `node scripts/sync-for-firefox.js` manually or by   
running
- `npm run build:firefox`
- `npm run run:firefox`
- `npm run run:firefox-mobile`

those scripts will, among other things generate `./browser-extension/popup`

to build the extension and the corresponding zip file to upload to addons.mozilla.org   
run `npm run build:firefox`.   
the zip file will be in `./dist/firefox`

### Step by Step for Firefox Addon
- Install [node LTS](https://nodejs.org/en/)   
  npm should be included
- run `npm i`
- run `npm run build:firefox` to create the zip file with the extension.

## Deployment
- Mozilla Addon -> tbd
- Web App -> https://shop-searcher.netlify.app/