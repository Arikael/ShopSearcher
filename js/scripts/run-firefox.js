import * as dotenv from 'dotenv'
import child_process from 'child_process'

dotenv.config()

const browser = process.argv[3].toLowerCase()

if (browser === 'firefox') {
    child_process.execSync('web-ext run --source-dir ./browser-extension', {stdio: 'inherit'})
} else if (browser === 'firefox-mobile') {
    if (!process.env.ANDROID_DEVICE) {
        throw new Error('ANDROID_DEVICE is not defined. At a .env file and put in the device (see readme)')
    }

    child_process.execSync(`web-ext run --source-dir ./browser-extension -t firefox-android --android-device ${process.env.ANDROID_DEVICE}`,
        {stdio: 'inherit'})
}