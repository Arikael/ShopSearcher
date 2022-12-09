<script>
     fetch('configs/swiss_wine.json').then(x => console.log(x))


    const search = () => {

    }
</script>

<input type="text">
<button on:click={search}>
    Search
</button>