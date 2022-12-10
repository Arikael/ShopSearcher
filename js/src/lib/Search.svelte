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

    let selectedShopFile = shopFiles.length > 0 ? shopFiles[0] : ''
    let shopCount = 0
    let searchTerm = ''
    let shops: Shop[] = []
    let shopsWithSearchTerms: Shop[] = []

    const onSelectShop = async () => {
        if(selectedShopFile) {
            shops = await getShops(selectedShopFile)
            shopCount = shops.length
        }
    }

    onMount(() => {
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
        if(e.which === 0) {
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

<div class="">
    Shop
    <select class="search-select" bind:value={selectedShopFile}>
        {#each shopFiles as shopFile}
            <option value="{shopFile}">{getDisplayConfigName(shopFile)}</option>
        {/each}
    </select><br/>
    <input type="text" placeholder="enter search term" class="search-input"
           bind:value="{searchTerm}" on:keypress="{onSearchKeyUp}"/>
    <br/>
    <button class="search" disabled="{shopsWithSearchTerms.length > 0 && searchTerm ? '' : 'disabled'}" on:click={onSearch}>
        Search
    </button>
</div>
<div class="test">

</div>
{#if shopCount > 0}
    <div class="box box--info">
        <div class="small search-info">
            Please be aware that searching will open {shopCount} windows/tabs simultaneously<br/>
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