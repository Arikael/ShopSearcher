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
    import {shopFiles} from '../config.ts';
    import {createShopsWithSearchTerm, getShops, getDisplayConfigName, Shop} from './shop';
    import {onMount} from 'svelte';
    import {isMobile} from "./utils.js";

    let selectedShopFile = shopFiles.length > 0 ? shopFiles[0] : ''
    let shopCount = 0
    let searchTerm = ''
    let shops: Shop[] = []
    let shopsWithSearchTerms: Shop[] = []
    let mobile = true

    const onSelectShop = async () => {
        if (selectedShopFile) {
            shops = await getShops(selectedShopFile)
            shopCount = shops.length
        }
    }

    onMount(() => {
        mobile = isMobile();
        toggleShopList(false)
        onSelectShop()

        if (selectedShopFile) {
            document.querySelector('.search-input').focus()
        }
    })

    const onSearch = () => {
        if (searchTerm && searchTerm.length > 0) {
            document.querySelectorAll('a.shop-link').forEach((link) => {
                window.open(link.href, '_blank')
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

    let shopListVisible = false
    let shopListHeader = ''
    const toggleShopList = (toggle: boolean = null) => {
        if (toggle === null) {
            toggle = !shopListVisible
        }

        shopListVisible = toggle
        shopListHeader = toggle ? '- Hide shops' : '+ Show shops'
    }

    const onShopListHeaderClick = () => {
        toggleShopList()
    }
</script>

{#if mobile}
    <p>
        <strong>It seems you are using a mobile browser<br/>
            Unfortunately the search probably won't work correctly<br/>
            Browser prevent (rightfully so) opening multiple tabs/windows programatically<br/>and
            while this can be allowed on a per page base on desktop it seems mobile browser always prevent it
        </strong>
    </p>
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
    <button class="search" disabled="{shopsWithSearchTerms.length > 0 && searchTerm ? '' : 'disabled'}"
            on:click={onSearch}>
        Search
    </button>
</div>
{#if shopCount > 0}
    <div class="box box--info">
        <div class="small search-info">
            Please be aware that searching will open <strong>{shopCount} windows/tabs</strong> simultaneously<br/>
            For security reasons you need to allow popups when prompted by the browser (on top of the page)
        </div>
        <button class="text-only" on:click="{onShopListHeaderClick}">{shopListHeader}</button>
        <div class="shop-list {shopListVisible ? '' : 'hidden'}">
            <ul class="flat">
                {#each shopsWithSearchTerms as shop}
                    <li><a class="shop-link" target="_blank" href="{shop.searchUrl}">{shop.name}</a></li>
                {/each}
            </ul>
        </div>
    </div>
{/if}