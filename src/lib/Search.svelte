<style>
    .search-select,
    .search,
    .search-input {
        margin-bottom: 1rem;
    }

    .search,
    .search-input {
        font-size: 1.5rem;
        padding: 1rem;
    }

    .search-input {
        width: 100%;
    }

    .search-info {
        margin-bottom: 1rem;
    }
</style>
<script lang="ts">
    import {isWebExtension, shopFiles} from '../config.ts';
    import {createShopsWithSearchTerm, getShops, getDisplayConfigName, Shop} from './shop';
    import {onMount} from 'svelte';
    import {isMobile} from "./utils.js";
    import ShopList from "./ShopList.svelte";
    import MobileInfo from "./MobileInfo.svelte";

    let selectedShopFile = shopFiles.length > 0 ? shopFiles[0] : ''
    let shopCount = 0
    let searchTerm = ''
    let oldSearchTerm = searchTerm
    let shops: Shop[] = []
    let shopsWithSearchTerms: Shop[] = []
    let mobile = true

    $: canSearch = shopsWithSearchTerms.length > 0 && searchTerm && (oldSearchTerm.length > 0
        && searchTerm != oldSearchTerm || oldSearchTerm.length == 0)

    const onSelectShop = async () => {
        if (selectedShopFile) {
            shops = await getShops(selectedShopFile)
            shopCount = shops.length
        }
    }

    onMount(() => {
        mobile = isMobile();
        onSelectShop()

        if (selectedShopFile) {
            document.querySelector('.search-input').focus()
        }
    })

    const onSearch = () => {
        if (searchTerm && searchTerm.length > 0) {
            if (isWebExtension && isMobile()) {
                window.close()
            }

            oldSearchTerm = searchTerm
            shopsWithSearchTerms.filter(shop => shop.enabled).forEach((shop) => {
                if (isWebExtension) {
                    browser.tabs.create({
                        url: shop.url,
                        active: false
                    })
                } else {
                    window.open(shop.url, '_blank')
                }
            });
        }
    }

    let timeout: NodeJS.Timeout
    const delay = 300
    const onSearchKeyUp = (e) => {
        if (e.which === 0) {
            return
        }

        if (timeout) {
            clearTimeout(timeout)
        }
        timeout = setTimeout(() => {
            shopsWithSearchTerms = createShopsWithSearchTerm(searchTerm, shops, shopsWithSearchTerms)
        }, delay)
    }
</script>
{#if mobile && !isWebExtension}
    <MobileInfo/>
{/if}
<div>
    Shop
    <select class="search-select" bind:value={selectedShopFile}>
        {#each shopFiles as shopFile}
            <option value="{shopFile}">{getDisplayConfigName(shopFile)}</option>
        {/each}
    </select><br/>
    <input type="text" placeholder="enter search term" class="search-input"
           bind:value="{searchTerm}" on:keypress="{onSearchKeyUp}"/>
    <br/>
    <button class="search" disabled="{canSearch ? '' : 'disabled'}"
            on:click={onSearch}>
        Search
    </button>
</div>
<div class="small search-info">
    Please be aware that searching will open up to <strong>{shopCount} windows/tabs</strong> simultaneously
    {#if !isWebExtension}
        <br/>For security reasons you need to allow popups when prompted by the browser (on top of the page)
    {/if}
    {#if isWebExtension && isMobile()}
        <br/>this window will close itself, when hitting "search"
    {/if}
</div>
{#if shopCount > 0}
    <div class="box box--info">
        <div class="box-header">
            Available Shops
        </div>
        {#if searchTerm.length === 0}
            <div class="default-padding">
                enter a search term to see the shops
            </div>
        {:else}
            <ShopList shopsWithSearchTerms={shopsWithSearchTerms}/>
        {/if}
    </div>
{/if}