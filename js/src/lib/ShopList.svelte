<script lang="ts">
    import {Shop} from "./shop";
    import {onMount} from "svelte";
    export let shopsWithSearchTerms: Shop[] = []
    let shopListVisible = false
    let shopListHeader = ''

    onMount(() => {
        toggleShopList(false)
    })

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

<button class="text-only" on:click="{onShopListHeaderClick}">{shopListHeader}</button>
<div class="shop-list {shopListVisible ? '' : 'hidden'}">
    <ul class="flat">
        {#each shopsWithSearchTerms as shop}
            <li><a class="shop-link" target="_blank" href="{shop.searchUrl}">{shop.name}</a></li>
        {/each}
    </ul>
</div>