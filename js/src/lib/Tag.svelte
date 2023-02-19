<script lang="ts">
    import {createEventDispatcher} from 'svelte';

    export let isClickable = false
    export let tag = ''
    const dispatch = createEventDispatcher()

    function toggleTag(tag: string) {
        if(!isClickable) {
            return
        }

        dispatch('toggleTag', {
            tag
        })
    }
</script>

<span class:tag-clickable="{isClickable}" class="tag" on:click={() => toggleTag(tag)}>
    {tag}
</span>