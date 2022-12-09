import {searchTermPattern, shopDir} from '../config';

export class Shop {
    name: string
    searchUrl: string

    constructor(url: string = '') {
        if (url) {
            this.searchUrl = url
            this.name = new URL(url).host.replace('www.', '')
        }
    }
}

export async function getShops(configName: string): Promise<Shop[]> {
    const response = await fetch(`${shopDir}/${configName}`)
    const json = await response.json()

    return json['shops'].map(shop => new Shop(shop)).sort((a, b) => a.name.localeCompare(b.name))
}

export function constructSearchUrls(searchTerm: string, shops: Shop[]): Shop[] {
    return shops.map(shop => {
        const newShop = new Shop()
        newShop.name = shop.name
        newShop.searchUrl = shop.searchUrl.replace(searchTermPattern, searchTerm)

        return newShop
    })
}

export function getDisplayConfigName(shop: string): string {
    let changedName = shop.replace('_', ' ')
    return changedName.substring(0, changedName.lastIndexOf('.'))
}
