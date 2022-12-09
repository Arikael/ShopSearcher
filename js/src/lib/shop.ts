import {searchTermPattern, shopDir} from '../config';

export class Shop {
    name: string
    searchUrl: string

    constructor(url: string = '') {
        if (url) {
            this.searchUrl = url
            this.name = new URL(url).host
        }
    }
}

export async function getShops(configName: string): Promise<Shop[]> {
    const response = await fetch(`${shopDir}/${configName}`)
    const json = await response.json()

    return json['shops'].map(shop => new Shop(shop))
}

export function constructSearchUrls(searchTerm: string, shops: Shop[]): Shop[] {
    return shops.map(shop => {
        const newShop = new Shop()
        newShop.name = shop.name
        newShop.searchUrl = shop.searchUrl.replace(searchTermPattern, searchTerm)

        return newShop
    })
}

export function getShopName(shop: string): string {
    let changedName = shop.replace('_', ' ')
    return changedName.substring(0, changedName.lastIndexOf('.'))
}
