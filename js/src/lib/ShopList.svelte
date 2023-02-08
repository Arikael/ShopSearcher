<script lang="ts">
    import {Shop} from "./shop";
    import {onMount} from "svelte";

    let tags: string[] = []
    export let shopsWithSearchTerms: Shop[] = []
    $: {
        tags = [...new Set(shopsWithSearchTerms.map(x => x.tags).flat())]
    }

    let shopListHeader = ''
    let allShopsToggle = false

    const onToggleAllShops = () => {
        allShopsToggle = !allShopsToggle
        shopsWithSearchTerms.forEach(shop => shop.enabled = allShopsToggle)
        shopsWithSearchTerms = shopsWithSearchTerms
    }

    const toggleTag = (tag) => {
        shopsWithSearchTerms.filter(shop => shop.tags.includes(tag))
            .forEach(shop => shop.enabled = !shop.enabled)
        shopsWithSearchTerms = shopsWithSearchTerms
    }
</script>
<div class="shop-list">
    <div class="tag-list">
        Special Tags:
        <div class="tags">
            {#each tags as tag}
                <span class="tag tag-clickable" on:click={() => toggleTag(tag)}>{tag}</span>
            {/each}
        </div>
    </div>
    <table>
        <tr>
            <th>
                <input type="checkbox" on:change={onToggleAllShops}/>
            </th>
            <th>Shop</th>
            <th>Tags</th>
        </tr>
        {#each shopsWithSearchTerms as shop}
            <tr>
                <td>
                    <input type="checkbox" bind:checked={shop.enabled} value={shop.name}/>
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