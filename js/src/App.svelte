<script lang="ts">
    import Search from "./lib/Search.svelte";
    import {isWebExtension} from "./config.js";
</script>

<main>
    {#if !isWebExtension }
        <h1>Search Shops</h1>
    {/if}
    <div class="card">
        <Search/>
    </div>
</main>