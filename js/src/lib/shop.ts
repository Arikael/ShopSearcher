import {searchTermPattern, shopDir} from '../config';

export class Shop {
    name: string
    url: string
    enabled = true
    tags: string[] = []

    constructor(shop: any | Shop) {
        if (shop) {
            this.url = shop.url
            this.name = new URL(this.url).host.replace('www.', '')
            this.enabled = shop.enabled
            this.tags = shop.tags
        }
    }
}

export async function getShops(configName: string): Promise<Shop[]> {
    const response = await fetch(`${shopDir}/${configName}`)
    const json = await response.json()

    return json['shops'].map(shop => new Shop(shop)).sort((a, b) => a.name.localeCompare(b.name))
}

export function createShopsWithSearchTerm(searchTerm: string, shops: Shop[]): Shop[] {
    return shops.map(shop => {
        const newShop = new Shop(shop)
        newShop.url = shop.url.replace(searchTermPattern, searchTerm)

        return newShop
    })
}

export function getDisplayConfigName(shop: string): string {
    let changedName = shop.replace('_', ' ')
    return changedName.substring(0, changedName.lastIndexOf('.'))
}
