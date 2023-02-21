<script lang="ts">
    import {Shop} from './shop'
    import Tag from './Tag.svelte'

    let tags: string[]
    let tagStates: Record<string, boolean> = {}
    export let shopsWithSearchTerms: Shop[] = []

    $: {
        tags = [...new Set(shopsWithSearchTerms.map(x => x.tags).flat())]
        for (let tag in tags) {
            tagStates[tag] = false
        }
    }

    let allShopsToggle = false

    const isTagEnabled = (tag: string): boolean => {
        return tagStates[tag]
    }

    const onToggleAllShops = () => {
        allShopsToggle = !allShopsToggle
        shopsWithSearchTerms.forEach(shop => shop.enabled = allShopsToggle)
        shopsWithSearchTerms = shopsWithSearchTerms
    }

    const toggleTag = (event) => {
        const tag = event.detail.tag
        const currentTagState = isTagEnabled(tag)
        const newTagState = !currentTagState

        if (!currentTagState) {
            shopsWithSearchTerms.filter(shop => shop.tags.includes(tag) && shop.enabled === false)
                .forEach(shop => shop.enabled = true)
        } else {
            shopsWithSearchTerms.filter(shop => shop.tags.includes(tag) && shop.enabled === true &&
                !shop.tags.filter(t => t !== tag).some(t => isTagEnabled(t)))
                .forEach(shop => shop.enabled = false)
        }

        shopsWithSearchTerms = shopsWithSearchTerms
        tagStates[tag] = newTagState
        tagStates = tagStates
    }
</script>
<div class="shop-list">
    <div class="tag-list">
        Special Tags:
        <div class="tags">
            {#each tags as tag}
                <Tag isClickable="true" isSelected="{isTagEnabled(tag)}" tag="{tag}" on:toggleTag="{toggleTag}"></Tag>
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
                        <Tag tag="{tag}"></Tag>
                    {/each}
                </td>
            </tr>
        {/each}
    </table>
</div>