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
            if(isWebExtension && isMobile()) {
                window.close()
            }

            oldSearchTerm = searchTerm
            document.querySelectorAll('a.shop-link').forEach((link) => {
                if (isWebExtension) {
                    browser.tabs.create({
                        url: link.href,
                        active: false
                    })
                } else {
                    window.open(link.href, '_blank')
                }
            })
        }
    }

    let timeout: NodeJS.Timeout
    const delay = 300
    const onSearchKeyUp = (e) => {
        if (e.which === 0) {
            return
        }

        shopsWithSearchTerms = []
        if (timeout) {
            clearTimeout(timeout)
        }
        timeout = setTimeout(() => {
            shopsWithSearchTerms = createShopsWithSearchTerm(searchTerm, shops)
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
{#if shopCount > 0}
    <div class="box box--info">
        <div class="small search-info">
            Please be aware that searching will open <strong>{shopCount} windows/tabs</strong> simultaneously
            {#if !isWebExtension}
                <br />For security reasons you need to allow popups when prompted by the browser (on top of the page)
            {/if}
            {#if isWebExtension && isMobile()}
                <br />this window will close itself, when hitting "search"
            {/if}
        </div>
        <ShopList shopsWithSearchTerms={shopsWithSearchTerms}/>
    </div>
{/if}