<script lang="ts">
    import {Shop} from "./shop";
    import {onMount} from "svelte";

    export let shopsWithSearchTerms: Shop[] = []
    let shopListVisible = false
    let shopListHeader = ''
    let allShopsToggle = false

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

    const onToggleAllShops = () => {
        allShopsToggle = !allShopsToggle
        shopsWithSearchTerms.forEach(shop => shop.enabled = allShopsToggle)
        shopsWithSearchTerms = shopsWithSearchTerms
    }
</script>
<button class="text-only" on:click="{onShopListHeaderClick}">{shopListHeader}</button>
<div class="shop-list {shopListVisible ? '' : 'hidden'}">
    <table>
        <tr>
            <th>
                <input type="checkbox" on:change={onToggleAllShops} />
            </th>
            <th>Shop</th>
            <th>Tags</th>
        </tr>
        {#each shopsWithSearchTerms as shop}
            <tr>
                <td>
                    <input type="checkbox" bind:checked={shop.enabled} value={shop.name} />
                </td>
                <td>
                    <a class="shop-link" target="_blank" href="{shop.url}">{shop.name}</a>
                </td>
                <td>
                    {#each shop.tags as tag}
                        <span class="tag">{tag}</span>
                    {/each}
                </td>
            </tr>
        {/each}
    </table>
</div>