<style>
    .tag {
        padding: 4px 8px;
        background: rebeccapurple;
        border-radius: var(--default-border-radius);
        margin: 0 4px;
    }

    .tag-clickable:hover {
        cursor: pointer;
        background: #8e76a7;
    }

    .tag.tag-selected {
        background: #a05389;
    }

    @media (prefers-color-scheme: light) {
        .tag {
            background: #744d9f;
            color: #fff;
        }
    }
</style>
<script lang="ts">
    import {createEventDispatcher} from 'svelte';

    export let isClickable = false
    export let tag = ''
    export let isSelected = false
    const dispatch = createEventDispatcher()

    function toggleTag(tag: string) {
        if (!isClickable) {
            return
        }

        dispatch('toggleTag', {
            tag
        })
    }
</script>

<span class:tag-clickable="{isClickable}" class:tag-selected="{isSelected}" class="tag" on:click={() => toggleTag(tag)}>
    {tag}
</span>